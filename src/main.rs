use clap::Parser;
use std::io::Read;
pub mod cfg;
pub mod parser;
pub mod render;

#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = "as2cfg")]
struct Args {
    /// input file path, None means stdin
    path: Option<String>,
}

fn main() {
    let args = Args::parse();
    let idata = if let Some(filepath) = args.path {
        std::fs::read_to_string(filepath).expect("read_to_string error")
    } else {
        let mut data = String::new();
        std::io::stdin()
            .read_to_string(&mut data)
            .expect("read from stdin error");
        data
    };
    let mut idata = idata.trim().to_string();
    idata.push('\n');

    let insts = parser::x86asm::InstsParser::new()
        .parse(&idata)
        .expect("parse error");

    let graph = cfg::insts_cfg(insts);
    println!("{}", render::dot_render(&graph));
}
