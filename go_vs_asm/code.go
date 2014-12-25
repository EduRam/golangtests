// check example on 162 Chapter 3 Machine-Level Representation of Programs
//
package test

var accum uint64 = 0

func sum(x, y uint64) uint64 {
	var t = x + y
	accum += t
	return t
}

/**

There are 2 outputs:

(1) disable optimizations

[e@localhost go_vs_asm]$ go tool 6g -N -S -S code.go

"".sum t=1 size=80 value=0 args=0x10 locals=0x8
	0x0000 00000 (code.go:7)	TEXT	"".sum+0(SB),4,$8-16
	0x0000 00000 (code.go:7)	SUBQ	$8,SP
	0x0004 00004 (code.go:7)	FUNCDATA	$0,gclocals·9308e7ef08d2cc2f72ae1228688dacf9+0(SB)
	0x0004 00004 (code.go:7)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0004 00004 (code.go:7)	MOVB	$0,"".~r2+24(FP)
	0x0009 00009 (code.go:8)	MOVB	"".x+16(FP),BL
	0x000d 00013 (code.go:8)	MOVB	"".y+17(FP),BPB
	0x0012 00018 (code.go:8)	ADDB	BPB,BL
	0x0015 00021 (code.go:8)	MOVB	BL,"".t+6(SP)
	0x0019 00025 (code.go:9)	MOVB	"".accum+0(SB),BL
	0x001f 00031 (code.go:9)	MOVB	BL,"".autotmp_0000+7(SP)
	0x0023 00035 (code.go:9)	MOVB	"".autotmp_0000+7(SP),BL
	0x0027 00039 (code.go:9)	MOVB	"".t+6(SP),BPB
	0x002c 00044 (code.go:9)	ADDB	BPB,BL
	0x002f 00047 (code.go:9)	MOVB	BL,"".accum+0(SB)
	0x0035 00053 (code.go:10)	MOVB	"".t+6(SP),BL
	0x0039 00057 (code.go:10)	MOVB	BL,"".~r2+24(FP)
	0x003d 00061 (code.go:10)	ADDQ	$8,SP
	0x0041 00065 (code.go:10)	RET	,
	0x0000 48 83 ec 08 c6 44 24 18 00 8a 5c 24 10 40 8a 6c  H....D$...\$.@.l
	0x0010 24 11 40 00 eb 88 5c 24 06 8a 1d 00 00 00 00 88  $.@...\$........
	0x0020 5c 24 07 8a 5c 24 07 40 8a 6c 24 06 40 00 eb 88  \$..\$.@.l$.@...
	0x0030 1d 00 00 00 00 8a 5c 24 06 88 5c 24 18 48 83 c4  ......\$..\$.H..
	0x0040 08 c3                                            ..
	rel 27+4 t=7 "".accum+0
	rel 49+4 t=7 "".accum+0


gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9308e7ef08d2cc2f72ae1228688dacf9 t=7 dupok size=12 value=0
	0x0000 01 00 00 00 04 00 00 00 01 00 00 00              ............
"".accum t=22 size=1 value=0
"".sum·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".sum+0
runtime.throwreturn·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.throwreturn+0
go.string."runtime" t=7 dupok size=24 value=0
	0x0000 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00  ................
	0x0010 72 75 6e 74 69 6d 65 00                          runtime.
	rel 0+8 t=1 go.string."runtime"+16
go.importpath.runtime. t=7 dupok size=16 value=0
	0x0000 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."runtime"+16
runtime.zerovalue t=7 dupok size=0 value=0



(2) optimization enable

[e@localhost go_vs_asm]$ go tool 6g -S -S code.go

"".sum t=1 size=48 value=0 args=0x10 locals=0x0
	0x0000 00000 (code.go:7)	TEXT	"".sum+0(SB),4,$0-16
	0x0000 00000 (code.go:7)	NOP	,
	0x0000 00000 (code.go:7)	NOP	,
	0x0000 00000 (code.go:7)	FUNCDATA	$0,gclocals·9308e7ef08d2cc2f72ae1228688dacf9+0(SB)
	0x0000 00000 (code.go:7)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0000 00000 (code.go:8)	MOVBQZX	"".x+8(FP),BX
	0x0005 00005 (code.go:8)	MOVBQZX	"".y+9(FP),BP
	0x000a 00010 (code.go:8)	ADDQ	BP,BX
	0x000d 00013 (code.go:8)	MOVQ	BX,AX
	0x0010 00016 (code.go:9)	MOVBQZX	"".accum+0(SB),BX
	0x0017 00023 (code.go:9)	ADDQ	AX,BX
	0x001a 00026 (code.go:9)	MOVB	BL,"".accum+0(SB)
	0x0020 00032 (code.go:9)	NOP	,
	0x0020 00032 (code.go:10)	MOVB	AL,"".~r2+16(FP)
	0x0024 00036 (code.go:10)	RET	,
	0x0000 0f b6 5c 24 08 0f b6 6c 24 09 48 01 eb 48 89 d8  ..\$...l$.H..H..
	0x0010 0f b6 1d 00 00 00 00 48 01 c3 88 1d 00 00 00 00  .......H........
	0x0020 88 44 24 10 c3                                   .D$..
	rel 19+4 t=7 "".accum+0
	rel 28+4 t=7 "".accum+0
gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9308e7ef08d2cc2f72ae1228688dacf9 t=7 dupok size=12 value=0
	0x0000 01 00 00 00 04 00 00 00 01 00 00 00              ............
"".accum t=22 size=1 value=0
"".sum·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".sum+0
runtime.throwreturn·f t=7 dupok size=8 value=0
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.throwreturn+0
go.string."runtime" t=7 dupok size=24 value=0
	0x0000 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00  ................
	0x0010 72 75 6e 74 69 6d 65 00                          runtime.
	rel 0+8 t=1 go.string."runtime"+16
go.importpath.runtime. t=7 dupok size=16 value=0
	0x0000 00 00 00 00 00 00 00 00 07 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 go.string."runtime"+16
runtime.zerovalue t=7 dupok size=0 value=0
[e@localhost go_vs_asm]$


*/
