// 3.4.3.main.go
package main

func exchange(xxxxp *int, yyyy int) int {
	var xxxx int = *xxxxp
	*xxxxp = yyyy
	return xxxx
}

func main() {
	var iiii int = 1111
	exchange(&iiii, 2222)
}

/*
[e@localhost 3.4.3]$ go tool 6g -S 3.4.3.go
"".exchange t=1 size=32 value=0 args=0x18 locals=0x0
	0x0000 00000 (3.4.3.go:4)	TEXT	"".exchange+0(SB),4,$0-24
	0x0000 00000 (3.4.3.go:4)	NOP	,
	0x0000 00000 (3.4.3.go:4)	NOP	,
	0x0000 00000 (3.4.3.go:4)	MOVQ	"".xxxxp+8(FP),CX
	0x0005 00005 (3.4.3.go:4)	FUNCDATA	$0,gclocals·64b411f0f44be3f38c26e84fc3239091+0(SB)
	0x0005 00005 (3.4.3.go:4)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0005 00005 (3.4.3.go:5)	NOP	,
	0x0005 00005 (3.4.3.go:5)	MOVQ	(CX),AX
	0x0008 00008 (3.4.3.go:6)	NOP	,
	0x0008 00008 (3.4.3.go:6)	MOVQ	"".yyyy+16(FP),BP
	0x000d 00013 (3.4.3.go:6)	MOVQ	BP,(CX)
	0x0010 00016 (3.4.3.go:7)	MOVQ	AX,"".~r2+24(FP)
	0x0015 00021 (3.4.3.go:7)	RET	,
	0x0000 48 8b 4c 24 08 48 8b 01 48 8b 6c 24 10 48 89 29  H.L$.H..H.l$.H.)
	0x0010 48 89 44 24 18 c3                                H.D$..
"".main t=1 size=48 value=0 args=0x0 locals=0x8
	0x0000 00000 (3.4.3.go:10)	TEXT	"".main+0(SB),4,$8-0
	0x0000 00000 (3.4.3.go:10)	SUBQ	$8,SP
	0x0004 00004 (3.4.3.go:10)	FUNCDATA	$0,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0004 00004 (3.4.3.go:10)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0004 00004 (3.4.3.go:11)	MOVQ	$1111,"".iiii+0(SP)
	0x000c 00012 (3.4.3.go:12)	LEAQ	"".iiii+0(SP),BX
	0x0010 00016 (3.4.3.go:12)	MOVQ	$2222,DX
	0x0017 00023 (3.4.3.go:12)	NOP	,
	0x0017 00023 (3.4.3.go:12)	MOVQ	(BX),CX
	0x001a 00026 (3.4.3.go:12)	NOP	,
	0x001a 00026 (3.4.3.go:12)	MOVQ	DX,(BX)
	0x001d 00029 (3.4.3.go:13)	ADDQ	$8,SP
	0x0021 00033 (3.4.3.go:13)	RET	,
*/
