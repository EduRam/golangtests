package main

func absdiff(xxxx, yyyy int) int {
	if xxxx < yyyy {
		return yyyy - xxxx
	}
	return xxxx - yyyy
}

func main() {
	absdiff(1111, 2222)
}

/**
[e@localhost 3.6.4]$ go tool 6g -S main.go
"".absdiff t=1 size=48 value=0 args=0x18 locals=0x0
	0x0000 00000 (main.go:3)	TEXT	"".absdiff+0(SB),4,$0-24
	0x0000 00000 (main.go:3)	NOP	,
	0x0000 00000 (main.go:3)	NOP	,
	0x0000 00000 (main.go:3)	MOVQ	"".xxxx+8(FP),CX
	0x0005 00005 (main.go:3)	MOVQ	"".yyyy+16(FP),AX
	0x000a 00010 (main.go:3)	FUNCDATA	$0,gclocals·f90cfd099b5ec2b453c391fece9d42bb+0(SB)
	0x000a 00010 (main.go:3)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x000a 00010 (main.go:4)	CMPQ	CX,AX
	0x000d 00013 (main.go:4)	JGE	,27
	0x000f 00015 (main.go:5)	MOVQ	AX,BX
	0x0012 00018 (main.go:5)	SUBQ	CX,BX
	0x0015 00021 (main.go:5)	MOVQ	BX,"".~r2+24(FP)
	0x001a 00026 (main.go:5)	RET	,
	0x001b 00027 (main.go:7)	MOVQ	CX,BX
	0x001e 00030 (main.go:7)	SUBQ	AX,BX
	0x0021 00033 (main.go:7)	MOVQ	BX,"".~r2+24(FP)
	0x0026 00038 (main.go:7)	RET	,
*/
