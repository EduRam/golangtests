// 3.2.3.main.go
package main

func simple(xxxxp *int, yyyy int) int {
	var t int = *xxxxp + yyyy
	*xxxxp = t
	return t
}

func main() {
	var iiii int = 1111
	simple(&iiii, 2222)
}

/**
[e@localhost 3.2.3]$ go tool 6g -S 3.2.3.main.go

"".simple t=1 size=32 value=0 args=0x18 locals=0x0
	0x0000 00000 (3.2.3.main.go:4)	TEXT	"".simple+0(SB),4,$0-24
	0x0000 00000 (3.2.3.main.go:4)	NOP	,
	0x0000 00000 (3.2.3.main.go:4)	NOP	,
	0x0000 00000 (3.2.3.main.go:4)	MOVQ	"".xxxxp+8(FP),CX
	0x0005 00005 (3.2.3.main.go:4)	FUNCDATA	$0,gclocals·64b411f0f44be3f38c26e84fc3239091+0(SB)
	0x0005 00005 (3.2.3.main.go:4)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0005 00005 (3.2.3.main.go:5)	NOP	,
	0x0005 00005 (3.2.3.main.go:5)	MOVQ	(CX),AX
	0x0008 00008 (3.2.3.main.go:5)	MOVQ	"".yyyy+16(FP),BP
	0x000d 00013 (3.2.3.main.go:5)	ADDQ	BP,AX
	0x0010 00016 (3.2.3.main.go:6)	NOP	,
	0x0010 00016 (3.2.3.main.go:6)	MOVQ	AX,(CX)
	0x0013 00019 (3.2.3.main.go:7)	MOVQ	AX,"".~r2+24(FP)
	0x0018 00024 (3.2.3.main.go:7)	RET	,
	0x0000 48 8b 4c 24 08 48 8b 01 48 8b 6c 24 10 48 01 e8  H.L$.H..H.l$.H..
	0x0010 48 89 01 48 89 44 24 18 c3                       H..H.D$..
"".main t=1 size=48 value=0 args=0x0 locals=0x8
	0x0000 00000 (3.2.3.main.go:10)	TEXT	"".main+0(SB),4,$8-0
	0x0000 00000 (3.2.3.main.go:10)	SUBQ	$8,SP
	0x0004 00004 (3.2.3.main.go:10)	FUNCDATA	$0,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0004 00004 (3.2.3.main.go:10)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x0004 00004 (3.2.3.main.go:11)	MOVQ	$1111,"".iiii+0(SP)
	0x000c 00012 (3.2.3.main.go:12)	LEAQ	"".iiii+0(SP),BX
	0x0010 00016 (3.2.3.main.go:12)	MOVQ	BX,CX
	0x0013 00019 (3.2.3.main.go:12)	MOVQ	$2222,AX
	0x001a 00026 (3.2.3.main.go:12)	NOP	,
	0x001a 00026 (3.2.3.main.go:12)	MOVQ	(BX),BX
	0x001d 00029 (3.2.3.main.go:12)	ADDQ	AX,BX
	0x0020 00032 (3.2.3.main.go:12)	NOP	,
	0x0020 00032 (3.2.3.main.go:12)	MOVQ	BX,(CX)
	0x0023 00035 (3.2.3.main.go:13)	ADDQ	$8,SP
	0x0027 00039 (3.2.3.main.go:13)	RET	,
	0x0000 48 83 ec 08 48 c7 04 24 57 04 00 00 48 8d 1c 24  H...H..$W...H..$
	0x0010 48 89 d9 48 c7 c0 ae 08 00 00 48 8b 1b 48 01 c3  H..H......H..H..
	0x0020 48 89 19 48 83 c4 08 c3                          H..H....
"".init t=1 size=80 value=0 args=0x0 locals=0x0
	0x0000 00000 (3.2.3.main.go:13)	TEXT	"".init+0(SB),$0-0
...




[e@localhost 3.2.3]$ objdump -d 3.2.3.main > objdump.txt

3.2.3.main:     file format elf64-x86-64

Disassembly of section .text:

0000000000400c00 <main.main>:
  400c00:	48 83 ec 08          	sub    $0x8,%rsp	// find above "48 83 ec 08" and "SUBQ $8,SP"
  400c04:	48 c7 04 24 01 00 00 	movq   $0x1,(%rsp)
  400c0b:	00
  400c0c:	48 8d 1c 24          	lea    (%rsp),%rbx
  400c10:	48 89 d9             	mov    %rbx,%rcx
  400c13:	48 c7 c0 02 00 00 00 	mov    $0x2,%rax
  400c1a:	48 8b 1b             	mov    (%rbx),%rbx
  400c1d:	48 01 c3             	add    %rax,%rbx
  400c20:	48 89 19             	mov    %rbx,(%rcx)
  400c23:	48 83 c4 08          	add    $0x8,%rsp
  400c27:	c3                   	retq

*/
