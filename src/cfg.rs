use crate::parser::AddressedInst;
use std::collections::HashMap;

#[derive(Copy, Clone)]
enum Eflags {
    OF,
    SF,
    ZF,
    AF,
    CF,
    PF,
    MAX,
}

trait Expr: std::string::ToString {
    // 返回 !(Expr) 的简化版本.
    fn not(&self) -> Box<dyn Expr>;
}

enum CompOp {
    L,  // <
    LE, // <=
    E,  // ==
    NE, // !=
    G,  // >
    GE, // >=
}

// compare expression
struct CompExpr {
    op: CompOp,
    // 本来 left, right 类型是 Box<dyn Expr> 后来一想没必要, 汇编嘛, 哪有嵌套表达式.
    left: String,
    right: String,
}

impl std::string::ToString for CompExpr {
    fn to_string(&self) -> String {
        let op = match self.op {
            CompOp::L => "<",
            CompOp::LE => "<=",
            CompOp::E => "==",
            CompOp::NE => "!=",
            CompOp::G => ">",
            CompOp::GE => ">=",
        };
        format!("{} {} {}", self.left, op, self.right)
    }
}

impl Expr for CompExpr {
    fn not(&self) -> Box<dyn Expr> {
        let op = match self.op {
            CompOp::L => CompOp::GE,
            CompOp::LE => CompOp::G,
            CompOp::E => CompOp::NE,
            CompOp::NE => CompOp::E,
            CompOp::G => CompOp::LE,
            CompOp::GE => CompOp::L,
        };
        Box::new(CompExpr {
            op,
            left: self.left.clone(),
            right: self.right.clone(),
        })
    }
}

impl CompExpr {
    fn new(op: CompOp, left: String, right: String) -> Self {
        Self { op, left, right }
    }
}
struct EflagsWriter {
    eflags: &'static [Eflags],
}

const FULL_EFLAGS: [Eflags; Eflags::MAX as usize] = [
    Eflags::OF,
    Eflags::SF,
    Eflags::ZF,
    Eflags::AF,
    Eflags::CF,
    Eflags::PF,
];

// 按照 inst 排序, 可以二分查找.
// 像 add/imul 这种指令也会修改 eflags, 但这里并未包含, 意味着:
// cmp rax, rbx,
// add rax, rdi
// jbe label  # 这里我们会使用 cmp 生成表达式, 但运行时是依据 add 结果来跳转.
// 或许以后可以修正下这个.
const EFLAGS_WRITER: [(&'static str, EflagsWriter); 4] = [
    (
        "and",
        EflagsWriter {
            eflags: &FULL_EFLAGS,
        },
    ),
    (
        "cmp",
        EflagsWriter {
            eflags: &FULL_EFLAGS,
        },
    ),
    (
        "sub",
        EflagsWriter {
            eflags: &FULL_EFLAGS,
        },
    ),
    (
        "test",
        EflagsWriter {
            eflags: &FULL_EFLAGS,
        },
    ),
];

fn get_eflags_writer(inst: &str) -> Option<&'static EflagsWriter> {
    for instattr in &EFLAGS_WRITER {
        if instattr.0 == inst {
            return Some(&instattr.1);
        }
    }
    return None;
}

struct JccAttr {
    eflags: &'static [Eflags],
}

// 按照 inst 从小到大.
const JCC_ATTR: [(&'static str, JccAttr); 33] = [
    (
        "ja",
        JccAttr {
            eflags: &[Eflags::CF, Eflags::ZF],
        },
    ),
    (
        "jae",
        JccAttr {
            eflags: &[Eflags::CF],
        },
    ),
    (
        "jb",
        JccAttr {
            eflags: &[Eflags::CF],
        },
    ),
    (
        "jbe",
        JccAttr {
            eflags: &[Eflags::CF, Eflags::ZF],
        },
    ),
    (
        "jc",
        JccAttr {
            eflags: &[Eflags::CF],
        },
    ),
    ("jcxz", JccAttr { eflags: &[] }), // CX
    (
        "je",
        JccAttr {
            eflags: &[Eflags::ZF],
        },
    ),
    ("jecxz", JccAttr { eflags: &[] }), // ECX
    (
        "jg",
        JccAttr {
            eflags: &[Eflags::ZF, Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jge",
        JccAttr {
            eflags: &[Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jl",
        JccAttr {
            eflags: &[Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jle",
        JccAttr {
            eflags: &[Eflags::ZF, Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jna",
        JccAttr {
            eflags: &[Eflags::CF, Eflags::ZF],
        },
    ),
    (
        "jnae",
        JccAttr {
            eflags: &[Eflags::CF],
        },
    ),
    (
        "jnb",
        JccAttr {
            eflags: &[Eflags::CF],
        },
    ),
    (
        "jnbe",
        JccAttr {
            eflags: &[Eflags::CF, Eflags::ZF],
        },
    ),
    (
        "jnc",
        JccAttr {
            eflags: &[Eflags::CF],
        },
    ),
    (
        "jne",
        JccAttr {
            eflags: &[Eflags::ZF],
        },
    ),
    (
        "jng",
        JccAttr {
            eflags: &[Eflags::ZF, Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jnge",
        JccAttr {
            eflags: &[Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jnl",
        JccAttr {
            eflags: &[Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jnle",
        JccAttr {
            eflags: &[Eflags::ZF, Eflags::SF, Eflags::OF],
        },
    ),
    (
        "jno",
        JccAttr {
            eflags: &[Eflags::OF],
        },
    ),
    (
        "jnp",
        JccAttr {
            eflags: &[Eflags::PF],
        },
    ),
    (
        "jns",
        JccAttr {
            eflags: &[Eflags::SF],
        },
    ),
    (
        "jnz",
        JccAttr {
            eflags: &[Eflags::ZF],
        },
    ),
    (
        "jo",
        JccAttr {
            eflags: &[Eflags::OF],
        },
    ),
    (
        "jp",
        JccAttr {
            eflags: &[Eflags::PF],
        },
    ),
    (
        "jpe",
        JccAttr {
            eflags: &[Eflags::PF],
        },
    ),
    (
        "jpo",
        JccAttr {
            eflags: &[Eflags::PF],
        },
    ),
    ("jrcxz", JccAttr { eflags: &[] }), // RCX
    (
        "js",
        JccAttr {
            eflags: &[Eflags::SF],
        },
    ),
    (
        "jz",
        JccAttr {
            eflags: &[Eflags::ZF],
        },
    ),
];

fn get_jcc_attr(inst: &str) -> Option<&'static JccAttr> {
    for (instmnem, attr) in &JCC_ATTR {
        if *instmnem == inst {
            return Some(attr);
        }
    }
    return None;
}

pub enum EdgeKind {
    // jmp 意味着控制流总是无条件地从 from 跳转到 to.
    Jmp,
    // Jcc, 意味着控制流在一定条件成立时会从 from 跳转到 to, 这里 desc 用于描述这个条件.
    Jcc(/* desc */ String),
    // Serial, 意味着 from, to 是邻近的, 控制流在一定条件成立时会从 from 执行到 to, desc 用于描述这个条件.
    Serial(/* desc */ String),
}

// ID 等于 node/edge/inst 在 Graph 中的 index.
pub type ID = usize;

/* a CFGEdge means that the control flow will go to `to` from `from` when the `expr` is true. */
pub struct Edge {
    pub from: ID,
    pub to: ID,
    pub kind: EdgeKind,
}

pub struct Node {
    pub id: ID, // 从 0 开始
    // 正常情况下, insts 总是不为空. 如果 insts 为空, 意味着 Node 是个匿名块,
    // 比如 jmp rax 这种目的 block 就对应着一个匿名块.
    pub insts: Vec<ID>, // 指令在 Graph::insts 中的下标.
}

pub struct Graph {
    pub insts: Vec<AddressedInst>,
    pub nodes: Vec<Node>,
    pub edges: Vec<Edge>,
}

impl Graph {
    fn new() -> Self {
        Self {
            insts: Vec::new(),
            nodes: Vec::new(),
            edges: Vec::new(),
        }
    }

    fn contain(&self, addr: u64) -> bool {
        let first_addr = if let Some(inst) = self.insts.first() {
            inst.addr
        } else {
            return false;
        };
        let last_addr = if let Some(inst) = self.insts.last() {
            inst.addr
        } else {
            return false;
        };
        return addr >= first_addr && addr <= last_addr;
    }

    fn new_node(&mut self) -> ID {
        let id = self.nodes.len();
        self.nodes.push(Node {
            id,
            insts: Vec::new(),
        });
        return id;
    }
    fn add_inst(&mut self, nodeid: ID, instid: ID) {
        self.nodes[nodeid].insts.push(instid);
    }
    fn add_edge(&mut self, from: ID, to: ID, kind: EdgeKind) {
        self.edges.push(Edge { from, to, kind });
    }
    fn contain_addr(&self, node: ID, addr: u64) -> bool {
        let first_addr = if let Some(&inst) = self.nodes[node].insts.first() {
            self.insts[inst].addr
        } else {
            return false;
        };
        let last_addr = if let Some(&inst) = self.nodes[node].insts.last() {
            self.insts[inst].addr
        } else {
            return false;
        };
        return addr >= first_addr && addr <= last_addr;
    }

    fn find_inst_idx(&self, node: ID, addr: u64) -> Option<usize> {
        for (inst_idx, &instid) in self.nodes[node].insts.iter().enumerate() {
            if self.insts[instid].addr == addr {
                return Some(inst_idx);
            }
        }
        return None;
    }

    fn find_node(&self, addr: u64) -> Option<ID> {
        for node in 0..self.nodes.len() {
            if self.contain_addr(node, addr) {
                return Some(node);
            }
        }
        return None;
    }

    fn update_edge_from(&mut self, oldfrom: ID, newfrom: ID) {
        for edge in &mut self.edges {
            if edge.from == oldfrom {
                edge.from = newfrom;
            }
        }
    }

    fn split_node(&mut self, container: ID, addr: u64) -> ID {
        let inst_idx = self.find_inst_idx(container, addr).expect(&format!(
            "split_node: container={} addr={}",
            container, addr
        ));
        if inst_idx == 0 {
            return container;
        }
        let newnode = self.new_node();
        self.nodes[newnode].insts = self.nodes[container].insts.split_off(inst_idx);
        self.update_edge_from(container, newnode);
        self.add_edge(container, newnode, EdgeKind::Jmp);
        return newnode;
    }
}

// How to construct a CFGEdge? we use status flags in EFLAGS such as CF, ZF, PF, etc. because of the existence of
// SETC-like instruction which can modify status flags individually, we can't use the eflags as a whole..
// So now we handle the status flags individually, we add each status flag which will be modified by the instruction
// to the Output field of the Instruction, and add any status flag to Input if the instruction
// will use this status flag.
//
// We will attempt to beautify the expression used in CFGEdge:
//
// ```
// cfgnode1:
//  cmp rax, rbx
//  jg cfgnode2
// ```
//
// the expression of CFGEdge introduced by JG is `rax > rbx`, not `ZF = 0 && (SF = OF)`,
// `rax > rbx` is more meanful.
//
// There is a lot of states when we construct the CFG, so carefully here!
struct Context {
    // current CFGNode, instruction will be pushed in this node.
    // None means that we should construct a new CFGNode.
    curnode: Option<ID>,
    // if pending_edge isn't nil, it means that the next new CFGNode will be `to` field of pending_edge,
    // and the pending_edge should be added to CFGGraph.
    pending_edge: Option<(/* from */ ID, EdgeKind)>,
    // if newnodes[Addr] exists, it means that instruction after Addr should be pushed to newnodes[Addr].
    // newnodes 中也可能存放着不属于当前函数的指令地址, clang 可能会利用 jmp 实现 call 效果, 此时
    // jmp dstaddr 指向着另外一个函数, 我们也会将这种 dstaddr 放入 newnodes 中, 这样如下 blk0/blk1 会指向着同一个 block, 效果好一点.
    //   blk0:
    //    mov rax, rbx
    //    jmp f1
    //   blk1:
    //    mov rcx, rbx
    //    jmp f1
    newnodes: HashMap<u64, ID>,
    // last_writer[eflags] 记录着最后一条修改 eflags 的指令.
    last_writer: [Option<ID>; Eflags::MAX as usize],
    graph: Graph,
}

impl Context {
    fn new() -> Self {
        Context {
            curnode: None,
            pending_edge: None,
            newnodes: HashMap::new(),
            last_writer: [None; Eflags::MAX as usize],
            graph: Graph::new(),
        }
    }

    fn begin_node(&mut self, node: ID) {
        self.curnode = Some(node);
        if let Some(pending_edge) = self.pending_edge.take() {
            self.graph.add_edge(pending_edge.0, node, pending_edge.1);
        }
        // 本来我以为 last_writer 要清空, 设想如下例子:
        // blk0:
        //   cmp rdi, rax
        //   jae blk1
        // blk3:
        //   mov rax, rsi
        //   # 此时这里 jb 依赖的 eflags 来自于 blk0/blk3,
        //   # 如果我们不清空 last_writer, as2cfg 会使用 blk0 的 cmp 来作为这里 jb 依据, 将生成错误的表达式.
        //   jb blk2
        // blk1:
        //   cmp rbx, rax
        //   jmp blk3
        //
        // 但后来一想, gcc 应该不会生成这样糟糕的代码, 就算 blk2 中的 jb 依赖于 blk0/blk3 中的 cmp,
        // 编译器这里生成的 cmp 应该也具有相同的操作数, 所以 as2cfg 使用 blk0 的 cmp 生成表达式倒也能看.
        // self.last_writer = [None; Eflags::MAX as usize];
    }

    fn end_node(&mut self) {
        self.curnode = None;
    }

    fn feed_inst(&mut self, inst_id: ID) {
        let inst_addr = self.graph.insts[inst_id].addr;
        let inst_mnem = self.graph.insts[inst_id].inst.mnem;
        let curnode = if let Some(&newnode) = self.newnodes.get(&inst_addr) {
            if let Some(curnode) = self.curnode {
                self.graph.add_edge(curnode, newnode, EdgeKind::Jmp);
            }
            self.begin_node(newnode);
            self.newnodes.remove(&inst_addr);
            newnode
        } else if let Some(curnode) = self.curnode {
            curnode
        } else {
            let nodeid = self.graph.new_node();
            self.begin_node(nodeid);
            nodeid
        };
        self.graph.add_inst(curnode, inst_id);
        if let Some(eflags_writer) = get_eflags_writer(inst_mnem) {
            for &eflag in eflags_writer.eflags {
                self.last_writer[eflag as usize] = Some(inst_id);
            }
        }
    }

    // unknown jmp 意味着目的地地址不明确, 比如 `jmp rax` 这种.
    fn add_unknown_jmp(&mut self, kind: EdgeKind) {
        let dstnode = self.graph.new_node();
        self.graph.add_edge(self.curnode.unwrap(), dstnode, kind);
    }

    fn add_jmp(&mut self, cur_addr: u64, dst_addr: u64, kind: EdgeKind) {
        let edge_to = if dst_addr > cur_addr || !self.graph.contain(dst_addr) {
            // clang 有些时候会使用 jmp 实现函数调用, 此时 dst_addr 不属于当前函数.
            if let Some(&newnode) = self.newnodes.get(&dst_addr) {
                newnode
            } else {
                let newnode = self.graph.new_node();
                self.newnodes.insert(dst_addr, newnode);
                newnode
            }
        } else {
            let oldnode = self
                .graph
                .find_node(dst_addr)
                .expect(&format!("invalid jmpdst: {}", dst_addr));
            let newnode = self.graph.split_node(oldnode, dst_addr);
            if self.curnode == Some(oldnode) {
                self.curnode = Some(newnode);
            }
            newnode
        };
        self.graph.add_edge(self.curnode.unwrap(), edge_to, kind);
    }

    fn add_jmp_helper(&mut self, jcc_inst_id: ID, kind: EdgeKind) {
        let inst_addr = self.graph.insts[jcc_inst_id].addr;
        if let Some(dst_addr) = self.graph.insts[jcc_inst_id].inst.op0_as_imme() {
            self.add_jmp(inst_addr, dst_addr, kind);
        } else {
            self.add_unknown_jmp(kind);
        }
    }
}

fn is_end(inst: &str) -> bool {
    return inst == "ret";
}

fn is_jmp(inst: &str) -> bool {
    return inst == "jmp";
}

fn is_test(inst: &str) -> bool {
    return inst == "test" || inst == "and";
}

fn is_cmp(inst: &str) -> bool {
    return inst == "cmp" || inst == "sub";
}

fn expr2str(expr: &Option<Box<dyn Expr>>) -> String {
    expr.as_ref().map(|v| v.to_string()).unwrap_or_default()
}

fn test_beautify(inst: &AddressedInst, is_equal: bool) -> Option<Box<dyn Expr>> {
    let op = inst.inst.op01_as_single()?;
    let cmpop = if is_equal { CompOp::E } else { CompOp::NE };
    return Some(Box::new(CompExpr::new(
        cmpop,
        op.to_string(),
        "0".to_string(),
    )));
}

fn cmp_beautify(inst: &AddressedInst, cmpop: CompOp) -> Option<Box<dyn Expr>> {
    return Some(Box::new(CompExpr::new(
        cmpop,
        inst.inst.op0_as_str().unwrap(),
        inst.inst.op1_as_str().unwrap(),
    )));
}

fn expr_beautify(
    eflag_writer: Option<&AddressedInst>,
    jcc: &AddressedInst,
) -> Option<Box<dyn Expr>> {
    match jcc.inst.mnem {
        "je" | "jz" => {
            // ZF = 1
            let inst = eflag_writer?;
            if is_test(inst.inst.mnem) {
                return test_beautify(inst, true);
            }
            if is_cmp(inst.inst.mnem) {
                return cmp_beautify(inst, CompOp::E);
            }
            return None;
        }
        "jne" | "jnz" => {
            // ZF = 0
            let inst = eflag_writer?;
            if is_test(inst.inst.mnem) {
                return test_beautify(inst, false);
            }
            if is_cmp(inst.inst.mnem) {
                return cmp_beautify(inst, CompOp::NE);
            }
            return None;
        }
        "jcxz" => {
            return Some(Box::new(CompExpr::new(
                CompOp::E,
                "cx".to_string(),
                "0".to_string(),
            )));
        }
        "jecxz" => {
            return Some(Box::new(CompExpr::new(
                CompOp::E,
                "ecx".to_string(),
                "0".to_string(),
            )));
        }
        "jrcxz" => {
            return Some(Box::new(CompExpr::new(
                CompOp::E,
                "rcx".to_string(),
                "0".to_string(),
            )));
        }
        _ => {
            let inst = eflag_writer?;
            if !is_cmp(inst.inst.mnem) {
                return None;
            }
            // 可能不太对==
            let cmpop = match jcc.inst.mnem {
                "ja" | "jg" | "jnbe" | "jnle" => CompOp::G,
                "jae" | "jge" | "jnb" | "jnl" => CompOp::GE,
                "jb" | "jl" | "jnae" | "jnge" => CompOp::L,
                "jbe" | "jle" | "jna" | "jng" => CompOp::LE,
                _ => return None,
            };
            return cmp_beautify(inst, cmpop);
        }
    }
}

pub fn insts_cfg(insts: Vec<AddressedInst>) -> Graph {
    let mut ctx = Context::new();
    ctx.graph.insts = insts;
    for inst_id in 0..ctx.graph.insts.len() {
        ctx.feed_inst(inst_id);
        let inst_mnem = ctx.graph.insts[inst_id].inst.mnem;
        if is_end(inst_mnem) {
            ctx.end_node();
            continue;
        }
        if is_jmp(inst_mnem) {
            ctx.add_jmp_helper(inst_id, EdgeKind::Jmp);
            ctx.end_node();
            continue;
        }
        if let Some(jccattr) = get_jcc_attr(inst_mnem) {
            let writer_inst_id = {
                let mut inst_id = None;
                for &eflag in jccattr.eflags {
                    let last_writer = ctx.last_writer[eflag as usize];
                    if inst_id == last_writer {
                        continue;
                    }
                    if inst_id.is_none() {
                        inst_id = last_writer;
                    } else {
                        // 此时意味着 jcc 依赖着的 eflag 被不同的指令修改, 这时我们无能为力生成表达式.
                        inst_id = None;
                        break;
                    }
                }
                inst_id
            };
            let jmpexpr = expr_beautify(
                writer_inst_id.map(|v| &ctx.graph.insts[v]),
                &ctx.graph.insts[inst_id],
            );
            ctx.add_jmp_helper(inst_id, EdgeKind::Jcc(expr2str(&jmpexpr)));
            let notjmpexpr = jmpexpr.map(|v| v.not());
            ctx.pending_edge = Some((
                ctx.curnode.unwrap(),
                EdgeKind::Serial(expr2str(&notjmpexpr)),
            ));
            ctx.end_node();
            continue;
        }
    }
    if let Some((fromnode, kind)) = ctx.pending_edge.take() {
        // 这种情况意味着最后一条指令是 `Jcc dst`, 不合法的输入, 这里稍微兼容一下.
        let dstnode = ctx.graph.new_node();
        ctx.graph.add_edge(fromnode, dstnode, kind);
    }
    return ctx.graph;
}
