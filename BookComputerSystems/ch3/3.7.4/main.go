package main

import "os"

func swapAddr(xp, yp *int) int {
	x := *xp
	y := *yp

	*xp = y
	*yp = x

	return x + y
}

func main() {
	var arg1 int = 534
	var arg2 int = 1057
	sum := swapAddr(&arg1, &arg2)
	diff := arg1 - arg2
	os.Exit(sum * diff)
}

/**
[e@localhost 3.7.4]$ go tool 6g -S main.go
"".swapAddr t=1 size=32 value=0 args=0x18 locals=0x0
	0x0000 00000 (main.go:5)	TEXT	"".swapAddr+0(SB),4,$0-24
	0x0000 00000 (main.go:5)	NOP	,
	0x0000 00000 (main.go:5)	NOP	,
	0x0000 00000 (main.go:5)	MOVQ	"".xp+8(FP),SI
	0x0005 00005 (main.go:5)	MOVQ	"".yp+16(FP),DX
	0x000a 00010 (main.go:5)	FUNCDATA	$0,gclocals·d3486bc7ce1948dc22d7ad1c0be2887a+0(SB)
	0x000a 00010 (main.go:5)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x000a 00010 (main.go:6)	NOP	,
	0x000a 00010 (main.go:6)	MOVQ	(SI),BX
	0x000d 00013 (main.go:7)	NOP	,
	0x000d 00013 (main.go:7)	MOVQ	(DX),AX
	0x0010 00016 (main.go:9)	NOP	,
	0x0010 00016 (main.go:9)	MOVQ	AX,(SI)
	0x0013 00019 (main.go:10)	NOP	,
	0x0013 00019 (main.go:10)	MOVQ	BX,(DX)
	0x0016 00022 (main.go:12)	ADDQ	AX,BX
	0x0019 00025 (main.go:12)	MOVQ	BX,"".~r2+24(FP)
	0x001e 00030 (main.go:12)	RET	,
	0x0000 48 8b 74 24 08 48 8b 54 24 10 48 8b 1e 48 8b 02  H.t$.H.T$.H..H..
	0x0010 48 89 06 48 89 1a 48 01 c3 48 89 5c 24 18 c3     H..H..H..H.\$..
"".main t=1 size=112 value=0 args=0x0 locals=0x18
	0x0000 00000 (main.go:15)	TEXT	"".main+0(SB),$24-0
	0x0000 00000 (main.go:15)	MOVQ	(TLS),CX
	0x0009 00009 (main.go:15)	CMPQ	SP,16(CX)
	0x000d 00013 (main.go:15)	JHI	,22
	0x000f 00015 (main.go:15)	CALL	,runtime.morestack_noctxt(SB)
	0x0014 00020 (main.go:15)	JMP	,0

	0x0016 00022 (main.go:15)	SUBQ	$24,SP  <-------------Allocate 24 bytes on stack, by decrementing

	0x001a 00026 (main.go:15)	FUNCDATA	$0,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x001a 00026 (main.go:15)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x001a 00026 (main.go:16)	MOVQ	$534,"".arg1+16(SP)
	0x0023 00035 (main.go:17)	MOVQ	$1057,"".arg2+8(SP)
	0x002c 00044 (main.go:18)	LEAQ	"".arg1+16(SP),BX
	0x0031 00049 (main.go:18)	MOVQ	BX,SI
	0x0034 00052 (main.go:18)	LEAQ	"".arg2+8(SP),BX
	0x0039 00057 (main.go:18)	MOVQ	BX,DX
	0x003c 00060 (main.go:18)	NOP	,
	0x003c 00060 (main.go:18)	MOVQ	(SI),BX
	0x003f 00063 (main.go:18)	NOP	,
	0x003f 00063 (main.go:18)	MOVQ	(DX),AX
	0x0042 00066 (main.go:18)	NOP	,
	0x0042 00066 (main.go:18)	MOVQ	AX,(SI)
	0x0045 00069 (main.go:18)	NOP	,
	0x0045 00069 (main.go:18)	MOVQ	BX,(DX)
	0x0048 00072 (main.go:18)	ADDQ	AX,BX
	0x004b 00075 (main.go:19)	MOVQ	"".arg1+16(SP),AX
	0x0050 00080 (main.go:19)	MOVQ	"".arg2+8(SP),BP
	0x0055 00085 (main.go:19)	SUBQ	BP,AX
	0x0058 00088 (main.go:20)	IMULQ	AX,BX
	0x005c 00092 (main.go:20)	MOVQ	BX,(SP)
	0x0060 00096 (main.go:20)	PCDATA	$0,$0
	0x0060 00096 (main.go:20)	CALL	,os.Exit(SB)
	0x0065 00101 (main.go:21)	ADDQ	$24,SP
	0x0069 00105 (main.go:21)	RET	,
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 77 07 e8  dH..%....H;a.w..
	0x0010 00 00 00 00 eb ea 48 83 ec 18 48 c7 44 24 10 16  ......H...H.D$..
	0x0020 02 00 00 48 c7 44 24 08 21 04 00 00 48 8d 5c 24  ...H.D$.!...H.\$
	0x0030 10 48 89 de 48 8d 5c 24 08 48 89 da 48 8b 1e 48  .H..H.\$.H..H..H
	0x0040 8b 02 48 89 06 48 89 1a 48 01 c3 48 8b 44 24 10  ..H..H..H..H.D$.
	0x0050 48 8b 6c 24 08 48 29 e8 48 0f af d8 48 89 1c 24  H.l$.H).H...H..$
	0x0060 e8 00 00 00 00 48 83 c4 18 c3                    .....H....
	rel 5+4 t=9 +0
	rel 16+4 t=3 runtime.morestack_noctxt+0
	rel 97+4 t=3 os.Exit+0
"".init t=1 size=80 value=0 args=0x0 locals=0x0
	0x0000 00000 (main.go:21)	TEXT	"".init+0(SB),$0-0
	0x0000 00000 (main.go:21)	MOVQ	(TLS),CX
	0x0009 00009 (main.go:21)	CMPQ	SP,16(CX)
	0x000d 00013 (main.go:21)	JHI	,22
	0x000f 00015 (main.go:21)	CALL	,runtime.morestack_noctxt(SB)
	0x0014 00020 (main.go:21)	JMP	,0
	0x0016 00022 (main.go:21)	NOP	,
	0x0016 00022 (main.go:21)	NOP	,
	0x0016 00022 (main.go:21)	FUNCDATA	$0,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0016 00022 (main.go:21)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0016 00022 (main.go:21)	MOVBQZX	"".initdone·+0(SB),BX
	0x001d 00029 (main.go:21)	CMPB	BL,$0
	0x0020 00032 (main.go:21)	JEQ	,54
	0x0022 00034 (main.go:21)	MOVBQZX	"".initdone·+0(SB),BX
	0x0029 00041 (main.go:21)	CMPB	BL,$2
	0x002c 00044 (main.go:21)	JNE	,47
	0x002e 00046 (main.go:21)	RET	,
	0x002f 00047 (main.go:21)	PCDATA	$0,$0
	0x002f 00047 (main.go:21)	CALL	,runtime.throwinit(SB)
	0x0034 00052 (main.go:21)	UNDEF	,
	0x0036 00054 (main.go:21)	MOVB	$1,"".initdone·+0(SB)
	0x003d 00061 (main.go:21)	PCDATA	$0,$0
	0x003d 00061 (main.go:21)	CALL	,os.init(SB)
	0x0042 00066 (main.go:21)	MOVB	$2,"".initdone·+0(SB)
	0x0049 00073 (main.go:21)	RET	,
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 77 07 e8  dH..%....H;a.w..
	0x0010 00 00 00 00 eb ea 0f b6 1d 00 00 00 00 80 fb 00  ................
	0x0020 74 14 0f b6 1d 00 00 00 00 80 fb 02 75 01 c3 e8  t...........u...
	0x0030 00 00 00 00 0f 0b c6 05 00 00 00 00 01 e8 00 00  ................
	0x0040 00 00 c6 05 00 00 00 00 02 c3                    ..........
	rel 5+4 t=9 +0
	rel 16+4 t=3 runtime.morestack_noctxt+0
	rel 25+4 t=7 "".initdone·+0
	rel 37+4 t=7 "".initdone·+0
	rel 48+4 t=3 runtime.throwinit+0
	rel 56+4 t=7 "".initdone·+-1
	rel 62+4 t=3 os.init+0
	rel 68+4 t=7 "".initdone·+-1
gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·d3486bc7ce1948dc22d7ad1c0be2887a t=7 dupok size=12 value=0
	0x0000 01 00 00 00 06 00 00 00 0a 00 00 00              ............
gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
"".initdone· t=22 size=1 value=0
"".swapAddr·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".swapAddr+0
runtime.throwreturn·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.throwreturn+0
"".main·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main+0
os.Exit·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 os.Exit+0
"".init·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".init+0
runtime.throwinit·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.throwinit+0
os.init·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 os.init+0
go.string."runtime" t=7 dupok size=24 value=0
	0x0000 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00  ................
	0x0010 72 75 6e 74 69 6d 65 00                          runtime.
	rel 0+8 t=1 go.string."runtime"+16
go.importpath.runtime. t=7 dupok size=16 value=0
	0x0000 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."runtime"+16
go.string."os" t=7 dupok size=24 value=0
	0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0010 6f 73 00                                         os.
	rel 0+8 t=1 go.string."os"+16
go.importpath.os. t=7 dupok size=16 value=0
	0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."os"+16
runtime.zerovalue t=7 dupok size=0 value=0
*/
