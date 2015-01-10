// 3.5.3.go
package main

func shiftXXXX(xxxx int, yyyy uint, nnnn uint) (int, uint) {
	xxxx <<= 2
	xxxx >>= nnnn
	yyyy >>= nnnn
	return xxxx, yyyy
}

func main() {
	shiftXXXX(1111, 2222, 3)
}

/**

// 0x0022 00034 (3.5.3.go:12)	SARQ	CX,BP
// This is SARQ (arithmetic right sift) because xxxx is signed int

// 0x0031 00049 (3.5.3.go:12)	SHRQ	CX,BP
// This is SARQ (logical right sift) because yyyy is uint



[e@localhost 3.5.3]$ go tool 6g -S 3.5.3.go

"".shiftXXXX t=1 size=80 value=0 args=0x28 locals=0x0
	0x0000 00000 (3.5.3.go:4)	TEXT	"".shiftXXXX+0(SB),4,$0-40
	0x0000 00000 (3.5.3.go:4)	NOP	,
	0x0000 00000 (3.5.3.go:4)	NOP	,
	0x0000 00000 (3.5.3.go:4)	MOVQ	"".nnnn+24(FP),SI
	0x0005 00005 (3.5.3.go:4)	FUNCDATA	$0,gclocals·754250e8590c282610f2a6c293641cbe+0(SB)
	0x0005 00005 (3.5.3.go:4)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0005 00005 (3.5.3.go:5)	MOVQ	"".xxxx+8(FP),BP
	0x000a 00010 (3.5.3.go:5)	SHLQ	$2,BP
	0x000e 00014 (3.5.3.go:5)	NOP	,
	0x000e 00014 (3.5.3.go:6)	MOVQ	SI,CX
	0x0011 00017 (3.5.3.go:6)	CMPQ	SI,$64
	0x0015 00021 (3.5.3.go:6)	JCC	$1,61
	0x0017 00023 (3.5.3.go:6)	SARQ	CX,BP
	0x001a 00026 (3.5.3.go:6)	MOVQ	BP,DX
	0x001d 00029 (3.5.3.go:6)	NOP	,
	0x001d 00029 (3.5.3.go:7)	MOVQ	"".yyyy+16(FP),BP
	0x0022 00034 (3.5.3.go:7)	MOVQ	SI,CX
	0x0025 00037 (3.5.3.go:7)	CMPQ	SI,$64
	0x0029 00041 (3.5.3.go:7)	JCC	$1,57
	0x002b 00043 (3.5.3.go:7)	SHRQ	CX,BP
	0x002e 00046 (3.5.3.go:7)	NOP	,
	0x002e 00046 (3.5.3.go:8)	MOVQ	DX,"".~r3+32(FP)
	0x0033 00051 (3.5.3.go:8)	MOVQ	BP,"".~r4+40(FP)
	0x0038 00056 (3.5.3.go:8)	RET	,
	0x0039 00057 (3.5.3.go:7)	MOVQ	$0,BP
	0x003b 00059 (3.5.3.go:7)	JMP	,43
	0x003d 00061 (3.5.3.go:6)	SARQ	$63,BP
	0x0041 00065 (3.5.3.go:6)	JMP	,23
	0x0000 48 8b 74 24 18 48 8b 6c 24 08 48 c1 e5 02 48 89  H.t$.H.l$.H...H.
	0x0010 f1 48 83 fe 40 73 26 48 d3 fd 48 89 ea 48 8b 6c  .H..@s&H..H..H.l
	0x0020 24 10 48 89 f1 48 83 fe 40 73 0e 48 d3 ed 48 89  $.H..H..@s.H..H.
	0x0030 54 24 20 48 89 6c 24 28 c3 31 ed eb ee 48 c1 fd  T$ H.l$(.1...H..
	0x0040 3f eb d4                                         ?..
"".main t=1 size=64 value=0 args=0x0 locals=0x0
	0x0000 00000 (3.5.3.go:11)	TEXT	"".main+0(SB),4,$0-0
	0x0000 00000 (3.5.3.go:11)	NOP	,
	0x0000 00000 (3.5.3.go:11)	NOP	,
	0x0000 00000 (3.5.3.go:11)	FUNCDATA	$0,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0000 00000 (3.5.3.go:11)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0000 00000 (3.5.3.go:12)	MOVQ	$1111,BP
	0x0007 00007 (3.5.3.go:12)	MOVQ	$2222,SI
	0x000e 00014 (3.5.3.go:12)	MOVQ	$3,AX
	0x0015 00021 (3.5.3.go:12)	SHLQ	$2,BP
	0x0019 00025 (3.5.3.go:12)	NOP	,
	0x0019 00025 (3.5.3.go:12)	MOVQ	AX,CX
	0x001c 00028 (3.5.3.go:12)	CMPQ	AX,$64
	0x0020 00032 (3.5.3.go:12)	JCC	$1,57
	0x0022 00034 (3.5.3.go:12)	SARQ	CX,BP
	0x0025 00037 (3.5.3.go:12)	NOP	,
	0x0025 00037 (3.5.3.go:12)	MOVQ	SI,BP
	0x0028 00040 (3.5.3.go:12)	MOVQ	AX,CX
	0x002b 00043 (3.5.3.go:12)	CMPQ	AX,$64
	0x002f 00047 (3.5.3.go:12)	JCC	$1,53
	0x0031 00049 (3.5.3.go:12)	SHRQ	CX,BP
	0x0034 00052 (3.5.3.go:12)	NOP	,
	0x0034 00052 (3.5.3.go:13)	RET	,
	0x0035 00053 (3.5.3.go:12)	MOVQ	$0,BP
	0x0037 00055 (3.5.3.go:12)	JMP	,49
	0x0039 00057 (3.5.3.go:12)	SARQ	$63,BP
	0x003d 00061 (3.5.3.go:12)	JMP	,34
	0x0000 48 c7 c5 57 04 00 00 48 c7 c6 ae 08 00 00 48 c7  H..W...H......H.
	0x0010 c0 03 00 00 00 48 c1 e5 02 48 89 c1 48 83 f8 40  .....H...H..H..@
	0x0020 73 17 48 d3 fd 48 89 f5 48 89 c1 48 83 f8 40 73  s.H..H..H..H..@s
	0x0030 04 48 d3 ed c3 31 ed eb f8 48 c1 fd 3f eb e3     .H...1...H..?..
"".init t=1 size=80 value=0 args=0x0 locals=0x0

*/
