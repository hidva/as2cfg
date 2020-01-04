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

The operand in instruction is represented by its SSA name, it means that if two operands in the same block have the same SSA name, they are the same operand.

And we will attempt to generate more meaningful expression for edge constructed by Jcc(JE, JNE, etc.) instruction, such as that the expression for the edge constructed by the 'je(0x400590)' after 'cmp(dl_1,0x9)' is `dl_1 == 0x9` and `dl_1 != 0x9`, not just `ZF = 1` and `ZF = 0`.
