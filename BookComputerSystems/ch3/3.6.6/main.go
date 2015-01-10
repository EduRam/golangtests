package main

func cmovdiff(xxxx, yyyy int) int {
	tval := yyyy - xxxx
	rval := xxxx - yyyy
	b := xxxx < yyyy
	if b {
		rval = tval
	}
	return rval
}

func main() {
	cmovdiff(1111, 2222)
}

/**

[e@localhost 3.6.6]$ go tool 6g -S main.go
"".cmovdiff t=1 size=64 value=0 args=0x18 locals=0x0
	0x0000 00000 (main.go:3)	TEXT	"".cmovdiff+0(SB),4,$0-24
	0x0000 00000 (main.go:3)	NOP	,
	0x0000 00000 (main.go:3)	NOP	,
	0x0000 00000 (main.go:3)	MOVQ	"".yyyy+16(FP),SI
	0x0005 00005 (main.go:3)	MOVQ	"".xxxx+8(FP),AX
	0x000a 00010 (main.go:3)	FUNCDATA	$0,gclocals·f90cfd099b5ec2b453c391fece9d42bb+0(SB)
	0x000a 00010 (main.go:3)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x000a 00010 (main.go:4)	MOVQ	SI,DX
	0x000d 00013 (main.go:4)	SUBQ	AX,DX
	0x0010 00016 (main.go:5)	MOVQ	AX,CX
	0x0013 00019 (main.go:5)	SUBQ	SI,CX
	0x0016 00022 (main.go:6)	CMPQ	AX,SI
	0x0019 00025 (main.go:6)	JLT	,42
	0x001b 00027 (main.go:6)	MOVQ	$0,AX
	0x001d 00029 (main.go:7)	CMPB	AL,$0
	0x001f 00031 (main.go:7)	JEQ	,36
	0x0021 00033 (main.go:8)	MOVQ	DX,CX
	0x0024 00036 (main.go:10)	MOVQ	CX,"".~r2+24(FP)
	0x0029 00041 (main.go:10)	RET	,
	0x002a 00042 (main.go:6)	MOVQ	$1,AX
	0x0031 00049 (main.go:7)	JMP	,29

*/
