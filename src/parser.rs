use lalrpop_util::lalrpop_mod;
use std::num::NonZeroU32;

// Mnemonics and Register is always lower case.
// base, index 总是小写的寄存器名.
#[derive(Copy, Clone, PartialEq, Eq)]
pub struct MemOffset {
    disp: Option<i64>,
    base: Option<&'static str>,
    // index, scale 公用一个 Option 就可以, 反正他俩要么同时为 Some, 要么同时为 None.
    index_scale: Option<(&'static str, NonZeroU32)>,
}

#[derive(Copy, Clone, PartialEq, Eq)]
pub enum Operand {
    Reg(&'static str), // 寄存器操作数, 其内存放着小写的寄存器名.
    Imme(u64),         // 立即数
    Mem {
        // 内存操作数
        // Segment Selector, None means that we don't specify a segment selector explicitly.
        seg_sel: Option<&'static str>,
        offset: MemOffset,
        // opsize, 内存操作数的大小, 比如 "BYTE PTR" 等.
        // None 意味着未指定, 比如 `lea rdi,[r12+0x82]` 这里
        opsize: Option<&'static str>,
    },
}

impl std::string::ToString for MemOffset {
    // 参考 lalrpop onlydispoffset/otheroffset rule
    fn to_string(&self) -> String {
        if self.base.is_none() && self.index_scale.is_none() {
            return format!("{:#x}", self.disp.unwrap_or(0));
        }
        let mut ret = String::new();
        ret.push('[');
        if let Some(base) = self.base {
            ret.push_str(base);
        }
        if let Some((index, scale)) = self.index_scale {
            if self.base.is_some() {
                ret.push('+');
            }
            ret.push_str(&format!("{}*{}", index, scale));
        }
        if let Some(disp) = self.disp {
            if disp > 0 {
                ret.push_str(&format!("+{:#x}", disp));
            } else if disp < 0 {
                ret.push_str(&format!("-{:#x}", -disp));
            }
        }
        ret.push(']');
        return ret;
    }
}

impl std::string::ToString for Operand {
    fn to_string(&self) -> String {
        match self {
            &Operand::Reg(r) => r.to_string(),
            &Operand::Imme(v) => format!("{:#x}", v),
            Operand::Mem {
                seg_sel,
                offset,
                opsize,
            } => {
                let mut ret = String::new();
                if let &Some(opsize) = opsize {
                    ret.push_str(opsize);
                    ret.push(' ');
                }
                if let &Some(segsel) = seg_sel {
                    ret.push_str(segsel);
                    ret.push(':');
                }
                ret.push_str(&offset.to_string());
                ret
            }
        }
    }
}

#[derive(Copy, Clone)]
pub struct Instruction {
    prefix: Option<&'static str>,
    pub mnem: &'static str,
    // The INTEL MANUAL says: There may be from zero to three operands, depending on the opcode.
    pub op0: Option<Operand>,
    pub op1: Option<Operand>,
    op2: Option<Operand>,
}

impl std::string::ToString for Instruction {
    fn to_string(&self) -> String {
        let mut ret = String::new();
        if let Some(prefix) = self.prefix {
            ret.push_str(prefix);
            ret.push(' ');
        }
        ret.push_str(self.mnem);
        ret.push(' ');
        if let Some(op) = self.op0 {
            ret.push_str(&op.to_string());
        }
        if let Some(op) = self.op1 {
            ret.push(',');
            ret.push_str(&op.to_string());
        }
        if let Some(op) = self.op2 {
            ret.push(',');
            ret.push_str(&op.to_string());
        }
        return ret;
    }
}

impl Instruction {
    fn new(mnem: &'static str, ops: &[Operand]) -> Self {
        Instruction {
            prefix: None,
            mnem,
            op0: ops.get(0).map(|v| *v),
            op1: ops.get(1).map(|v| *v),
            op2: ops.get(2).map(|v| *v),
        }
    }

    fn new_prefix(prefix: &'static str, mnem: &'static str, ops: &[Operand]) -> Self {
        let mut inst = Self::new(mnem, ops);
        inst.prefix = Some(prefix);
        inst
    }

    pub fn op0_as_imme(&self) -> Option<u64> {
        if let Some(Operand::Imme(val)) = &self.op0 {
            return Some(*val);
        } else {
            return None;
        }
    }

    pub fn op01_as_single(&self) -> Option<&Operand> {
        if self.op0 != self.op1 {
            return None;
        }
        return self.op0.as_ref();
    }

    pub fn op0_as_str(&self) -> Option<String> {
        return self.op0.as_ref().map(|v| v.to_string());
    }
    pub fn op1_as_str(&self) -> Option<String> {
        return self.op1.as_ref().map(|v| v.to_string());
    }
}

#[derive(Copy, Clone)]
pub struct AddressedInst {
    pub addr: u64,
    pub inst: Instruction,
    // true 意味着指令前面有 '=>' 标识, 比如
    // => ADDR <+829>: mov    rax,QWORD PTR [rdi+0x8]
    pub interested: bool,
}

lalrpop_mod!(pub x86asm); // synthesized by LALRPOP
