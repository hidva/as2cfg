AS2CFG, a utility that converting the assembly code output by GDB disassembler to CFG(Control Flow Graph).

Currently, the usage is very simple, just type:

```
gdb --batch -ex 'disas MyAtoI' atoi.out  | grep -F 0x | as2cfg | dot -Tsvg > atoi.cfg.svg
```

`atoi.out` is compiled from:

```c
int MyAtoI(const char *input) {
    int res = 0;
    int sign = 1;

    while (*input == ' ' || *input == '\t')
        ++input;

    if (*input == '-') {
        sign = -1;
        ++input;
    } else if (*input == '+') {
        ++input;
    }

    for (; *input != '\0'; ++input) {
        if (*input >= '0' && *input <= '9') {
            res = res * 10 + *input - '0';
        } else {
            break;
        }
    }

    return sign * res;
}
```

And `atoi.cfg.svg` looks like:

![atoi.cfg.svg](https://github.com/hidva/as2cfg/blob/master/atoi.cfg.svg)

~~The operand in instruction is represented by its SSA name, it means that if two operands in the same block have the same SSA name, they are the same operand.~~

And we will attempt to generate more meaningful expression for edge constructed by Jcc(JE, JNE, etc.) instruction, such as that the expression for the edge constructed by the 'jbe 0x400500' after 'cmp cl,0x9' is `cl > 0x9` and `cl <= 0x9`, not just `CF = 1 or ZF = 1`.

Instructions marked with an '=>' will be more prominent when displayed:

```
   0x00000000004004f8 <+56>:	ja     0x40051a <MyAtoI+90>
   0x00000000004004fa <+58>:	nop    WORD PTR [rax+rax*1+0x0]
=> 0x0000000000400500 <+64>:	lea    ecx,[rax+rax*4]
   0x0000000000400503 <+67>:	add    rdi,0x1
   0x0000000000400507 <+71>:	lea    eax,[rdx+rcx*2-0x30]
```

![atoi.cfg.mark.jpg](https://github.com/hidva/as2cfg/blob/master/atoi.cfg.mark.jpg)


![h](https://blog.hidva.com/assets/followme.gif?f=GITHUBas2cfg)
