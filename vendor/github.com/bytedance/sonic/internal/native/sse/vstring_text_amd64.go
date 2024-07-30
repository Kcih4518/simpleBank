//go:build amd64
// +build amd64

// Code generated by asm2asm, DO NOT EDIT.

package sse

var _text_vstring = []byte{
	// .p2align 4, 0x00
	// LCPI0_0
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, // QUAD $0x2222222222222222; QUAD $0x2222222222222222  // .space 16, '""""""""""""""""'
	//0x00000010 LCPI0_1
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, //0x00000010 QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	//0x00000020 LCPI0_2
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, //0x00000020 QUAD $0x2020202020202020; QUAD $0x2020202020202020  // .space 16, '                '
	//0x00000030 .p2align 4, 0x90
	//0x00000030 _vstring
	0x55,             //0x00000030 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000031 movq         %rsp, %rbp
	0x41, 0x57, //0x00000034 pushq        %r15
	0x41, 0x56, //0x00000036 pushq        %r14
	0x41, 0x55, //0x00000038 pushq        %r13
	0x41, 0x54, //0x0000003a pushq        %r12
	0x53,                   //0x0000003c pushq        %rbx
	0x48, 0x83, 0xec, 0x38, //0x0000003d subq         $56, %rsp
	0x48, 0x89, 0x55, 0xc0, //0x00000041 movq         %rdx, $-64(%rbp)
	0x48, 0x89, 0x75, 0xc8, //0x00000045 movq         %rsi, $-56(%rbp)
	0x48, 0x8b, 0x1e, //0x00000049 movq         (%rsi), %rbx
	0xf6, 0xc1, 0x20, //0x0000004c testb        $32, %cl
	0x48, 0x89, 0x5d, 0xb0, //0x0000004f movq         %rbx, $-80(%rbp)
	0x0f, 0x85, 0x97, 0x01, 0x00, 0x00, //0x00000053 jne          LBB0_13
	0x4c, 0x8b, 0x6f, 0x08, //0x00000059 movq         $8(%rdi), %r13
	0x4c, 0x89, 0x6d, 0xb8, //0x0000005d movq         %r13, $-72(%rbp)
	0x49, 0x29, 0xdd, //0x00000061 subq         %rbx, %r13
	0x0f, 0x84, 0x57, 0x05, 0x00, 0x00, //0x00000064 je           LBB0_42
	0x4c, 0x8b, 0x27, //0x0000006a movq         (%rdi), %r12
	0x49, 0x83, 0xfd, 0x40, //0x0000006d cmpq         $64, %r13
	0x0f, 0x82, 0x56, 0x05, 0x00, 0x00, //0x00000071 jb           LBB0_43
	0x49, 0x89, 0xd8, //0x00000077 movq         %rbx, %r8
	0x48, 0xf7, 0xd3, //0x0000007a notq         %rbx
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x0000007d movq         $-1, $-48(%rbp)
	0x45, 0x31, 0xdb, //0x00000085 xorl         %r11d, %r11d
	0xf3, 0x0f, 0x6f, 0x05, 0x70, 0xff, 0xff, 0xff, //0x00000088 movdqu       $-144(%rip), %xmm0  /* LCPI0_0+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x0d, 0x78, 0xff, 0xff, 0xff, //0x00000090 movdqu       $-136(%rip), %xmm1  /* LCPI0_1+0(%rip) */
	0x49, 0xbf, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, //0x00000098 movabsq      $6148914691236517205, %r15
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000000a2 .p2align 4, 0x90
	//0x000000b0 LBB0_4
	0xf3, 0x43, 0x0f, 0x6f, 0x14, 0x04, //0x000000b0 movdqu       (%r12,%r8), %xmm2
	0xf3, 0x43, 0x0f, 0x6f, 0x5c, 0x04, 0x10, //0x000000b6 movdqu       $16(%r12,%r8), %xmm3
	0xf3, 0x43, 0x0f, 0x6f, 0x64, 0x04, 0x20, //0x000000bd movdqu       $32(%r12,%r8), %xmm4
	0xf3, 0x43, 0x0f, 0x6f, 0x6c, 0x04, 0x30, //0x000000c4 movdqu       $48(%r12,%r8), %xmm5
	0x66, 0x0f, 0x6f, 0xf2, //0x000000cb movdqa       %xmm2, %xmm6
	0x66, 0x0f, 0x74, 0xf0, //0x000000cf pcmpeqb      %xmm0, %xmm6
	0x66, 0x44, 0x0f, 0xd7, 0xce, //0x000000d3 pmovmskb     %xmm6, %r9d
	0x66, 0x0f, 0x6f, 0xf3, //0x000000d8 movdqa       %xmm3, %xmm6
	0x66, 0x0f, 0x74, 0xf0, //0x000000dc pcmpeqb      %xmm0, %xmm6
	0x66, 0x0f, 0xd7, 0xf6, //0x000000e0 pmovmskb     %xmm6, %esi
	0x66, 0x0f, 0x6f, 0xf4, //0x000000e4 movdqa       %xmm4, %xmm6
	0x66, 0x0f, 0x74, 0xf0, //0x000000e8 pcmpeqb      %xmm0, %xmm6
	0x66, 0x0f, 0xd7, 0xce, //0x000000ec pmovmskb     %xmm6, %ecx
	0x66, 0x0f, 0x6f, 0xf5, //0x000000f0 movdqa       %xmm5, %xmm6
	0x66, 0x0f, 0x74, 0xf0, //0x000000f4 pcmpeqb      %xmm0, %xmm6
	0x66, 0x0f, 0xd7, 0xfe, //0x000000f8 pmovmskb     %xmm6, %edi
	0x66, 0x0f, 0x74, 0xd1, //0x000000fc pcmpeqb      %xmm1, %xmm2
	0x66, 0x44, 0x0f, 0xd7, 0xd2, //0x00000100 pmovmskb     %xmm2, %r10d
	0x66, 0x0f, 0x74, 0xd9, //0x00000105 pcmpeqb      %xmm1, %xmm3
	0x66, 0x0f, 0xd7, 0xd3, //0x00000109 pmovmskb     %xmm3, %edx
	0x66, 0x0f, 0x74, 0xe1, //0x0000010d pcmpeqb      %xmm1, %xmm4
	0x66, 0x0f, 0xd7, 0xc4, //0x00000111 pmovmskb     %xmm4, %eax
	0x66, 0x0f, 0x74, 0xe9, //0x00000115 pcmpeqb      %xmm1, %xmm5
	0x66, 0x44, 0x0f, 0xd7, 0xf5, //0x00000119 pmovmskb     %xmm5, %r14d
	0x48, 0xc1, 0xe7, 0x30, //0x0000011e shlq         $48, %rdi
	0x48, 0xc1, 0xe1, 0x20, //0x00000122 shlq         $32, %rcx
	0x48, 0x09, 0xf9, //0x00000126 orq          %rdi, %rcx
	0x48, 0xc1, 0xe6, 0x10, //0x00000129 shlq         $16, %rsi
	0x48, 0x09, 0xce, //0x0000012d orq          %rcx, %rsi
	0x49, 0x09, 0xf1, //0x00000130 orq          %rsi, %r9
	0x49, 0xc1, 0xe6, 0x30, //0x00000133 shlq         $48, %r14
	0x48, 0xc1, 0xe0, 0x20, //0x00000137 shlq         $32, %rax
	0x4c, 0x09, 0xf0, //0x0000013b orq          %r14, %rax
	0x48, 0xc1, 0xe2, 0x10, //0x0000013e shlq         $16, %rdx
	0x48, 0x09, 0xc2, //0x00000142 orq          %rax, %rdx
	0x49, 0x09, 0xd2, //0x00000145 orq          %rdx, %r10
	0x0f, 0x85, 0x30, 0x00, 0x00, 0x00, //0x00000148 jne          LBB0_8
	0x4d, 0x85, 0xdb, //0x0000014e testq        %r11, %r11
	0x0f, 0x85, 0x3d, 0x00, 0x00, 0x00, //0x00000151 jne          LBB0_10
	0x45, 0x31, 0xdb, //0x00000157 xorl         %r11d, %r11d
	0x4d, 0x85, 0xc9, //0x0000015a testq        %r9, %r9
	0x0f, 0x85, 0x79, 0x00, 0x00, 0x00, //0x0000015d jne          LBB0_11
	//0x00000163 LBB0_7
	0x49, 0x83, 0xc5, 0xc0, //0x00000163 addq         $-64, %r13
	0x48, 0x83, 0xc3, 0xc0, //0x00000167 addq         $-64, %rbx
	0x49, 0x83, 0xc0, 0x40, //0x0000016b addq         $64, %r8
	0x49, 0x83, 0xfd, 0x3f, //0x0000016f cmpq         $63, %r13
	0x0f, 0x87, 0x37, 0xff, 0xff, 0xff, //0x00000173 ja           LBB0_4
	0xe9, 0xf9, 0x02, 0x00, 0x00, //0x00000179 jmp          LBB0_32
	//0x0000017e LBB0_8
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x0000017e cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x0b, 0x00, 0x00, 0x00, //0x00000183 jne          LBB0_10
	0x49, 0x0f, 0xbc, 0xc2, //0x00000189 bsfq         %r10, %rax
	0x4c, 0x01, 0xc0, //0x0000018d addq         %r8, %rax
	0x48, 0x89, 0x45, 0xd0, //0x00000190 movq         %rax, $-48(%rbp)
	//0x00000194 LBB0_10
	0x4c, 0x89, 0xd8, //0x00000194 movq         %r11, %rax
	0x48, 0xf7, 0xd0, //0x00000197 notq         %rax
	0x4c, 0x21, 0xd0, //0x0000019a andq         %r10, %rax
	0x48, 0x8d, 0x0c, 0x00, //0x0000019d leaq         (%rax,%rax), %rcx
	0x4c, 0x09, 0xd9, //0x000001a1 orq          %r11, %rcx
	0x48, 0x89, 0xca, //0x000001a4 movq         %rcx, %rdx
	0x48, 0xf7, 0xd2, //0x000001a7 notq         %rdx
	0x4c, 0x21, 0xd2, //0x000001aa andq         %r10, %rdx
	0x48, 0xbe, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, //0x000001ad movabsq      $-6148914691236517206, %rsi
	0x48, 0x21, 0xf2, //0x000001b7 andq         %rsi, %rdx
	0x45, 0x31, 0xdb, //0x000001ba xorl         %r11d, %r11d
	0x48, 0x01, 0xc2, //0x000001bd addq         %rax, %rdx
	0x41, 0x0f, 0x92, 0xc3, //0x000001c0 setb         %r11b
	0x48, 0x01, 0xd2, //0x000001c4 addq         %rdx, %rdx
	0x4c, 0x31, 0xfa, //0x000001c7 xorq         %r15, %rdx
	0x48, 0x21, 0xca, //0x000001ca andq         %rcx, %rdx
	0x48, 0xf7, 0xd2, //0x000001cd notq         %rdx
	0x49, 0x21, 0xd1, //0x000001d0 andq         %rdx, %r9
	0x4d, 0x85, 0xc9, //0x000001d3 testq        %r9, %r9
	0x0f, 0x84, 0x87, 0xff, 0xff, 0xff, //0x000001d6 je           LBB0_7
	//0x000001dc LBB0_11
	0x4d, 0x0f, 0xbc, 0xd9, //0x000001dc bsfq         %r9, %r11
	0x49, 0x29, 0xdb, //0x000001e0 subq         %rbx, %r11
	//0x000001e3 LBB0_12
	0x48, 0x8b, 0x75, 0xc0, //0x000001e3 movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x7d, 0xc8, //0x000001e7 movq         $-56(%rbp), %rdi
	0xe9, 0x51, 0x02, 0x00, 0x00, //0x000001eb jmp          LBB0_29
	//0x000001f0 LBB0_13
	0x4c, 0x8b, 0x6f, 0x08, //0x000001f0 movq         $8(%rdi), %r13
	0x4c, 0x89, 0x6d, 0xb8, //0x000001f4 movq         %r13, $-72(%rbp)
	0x49, 0x29, 0xdd, //0x000001f8 subq         %rbx, %r13
	0x0f, 0x84, 0xc0, 0x03, 0x00, 0x00, //0x000001fb je           LBB0_42
	0x48, 0x8b, 0x07, //0x00000201 movq         (%rdi), %rax
	0x48, 0x89, 0x45, 0xa8, //0x00000204 movq         %rax, $-88(%rbp)
	0x49, 0x83, 0xfd, 0x40, //0x00000208 cmpq         $64, %r13
	0x0f, 0x82, 0xd9, 0x03, 0x00, 0x00, //0x0000020c jb           LBB0_44
	0x48, 0x89, 0xd8, //0x00000212 movq         %rbx, %rax
	0x48, 0xf7, 0xd3, //0x00000215 notq         %rbx
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x00000218 movq         $-1, $-48(%rbp)
	0x31, 0xf6, //0x00000220 xorl         %esi, %esi
	0xf3, 0x44, 0x0f, 0x6f, 0x05, 0xd5, 0xfd, 0xff, 0xff, //0x00000222 movdqu       $-555(%rip), %xmm8  /* LCPI0_0+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x0d, 0xdd, 0xfd, 0xff, 0xff, //0x0000022b movdqu       $-547(%rip), %xmm1  /* LCPI0_1+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x15, 0xe5, 0xfd, 0xff, 0xff, //0x00000233 movdqu       $-539(%rip), %xmm2  /* LCPI0_2+0(%rip) */
	0x66, 0x0f, 0x76, 0xdb, //0x0000023b pcmpeqd      %xmm3, %xmm3
	0x90, //0x0000023f .p2align 4, 0x90
	//0x00000240 LBB0_16
	0x4c, 0x89, 0x6d, 0xa0, //0x00000240 movq         %r13, $-96(%rbp)
	0x49, 0x89, 0xf5, //0x00000244 movq         %rsi, %r13
	0x48, 0x8b, 0x4d, 0xa8, //0x00000247 movq         $-88(%rbp), %rcx
	0xf3, 0x0f, 0x6f, 0x04, 0x01, //0x0000024b movdqu       (%rcx,%rax), %xmm0
	0xf3, 0x0f, 0x6f, 0x7c, 0x01, 0x10, //0x00000250 movdqu       $16(%rcx,%rax), %xmm7
	0xf3, 0x0f, 0x6f, 0x74, 0x01, 0x20, //0x00000256 movdqu       $32(%rcx,%rax), %xmm6
	0xf3, 0x0f, 0x6f, 0x6c, 0x01, 0x30, //0x0000025c movdqu       $48(%rcx,%rax), %xmm5
	0x66, 0x0f, 0x6f, 0xe0, //0x00000262 movdqa       %xmm0, %xmm4
	0x66, 0x41, 0x0f, 0x74, 0xe0, //0x00000266 pcmpeqb      %xmm8, %xmm4
	0x66, 0x0f, 0xd7, 0xfc, //0x0000026b pmovmskb     %xmm4, %edi
	0x66, 0x0f, 0x6f, 0xe7, //0x0000026f movdqa       %xmm7, %xmm4
	0x66, 0x41, 0x0f, 0x74, 0xe0, //0x00000273 pcmpeqb      %xmm8, %xmm4
	0x66, 0x44, 0x0f, 0xd7, 0xc4, //0x00000278 pmovmskb     %xmm4, %r8d
	0x66, 0x0f, 0x6f, 0xe6, //0x0000027d movdqa       %xmm6, %xmm4
	0x66, 0x41, 0x0f, 0x74, 0xe0, //0x00000281 pcmpeqb      %xmm8, %xmm4
	0x66, 0x0f, 0xd7, 0xf4, //0x00000286 pmovmskb     %xmm4, %esi
	0x66, 0x0f, 0x6f, 0xe5, //0x0000028a movdqa       %xmm5, %xmm4
	0x66, 0x41, 0x0f, 0x74, 0xe0, //0x0000028e pcmpeqb      %xmm8, %xmm4
	0x66, 0x44, 0x0f, 0xd7, 0xdc, //0x00000293 pmovmskb     %xmm4, %r11d
	0x66, 0x0f, 0x6f, 0xe0, //0x00000298 movdqa       %xmm0, %xmm4
	0x66, 0x0f, 0x74, 0xe1, //0x0000029c pcmpeqb      %xmm1, %xmm4
	0x66, 0x44, 0x0f, 0xd7, 0xf4, //0x000002a0 pmovmskb     %xmm4, %r14d
	0x66, 0x0f, 0x6f, 0xe7, //0x000002a5 movdqa       %xmm7, %xmm4
	0x66, 0x0f, 0x74, 0xe1, //0x000002a9 pcmpeqb      %xmm1, %xmm4
	0x66, 0x44, 0x0f, 0xd7, 0xfc, //0x000002ad pmovmskb     %xmm4, %r15d
	0x66, 0x0f, 0x6f, 0xe6, //0x000002b2 movdqa       %xmm6, %xmm4
	0x66, 0x0f, 0x74, 0xe1, //0x000002b6 pcmpeqb      %xmm1, %xmm4
	0x66, 0x44, 0x0f, 0xd7, 0xd4, //0x000002ba pmovmskb     %xmm4, %r10d
	0x66, 0x0f, 0x6f, 0xe5, //0x000002bf movdqa       %xmm5, %xmm4
	0x66, 0x0f, 0x74, 0xe1, //0x000002c3 pcmpeqb      %xmm1, %xmm4
	0x66, 0x0f, 0xd7, 0xd4, //0x000002c7 pmovmskb     %xmm4, %edx
	0x66, 0x0f, 0x6f, 0xe2, //0x000002cb movdqa       %xmm2, %xmm4
	0x66, 0x0f, 0x64, 0xe7, //0x000002cf pcmpgtb      %xmm7, %xmm4
	0x66, 0x0f, 0x64, 0xfb, //0x000002d3 pcmpgtb      %xmm3, %xmm7
	0x66, 0x0f, 0xdb, 0xfc, //0x000002d7 pand         %xmm4, %xmm7
	0x66, 0x0f, 0xd7, 0xcf, //0x000002db pmovmskb     %xmm7, %ecx
	0x66, 0x0f, 0x6f, 0xe2, //0x000002df movdqa       %xmm2, %xmm4
	0x66, 0x0f, 0x64, 0xe6, //0x000002e3 pcmpgtb      %xmm6, %xmm4
	0x66, 0x0f, 0x64, 0xf3, //0x000002e7 pcmpgtb      %xmm3, %xmm6
	0x66, 0x0f, 0xdb, 0xf4, //0x000002eb pand         %xmm4, %xmm6
	0x66, 0x44, 0x0f, 0xd7, 0xce, //0x000002ef pmovmskb     %xmm6, %r9d
	0x66, 0x0f, 0x6f, 0xe2, //0x000002f4 movdqa       %xmm2, %xmm4
	0x66, 0x0f, 0x64, 0xe5, //0x000002f8 pcmpgtb      %xmm5, %xmm4
	0x66, 0x0f, 0x64, 0xeb, //0x000002fc pcmpgtb      %xmm3, %xmm5
	0x66, 0x0f, 0xdb, 0xec, //0x00000300 pand         %xmm4, %xmm5
	0x66, 0x44, 0x0f, 0xd7, 0xe5, //0x00000304 pmovmskb     %xmm5, %r12d
	0x49, 0xc1, 0xe3, 0x30, //0x00000309 shlq         $48, %r11
	0x48, 0xc1, 0xe6, 0x20, //0x0000030d shlq         $32, %rsi
	0x4c, 0x09, 0xde, //0x00000311 orq          %r11, %rsi
	0x49, 0xc1, 0xe0, 0x10, //0x00000314 shlq         $16, %r8
	0x49, 0x09, 0xf0, //0x00000318 orq          %rsi, %r8
	0x4c, 0x09, 0xc7, //0x0000031b orq          %r8, %rdi
	0x48, 0xc1, 0xe2, 0x30, //0x0000031e shlq         $48, %rdx
	0x49, 0xc1, 0xe2, 0x20, //0x00000322 shlq         $32, %r10
	0x49, 0x09, 0xd2, //0x00000326 orq          %rdx, %r10
	0x49, 0xc1, 0xe7, 0x10, //0x00000329 shlq         $16, %r15
	0x4d, 0x09, 0xd7, //0x0000032d orq          %r10, %r15
	0x49, 0xc1, 0xe4, 0x30, //0x00000330 shlq         $48, %r12
	0x49, 0xc1, 0xe1, 0x20, //0x00000334 shlq         $32, %r9
	0x4d, 0x09, 0xe1, //0x00000338 orq          %r12, %r9
	0x48, 0xc1, 0xe1, 0x10, //0x0000033b shlq         $16, %rcx
	0x4c, 0x09, 0xc9, //0x0000033f orq          %r9, %rcx
	0x4d, 0x09, 0xfe, //0x00000342 orq          %r15, %r14
	0x0f, 0x85, 0x53, 0x00, 0x00, 0x00, //0x00000345 jne          LBB0_22
	0x4d, 0x85, 0xed, //0x0000034b testq        %r13, %r13
	0x0f, 0x85, 0x60, 0x00, 0x00, 0x00, //0x0000034e jne          LBB0_24
	0x31, 0xf6, //0x00000354 xorl         %esi, %esi
	//0x00000356 LBB0_19
	0x66, 0x0f, 0x6f, 0xe2, //0x00000356 movdqa       %xmm2, %xmm4
	0x66, 0x0f, 0x64, 0xe0, //0x0000035a pcmpgtb      %xmm0, %xmm4
	0x66, 0x0f, 0x64, 0xc3, //0x0000035e pcmpgtb      %xmm3, %xmm0
	0x66, 0x0f, 0xdb, 0xc4, //0x00000362 pand         %xmm4, %xmm0
	0x66, 0x0f, 0xd7, 0xd0, //0x00000366 pmovmskb     %xmm0, %edx
	0x48, 0x09, 0xd1, //0x0000036a orq          %rdx, %rcx
	0x48, 0x85, 0xff, //0x0000036d testq        %rdi, %rdi
	0x0f, 0x85, 0x8b, 0x00, 0x00, 0x00, //0x00000370 jne          LBB0_25
	0x48, 0x85, 0xc9, //0x00000376 testq        %rcx, %rcx
	0x0f, 0x85, 0xf6, 0x04, 0x00, 0x00, //0x00000379 jne          LBB0_84
	0x4c, 0x8b, 0x6d, 0xa0, //0x0000037f movq         $-96(%rbp), %r13
	0x49, 0x83, 0xc5, 0xc0, //0x00000383 addq         $-64, %r13
	0x48, 0x83, 0xc3, 0xc0, //0x00000387 addq         $-64, %rbx
	0x48, 0x83, 0xc0, 0x40, //0x0000038b addq         $64, %rax
	0x49, 0x83, 0xfd, 0x3f, //0x0000038f cmpq         $63, %r13
	0x0f, 0x87, 0xa7, 0xfe, 0xff, 0xff, //0x00000393 ja           LBB0_16
	0xe9, 0x66, 0x01, 0x00, 0x00, //0x00000399 jmp          LBB0_37
	//0x0000039e LBB0_22
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x0000039e cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x0b, 0x00, 0x00, 0x00, //0x000003a3 jne          LBB0_24
	0x49, 0x0f, 0xbc, 0xd6, //0x000003a9 bsfq         %r14, %rdx
	0x48, 0x01, 0xc2, //0x000003ad addq         %rax, %rdx
	0x48, 0x89, 0x55, 0xd0, //0x000003b0 movq         %rdx, $-48(%rbp)
	//0x000003b4 LBB0_24
	0x4d, 0x89, 0xe9, //0x000003b4 movq         %r13, %r9
	0x49, 0xf7, 0xd1, //0x000003b7 notq         %r9
	0x4d, 0x21, 0xf1, //0x000003ba andq         %r14, %r9
	0x4f, 0x8d, 0x04, 0x09, //0x000003bd leaq         (%r9,%r9), %r8
	0x4d, 0x09, 0xe8, //0x000003c1 orq          %r13, %r8
	0x4d, 0x89, 0xc2, //0x000003c4 movq         %r8, %r10
	0x49, 0xf7, 0xd2, //0x000003c7 notq         %r10
	0x4d, 0x21, 0xf2, //0x000003ca andq         %r14, %r10
	0x48, 0xbe, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, //0x000003cd movabsq      $-6148914691236517206, %rsi
	0x49, 0x21, 0xf2, //0x000003d7 andq         %rsi, %r10
	0x31, 0xf6, //0x000003da xorl         %esi, %esi
	0x4d, 0x01, 0xca, //0x000003dc addq         %r9, %r10
	0x40, 0x0f, 0x92, 0xc6, //0x000003df setb         %sil
	0x4d, 0x01, 0xd2, //0x000003e3 addq         %r10, %r10
	0x48, 0xba, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, //0x000003e6 movabsq      $6148914691236517205, %rdx
	0x49, 0x31, 0xd2, //0x000003f0 xorq         %rdx, %r10
	0x4d, 0x21, 0xc2, //0x000003f3 andq         %r8, %r10
	0x49, 0xf7, 0xd2, //0x000003f6 notq         %r10
	0x4c, 0x21, 0xd7, //0x000003f9 andq         %r10, %rdi
	0xe9, 0x55, 0xff, 0xff, 0xff, //0x000003fc jmp          LBB0_19
	//0x00000401 LBB0_25
	0x4c, 0x0f, 0xbc, 0xdf, //0x00000401 bsfq         %rdi, %r11
	0x48, 0x85, 0xc9, //0x00000405 testq        %rcx, %rcx
	0x0f, 0x84, 0x1a, 0x00, 0x00, 0x00, //0x00000408 je           LBB0_27
	0x48, 0x0f, 0xbc, 0xc1, //0x0000040e bsfq         %rcx, %rax
	0x48, 0x8b, 0x75, 0xc0, //0x00000412 movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x7d, 0xc8, //0x00000416 movq         $-56(%rbp), %rdi
	0x4c, 0x39, 0xd8, //0x0000041a cmpq         %r11, %rax
	0x0f, 0x83, 0x1b, 0x00, 0x00, 0x00, //0x0000041d jae          LBB0_28
	0xe9, 0xa4, 0x04, 0x00, 0x00, //0x00000423 jmp          LBB0_92
	//0x00000428 LBB0_27
	0xb8, 0x40, 0x00, 0x00, 0x00, //0x00000428 movl         $64, %eax
	0x48, 0x8b, 0x75, 0xc0, //0x0000042d movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x7d, 0xc8, //0x00000431 movq         $-56(%rbp), %rdi
	0x4c, 0x39, 0xd8, //0x00000435 cmpq         %r11, %rax
	0x0f, 0x82, 0x8e, 0x04, 0x00, 0x00, //0x00000438 jb           LBB0_92
	//0x0000043e LBB0_28
	0x49, 0x29, 0xdb, //0x0000043e subq         %rbx, %r11
	//0x00000441 LBB0_29
	0x48, 0x8b, 0x5d, 0xb0, //0x00000441 movq         $-80(%rbp), %rbx
	//0x00000445 LBB0_30
	0x4d, 0x85, 0xdb, //0x00000445 testq        %r11, %r11
	0x0f, 0x88, 0x85, 0x04, 0x00, 0x00, //0x00000448 js           LBB0_93
	0x4c, 0x89, 0x1f, //0x0000044e movq         %r11, (%rdi)
	0x48, 0x89, 0x5e, 0x10, //0x00000451 movq         %rbx, $16(%rsi)
	0x48, 0xc7, 0x06, 0x07, 0x00, 0x00, 0x00, //0x00000455 movq         $7, (%rsi)
	0x48, 0x8b, 0x4d, 0xd0, //0x0000045c movq         $-48(%rbp), %rcx
	0x4c, 0x39, 0xd9, //0x00000460 cmpq         %r11, %rcx
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x00000463 movq         $-1, %rax
	0x48, 0x0f, 0x4c, 0xc1, //0x0000046a cmovlq       %rcx, %rax
	0x48, 0x89, 0x46, 0x18, //0x0000046e movq         %rax, $24(%rsi)
	0xe9, 0x66, 0x04, 0x00, 0x00, //0x00000472 jmp          LBB0_95
	//0x00000477 LBB0_32
	0x4d, 0x01, 0xe0, //0x00000477 addq         %r12, %r8
	0x48, 0x8b, 0x5d, 0xb0, //0x0000047a movq         $-80(%rbp), %rbx
	0x49, 0x83, 0xfd, 0x20, //0x0000047e cmpq         $32, %r13
	0x0f, 0x82, 0xea, 0x01, 0x00, 0x00, //0x00000482 jb           LBB0_50
	//0x00000488 LBB0_33
	0xf3, 0x41, 0x0f, 0x6f, 0x00, //0x00000488 movdqu       (%r8), %xmm0
	0xf3, 0x41, 0x0f, 0x6f, 0x48, 0x10, //0x0000048d movdqu       $16(%r8), %xmm1
	0xf3, 0x0f, 0x6f, 0x15, 0x65, 0xfb, 0xff, 0xff, //0x00000493 movdqu       $-1179(%rip), %xmm2  /* LCPI0_0+0(%rip) */
	0xf3, 0x0f, 0x6f, 0x1d, 0x6d, 0xfb, 0xff, 0xff, //0x0000049b movdqu       $-1171(%rip), %xmm3  /* LCPI0_1+0(%rip) */
	0x66, 0x0f, 0x6f, 0xe0, //0x000004a3 movdqa       %xmm0, %xmm4
	0x66, 0x0f, 0x74, 0xe2, //0x000004a7 pcmpeqb      %xmm2, %xmm4
	0x66, 0x0f, 0xd7, 0xfc, //0x000004ab pmovmskb     %xmm4, %edi
	0x66, 0x0f, 0x74, 0xd1, //0x000004af pcmpeqb      %xmm1, %xmm2
	0x66, 0x0f, 0xd7, 0xca, //0x000004b3 pmovmskb     %xmm2, %ecx
	0x66, 0x0f, 0x74, 0xc3, //0x000004b7 pcmpeqb      %xmm3, %xmm0
	0x66, 0x0f, 0xd7, 0xc0, //0x000004bb pmovmskb     %xmm0, %eax
	0x66, 0x0f, 0x74, 0xcb, //0x000004bf pcmpeqb      %xmm3, %xmm1
	0x66, 0x0f, 0xd7, 0xd1, //0x000004c3 pmovmskb     %xmm1, %edx
	0x48, 0xc1, 0xe1, 0x10, //0x000004c7 shlq         $16, %rcx
	0x48, 0x09, 0xcf, //0x000004cb orq          %rcx, %rdi
	0x48, 0xc1, 0xe2, 0x10, //0x000004ce shlq         $16, %rdx
	0x48, 0x09, 0xd0, //0x000004d2 orq          %rdx, %rax
	0x0f, 0x85, 0x37, 0x01, 0x00, 0x00, //0x000004d5 jne          LBB0_46
	0x4d, 0x85, 0xdb, //0x000004db testq        %r11, %r11
	0x0f, 0x85, 0x4a, 0x01, 0x00, 0x00, //0x000004de jne          LBB0_48
	0x45, 0x31, 0xdb, //0x000004e4 xorl         %r11d, %r11d
	0x48, 0x85, 0xff, //0x000004e7 testq        %rdi, %rdi
	0x0f, 0x84, 0x7a, 0x01, 0x00, 0x00, //0x000004ea je           LBB0_49
	//0x000004f0 LBB0_36
	0x48, 0x0f, 0xbc, 0xc7, //0x000004f0 bsfq         %rdi, %rax
	0x4d, 0x29, 0xe0, //0x000004f4 subq         %r12, %r8
	0x4d, 0x8d, 0x1c, 0x00, //0x000004f7 leaq         (%r8,%rax), %r11
	0x49, 0x83, 0xc3, 0x01, //0x000004fb addq         $1, %r11
	0xe9, 0xbb, 0x03, 0x00, 0x00, //0x000004ff jmp          LBB0_91
	//0x00000504 LBB0_37
	0x48, 0x03, 0x45, 0xa8, //0x00000504 addq         $-88(%rbp), %rax
	0x4d, 0x89, 0xea, //0x00000508 movq         %r13, %r10
	0x49, 0x83, 0xfd, 0x20, //0x0000050b cmpq         $32, %r13
	0x0f, 0x82, 0xf4, 0x00, 0x00, 0x00, //0x0000050f jb           LBB0_45
	//0x00000515 LBB0_38
	0xf3, 0x0f, 0x6f, 0x00, //0x00000515 movdqu       (%rax), %xmm0
	0xf3, 0x0f, 0x6f, 0x48, 0x10, //0x00000519 movdqu       $16(%rax), %xmm1
	0xf3, 0x0f, 0x6f, 0x15, 0xda, 0xfa, 0xff, 0xff, //0x0000051e movdqu       $-1318(%rip), %xmm2  /* LCPI0_0+0(%rip) */
	0x66, 0x0f, 0x6f, 0xd8, //0x00000526 movdqa       %xmm0, %xmm3
	0x66, 0x0f, 0x74, 0xda, //0x0000052a pcmpeqb      %xmm2, %xmm3
	0x66, 0x0f, 0xd7, 0xcb, //0x0000052e pmovmskb     %xmm3, %ecx
	0x66, 0x0f, 0x74, 0xd1, //0x00000532 pcmpeqb      %xmm1, %xmm2
	0x66, 0x0f, 0xd7, 0xd2, //0x00000536 pmovmskb     %xmm2, %edx
	0xf3, 0x0f, 0x6f, 0x15, 0xce, 0xfa, 0xff, 0xff, //0x0000053a movdqu       $-1330(%rip), %xmm2  /* LCPI0_1+0(%rip) */
	0x66, 0x0f, 0x6f, 0xd8, //0x00000542 movdqa       %xmm0, %xmm3
	0x66, 0x0f, 0x74, 0xda, //0x00000546 pcmpeqb      %xmm2, %xmm3
	0x66, 0x44, 0x0f, 0xd7, 0xcb, //0x0000054a pmovmskb     %xmm3, %r9d
	0x66, 0x0f, 0x74, 0xd1, //0x0000054f pcmpeqb      %xmm1, %xmm2
	0x66, 0x0f, 0xd7, 0xda, //0x00000553 pmovmskb     %xmm2, %ebx
	0xf3, 0x0f, 0x6f, 0x15, 0xc1, 0xfa, 0xff, 0xff, //0x00000557 movdqu       $-1343(%rip), %xmm2  /* LCPI0_2+0(%rip) */
	0x66, 0x0f, 0x6f, 0xda, //0x0000055f movdqa       %xmm2, %xmm3
	0x66, 0x0f, 0x64, 0xd8, //0x00000563 pcmpgtb      %xmm0, %xmm3
	0x66, 0x0f, 0x76, 0xe4, //0x00000567 pcmpeqd      %xmm4, %xmm4
	0x66, 0x0f, 0x64, 0xc4, //0x0000056b pcmpgtb      %xmm4, %xmm0
	0x66, 0x0f, 0xdb, 0xc3, //0x0000056f pand         %xmm3, %xmm0
	0x66, 0x44, 0x0f, 0xd7, 0xf0, //0x00000573 pmovmskb     %xmm0, %r14d
	0x66, 0x0f, 0x64, 0xd1, //0x00000578 pcmpgtb      %xmm1, %xmm2
	0x66, 0x0f, 0x64, 0xcc, //0x0000057c pcmpgtb      %xmm4, %xmm1
	0x66, 0x0f, 0xdb, 0xca, //0x00000580 pand         %xmm2, %xmm1
	0x66, 0x0f, 0xd7, 0xf9, //0x00000584 pmovmskb     %xmm1, %edi
	0x48, 0xc1, 0xe2, 0x10, //0x00000588 shlq         $16, %rdx
	0x48, 0x09, 0xd1, //0x0000058c orq          %rdx, %rcx
	0x48, 0xc1, 0xe3, 0x10, //0x0000058f shlq         $16, %rbx
	0x49, 0x09, 0xd9, //0x00000593 orq          %rbx, %r9
	0x0f, 0x85, 0xa0, 0x01, 0x00, 0x00, //0x00000596 jne          LBB0_63
	0x48, 0x85, 0xf6, //0x0000059c testq        %rsi, %rsi
	0x0f, 0x85, 0xbc, 0x01, 0x00, 0x00, //0x0000059f jne          LBB0_65
	0x31, 0xf6, //0x000005a5 xorl         %esi, %esi
	0x48, 0xc1, 0xe7, 0x10, //0x000005a7 shlq         $16, %rdi
	0x48, 0x85, 0xc9, //0x000005ab testq        %rcx, %rcx
	0x48, 0x8b, 0x5d, 0xb0, //0x000005ae movq         $-80(%rbp), %rbx
	0x0f, 0x84, 0xee, 0x01, 0x00, 0x00, //0x000005b2 je           LBB0_67
	//0x000005b8 LBB0_41
	0x48, 0x0f, 0xbc, 0xd1, //0x000005b8 bsfq         %rcx, %rdx
	0xe9, 0xea, 0x01, 0x00, 0x00, //0x000005bc jmp          LBB0_68
	//0x000005c1 LBB0_42
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x000005c1 movq         $-1, %r11
	0xe9, 0xb3, 0x02, 0x00, 0x00, //0x000005c8 jmp          LBB0_86
	//0x000005cd LBB0_43
	0x4d, 0x8d, 0x04, 0x1c, //0x000005cd leaq         (%r12,%rbx), %r8
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x000005d1 movq         $-1, $-48(%rbp)
	0x45, 0x31, 0xdb, //0x000005d9 xorl         %r11d, %r11d
	0x49, 0x83, 0xfd, 0x20, //0x000005dc cmpq         $32, %r13
	0x0f, 0x83, 0xa2, 0xfe, 0xff, 0xff, //0x000005e0 jae          LBB0_33
	0xe9, 0x87, 0x00, 0x00, 0x00, //0x000005e6 jmp          LBB0_50
	//0x000005eb LBB0_44
	0x48, 0x8b, 0x45, 0xa8, //0x000005eb movq         $-88(%rbp), %rax
	0x48, 0x01, 0xd8, //0x000005ef addq         %rbx, %rax
	0x48, 0xc7, 0x45, 0xd0, 0xff, 0xff, 0xff, 0xff, //0x000005f2 movq         $-1, $-48(%rbp)
	0x31, 0xf6, //0x000005fa xorl         %esi, %esi
	0x4d, 0x89, 0xea, //0x000005fc movq         %r13, %r10
	0x49, 0x83, 0xfd, 0x20, //0x000005ff cmpq         $32, %r13
	0x0f, 0x83, 0x0c, 0xff, 0xff, 0xff, //0x00000603 jae          LBB0_38
	//0x00000609 LBB0_45
	0x48, 0x8b, 0x5d, 0xb0, //0x00000609 movq         $-80(%rbp), %rbx
	0xe9, 0xc8, 0x01, 0x00, 0x00, //0x0000060d jmp          LBB0_73
	//0x00000612 LBB0_46
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x00000612 cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x11, 0x00, 0x00, 0x00, //0x00000617 jne          LBB0_48
	0x4c, 0x89, 0xc1, //0x0000061d movq         %r8, %rcx
	0x4c, 0x29, 0xe1, //0x00000620 subq         %r12, %rcx
	0x48, 0x0f, 0xbc, 0xd0, //0x00000623 bsfq         %rax, %rdx
	0x48, 0x01, 0xca, //0x00000627 addq         %rcx, %rdx
	0x48, 0x89, 0x55, 0xd0, //0x0000062a movq         %rdx, $-48(%rbp)
	//0x0000062e LBB0_48
	0x44, 0x89, 0xd9, //0x0000062e movl         %r11d, %ecx
	0xf7, 0xd1, //0x00000631 notl         %ecx
	0x21, 0xc1, //0x00000633 andl         %eax, %ecx
	0x41, 0x8d, 0x14, 0x4b, //0x00000635 leal         (%r11,%rcx,2), %edx
	0x8d, 0x34, 0x09, //0x00000639 leal         (%rcx,%rcx), %esi
	0xf7, 0xd6, //0x0000063c notl         %esi
	0x21, 0xc6, //0x0000063e andl         %eax, %esi
	0x81, 0xe6, 0xaa, 0xaa, 0xaa, 0xaa, //0x00000640 andl         $-1431655766, %esi
	0x45, 0x31, 0xdb, //0x00000646 xorl         %r11d, %r11d
	0x01, 0xce, //0x00000649 addl         %ecx, %esi
	0x41, 0x0f, 0x92, 0xc3, //0x0000064b setb         %r11b
	0x01, 0xf6, //0x0000064f addl         %esi, %esi
	0x81, 0xf6, 0x55, 0x55, 0x55, 0x55, //0x00000651 xorl         $1431655765, %esi
	0x21, 0xd6, //0x00000657 andl         %edx, %esi
	0xf7, 0xd6, //0x00000659 notl         %esi
	0x21, 0xf7, //0x0000065b andl         %esi, %edi
	0x48, 0x8b, 0x5d, 0xb0, //0x0000065d movq         $-80(%rbp), %rbx
	0x48, 0x85, 0xff, //0x00000661 testq        %rdi, %rdi
	0x0f, 0x85, 0x86, 0xfe, 0xff, 0xff, //0x00000664 jne          LBB0_36
	//0x0000066a LBB0_49
	0x49, 0x83, 0xc0, 0x20, //0x0000066a addq         $32, %r8
	0x49, 0x83, 0xc5, 0xe0, //0x0000066e addq         $-32, %r13
	//0x00000672 LBB0_50
	0x4d, 0x85, 0xdb, //0x00000672 testq        %r11, %r11
	0x0f, 0x85, 0x71, 0x02, 0x00, 0x00, //0x00000675 jne          LBB0_96
	0x4c, 0x89, 0xe7, //0x0000067b movq         %r12, %rdi
	0x48, 0xf7, 0xd7, //0x0000067e notq         %rdi
	0x48, 0x8b, 0x5d, 0xd0, //0x00000681 movq         $-48(%rbp), %rbx
	0x4d, 0x85, 0xed, //0x00000685 testq        %r13, %r13
	0x0f, 0x84, 0x8b, 0x00, 0x00, 0x00, //0x00000688 je           LBB0_60
	//0x0000068e LBB0_52
	0x48, 0x83, 0xc7, 0x01, //0x0000068e addq         $1, %rdi
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x00000692 movq         $-1, %r11
	//0x00000699 LBB0_53
	0x31, 0xc0, //0x00000699 xorl         %eax, %eax
	//0x0000069b LBB0_54
	0x41, 0x0f, 0xb6, 0x14, 0x00, //0x0000069b movzbl       (%r8,%rax), %edx
	0x80, 0xfa, 0x22, //0x000006a0 cmpb         $34, %dl
	0x0f, 0x84, 0x69, 0x00, 0x00, 0x00, //0x000006a3 je           LBB0_59
	0x80, 0xfa, 0x5c, //0x000006a9 cmpb         $92, %dl
	0x0f, 0x84, 0x12, 0x00, 0x00, 0x00, //0x000006ac je           LBB0_57
	0x48, 0x83, 0xc0, 0x01, //0x000006b2 addq         $1, %rax
	0x49, 0x39, 0xc5, //0x000006b6 cmpq         %rax, %r13
	0x0f, 0x85, 0xdc, 0xff, 0xff, 0xff, //0x000006b9 jne          LBB0_54
	0xe9, 0x60, 0x00, 0x00, 0x00, //0x000006bf jmp          LBB0_61
	//0x000006c4 LBB0_57
	0x49, 0x8d, 0x4d, 0xff, //0x000006c4 leaq         $-1(%r13), %rcx
	0x48, 0x39, 0xc1, //0x000006c8 cmpq         %rax, %rcx
	0x0f, 0x84, 0xab, 0x01, 0x00, 0x00, //0x000006cb je           LBB0_85
	0x4a, 0x8d, 0x0c, 0x07, //0x000006d1 leaq         (%rdi,%r8), %rcx
	0x48, 0x01, 0xc1, //0x000006d5 addq         %rax, %rcx
	0x48, 0x83, 0xfb, 0xff, //0x000006d8 cmpq         $-1, %rbx
	0x48, 0x8b, 0x55, 0xd0, //0x000006dc movq         $-48(%rbp), %rdx
	0x48, 0x0f, 0x44, 0xd1, //0x000006e0 cmoveq       %rcx, %rdx
	0x48, 0x89, 0x55, 0xd0, //0x000006e4 movq         %rdx, $-48(%rbp)
	0x48, 0x0f, 0x44, 0xd9, //0x000006e8 cmoveq       %rcx, %rbx
	0x49, 0x01, 0xc0, //0x000006ec addq         %rax, %r8
	0x49, 0x83, 0xc0, 0x02, //0x000006ef addq         $2, %r8
	0x4c, 0x89, 0xe9, //0x000006f3 movq         %r13, %rcx
	0x48, 0x29, 0xc1, //0x000006f6 subq         %rax, %rcx
	0x48, 0x83, 0xc1, 0xfe, //0x000006f9 addq         $-2, %rcx
	0x49, 0x83, 0xc5, 0xfe, //0x000006fd addq         $-2, %r13
	0x49, 0x39, 0xc5, //0x00000701 cmpq         %rax, %r13
	0x49, 0x89, 0xcd, //0x00000704 movq         %rcx, %r13
	0x0f, 0x85, 0x8c, 0xff, 0xff, 0xff, //0x00000707 jne          LBB0_53
	0xe9, 0x6a, 0x01, 0x00, 0x00, //0x0000070d jmp          LBB0_85
	//0x00000712 LBB0_59
	0x49, 0x01, 0xc0, //0x00000712 addq         %rax, %r8
	0x49, 0x83, 0xc0, 0x01, //0x00000715 addq         $1, %r8
	//0x00000719 LBB0_60
	0x4d, 0x29, 0xe0, //0x00000719 subq         %r12, %r8
	0x4d, 0x89, 0xc3, //0x0000071c movq         %r8, %r11
	0xe9, 0xbf, 0xfa, 0xff, 0xff, //0x0000071f jmp          LBB0_12
	//0x00000724 LBB0_61
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x00000724 movq         $-1, %r11
	0x80, 0xfa, 0x22, //0x0000072b cmpb         $34, %dl
	0x0f, 0x85, 0x48, 0x01, 0x00, 0x00, //0x0000072e jne          LBB0_85
	0x4d, 0x01, 0xe8, //0x00000734 addq         %r13, %r8
	0xe9, 0xdd, 0xff, 0xff, 0xff, //0x00000737 jmp          LBB0_60
	//0x0000073c LBB0_63
	0x48, 0x89, 0xf2, //0x0000073c movq         %rsi, %rdx
	0x48, 0x83, 0x7d, 0xd0, 0xff, //0x0000073f cmpq         $-1, $-48(%rbp)
	0x0f, 0x85, 0x1a, 0x00, 0x00, 0x00, //0x00000744 jne          LBB0_66
	0x49, 0x89, 0xc0, //0x0000074a movq         %rax, %r8
	0x4c, 0x2b, 0x45, 0xa8, //0x0000074d subq         $-88(%rbp), %r8
	0x49, 0x0f, 0xbc, 0xf1, //0x00000751 bsfq         %r9, %rsi
	0x4c, 0x01, 0xc6, //0x00000755 addq         %r8, %rsi
	0x48, 0x89, 0x75, 0xd0, //0x00000758 movq         %rsi, $-48(%rbp)
	0xe9, 0x03, 0x00, 0x00, 0x00, //0x0000075c jmp          LBB0_66
	//0x00000761 LBB0_65
	0x48, 0x89, 0xf2, //0x00000761 movq         %rsi, %rdx
	//0x00000764 LBB0_66
	0x48, 0x89, 0xd6, //0x00000764 movq         %rdx, %rsi
	0xf7, 0xd2, //0x00000767 notl         %edx
	0x44, 0x21, 0xca, //0x00000769 andl         %r9d, %edx
	0x44, 0x8d, 0x04, 0x56, //0x0000076c leal         (%rsi,%rdx,2), %r8d
	0x8d, 0x1c, 0x12, //0x00000770 leal         (%rdx,%rdx), %ebx
	0xf7, 0xd3, //0x00000773 notl         %ebx
	0x44, 0x21, 0xcb, //0x00000775 andl         %r9d, %ebx
	0x81, 0xe3, 0xaa, 0xaa, 0xaa, 0xaa, //0x00000778 andl         $-1431655766, %ebx
	0x31, 0xf6, //0x0000077e xorl         %esi, %esi
	0x01, 0xd3, //0x00000780 addl         %edx, %ebx
	0x40, 0x0f, 0x92, 0xc6, //0x00000782 setb         %sil
	0x01, 0xdb, //0x00000786 addl         %ebx, %ebx
	0x81, 0xf3, 0x55, 0x55, 0x55, 0x55, //0x00000788 xorl         $1431655765, %ebx
	0x44, 0x21, 0xc3, //0x0000078e andl         %r8d, %ebx
	0xf7, 0xd3, //0x00000791 notl         %ebx
	0x21, 0xd9, //0x00000793 andl         %ebx, %ecx
	0x48, 0xc1, 0xe7, 0x10, //0x00000795 shlq         $16, %rdi
	0x48, 0x85, 0xc9, //0x00000799 testq        %rcx, %rcx
	0x48, 0x8b, 0x5d, 0xb0, //0x0000079c movq         $-80(%rbp), %rbx
	0x0f, 0x85, 0x12, 0xfe, 0xff, 0xff, //0x000007a0 jne          LBB0_41
	//0x000007a6 LBB0_67
	0xba, 0x40, 0x00, 0x00, 0x00, //0x000007a6 movl         $64, %edx
	//0x000007ab LBB0_68
	0x4c, 0x09, 0xf7, //0x000007ab orq          %r14, %rdi
	0x48, 0x85, 0xc9, //0x000007ae testq        %rcx, %rcx
	0x0f, 0x84, 0x12, 0x00, 0x00, 0x00, //0x000007b1 je           LBB0_71
	0x48, 0x85, 0xff, //0x000007b7 testq        %rdi, %rdi
	0x0f, 0x84, 0xcd, 0x00, 0x00, 0x00, //0x000007ba je           LBB0_87
	0x48, 0x0f, 0xbc, 0xcf, //0x000007c0 bsfq         %rdi, %rcx
	0xe9, 0xc9, 0x00, 0x00, 0x00, //0x000007c4 jmp          LBB0_88
	//0x000007c9 LBB0_71
	0x48, 0x85, 0xff, //0x000007c9 testq        %rdi, %rdi
	0x0f, 0x85, 0xa3, 0x00, 0x00, 0x00, //0x000007cc jne          LBB0_84
	0x48, 0x83, 0xc0, 0x20, //0x000007d2 addq         $32, %rax
	0x49, 0x83, 0xc2, 0xe0, //0x000007d6 addq         $-32, %r10
	//0x000007da LBB0_73
	0x48, 0x85, 0xf6, //0x000007da testq        %rsi, %rsi
	0x0f, 0x85, 0x49, 0x01, 0x00, 0x00, //0x000007dd jne          LBB0_98
	0x48, 0x8b, 0x4d, 0xd0, //0x000007e3 movq         $-48(%rbp), %rcx
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x000007e7 movq         $-1, %r11
	0x4d, 0x85, 0xd2, //0x000007ee testq        %r10, %r10
	0x0f, 0x84, 0x85, 0x00, 0x00, 0x00, //0x000007f1 je           LBB0_85
	//0x000007f7 LBB0_75
	0x4c, 0x89, 0xd7, //0x000007f7 movq         %r10, %rdi
	//0x000007fa LBB0_76
	0x0f, 0xb6, 0x10, //0x000007fa movzbl       (%rax), %edx
	0x80, 0xfa, 0x22, //0x000007fd cmpb         $34, %dl
	0x0f, 0x84, 0xae, 0x00, 0x00, 0x00, //0x00000800 je           LBB0_90
	0x80, 0xfa, 0x5c, //0x00000806 cmpb         $92, %dl
	0x0f, 0x84, 0x26, 0x00, 0x00, 0x00, //0x00000809 je           LBB0_80
	0x80, 0xfa, 0x20, //0x0000080f cmpb         $32, %dl
	0x0f, 0x82, 0x5d, 0x00, 0x00, 0x00, //0x00000812 jb           LBB0_84
	0x48, 0xc7, 0xc2, 0xff, 0xff, 0xff, 0xff, //0x00000818 movq         $-1, %rdx
	0xbe, 0x01, 0x00, 0x00, 0x00, //0x0000081f movl         $1, %esi
	0x48, 0x01, 0xf0, //0x00000824 addq         %rsi, %rax
	0x48, 0x01, 0xd7, //0x00000827 addq         %rdx, %rdi
	0x0f, 0x85, 0xca, 0xff, 0xff, 0xff, //0x0000082a jne          LBB0_76
	0xe9, 0x47, 0x00, 0x00, 0x00, //0x00000830 jmp          LBB0_85
	//0x00000835 LBB0_80
	0x48, 0x83, 0xff, 0x01, //0x00000835 cmpq         $1, %rdi
	0x0f, 0x84, 0x3d, 0x00, 0x00, 0x00, //0x00000839 je           LBB0_85
	0x48, 0xc7, 0xc2, 0xfe, 0xff, 0xff, 0xff, //0x0000083f movq         $-2, %rdx
	0xbe, 0x02, 0x00, 0x00, 0x00, //0x00000846 movl         $2, %esi
	0x48, 0x83, 0xf9, 0xff, //0x0000084b cmpq         $-1, %rcx
	0x0f, 0x85, 0x0b, 0x00, 0x00, 0x00, //0x0000084f jne          LBB0_83
	0x48, 0x89, 0xc1, //0x00000855 movq         %rax, %rcx
	0x48, 0x2b, 0x4d, 0xa8, //0x00000858 subq         $-88(%rbp), %rcx
	0x48, 0x89, 0x4d, 0xd0, //0x0000085c movq         %rcx, $-48(%rbp)
	//0x00000860 LBB0_83
	0x48, 0x8b, 0x5d, 0xb0, //0x00000860 movq         $-80(%rbp), %rbx
	0x48, 0x01, 0xf0, //0x00000864 addq         %rsi, %rax
	0x48, 0x01, 0xd7, //0x00000867 addq         %rdx, %rdi
	0x0f, 0x85, 0x8a, 0xff, 0xff, 0xff, //0x0000086a jne          LBB0_76
	0xe9, 0x07, 0x00, 0x00, 0x00, //0x00000870 jmp          LBB0_85
	//0x00000875 LBB0_84
	0x49, 0xc7, 0xc3, 0xfe, 0xff, 0xff, 0xff, //0x00000875 movq         $-2, %r11
	//0x0000087c LBB0_85
	0x48, 0x8b, 0x5d, 0xb8, //0x0000087c movq         $-72(%rbp), %rbx
	//0x00000880 LBB0_86
	0x48, 0x8b, 0x75, 0xc0, //0x00000880 movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x7d, 0xc8, //0x00000884 movq         $-56(%rbp), %rdi
	0xe9, 0x4a, 0x00, 0x00, 0x00, //0x00000888 jmp          LBB0_94
	//0x0000088d LBB0_87
	0xb9, 0x40, 0x00, 0x00, 0x00, //0x0000088d movl         $64, %ecx
	//0x00000892 LBB0_88
	0x48, 0x8b, 0x75, 0xc0, //0x00000892 movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x7d, 0xc8, //0x00000896 movq         $-56(%rbp), %rdi
	0x48, 0x39, 0xd1, //0x0000089a cmpq         %rdx, %rcx
	0x0f, 0x82, 0x29, 0x00, 0x00, 0x00, //0x0000089d jb           LBB0_92
	0x48, 0x2b, 0x45, 0xa8, //0x000008a3 subq         $-88(%rbp), %rax
	0x4c, 0x8d, 0x1c, 0x10, //0x000008a7 leaq         (%rax,%rdx), %r11
	0x49, 0x83, 0xc3, 0x01, //0x000008ab addq         $1, %r11
	0xe9, 0x91, 0xfb, 0xff, 0xff, //0x000008af jmp          LBB0_30
	//0x000008b4 LBB0_90
	0x48, 0x2b, 0x45, 0xa8, //0x000008b4 subq         $-88(%rbp), %rax
	0x48, 0x83, 0xc0, 0x01, //0x000008b8 addq         $1, %rax
	0x49, 0x89, 0xc3, //0x000008bc movq         %rax, %r11
	//0x000008bf LBB0_91
	0x48, 0x8b, 0x75, 0xc0, //0x000008bf movq         $-64(%rbp), %rsi
	0x48, 0x8b, 0x7d, 0xc8, //0x000008c3 movq         $-56(%rbp), %rdi
	0xe9, 0x79, 0xfb, 0xff, 0xff, //0x000008c7 jmp          LBB0_30
	//0x000008cc LBB0_92
	0x49, 0xc7, 0xc3, 0xfe, 0xff, 0xff, 0xff, //0x000008cc movq         $-2, %r11
	//0x000008d3 LBB0_93
	0x48, 0x8b, 0x5d, 0xb8, //0x000008d3 movq         $-72(%rbp), %rbx
	//0x000008d7 LBB0_94
	0x48, 0x89, 0x1f, //0x000008d7 movq         %rbx, (%rdi)
	0x4c, 0x89, 0x1e, //0x000008da movq         %r11, (%rsi)
	//0x000008dd LBB0_95
	0x48, 0x83, 0xc4, 0x38, //0x000008dd addq         $56, %rsp
	0x5b,       //0x000008e1 popq         %rbx
	0x41, 0x5c, //0x000008e2 popq         %r12
	0x41, 0x5d, //0x000008e4 popq         %r13
	0x41, 0x5e, //0x000008e6 popq         %r14
	0x41, 0x5f, //0x000008e8 popq         %r15
	0x5d, //0x000008ea popq         %rbp
	0xc3, //0x000008eb retq
	//0x000008ec LBB0_96
	0x4d, 0x85, 0xed, //0x000008ec testq        %r13, %r13
	0x0f, 0x84, 0x88, 0x00, 0x00, 0x00, //0x000008ef je           LBB0_101
	0x4c, 0x89, 0xe7, //0x000008f5 movq         %r12, %rdi
	0x48, 0xf7, 0xd7, //0x000008f8 notq         %rdi
	0x49, 0x8d, 0x04, 0x38, //0x000008fb leaq         (%r8,%rdi), %rax
	0x48, 0x8b, 0x4d, 0xd0, //0x000008ff movq         $-48(%rbp), %rcx
	0x48, 0x83, 0xf9, 0xff, //0x00000903 cmpq         $-1, %rcx
	0x48, 0x89, 0xcb, //0x00000907 movq         %rcx, %rbx
	0x48, 0x0f, 0x44, 0xc8, //0x0000090a cmoveq       %rax, %rcx
	0x48, 0x0f, 0x44, 0xd8, //0x0000090e cmoveq       %rax, %rbx
	0x49, 0x83, 0xc0, 0x01, //0x00000912 addq         $1, %r8
	0x49, 0x83, 0xc5, 0xff, //0x00000916 addq         $-1, %r13
	0x48, 0x89, 0x4d, 0xd0, //0x0000091a movq         %rcx, $-48(%rbp)
	0x4d, 0x85, 0xed, //0x0000091e testq        %r13, %r13
	0x0f, 0x85, 0x67, 0xfd, 0xff, 0xff, //0x00000921 jne          LBB0_52
	0xe9, 0xed, 0xfd, 0xff, 0xff, //0x00000927 jmp          LBB0_60
	//0x0000092c LBB0_98
	0x4d, 0x85, 0xd2, //0x0000092c testq        %r10, %r10
	0x0f, 0x84, 0x48, 0x00, 0x00, 0x00, //0x0000092f je           LBB0_101
	0x4c, 0x89, 0xd7, //0x00000935 movq         %r10, %rdi
	0x48, 0x8b, 0x4d, 0xa8, //0x00000938 movq         $-88(%rbp), %rcx
	0x48, 0xf7, 0xd1, //0x0000093c notq         %rcx
	0x48, 0x01, 0xc1, //0x0000093f addq         %rax, %rcx
	0x48, 0x8b, 0x75, 0xd0, //0x00000942 movq         $-48(%rbp), %rsi
	0x48, 0x83, 0xfe, 0xff, //0x00000946 cmpq         $-1, %rsi
	0x48, 0x89, 0xf2, //0x0000094a movq         %rsi, %rdx
	0x48, 0x0f, 0x44, 0xd1, //0x0000094d cmoveq       %rcx, %rdx
	0x48, 0x0f, 0x45, 0xce, //0x00000951 cmovneq      %rsi, %rcx
	0x48, 0x83, 0xc0, 0x01, //0x00000955 addq         $1, %rax
	0x48, 0x83, 0xc7, 0xff, //0x00000959 addq         $-1, %rdi
	0x48, 0x89, 0x55, 0xd0, //0x0000095d movq         %rdx, $-48(%rbp)
	0x49, 0x89, 0xfa, //0x00000961 movq         %rdi, %r10
	0x48, 0x8b, 0x5d, 0xb0, //0x00000964 movq         $-80(%rbp), %rbx
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x00000968 movq         $-1, %r11
	0x4d, 0x85, 0xd2, //0x0000096f testq        %r10, %r10
	0x0f, 0x85, 0x7f, 0xfe, 0xff, 0xff, //0x00000972 jne          LBB0_75
	0xe9, 0xff, 0xfe, 0xff, 0xff, //0x00000978 jmp          LBB0_85
	//0x0000097d LBB0_101
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x0000097d movq         $-1, %r11
	0xe9, 0xf3, 0xfe, 0xff, 0xff, //0x00000984 jmp          LBB0_85
	0x00, 0x00, 0x00, //0x00000989 .p2align 2, 0x00
	//0x0000098c _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x0000098c .long 2
}