package main

import "os"

func main() {
	var resultZZZ int
	resultZZZ = SumZZZ(2, 3)
	os.Exit(resultZZZ)
}

/**

[e@localhost ch3]$ go tool objdump ch3 > gobjdump.txt

TEXT main.main(SB) /home/e/workspace/gopath/src/github.com/EduRam/golangtests/BookComputerSystems/ch3/3.2.2.main.go
	3.2.2.main.go:5	0x400c00	64488b0c25f0ffffff	FS MOVQ FS:0xfffffff0, CX
	3.2.2.main.go:5	0x400c09	483b6110		CMPQ 0x10(CX), SP
	3.2.2.main.go:5	0x400c0d	7707			JA 0x400c16
	3.2.2.main.go:5	0x400c0f	e82c0a0300		CALL runtime.morestack_noctxt(SB)
	3.2.2.main.go:5	0x400c14	ebea			JMP main.main(SB)
	3.2.2.main.go:5	0x400c16	4883ec08		SUBQ $0x8, SP
	3.2.2.main.go:7	0x400c1a	48c7c302000000	MOVQ $0x2, BX
	3.2.2.main.go:7	0x400c21	48c7c003000000	MOVQ $0x3, AX
	3.2.2.main.go:7	0x400c28	4801c3			ADDQ AX, BX
	3.2.2.main.go:7	0x400c2b	4889d8			MOVQ BX, AX
	3.2.2.main.go:7	0x400c2e	488b1da39b0b00	MOVQ 0xb9ba3(IP), BX
	3.2.2.main.go:7	0x400c35	4801c3			ADDQ AX, BX
	3.2.2.main.go:7	0x400c38	48891d999b0b00	MOVQ BX, 0xb9b99(IP)
	3.2.2.main.go:8	0x400c3f	48890424		MOVQ AX, 0(SP)
	3.2.2.main.go:8	0x400c43	e8886c0300		CALL os.Exit(SB)
	3.2.2.main.go:9	0x400c48	4883c408		ADDQ $0x8, SP
	3.2.2.main.go:9	0x400c4c	c3				RET
	3.2.2.main.go:9	0x400c4d	0000			ADDL AL, 0(AX)
	3.2.2.main.go:9	0x400c4f	00				?


*/
