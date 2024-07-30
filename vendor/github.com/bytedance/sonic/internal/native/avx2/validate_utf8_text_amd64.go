//go:build amd64
// +build amd64

// Code generated by asm2asm, DO NOT EDIT.

package avx2

var _text_validate_utf8 = []byte{
	// .p2align 4, 0x90
	// _validate_utf8
	0x55,             // pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000001 movq         %rsp, %rbp
	0x41, 0x57, //0x00000004 pushq        %r15
	0x41, 0x56, //0x00000006 pushq        %r14
	0x41, 0x54, //0x00000008 pushq        %r12
	0x53,             //0x0000000a pushq        %rbx
	0x50,             //0x0000000b pushq        %rax
	0x4c, 0x8b, 0x17, //0x0000000c movq         (%rdi), %r10
	0x4c, 0x8b, 0x5f, 0x08, //0x0000000f movq         $8(%rdi), %r11
	0x48, 0x8b, 0x0e, //0x00000013 movq         (%rsi), %rcx
	0x4c, 0x01, 0xd1, //0x00000016 addq         %r10, %rcx
	0x4f, 0x8d, 0x04, 0x1a, //0x00000019 leaq         (%r10,%r11), %r8
	0x49, 0x83, 0xc0, 0xfd, //0x0000001d addq         $-3, %r8
	0xe9, 0x0d, 0x00, 0x00, 0x00, //0x00000021 jmp          LBB0_1
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000026 .p2align 4, 0x90
	//0x00000030 LBB0_19
	0x48, 0x01, 0xd9, //0x00000030 addq         %rbx, %rcx
	//0x00000033 LBB0_1
	0x4c, 0x39, 0xc1, //0x00000033 cmpq         %r8, %rcx
	0x0f, 0x83, 0xe1, 0x00, 0x00, 0x00, //0x00000036 jae          LBB0_2
	0xbb, 0x01, 0x00, 0x00, 0x00, //0x0000003c movl         $1, %ebx
	0x80, 0x39, 0x00, //0x00000041 cmpb         $0, (%rcx)
	0x0f, 0x89, 0xe6, 0xff, 0xff, 0xff, //0x00000044 jns          LBB0_19
	0x8b, 0x01, //0x0000004a movl         (%rcx), %eax
	0x89, 0xc7, //0x0000004c movl         %eax, %edi
	0x81, 0xe7, 0xf0, 0xc0, 0xc0, 0x00, //0x0000004e andl         $12632304, %edi
	0x81, 0xff, 0xe0, 0x80, 0x80, 0x00, //0x00000054 cmpl         $8421600, %edi
	0x0f, 0x85, 0x30, 0x00, 0x00, 0x00, //0x0000005a jne          LBB0_10
	0x89, 0xc7, //0x00000060 movl         %eax, %edi
	0x81, 0xe7, 0x0f, 0x20, 0x00, 0x00, //0x00000062 andl         $8207, %edi
	0x81, 0xff, 0x0d, 0x20, 0x00, 0x00, //0x00000068 cmpl         $8205, %edi
	0x0f, 0x84, 0x1c, 0x00, 0x00, 0x00, //0x0000006e je           LBB0_10
	0xbb, 0x03, 0x00, 0x00, 0x00, //0x00000074 movl         $3, %ebx
	0x85, 0xff, //0x00000079 testl        %edi, %edi
	0x0f, 0x85, 0xaf, 0xff, 0xff, 0xff, //0x0000007b jne          LBB0_19
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000081 .p2align 4, 0x90
	//0x00000090 LBB0_10
	0x89, 0xc7, //0x00000090 movl         %eax, %edi
	0x81, 0xe7, 0xe0, 0xc0, 0x00, 0x00, //0x00000092 andl         $49376, %edi
	0x81, 0xff, 0xc0, 0x80, 0x00, 0x00, //0x00000098 cmpl         $32960, %edi
	0x0f, 0x85, 0x10, 0x00, 0x00, 0x00, //0x0000009e jne          LBB0_12
	0x89, 0xc7, //0x000000a4 movl         %eax, %edi
	0xbb, 0x02, 0x00, 0x00, 0x00, //0x000000a6 movl         $2, %ebx
	0x83, 0xe7, 0x1e, //0x000000ab andl         $30, %edi
	0x0f, 0x85, 0x7c, 0xff, 0xff, 0xff, //0x000000ae jne          LBB0_19
	//0x000000b4 LBB0_12
	0x89, 0xc7, //0x000000b4 movl         %eax, %edi
	0x81, 0xe7, 0xf8, 0xc0, 0xc0, 0xc0, //0x000000b6 andl         $-1061109512, %edi
	0x81, 0xff, 0xf0, 0x80, 0x80, 0x80, //0x000000bc cmpl         $-2139062032, %edi
	0x0f, 0x85, 0x26, 0x00, 0x00, 0x00, //0x000000c2 jne          LBB0_16
	0x89, 0xc7, //0x000000c8 movl         %eax, %edi
	0x81, 0xe7, 0x07, 0x30, 0x00, 0x00, //0x000000ca andl         $12295, %edi
	0x0f, 0x84, 0x18, 0x00, 0x00, 0x00, //0x000000d0 je           LBB0_16
	0xbb, 0x04, 0x00, 0x00, 0x00, //0x000000d6 movl         $4, %ebx
	0xa8, 0x04, //0x000000db testb        $4, %al
	0x0f, 0x84, 0x4d, 0xff, 0xff, 0xff, //0x000000dd je           LBB0_19
	0x25, 0x03, 0x30, 0x00, 0x00, //0x000000e3 andl         $12291, %eax
	0x0f, 0x84, 0x42, 0xff, 0xff, 0xff, //0x000000e8 je           LBB0_19
	//0x000000ee LBB0_16
	0x48, 0x89, 0xcf, //0x000000ee movq         %rcx, %rdi
	0x4c, 0x29, 0xd7, //0x000000f1 subq         %r10, %rdi
	0x48, 0x8b, 0x1a, //0x000000f4 movq         (%rdx), %rbx
	0x48, 0x81, 0xfb, 0x00, 0x10, 0x00, 0x00, //0x000000f7 cmpq         $4096, %rbx
	0x0f, 0x83, 0x97, 0x01, 0x00, 0x00, //0x000000fe jae          LBB0_17
	0x48, 0x63, 0xc7, //0x00000104 movslq       %edi, %rax
	0x48, 0x8d, 0x7b, 0x01, //0x00000107 leaq         $1(%rbx), %rdi
	0x48, 0x89, 0x3a, //0x0000010b movq         %rdi, (%rdx)
	0x48, 0x89, 0x44, 0xda, 0x08, //0x0000010e movq         %rax, $8(%rdx,%rbx,8)
	0xbb, 0x01, 0x00, 0x00, 0x00, //0x00000113 movl         $1, %ebx
	0xe9, 0x13, 0xff, 0xff, 0xff, //0x00000118 jmp          LBB0_19
	//0x0000011d LBB0_2
	0x4d, 0x01, 0xd3, //0x0000011d addq         %r10, %r11
	0x4c, 0x39, 0xd9, //0x00000120 cmpq         %r11, %rcx
	0x0f, 0x83, 0x4e, 0x01, 0x00, 0x00, //0x00000123 jae          LBB0_36
	0x4c, 0x8d, 0x45, 0xdc, //0x00000129 leaq         $-36(%rbp), %r8
	0x4c, 0x8d, 0x4d, 0xda, //0x0000012d leaq         $-38(%rbp), %r9
	0xe9, 0x17, 0x00, 0x00, 0x00, //0x00000131 jmp          LBB0_4
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000136 .p2align 4, 0x90
	//0x00000140 LBB0_5
	0x48, 0x83, 0xc1, 0x01, //0x00000140 addq         $1, %rcx
	0x4c, 0x39, 0xd9, //0x00000144 cmpq         %r11, %rcx
	0x0f, 0x83, 0x2a, 0x01, 0x00, 0x00, //0x00000147 jae          LBB0_36
	//0x0000014d LBB0_4
	0x80, 0x39, 0x00, //0x0000014d cmpb         $0, (%rcx)
	0x0f, 0x89, 0xea, 0xff, 0xff, 0xff, //0x00000150 jns          LBB0_5
	0xc6, 0x45, 0xdc, 0x00, //0x00000156 movb         $0, $-36(%rbp)
	0xc6, 0x45, 0xda, 0x00, //0x0000015a movb         $0, $-38(%rbp)
	0x4c, 0x89, 0xdb, //0x0000015e movq         %r11, %rbx
	0x48, 0x29, 0xcb, //0x00000161 subq         %rcx, %rbx
	0x48, 0x83, 0xfb, 0x02, //0x00000164 cmpq         $2, %rbx
	0x0f, 0x82, 0x35, 0x00, 0x00, 0x00, //0x00000168 jb           LBB0_21
	0x44, 0x0f, 0xb6, 0x21, //0x0000016e movzbl       (%rcx), %r12d
	0x44, 0x0f, 0xb6, 0x71, 0x01, //0x00000172 movzbl       $1(%rcx), %r14d
	0x44, 0x88, 0x65, 0xdc, //0x00000177 movb         %r12b, $-36(%rbp)
	0x4c, 0x8d, 0x79, 0x02, //0x0000017b leaq         $2(%rcx), %r15
	0x48, 0x83, 0xc3, 0xfe, //0x0000017f addq         $-2, %rbx
	0x4c, 0x89, 0xcf, //0x00000183 movq         %r9, %rdi
	0x48, 0x85, 0xdb, //0x00000186 testq        %rbx, %rbx
	0x0f, 0x84, 0x29, 0x00, 0x00, 0x00, //0x00000189 je           LBB0_24
	//0x0000018f LBB0_25
	0x41, 0x0f, 0xb6, 0x07, //0x0000018f movzbl       (%r15), %eax
	0x88, 0x07, //0x00000193 movb         %al, (%rdi)
	0x44, 0x0f, 0xb6, 0x65, 0xdc, //0x00000195 movzbl       $-36(%rbp), %r12d
	0x0f, 0xb6, 0x7d, 0xda, //0x0000019a movzbl       $-38(%rbp), %edi
	0xe9, 0x17, 0x00, 0x00, 0x00, //0x0000019e jmp          LBB0_26
	//0x000001a3 LBB0_21
	0x45, 0x31, 0xe4, //0x000001a3 xorl         %r12d, %r12d
	0x45, 0x31, 0xf6, //0x000001a6 xorl         %r14d, %r14d
	0x4c, 0x89, 0xc7, //0x000001a9 movq         %r8, %rdi
	0x49, 0x89, 0xcf, //0x000001ac movq         %rcx, %r15
	0x48, 0x85, 0xdb, //0x000001af testq        %rbx, %rbx
	0x0f, 0x85, 0xd7, 0xff, 0xff, 0xff, //0x000001b2 jne          LBB0_25
	//0x000001b8 LBB0_24
	0x31, 0xff, //0x000001b8 xorl         %edi, %edi
	//0x000001ba LBB0_26
	0x40, 0x0f, 0xb6, 0xc7, //0x000001ba movzbl       %dil, %eax
	0xc1, 0xe0, 0x10, //0x000001be shll         $16, %eax
	0x41, 0x0f, 0xb6, 0xde, //0x000001c1 movzbl       %r14b, %ebx
	0xc1, 0xe3, 0x08, //0x000001c5 shll         $8, %ebx
	0x09, 0xc3, //0x000001c8 orl          %eax, %ebx
	0x41, 0x0f, 0xb6, 0xfc, //0x000001ca movzbl       %r12b, %edi
	0x09, 0xdf, //0x000001ce orl          %ebx, %edi
	0x89, 0xf8, //0x000001d0 movl         %edi, %eax
	0x25, 0xf0, 0xc0, 0xc0, 0x00, //0x000001d2 andl         $12632304, %eax
	0x3d, 0xe0, 0x80, 0x80, 0x00, //0x000001d7 cmpl         $8421600, %eax
	0x0f, 0x85, 0x2e, 0x00, 0x00, 0x00, //0x000001dc jne          LBB0_29
	0x89, 0xf8, //0x000001e2 movl         %edi, %eax
	0x25, 0x0f, 0x20, 0x00, 0x00, //0x000001e4 andl         $8207, %eax
	0x3d, 0x0d, 0x20, 0x00, 0x00, //0x000001e9 cmpl         $8205, %eax
	0x0f, 0x84, 0x1c, 0x00, 0x00, 0x00, //0x000001ee je           LBB0_29
	0xbb, 0x03, 0x00, 0x00, 0x00, //0x000001f4 movl         $3, %ebx
	0x85, 0xc0, //0x000001f9 testl        %eax, %eax
	0x0f, 0x85, 0x30, 0x00, 0x00, 0x00, //0x000001fb jne          LBB0_34
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000201 .p2align 4, 0x90
	//0x00000210 LBB0_29
	0x41, 0xf6, 0xc4, 0x1e, //0x00000210 testb        $30, %r12b
	0x0f, 0x84, 0x28, 0x00, 0x00, 0x00, //0x00000214 je           LBB0_31
	0x81, 0xe7, 0xe0, 0xc0, 0x00, 0x00, //0x0000021a andl         $49376, %edi
	0xbb, 0x02, 0x00, 0x00, 0x00, //0x00000220 movl         $2, %ebx
	0x81, 0xff, 0xc0, 0x80, 0x00, 0x00, //0x00000225 cmpl         $32960, %edi
	0x0f, 0x85, 0x11, 0x00, 0x00, 0x00, //0x0000022b jne          LBB0_31
	//0x00000231 LBB0_34
	0x48, 0x01, 0xd9, //0x00000231 addq         %rbx, %rcx
	0x4c, 0x39, 0xd9, //0x00000234 cmpq         %r11, %rcx
	0x0f, 0x82, 0x10, 0xff, 0xff, 0xff, //0x00000237 jb           LBB0_4
	0xe9, 0x35, 0x00, 0x00, 0x00, //0x0000023d jmp          LBB0_36
	//0x00000242 LBB0_31
	0x48, 0x89, 0xc8, //0x00000242 movq         %rcx, %rax
	0x4c, 0x29, 0xd0, //0x00000245 subq         %r10, %rax
	0x48, 0x8b, 0x3a, //0x00000248 movq         (%rdx), %rdi
	0x48, 0x81, 0xff, 0x00, 0x10, 0x00, 0x00, //0x0000024b cmpq         $4096, %rdi
	0x0f, 0x83, 0x34, 0x00, 0x00, 0x00, //0x00000252 jae          LBB0_32
	0x48, 0x98, //0x00000258 cltq
	0x48, 0x8d, 0x5f, 0x01, //0x0000025a leaq         $1(%rdi), %rbx
	0x48, 0x89, 0x1a, //0x0000025e movq         %rbx, (%rdx)
	0x48, 0x89, 0x44, 0xfa, 0x08, //0x00000261 movq         %rax, $8(%rdx,%rdi,8)
	0xbb, 0x01, 0x00, 0x00, 0x00, //0x00000266 movl         $1, %ebx
	0x48, 0x01, 0xd9, //0x0000026b addq         %rbx, %rcx
	0x4c, 0x39, 0xd9, //0x0000026e cmpq         %r11, %rcx
	0x0f, 0x82, 0xd6, 0xfe, 0xff, 0xff, //0x00000271 jb           LBB0_4
	//0x00000277 LBB0_36
	0x4c, 0x29, 0xd1, //0x00000277 subq         %r10, %rcx
	0x48, 0x89, 0x0e, //0x0000027a movq         %rcx, (%rsi)
	0x31, 0xc0, //0x0000027d xorl         %eax, %eax
	//0x0000027f LBB0_37
	0x48, 0x83, 0xc4, 0x08, //0x0000027f addq         $8, %rsp
	0x5b,       //0x00000283 popq         %rbx
	0x41, 0x5c, //0x00000284 popq         %r12
	0x41, 0x5e, //0x00000286 popq         %r14
	0x41, 0x5f, //0x00000288 popq         %r15
	0x5d, //0x0000028a popq         %rbp
	0xc3, //0x0000028b retq
	//0x0000028c LBB0_32
	0x48, 0x89, 0x06, //0x0000028c movq         %rax, (%rsi)
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x0000028f movq         $-1, %rax
	0xe9, 0xe4, 0xff, 0xff, 0xff, //0x00000296 jmp          LBB0_37
	//0x0000029b LBB0_17
	0x48, 0x89, 0x3e, //0x0000029b movq         %rdi, (%rsi)
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x0000029e movq         $-1, %rax
	0xe9, 0xd5, 0xff, 0xff, 0xff, //0x000002a5 jmp          LBB0_37
	0x00, 0x00, //0x000002aa .p2align 2, 0x00
	//0x000002ac _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x000002ac .long 2
}
