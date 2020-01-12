
sizedict = {}
relateddict = {}

# r[8-15][bwd ]
suffixs = ('b', 'w', 'd', '')
for rd in xrange(8, 16):
    for suffix in suffixs:
        sufsizemap = {'b': 8, 'w': 16, 'd': 32, '': 64}
        reg = 'r%s%s' % (rd, suffix)
        sizedict[reg] = sufsizemap[suffix]
        suffixset = set(suffixs)
        suffixset.remove(suffix)
        relateddict[reg] = ['r%s%s' % (rd, s) for s in suffixset]

# [xyz]mm[0-31]
prefix_size_map = {'x': 128, 'y': 256, 'z': 512}
prefixs = set(['x', 'y', 'z'])
for p in prefixs:
    for d in xrange(0, 33):
        reg = '%smm%s' % (p, d)
        sizedict[reg] = prefix_size_map[p]
        tmpprefix = prefixs.copy()
        tmpprefix.remove(p)
        relateddict[reg] = ['%smm%s' % (tp, d) for tp in tmpprefix]


# 'mm0', 'mm1', 'mm2', 'mm3', 'mm4', 'mm5', 'mm6', 'mm7'
for i in xrange(0, 8):
    reg = 'mm%s' % i
    sizedict[reg] = 64

# register bnd[0-3]
for i in xrange(0, 4):
    reg = 'bnd%d' % i
    sizedict[reg] = 128

# dil, di, edi, rdi,
# sil, si, esi, rsi,
# bpl, bp, ebp, rbp,
# spl, sp, esp, rsp,
for r in ['di', 'si', 'bp', 'sp']:
    rl = '%sl' % r
    er = 'e%s' % r
    rr = 'r%s' % r

    sizedict[rl] = 8
    sizedict[r] = 16
    sizedict[er] = 32
    sizedict[rr] = 64    

    regs = set([r, rl, er, rr])
    for reg in regs:
        tmpregs = set(list(regs))
        tmpregs.remove(reg)
        relateddict[reg] = list(tmpregs)

# cs, ds, es, fs, gs, ss,
for r in ['c', 'd', 'e', 'f', 'g', 's']:
    rs = '%ss' % r
    sizedict[rs] = 16

# ah, al, ax, eax, rax
# bh, bl, bx, ebx, rbx
# ch, cl, cx, ecx, rcx
# dh, dl, dx, edx, rdx
for prefix in ['a', 'b', 'c', 'd']:
    rh = '%sh' % prefix
    rl = '%sl' % prefix
    rx = '%sx' % prefix
    erx = 'e%sx' % prefix
    rrx = 'r%sx' % prefix

    sizedict[rh] = 8
    sizedict[rl] = 8
    sizedict[rx] = 16
    sizedict[erx] = 32
    sizedict[rrx] = 64

    relateddict[rh] = (rx, erx, rrx)
    relateddict[rl] = (rx, erx, rrx)
    relateddict[rx] = (rh, rl, erx, rrx)
    relateddict[erx] = (rh, rl, rx, rrx)
    relateddict[rrx] = (rh, rl, erx, rx)

for r1 in relateddict:
    for r2 in relateddict[r1]:
        if r1 not in relateddict[r2]:
            print "ERROR: %s, %s" % (r1, r2)

for reg in sizedict:
    print 'REG_%s = "%s"' % (reg.upper(), reg)

regattrtemp = """\
REG_%s: {
    Size: %s,
    RelatedRegs: map[string]bool{
        %s
    },
},
"""

for reg in sizedict:
    relatedregs = relateddict.get(reg, [])
    relatedregstr = '\n'.join(["REG_%s: true," % rreg.upper() for rreg in relatedregs])
    print regattrtemp % (reg.upper(), sizedict[reg], relatedregstr)