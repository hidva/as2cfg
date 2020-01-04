as2cfg.out: parser/as2cfg.l.go parser/as2cfg.y.go virtualall
	go build -o as2cfg.out  github.com/hidva/as2cfg/as2cfg

parser/as2cfg.l.go: parser/as2cfg.l
	golex -o parser/as2cfg.l.go parser/as2cfg.l 
parser/as2cfg.y.go: parser/as2cfg.y
	goyacc -o parser/as2cfg.y.go parser/as2cfg.y

.PHONY: virtualall
virtualall:
	@echo ""	
