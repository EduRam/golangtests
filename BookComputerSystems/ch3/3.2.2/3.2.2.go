package main

var accumZZZ int = 0

func SumZZZ(xZZZ, yZZZ int) int {
	var tZZZ = xZZZ + yZZZ
	accumZZZ += tZZZ
	return tZZZ
}

/**
Search for ZZZ

[e@localhost ch3]$ go tool 6g -S -S 3.2.2.go
"".sum t=1 size=48 value=0 args=0x10 locals=0x0
	0x0000 00000 (3.2.2.go:5)	TEXT	"".sum+0(SB),4,$0-16
	0x0000 00000 (3.2.2.go:5)	NOP	,
	0x0000 00000 (3.2.2.go:5)	NOP	,
	0x0000 00000 (3.2.2.go:5)	FUNCDATA	$0,gclocals·9308e7ef08d2cc2f72ae1228688dacf9+0(SB)
	0x0000 00000 (3.2.2.go:5)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0000 00000 (3.2.2.go:6)	MOVBQZX	"".xZZZ+8(FP),BX
	0x0005 00005 (3.2.2.go:6)	MOVBQZX	"".yZZZ+9(FP),BP
	0x000a 00010 (3.2.2.go:6)	ADDQ	BP,BX
	0x000d 00013 (3.2.2.go:6)	MOVQ	BX,AX
	0x0010 00016 (3.2.2.go:7)	MOVBQZX	"".accumZZZ+0(SB),BX
	0x0017 00023 (3.2.2.go:7)	ADDQ	AX,BX
	0x001a 00026 (3.2.2.go:7)	MOVB	BL,"".accumZZZ+0(SB)
	0x0020 00032 (3.2.2.go:7)	NOP	,
	0x0020 00032 (3.2.2.go:8)	MOVB	AL,"".~r2+16(FP)
	0x0024 00036 (3.2.2.go:8)	RET	,
	0x0000 0f b6 5c 24 08 0f b6 6c 24 09 48 01 eb 48 89 d8  ..\$...l$.H..H..
	0x0010 0f b6 1d 00 00 00 00 48 01 c3 88 1d 00 00 00 00  .......H........
	0x0020 88 44 24 10 c3                                   .D$..
	rel 19+4 t=7 "".accumZZZ+0
	rel 28+4 t=7 "".accumZZZ+0
gclocals·3280bececceccd33cb74587feedb1f9f t=7 dupok size=8 value=0
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·9308e7ef08d2cc2f72ae1228688dacf9 t=7 dupok size=12 value=0
	0x0000 01 00 00 00 04 00 00 00 01 00 00 00              ............
"".accumZZZ t=22 size=1 value=0
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


*/
