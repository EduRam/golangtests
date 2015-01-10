package main

func Absdiff(xxxx, yyyy int) int {
	var result int = 0
	if xxxx >= yyyy {
		goto x_ge_y
	} else {
		result = yyyy - xxxx
		goto done
	}

x_ge_y:
	result = xxxx - yyyy
done:
	return result
}

func main() {
	Absdiff(1111, 2222)
}

/*
[e@localhost 3.6.4.b]$ go tool 6g -S main.go

"".Absdiff t=1 size=48 value=0 args=0x18 locals=0x0
	0x0000 00000 (main.go:3)	TEXT	"".Absdiff+0(SB),4,$0-24
	0x0000 00000 (main.go:3)	NOP	,
	0x0000 00000 (main.go:3)	NOP	,
	0x0000 00000 (main.go:3)	MOVQ	"".xxxx+8(FP),CX
	0x0005 00005 (main.go:3)	MOVQ	"".yyyy+16(FP),AX
	0x000a 00010 (main.go:3)	FUNCDATA	$0,gclocals·f90cfd099b5ec2b453c391fece9d42bb+0(SB)
	0x000a 00010 (main.go:3)	FUNCDATA	$1,gclocals·3280bececceccd33cb74587feedb1f9f+0(SB)
	0x000a 00010 (main.go:5)	CMPQ	CX,AX
	0x000d 00013 (main.go:5)	JLT	,30 <------------------ goto to label
	0x000f 00015 (main.go:13)	MOVQ	CX,BX
	0x0012 00018 (main.go:13)	SUBQ	AX,BX
	0x0015 00021 (main.go:13)	MOVQ	BX,AX
	0x0018 00024 (main.go:15)	MOVQ	AX,"".~r2+24(FP)
	0x001d 00029 (main.go:15)	RET	,
	0x001e 00030 (main.go:8)	SUBQ	CX,AX
	0x0021 00033 (main.go:15)	JMP	,24
	0x0000 48 8b 4c 24 08 48 8b 44 24 10 48 39 c1 7c 0f 48  H.L$.H.D$.H9.|.H
*/
