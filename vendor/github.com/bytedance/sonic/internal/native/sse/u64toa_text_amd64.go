//go:build amd64
// +build amd64

// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_u64toa = []byte{
	// .p2align 4, 0x00
	// LCPI0_0
	0x59, 0x17, 0xb7, 0xd1, 0x00, 0x00, 0x00, 0x00, // .quad 3518437209
	0x59, 0x17, 0xb7, 0xd1, 0x00, 0x00, 0x00, 0x00, //0x00000008 .quad 3518437209
	//0x00000010 LCPI0_1
	0xc5, 0x20, //0x00000010 .word 8389
	0x7b, 0x14, //0x00000012 .word 5243
	0x34, 0x33, //0x00000014 .word 13108
	0x00, 0x80, //0x00000016 .word 32768
	0xc5, 0x20, //0x00000018 .word 8389
	0x7b, 0x14, //0x0000001a .word 5243
	0x34, 0x33, //0x0000001c .word 13108
	0x00, 0x80, //0x0000001e .word 32768
	//0x00000020 LCPI0_2
	0x80, 0x00, //0x00000020 .word 128
	0x00, 0x08, //0x00000022 .word 2048
	0x00, 0x20, //0x00000024 .word 8192
	0x00, 0x80, //0x00000026 .word 32768
	0x80, 0x00, //0x00000028 .word 128
	0x00, 0x08, //0x0000002a .word 2048
	0x00, 0x20, //0x0000002c .word 8192
	0x00, 0x80, //0x0000002e .word 32768
	//0x00000030 LCPI0_3
	0x0a, 0x00, //0x00000030 .word 10
	0x0a, 0x00, //0x00000032 .word 10
	0x0a, 0x00, //0x00000034 .word 10
	0x0a, 0x00, //0x00000036 .word 10
	0x0a, 0x00, //0x00000038 .word 10
	0x0a, 0x00, //0x0000003a .word 10
	0x0a, 0x00, //0x0000003c .word 10
	0x0a, 0x00, //0x0000003e .word 10
	//0x00000040 LCPI0_4
	0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, //0x00000040 QUAD $0x3030303030303030; QUAD $0x3030303030303030  // .space 16, '0000000000000000'
	//0x00000050 .p2align 4, 0x90
	//0x00000050 _u64toa
	0x55,             //0x00000050 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000051 movq         %rsp, %rbp
	0x48, 0x81, 0xfe, 0x0f, 0x27, 0x00, 0x00, //0x00000054 cmpq         $9999, %rsi
	0x0f, 0x87, 0xa5, 0x00, 0x00, 0x00, //0x0000005b ja           LBB0_8
	0x0f, 0xb7, 0xc6, //0x00000061 movzwl       %si, %eax
	0xc1, 0xe8, 0x02, //0x00000064 shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x00000067 imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x0000006d shrl         $17, %eax
	0x48, 0x8d, 0x14, 0x00, //0x00000070 leaq         (%rax,%rax), %rdx
	0x6b, 0xc0, 0x64, //0x00000074 imull        $100, %eax, %eax
	0x89, 0xf1, //0x00000077 movl         %esi, %ecx
	0x29, 0xc1, //0x00000079 subl         %eax, %ecx
	0x0f, 0xb7, 0xc1, //0x0000007b movzwl       %cx, %eax
	0x48, 0x01, 0xc0, //0x0000007e addq         %rax, %rax
	0x81, 0xfe, 0xe8, 0x03, 0x00, 0x00, //0x00000081 cmpl         $1000, %esi
	0x0f, 0x82, 0x16, 0x00, 0x00, 0x00, //0x00000087 jb           LBB0_3
	0x48, 0x8d, 0x0d, 0x8c, 0x04, 0x00, 0x00, //0x0000008d leaq         $1164(%rip), %rcx  /* _Digits+0(%rip) */
	0x8a, 0x0c, 0x0a, //0x00000094 movb         (%rdx,%rcx), %cl
	0x88, 0x0f, //0x00000097 movb         %cl, (%rdi)
	0xb9, 0x01, 0x00, 0x00, 0x00, //0x00000099 movl         $1, %ecx
	0xe9, 0x0b, 0x00, 0x00, 0x00, //0x0000009e jmp          LBB0_4
	//0x000000a3 LBB0_3
	0x31, 0xc9, //0x000000a3 xorl         %ecx, %ecx
	0x83, 0xfe, 0x64, //0x000000a5 cmpl         $100, %esi
	0x0f, 0x82, 0x48, 0x00, 0x00, 0x00, //0x000000a8 jb           LBB0_5
	//0x000000ae LBB0_4
	0x0f, 0xb7, 0xd2, //0x000000ae movzwl       %dx, %edx
	0x48, 0x83, 0xca, 0x01, //0x000000b1 orq          $1, %rdx
	0x48, 0x8d, 0x35, 0x64, 0x04, 0x00, 0x00, //0x000000b5 leaq         $1124(%rip), %rsi  /* _Digits+0(%rip) */
	0x8a, 0x14, 0x32, //0x000000bc movb         (%rdx,%rsi), %dl
	0x89, 0xce, //0x000000bf movl         %ecx, %esi
	0x83, 0xc1, 0x01, //0x000000c1 addl         $1, %ecx
	0x88, 0x14, 0x37, //0x000000c4 movb         %dl, (%rdi,%rsi)
	//0x000000c7 LBB0_6
	0x48, 0x8d, 0x15, 0x52, 0x04, 0x00, 0x00, //0x000000c7 leaq         $1106(%rip), %rdx  /* _Digits+0(%rip) */
	0x8a, 0x14, 0x10, //0x000000ce movb         (%rax,%rdx), %dl
	0x89, 0xce, //0x000000d1 movl         %ecx, %esi
	0x83, 0xc1, 0x01, //0x000000d3 addl         $1, %ecx
	0x88, 0x14, 0x37, //0x000000d6 movb         %dl, (%rdi,%rsi)
	//0x000000d9 LBB0_7
	0x0f, 0xb7, 0xc0, //0x000000d9 movzwl       %ax, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000000dc orq          $1, %rax
	0x48, 0x8d, 0x15, 0x39, 0x04, 0x00, 0x00, //0x000000e0 leaq         $1081(%rip), %rdx  /* _Digits+0(%rip) */
	0x8a, 0x04, 0x10, //0x000000e7 movb         (%rax,%rdx), %al
	0x89, 0xca, //0x000000ea movl         %ecx, %edx
	0x83, 0xc1, 0x01, //0x000000ec addl         $1, %ecx
	0x88, 0x04, 0x17, //0x000000ef movb         %al, (%rdi,%rdx)
	0x89, 0xc8, //0x000000f2 movl         %ecx, %eax
	0x5d, //0x000000f4 popq         %rbp
	0xc3, //0x000000f5 retq
	//0x000000f6 LBB0_5
	0x31, 0xc9, //0x000000f6 xorl         %ecx, %ecx
	0x83, 0xfe, 0x0a, //0x000000f8 cmpl         $10, %esi
	0x0f, 0x83, 0xc6, 0xff, 0xff, 0xff, //0x000000fb jae          LBB0_6
	0xe9, 0xd3, 0xff, 0xff, 0xff, //0x00000101 jmp          LBB0_7
	//0x00000106 LBB0_8
	0x48, 0x81, 0xfe, 0xff, 0xe0, 0xf5, 0x05, //0x00000106 cmpq         $99999999, %rsi
	0x0f, 0x87, 0x20, 0x01, 0x00, 0x00, //0x0000010d ja           LBB0_16
	0x89, 0xf0, //0x00000113 movl         %esi, %eax
	0xba, 0x59, 0x17, 0xb7, 0xd1, //0x00000115 movl         $3518437209, %edx
	0x48, 0x0f, 0xaf, 0xd0, //0x0000011a imulq        %rax, %rdx
	0x48, 0xc1, 0xea, 0x2d, //0x0000011e shrq         $45, %rdx
	0x44, 0x69, 0xc2, 0x10, 0x27, 0x00, 0x00, //0x00000122 imull        $10000, %edx, %r8d
	0x89, 0xf1, //0x00000129 movl         %esi, %ecx
	0x44, 0x29, 0xc1, //0x0000012b subl         %r8d, %ecx
	0x4c, 0x69, 0xd0, 0x83, 0xde, 0x1b, 0x43, //0x0000012e imulq        $1125899907, %rax, %r10
	0x49, 0xc1, 0xea, 0x31, //0x00000135 shrq         $49, %r10
	0x41, 0x83, 0xe2, 0xfe, //0x00000139 andl         $-2, %r10d
	0x0f, 0xb7, 0xc2, //0x0000013d movzwl       %dx, %eax
	0xc1, 0xe8, 0x02, //0x00000140 shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x00000143 imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x00000149 shrl         $17, %eax
	0x6b, 0xc0, 0x64, //0x0000014c imull        $100, %eax, %eax
	0x29, 0xc2, //0x0000014f subl         %eax, %edx
	0x44, 0x0f, 0xb7, 0xca, //0x00000151 movzwl       %dx, %r9d
	0x4d, 0x01, 0xc9, //0x00000155 addq         %r9, %r9
	0x0f, 0xb7, 0xc1, //0x00000158 movzwl       %cx, %eax
	0xc1, 0xe8, 0x02, //0x0000015b shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x0000015e imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x00000164 shrl         $17, %eax
	0x4c, 0x8d, 0x04, 0x00, //0x00000167 leaq         (%rax,%rax), %r8
	0x6b, 0xc0, 0x64, //0x0000016b imull        $100, %eax, %eax
	0x29, 0xc1, //0x0000016e subl         %eax, %ecx
	0x44, 0x0f, 0xb7, 0xd9, //0x00000170 movzwl       %cx, %r11d
	0x4d, 0x01, 0xdb, //0x00000174 addq         %r11, %r11
	0x81, 0xfe, 0x80, 0x96, 0x98, 0x00, //0x00000177 cmpl         $10000000, %esi
	0x0f, 0x82, 0x17, 0x00, 0x00, 0x00, //0x0000017d jb           LBB0_11
	0x48, 0x8d, 0x05, 0x96, 0x03, 0x00, 0x00, //0x00000183 leaq         $918(%rip), %rax  /* _Digits+0(%rip) */
	0x41, 0x8a, 0x04, 0x02, //0x0000018a movb         (%r10,%rax), %al
	0x88, 0x07, //0x0000018e movb         %al, (%rdi)
	0xb9, 0x01, 0x00, 0x00, 0x00, //0x00000190 movl         $1, %ecx
	0xe9, 0x0e, 0x00, 0x00, 0x00, //0x00000195 jmp          LBB0_12
	//0x0000019a LBB0_11
	0x31, 0xc9, //0x0000019a xorl         %ecx, %ecx
	0x81, 0xfe, 0x40, 0x42, 0x0f, 0x00, //0x0000019c cmpl         $1000000, %esi
	0x0f, 0x82, 0x78, 0x00, 0x00, 0x00, //0x000001a2 jb           LBB0_13
	//0x000001a8 LBB0_12
	0x44, 0x89, 0xd0, //0x000001a8 movl         %r10d, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000001ab orq          $1, %rax
	0x48, 0x8d, 0x35, 0x6a, 0x03, 0x00, 0x00, //0x000001af leaq         $874(%rip), %rsi  /* _Digits+0(%rip) */
	0x8a, 0x04, 0x30, //0x000001b6 movb         (%rax,%rsi), %al
	0x89, 0xce, //0x000001b9 movl         %ecx, %esi
	0x83, 0xc1, 0x01, //0x000001bb addl         $1, %ecx
	0x88, 0x04, 0x37, //0x000001be movb         %al, (%rdi,%rsi)
	//0x000001c1 LBB0_14
	0x48, 0x8d, 0x05, 0x58, 0x03, 0x00, 0x00, //0x000001c1 leaq         $856(%rip), %rax  /* _Digits+0(%rip) */
	0x41, 0x8a, 0x04, 0x01, //0x000001c8 movb         (%r9,%rax), %al
	0x89, 0xce, //0x000001cc movl         %ecx, %esi
	0x83, 0xc1, 0x01, //0x000001ce addl         $1, %ecx
	0x88, 0x04, 0x37, //0x000001d1 movb         %al, (%rdi,%rsi)
	//0x000001d4 LBB0_15
	0x41, 0x0f, 0xb7, 0xc1, //0x000001d4 movzwl       %r9w, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000001d8 orq          $1, %rax
	0x48, 0x8d, 0x35, 0x3d, 0x03, 0x00, 0x00, //0x000001dc leaq         $829(%rip), %rsi  /* _Digits+0(%rip) */
	0x8a, 0x04, 0x30, //0x000001e3 movb         (%rax,%rsi), %al
	0x89, 0xca, //0x000001e6 movl         %ecx, %edx
	0x88, 0x04, 0x17, //0x000001e8 movb         %al, (%rdi,%rdx)
	0x41, 0x8a, 0x04, 0x30, //0x000001eb movb         (%r8,%rsi), %al
	0x88, 0x44, 0x17, 0x01, //0x000001ef movb         %al, $1(%rdi,%rdx)
	0x41, 0x0f, 0xb7, 0xc0, //0x000001f3 movzwl       %r8w, %eax
	0x48, 0x83, 0xc8, 0x01, //0x000001f7 orq          $1, %rax
	0x8a, 0x04, 0x30, //0x000001fb movb         (%rax,%rsi), %al
	0x88, 0x44, 0x17, 0x02, //0x000001fe movb         %al, $2(%rdi,%rdx)
	0x41, 0x8a, 0x04, 0x33, //0x00000202 movb         (%r11,%rsi), %al
	0x88, 0x44, 0x17, 0x03, //0x00000206 movb         %al, $3(%rdi,%rdx)
	0x41, 0x0f, 0xb7, 0xc3, //0x0000020a movzwl       %r11w, %eax
	0x48, 0x83, 0xc8, 0x01, //0x0000020e orq          $1, %rax
	0x8a, 0x04, 0x30, //0x00000212 movb         (%rax,%rsi), %al
	0x83, 0xc1, 0x05, //0x00000215 addl         $5, %ecx
	0x88, 0x44, 0x17, 0x04, //0x00000218 movb         %al, $4(%rdi,%rdx)
	0x89, 0xc8, //0x0000021c movl         %ecx, %eax
	0x5d, //0x0000021e popq         %rbp
	0xc3, //0x0000021f retq
	//0x00000220 LBB0_13
	0x31, 0xc9, //0x00000220 xorl         %ecx, %ecx
	0x81, 0xfe, 0xa0, 0x86, 0x01, 0x00, //0x00000222 cmpl         $100000, %esi
	0x0f, 0x83, 0x93, 0xff, 0xff, 0xff, //0x00000228 jae          LBB0_14
	0xe9, 0xa1, 0xff, 0xff, 0xff, //0x0000022e jmp          LBB0_15
	//0x00000233 LBB0_16
	0x48, 0xb8, 0xff, 0xff, 0xc0, 0x6f, 0xf2, 0x86, 0x23, 0x00, //0x00000233 movabsq      $9999999999999999, %rax
	0x48, 0x39, 0xc6, //0x0000023d cmpq         %rax, %rsi
	0x0f, 0x87, 0x12, 0x01, 0x00, 0x00, //0x00000240 ja           LBB0_18
	0x48, 0xb9, 0xfd, 0xce, 0x61, 0x84, 0x11, 0x77, 0xcc, 0xab, //0x00000246 movabsq      $-6067343680855748867, %rcx
	0x48, 0x89, 0xf0, //0x00000250 movq         %rsi, %rax
	0x48, 0xf7, 0xe1, //0x00000253 mulq         %rcx
	0x48, 0xc1, 0xea, 0x1a, //0x00000256 shrq         $26, %rdx
	0x69, 0xc2, 0x00, 0xe1, 0xf5, 0x05, //0x0000025a imull        $100000000, %edx, %eax
	0x29, 0xc6, //0x00000260 subl         %eax, %esi
	0x66, 0x0f, 0x6e, 0xc2, //0x00000262 movd         %edx, %xmm0
	0xf3, 0x0f, 0x6f, 0x0d, 0x92, 0xfd, 0xff, 0xff, //0x00000266 movdqu       $-622(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	0x66, 0x0f, 0x6f, 0xd0, //0x0000026e movdqa       %xmm0, %xmm2
	0x66, 0x0f, 0xf4, 0xd1, //0x00000272 pmuludq      %xmm1, %xmm2
	0x66, 0x0f, 0x73, 0xd2, 0x2d, //0x00000276 psrlq        $45, %xmm2
	0xb8, 0x10, 0x27, 0x00, 0x00, //0x0000027b movl         $10000, %eax
	0x66, 0x48, 0x0f, 0x6e, 0xd8, //0x00000280 movq         %rax, %xmm3
	0x66, 0x0f, 0x6f, 0xe2, //0x00000285 movdqa       %xmm2, %xmm4
	0x66, 0x0f, 0xf4, 0xe3, //0x00000289 pmuludq      %xmm3, %xmm4
	0x66, 0x0f, 0xfa, 0xc4, //0x0000028d psubd        %xmm4, %xmm0
	0x66, 0x0f, 0x61, 0xd0, //0x00000291 punpcklwd    %xmm0, %xmm2
	0x66, 0x0f, 0x73, 0xf2, 0x02, //0x00000295 psllq        $2, %xmm2
	0xf2, 0x0f, 0x70, 0xc2, 0x50, //0x0000029a pshuflw      $80, %xmm2, %xmm0
	0x66, 0x0f, 0x70, 0xc0, 0x50, //0x0000029f pshufd       $80, %xmm0, %xmm0
	0xf3, 0x0f, 0x6f, 0x15, 0x64, 0xfd, 0xff, 0xff, //0x000002a4 movdqu       $-668(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc2, //0x000002ac pmulhuw      %xmm2, %xmm0
	0xf3, 0x0f, 0x6f, 0x25, 0x68, 0xfd, 0xff, 0xff, //0x000002b0 movdqu       $-664(%rip), %xmm4  /* LCPI0_2+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc4, //0x000002b8 pmulhuw      %xmm4, %xmm0
	0xf3, 0x0f, 0x6f, 0x2d, 0x6c, 0xfd, 0xff, 0xff, //0x000002bc movdqu       $-660(%rip), %xmm5  /* LCPI0_3+0(%rip) */
	0x66, 0x0f, 0x6f, 0xf0, //0x000002c4 movdqa       %xmm0, %xmm6
	0x66, 0x0f, 0xd5, 0xf5, //0x000002c8 pmullw       %xmm5, %xmm6
	0x66, 0x0f, 0x73, 0xf6, 0x10, //0x000002cc psllq        $16, %xmm6
	0x66, 0x0f, 0xf9, 0xc6, //0x000002d1 psubw        %xmm6, %xmm0
	0x66, 0x0f, 0x6e, 0xf6, //0x000002d5 movd         %esi, %xmm6
	0x66, 0x0f, 0xf4, 0xce, //0x000002d9 pmuludq      %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xd1, 0x2d, //0x000002dd psrlq        $45, %xmm1
	0x66, 0x0f, 0xf4, 0xd9, //0x000002e2 pmuludq      %xmm1, %xmm3
	0x66, 0x0f, 0xfa, 0xf3, //0x000002e6 psubd        %xmm3, %xmm6
	0x66, 0x0f, 0x61, 0xce, //0x000002ea punpcklwd    %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xf1, 0x02, //0x000002ee psllq        $2, %xmm1
	0xf2, 0x0f, 0x70, 0xc9, 0x50, //0x000002f3 pshuflw      $80, %xmm1, %xmm1
	0x66, 0x0f, 0x70, 0xc9, 0x50, //0x000002f8 pshufd       $80, %xmm1, %xmm1
	0x66, 0x0f, 0xe4, 0xca, //0x000002fd pmulhuw      %xmm2, %xmm1
	0x66, 0x0f, 0xe4, 0xcc, //0x00000301 pmulhuw      %xmm4, %xmm1
	0x66, 0x0f, 0xd5, 0xe9, //0x00000305 pmullw       %xmm1, %xmm5
	0x66, 0x0f, 0x73, 0xf5, 0x10, //0x00000309 psllq        $16, %xmm5
	0x66, 0x0f, 0xf9, 0xcd, //0x0000030e psubw        %xmm5, %xmm1
	0x66, 0x0f, 0x67, 0xc1, //0x00000312 packuswb     %xmm1, %xmm0
	0xf3, 0x0f, 0x6f, 0x0d, 0x22, 0xfd, 0xff, 0xff, //0x00000316 movdqu       $-734(%rip), %xmm1  /* LCPI0_4+0(%rip) */
	0x66, 0x0f, 0xfc, 0xc8, //0x0000031e paddb        %xmm0, %xmm1
	0x66, 0x0f, 0xef, 0xd2, //0x00000322 pxor         %xmm2, %xmm2
	0x66, 0x0f, 0x74, 0xd0, //0x00000326 pcmpeqb      %xmm0, %xmm2
	0x66, 0x0f, 0xd7, 0xc2, //0x0000032a pmovmskb     %xmm2, %eax
	0xf7, 0xd0, //0x0000032e notl         %eax
	0x0d, 0x00, 0x80, 0x00, 0x00, //0x00000330 orl          $32768, %eax
	0x0f, 0xbc, 0xc0, //0x00000335 bsfl         %eax, %eax
	0xb9, 0x10, 0x00, 0x00, 0x00, //0x00000338 movl         $16, %ecx
	0x29, 0xc1, //0x0000033d subl         %eax, %ecx
	0x48, 0xc1, 0xe0, 0x04, //0x0000033f shlq         $4, %rax
	0x48, 0x8d, 0x15, 0xa6, 0x02, 0x00, 0x00, //0x00000343 leaq         $678(%rip), %rdx  /* _VecShiftShuffles+0(%rip) */
	0x66, 0x0f, 0x38, 0x00, 0x0c, 0x10, //0x0000034a pshufb       (%rax,%rdx), %xmm1
	0xf3, 0x0f, 0x7f, 0x0f, //0x00000350 movdqu       %xmm1, (%rdi)
	0x89, 0xc8, //0x00000354 movl         %ecx, %eax
	0x5d, //0x00000356 popq         %rbp
	0xc3, //0x00000357 retq
	//0x00000358 LBB0_18
	0x48, 0xb9, 0x57, 0x78, 0x13, 0xb1, 0x2f, 0x65, 0xa5, 0x39, //0x00000358 movabsq      $4153837486827862103, %rcx
	0x48, 0x89, 0xf0, //0x00000362 movq         %rsi, %rax
	0x48, 0xf7, 0xe1, //0x00000365 mulq         %rcx
	0x48, 0xc1, 0xea, 0x33, //0x00000368 shrq         $51, %rdx
	0x48, 0xb8, 0x00, 0x00, 0xc1, 0x6f, 0xf2, 0x86, 0x23, 0x00, //0x0000036c movabsq      $10000000000000000, %rax
	0x48, 0x0f, 0xaf, 0xc2, //0x00000376 imulq        %rdx, %rax
	0x48, 0x29, 0xc6, //0x0000037a subq         %rax, %rsi
	0x83, 0xfa, 0x09, //0x0000037d cmpl         $9, %edx
	0x0f, 0x87, 0x0f, 0x00, 0x00, 0x00, //0x00000380 ja           LBB0_20
	0x80, 0xc2, 0x30, //0x00000386 addb         $48, %dl
	0x88, 0x17, //0x00000389 movb         %dl, (%rdi)
	0xb9, 0x01, 0x00, 0x00, 0x00, //0x0000038b movl         $1, %ecx
	0xe9, 0xa5, 0x00, 0x00, 0x00, //0x00000390 jmp          LBB0_25
	//0x00000395 LBB0_20
	0x83, 0xfa, 0x63, //0x00000395 cmpl         $99, %edx
	0x0f, 0x87, 0x1a, 0x00, 0x00, 0x00, //0x00000398 ja           LBB0_22
	0x89, 0xd0, //0x0000039e movl         %edx, %eax
	0x48, 0x8d, 0x0d, 0x79, 0x01, 0x00, 0x00, //0x000003a0 leaq         $377(%rip), %rcx  /* _Digits+0(%rip) */
	0x0f, 0xb7, 0x04, 0x41, //0x000003a7 movzwl       (%rcx,%rax,2), %eax
	0x66, 0x89, 0x07, //0x000003ab movw         %ax, (%rdi)
	0xb9, 0x02, 0x00, 0x00, 0x00, //0x000003ae movl         $2, %ecx
	0xe9, 0x82, 0x00, 0x00, 0x00, //0x000003b3 jmp          LBB0_25
	//0x000003b8 LBB0_22
	0x89, 0xd0, //0x000003b8 movl         %edx, %eax
	0xc1, 0xe8, 0x02, //0x000003ba shrl         $2, %eax
	0x69, 0xc0, 0x7b, 0x14, 0x00, 0x00, //0x000003bd imull        $5243, %eax, %eax
	0xc1, 0xe8, 0x11, //0x000003c3 shrl         $17, %eax
	0x81, 0xfa, 0xe7, 0x03, 0x00, 0x00, //0x000003c6 cmpl         $999, %edx
	0x0f, 0x87, 0x37, 0x00, 0x00, 0x00, //0x000003cc ja           LBB0_24
	0x83, 0xc0, 0x30, //0x000003d2 addl         $48, %eax
	0x88, 0x07, //0x000003d5 movb         %al, (%rdi)
	0x0f, 0xb7, 0xc2, //0x000003d7 movzwl       %dx, %eax
	0x89, 0xc1, //0x000003da movl         %eax, %ecx
	0xc1, 0xe9, 0x02, //0x000003dc shrl         $2, %ecx
	0x69, 0xc9, 0x7b, 0x14, 0x00, 0x00, //0x000003df imull        $5243, %ecx, %ecx
	0xc1, 0xe9, 0x11, //0x000003e5 shrl         $17, %ecx
	0x6b, 0xc9, 0x64, //0x000003e8 imull        $100, %ecx, %ecx
	0x29, 0xc8, //0x000003eb subl         %ecx, %eax
	0x0f, 0xb7, 0xc0, //0x000003ed movzwl       %ax, %eax
	0x48, 0x8d, 0x0d, 0x29, 0x01, 0x00, 0x00, //0x000003f0 leaq         $297(%rip), %rcx  /* _Digits+0(%rip) */
	0x0f, 0xb7, 0x04, 0x41, //0x000003f7 movzwl       (%rcx,%rax,2), %eax
	0x66, 0x89, 0x47, 0x01, //0x000003fb movw         %ax, $1(%rdi)
	0xb9, 0x03, 0x00, 0x00, 0x00, //0x000003ff movl         $3, %ecx
	0xe9, 0x31, 0x00, 0x00, 0x00, //0x00000404 jmp          LBB0_25
	//0x00000409 LBB0_24
	0x6b, 0xc8, 0x64, //0x00000409 imull        $100, %eax, %ecx
	0x29, 0xca, //0x0000040c subl         %ecx, %edx
	0x0f, 0xb7, 0xc0, //0x0000040e movzwl       %ax, %eax
	0x48, 0x8d, 0x0d, 0x08, 0x01, 0x00, 0x00, //0x00000411 leaq         $264(%rip), %rcx  /* _Digits+0(%rip) */
	0x0f, 0xb7, 0x04, 0x41, //0x00000418 movzwl       (%rcx,%rax,2), %eax
	0x66, 0x89, 0x07, //0x0000041c movw         %ax, (%rdi)
	0x0f, 0xb7, 0xc2, //0x0000041f movzwl       %dx, %eax
	0x8a, 0x14, 0x41, //0x00000422 movb         (%rcx,%rax,2), %dl
	0x48, 0x01, 0xc0, //0x00000425 addq         %rax, %rax
	0x88, 0x57, 0x02, //0x00000428 movb         %dl, $2(%rdi)
	0x0f, 0xb7, 0xc0, //0x0000042b movzwl       %ax, %eax
	0x8a, 0x44, 0x08, 0x01, //0x0000042e movb         $1(%rax,%rcx), %al
	0x88, 0x47, 0x03, //0x00000432 movb         %al, $3(%rdi)
	0xb9, 0x04, 0x00, 0x00, 0x00, //0x00000435 movl         $4, %ecx
	//0x0000043a LBB0_25
	0x48, 0xba, 0xfd, 0xce, 0x61, 0x84, 0x11, 0x77, 0xcc, 0xab, //0x0000043a movabsq      $-6067343680855748867, %rdx
	0x48, 0x89, 0xf0, //0x00000444 movq         %rsi, %rax
	0x48, 0xf7, 0xe2, //0x00000447 mulq         %rdx
	0x48, 0xc1, 0xea, 0x1a, //0x0000044a shrq         $26, %rdx
	0x66, 0x0f, 0x6e, 0xc2, //0x0000044e movd         %edx, %xmm0
	0xf3, 0x0f, 0x6f, 0x0d, 0xa6, 0xfb, 0xff, 0xff, //0x00000452 movdqu       $-1114(%rip), %xmm1  /* LCPI0_0+0(%rip) */
	0x66, 0x0f, 0x6f, 0xd8, //0x0000045a movdqa       %xmm0, %xmm3
	0x66, 0x0f, 0xf4, 0xd9, //0x0000045e pmuludq      %xmm1, %xmm3
	0x66, 0x0f, 0x73, 0xd3, 0x2d, //0x00000462 psrlq        $45, %xmm3
	0xb8, 0x10, 0x27, 0x00, 0x00, //0x00000467 movl         $10000, %eax
	0x66, 0x48, 0x0f, 0x6e, 0xd0, //0x0000046c movq         %rax, %xmm2
	0x66, 0x0f, 0x6f, 0xe3, //0x00000471 movdqa       %xmm3, %xmm4
	0x66, 0x0f, 0xf4, 0xe2, //0x00000475 pmuludq      %xmm2, %xmm4
	0x66, 0x0f, 0xfa, 0xc4, //0x00000479 psubd        %xmm4, %xmm0
	0x66, 0x0f, 0x61, 0xd8, //0x0000047d punpcklwd    %xmm0, %xmm3
	0x66, 0x0f, 0x73, 0xf3, 0x02, //0x00000481 psllq        $2, %xmm3
	0xf2, 0x0f, 0x70, 0xc3, 0x50, //0x00000486 pshuflw      $80, %xmm3, %xmm0
	0x66, 0x0f, 0x70, 0xc0, 0x50, //0x0000048b pshufd       $80, %xmm0, %xmm0
	0xf3, 0x0f, 0x6f, 0x25, 0x78, 0xfb, 0xff, 0xff, //0x00000490 movdqu       $-1160(%rip), %xmm4  /* LCPI0_1+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc4, //0x00000498 pmulhuw      %xmm4, %xmm0
	0xf3, 0x0f, 0x6f, 0x2d, 0x7c, 0xfb, 0xff, 0xff, //0x0000049c movdqu       $-1156(%rip), %xmm5  /* LCPI0_2+0(%rip) */
	0x66, 0x0f, 0xe4, 0xc5, //0x000004a4 pmulhuw      %xmm5, %xmm0
	0xf3, 0x0f, 0x6f, 0x1d, 0x80, 0xfb, 0xff, 0xff, //0x000004a8 movdqu       $-1152(%rip), %xmm3  /* LCPI0_3+0(%rip) */
	0x66, 0x0f, 0x6f, 0xf0, //0x000004b0 movdqa       %xmm0, %xmm6
	0x66, 0x0f, 0xd5, 0xf3, //0x000004b4 pmullw       %xmm3, %xmm6
	0x66, 0x0f, 0x73, 0xf6, 0x10, //0x000004b8 psllq        $16, %xmm6
	0x66, 0x0f, 0xf9, 0xc6, //0x000004bd psubw        %xmm6, %xmm0
	0x69, 0xc2, 0x00, 0xe1, 0xf5, 0x05, //0x000004c1 imull        $100000000, %edx, %eax
	0x29, 0xc6, //0x000004c7 subl         %eax, %esi
	0x66, 0x0f, 0x6e, 0xf6, //0x000004c9 movd         %esi, %xmm6
	0x66, 0x0f, 0xf4, 0xce, //0x000004cd pmuludq      %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xd1, 0x2d, //0x000004d1 psrlq        $45, %xmm1
	0x66, 0x0f, 0xf4, 0xd1, //0x000004d6 pmuludq      %xmm1, %xmm2
	0x66, 0x0f, 0xfa, 0xf2, //0x000004da psubd        %xmm2, %xmm6
	0x66, 0x0f, 0x61, 0xce, //0x000004de punpcklwd    %xmm6, %xmm1
	0x66, 0x0f, 0x73, 0xf1, 0x02, //0x000004e2 psllq        $2, %xmm1
	0xf2, 0x0f, 0x70, 0xc9, 0x50, //0x000004e7 pshuflw      $80, %xmm1, %xmm1
	0x66, 0x0f, 0x70, 0xc9, 0x50, //0x000004ec pshufd       $80, %xmm1, %xmm1
	0x66, 0x0f, 0xe4, 0xcc, //0x000004f1 pmulhuw      %xmm4, %xmm1
	0x66, 0x0f, 0xe4, 0xcd, //0x000004f5 pmulhuw      %xmm5, %xmm1
	0x66, 0x0f, 0xd5, 0xd9, //0x000004f9 pmullw       %xmm1, %xmm3
	0x66, 0x0f, 0x73, 0xf3, 0x10, //0x000004fd psllq        $16, %xmm3
	0x66, 0x0f, 0xf9, 0xcb, //0x00000502 psubw        %xmm3, %xmm1
	0x66, 0x0f, 0x67, 0xc1, //0x00000506 packuswb     %xmm1, %xmm0
	0x66, 0x0f, 0xfc, 0x05, 0x2e, 0xfb, 0xff, 0xff, //0x0000050a paddb        $-1234(%rip), %xmm0  /* LCPI0_4+0(%rip) */
	0x89, 0xc8, //0x00000512 movl         %ecx, %eax
	0xf3, 0x0f, 0x7f, 0x04, 0x07, //0x00000514 movdqu       %xmm0, (%rdi,%rax)
	0x83, 0xc9, 0x10, //0x00000519 orl          $16, %ecx
	0x89, 0xc8, //0x0000051c movl         %ecx, %eax
	0x5d, //0x0000051e popq         %rbp
	0xc3, //0x0000051f retq
	//0x00000520 .p2align 4, 0x00
	//0x00000520 _Digits
	0x30, 0x30, 0x30, 0x31, 0x30, 0x32, 0x30, 0x33, 0x30, 0x34, 0x30, 0x35, 0x30, 0x36, 0x30, 0x37, //0x00000520 QUAD $0x3330323031303030; QUAD $0x3730363035303430  // .ascii 16, '0001020304050607'
	0x30, 0x38, 0x30, 0x39, 0x31, 0x30, 0x31, 0x31, 0x31, 0x32, 0x31, 0x33, 0x31, 0x34, 0x31, 0x35, //0x00000530 QUAD $0x3131303139303830; QUAD $0x3531343133313231  // .ascii 16, '0809101112131415'
	0x31, 0x36, 0x31, 0x37, 0x31, 0x38, 0x31, 0x39, 0x32, 0x30, 0x32, 0x31, 0x32, 0x32, 0x32, 0x33, //0x00000540 QUAD $0x3931383137313631; QUAD $0x3332323231323032  // .ascii 16, '1617181920212223'
	0x32, 0x34, 0x32, 0x35, 0x32, 0x36, 0x32, 0x37, 0x32, 0x38, 0x32, 0x39, 0x33, 0x30, 0x33, 0x31, //0x00000550 QUAD $0x3732363235323432; QUAD $0x3133303339323832  // .ascii 16, '2425262728293031'
	0x33, 0x32, 0x33, 0x33, 0x33, 0x34, 0x33, 0x35, 0x33, 0x36, 0x33, 0x37, 0x33, 0x38, 0x33, 0x39, //0x00000560 QUAD $0x3533343333333233; QUAD $0x3933383337333633  // .ascii 16, '3233343536373839'
	0x34, 0x30, 0x34, 0x31, 0x34, 0x32, 0x34, 0x33, 0x34, 0x34, 0x34, 0x35, 0x34, 0x36, 0x34, 0x37, //0x00000570 QUAD $0x3334323431343034; QUAD $0x3734363435343434  // .ascii 16, '4041424344454647'
	0x34, 0x38, 0x34, 0x39, 0x35, 0x30, 0x35, 0x31, 0x35, 0x32, 0x35, 0x33, 0x35, 0x34, 0x35, 0x35, //0x00000580 QUAD $0x3135303539343834; QUAD $0x3535343533353235  // .ascii 16, '4849505152535455'
	0x35, 0x36, 0x35, 0x37, 0x35, 0x38, 0x35, 0x39, 0x36, 0x30, 0x36, 0x31, 0x36, 0x32, 0x36, 0x33, //0x00000590 QUAD $0x3935383537353635; QUAD $0x3336323631363036  // .ascii 16, '5657585960616263'
	0x36, 0x34, 0x36, 0x35, 0x36, 0x36, 0x36, 0x37, 0x36, 0x38, 0x36, 0x39, 0x37, 0x30, 0x37, 0x31, //0x000005a0 QUAD $0x3736363635363436; QUAD $0x3137303739363836  // .ascii 16, '6465666768697071'
	0x37, 0x32, 0x37, 0x33, 0x37, 0x34, 0x37, 0x35, 0x37, 0x36, 0x37, 0x37, 0x37, 0x38, 0x37, 0x39, //0x000005b0 QUAD $0x3537343733373237; QUAD $0x3937383737373637  // .ascii 16, '7273747576777879'
	0x38, 0x30, 0x38, 0x31, 0x38, 0x32, 0x38, 0x33, 0x38, 0x34, 0x38, 0x35, 0x38, 0x36, 0x38, 0x37, //0x000005c0 QUAD $0x3338323831383038; QUAD $0x3738363835383438  // .ascii 16, '8081828384858687'
	0x38, 0x38, 0x38, 0x39, 0x39, 0x30, 0x39, 0x31, 0x39, 0x32, 0x39, 0x33, 0x39, 0x34, 0x39, 0x35, //0x000005d0 QUAD $0x3139303939383838; QUAD $0x3539343933393239  // .ascii 16, '8889909192939495'
	0x39, 0x36, 0x39, 0x37, 0x39, 0x38, 0x39, 0x39, //0x000005e0 QUAD $0x3939383937393639  // .ascii 8, '96979899'
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, //0x000005e8 .p2align 4, 0x00
	//0x000005f0 _VecShiftShuffles
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, //0x000005f0 QUAD $0x0706050403020100; QUAD $0x0f0e0d0c0b0a0908  // .ascii 16, '\x00\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f'
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, //0x00000600 QUAD $0x0807060504030201; QUAD $0xff0f0e0d0c0b0a09  // .ascii 16, '\x01\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff'
	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, //0x00000610 QUAD $0x0908070605040302; QUAD $0xffff0f0e0d0c0b0a  // .ascii 16, '\x02\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff'
	0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, //0x00000620 QUAD $0x0a09080706050403; QUAD $0xffffff0f0e0d0c0b  // .ascii 16, '\x03\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff'
	0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, //0x00000630 QUAD $0x0b0a090807060504; QUAD $0xffffffff0f0e0d0c  // .ascii 16, '\x04\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff'
	0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000640 QUAD $0x0c0b0a0908070605; QUAD $0xffffffffff0f0e0d  // .ascii 16, '\x05\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff'
	0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000650 QUAD $0x0d0c0b0a09080706; QUAD $0xffffffffffff0f0e  // .ascii 16, '\x06\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff'
	0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000660 QUAD $0x0e0d0c0b0a090807; QUAD $0xffffffffffffff0f  // .ascii 16, '\x07\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff\xff'
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, //0x00000670 QUAD $0x0f0e0d0c0b0a0908; QUAD $0xffffffffffffffff  // .ascii 16, '\x08\t\n\x0b\x0c\r\x0e\x0f\xff\xff\xff\xff\xff\xff\xff\xff'
}
