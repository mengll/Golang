### go 汇编指令
Go 的调用规约强制我们将所有参数都通过栈来进行传递。

.. Go语言有4个伪寄存器，实际是对内存位置的一个引用。

FP（frame pointer）: 帧指针，保存参数和本地变量
 编译器维护了一个虚拟的栈指针，使用对伪寄存器的offsets操作的形式，指向栈上的函数参数。 于是，0(FP)就是第一个参数，8(FP)就是第二个(64位机器)

PC：程序指针，负责跳转和流程控制

SB（static base）: 静态基指针，全局变量
SB伪寄存器用来表示全局的变量或者函数，比如foo(SB)用来表示foo的地址。加<>表示符号本文件内可见。

SP（stack pointer）：栈指针，栈顶
 用来指向栈帧本地的变量和为函数调用准备参数。它指向本地栈帧的顶部，所以一个对栈帧的引用必须是一个负值且范围在[-framesize:0]之间，例如: x-8(SP)，y-4(SP)。 0（SP）表示第一个局部变量

$0-16: $0 代表即将分配的栈帧大小；而 $16 指定了调用方传入的参数大小。

"".b+12(SP) 和 "".a+8(SP) 分别指向栈的低 12 字节和低 8 字节位置(记住: 栈是向低位地址方向增长的！)。
.a 和 .b 是分配给引用地址的任意别名

第一个变量 a 的地址并不是 0(SP)，而是在 8(SP)；这是因为调用方通过使用 CALL 伪指令，把其返回地址保存在了 0(SP) 位置。
参数是反序传入的；也就是说，第一个参数和栈顶距离最近。

0x0008 ADDL CX, AX
0x000a MOVL AX, "".~r2+16(SP)
0x000e MOVB $1, "".~r3+20(SP)

ADDL 进行实际的加法操作，L 这里代表 Long，4 字节的值，其将保存在 AX 和 CX 寄存器中的值进行相加，然后再保存进 AX 寄存器中。
这个结果之后被移动到 "".~r2+16(SP) 地址处，这是之前调用方专门为返回值预留的栈空间。这一次 "".~r2 同样没什么语义上的含义。

;; Declare global function symbol "".add (actually main.add once linked)
;; Do not insert stack-split preamble
;; 0 bytes of stack-frame, 16 bytes of arguments passed in
;; func add(a, b int32) (int32, bool)
0x0000 TEXT	"".add(SB), NOSPLIT, $0-16
  ;; ...omitted FUNCDATA stuff...
  0x0000 MOVL	"".b+12(SP), AX	    ;; move second Long-word (4B) argument from caller's stack-frame into AX
  0x0004 MOVL	"".a+8(SP), CX	    ;; move first Long-word (4B) argument from caller's stack-frame into CX
  0x0008 ADDL	CX, AX		    ;; compute AX=CX+AX
  0x000a MOVL	AX, "".~r2+16(SP)   ;; move addition result (AX) into caller's stack-frame
  0x000e MOVB	$1, "".~r3+20(SP)   ;; move `true` boolean (constant) into caller's stack-frame
  0x0013 RET			    ;; jump to return address stored at 0(SP)

  |    +-------------------------+ <-- 32(SP)              
  |    |                         |                         
G |    |                         |                         
R |    |                         |                         
O |    | main.main's saved       |                         
W |    |     frame-pointer (BP)  |                         
S |    |-------------------------| <-- 24(SP)              
  |    |      [alignment]        |                         
D |    | "".~r3 (bool) = 1/true  | <-- 21(SP)              
O |    |-------------------------| <-- 20(SP)              
W |    |                         |                         
N |    | "".~r2 (int32) = 42     |                         
W |    |-------------------------| <-- 16(SP)              
A |    |                         |                         
R |    | "".b (int32) = 32       |                         
D |    |-------------------------| <-- 12(SP)              
S |    |                         |                         
  |    | "".a (int32) = 10       |                         
  |    |-------------------------| <-- 8(SP)               
  |    |                         |                         
  |    |                         |                         
  |    |                         |                         
\ | /  | return address to       |                         
 \|/   |     main.main + 0x30    |                         
  -    +-------------------------+ <-- 0(SP) (TOP OF STACK)

  0x000f SUBQ     $24, SP
  将其栈帧大小增加了 24 个字节(回忆一下栈是向低地址方向增长，所以这里的 SUBQ 指令是将栈帧的大小调整得更大了)

  8 个字节(16(SP)-24(SP)) 用来存储当前帧指针 BP (这是一个实际存在的寄存器)的值，以支持栈的展开和方便调试
  1+3 个字节(12(SP)-16(SP)) 是预留出的给第二个返回值 (bool) 的空间，除了类型本身的 1 个字节，在 amd64 平台上还额外需要 3 个字节来做对齐
  4 个字节(8(SP)-12(SP)) 预留给第一个返回值 (int32)
  4 个字节(4(SP)-8(SP)) 是预留给传给被调用函数的参数 b (int32)
  4 个字节(0(SP)-4(SP)) 预留给传入参数 a (int32)
  最后，跟着栈的增长，LEAQ 指令计算出帧指针的新地址，并将其存储到 BP 寄存器中。

  注意 CALL 指令还会将函数的返回地址(8 字节值)也推到栈顶；所以每次我们在 add 函数中引用 SP 寄存器的时候还需要额外偏移 8 个字节！
  例如，"".a 现在不是 0(SP) 了，而是在 8(SP) 位置。
  可以看到，栈分裂(stack-split)前导码被分成 prologue 和 epilogue 两个部分:

  ```
  0x0000 TEXT	"".main(SB), $24-0
  ;; stack-split prologue
  0x0000 MOVQ	(TLS), CX
  0x0009 CMPQ	SP, 16(CX)
  0x000d JLS	58

  0x000f SUBQ	$24, SP
  0x0013 MOVQ	BP, 16(SP)
  0x0018 LEAQ	16(SP), BP
  ;; ...omitted FUNCDATA stuff...
  0x001d MOVQ	$137438953482, AX
  0x0027 MOVQ	AX, (SP)
  ;; ...omitted PCDATA stuff...
  0x002b CALL	"".add(SB)
  0x0030 MOVQ	16(SP), BP
  0x0035 ADDQ	$24, SP
  0x0039 RET

  ;; stack-split epilogue
  0x003a NOP
  ;; ...omitted PCDATA stuff...
  0x003a CALL	runtime.morestack_noctxt(SB)
  0x003f JMP	0

  ```

  prologue 会检查当前 goroutine 是否已经用完了所有的空间，然后如果确实用完了的话，会直接跳转到后部。
  epilogue 会触发栈增长(stack-growth)，然后再跳回到前部

```
  0x0000 MOVQ	(TLS), CX   ;; store current *g in CX
  0x0009 CMPQ	SP, 16(CX)  ;; compare SP and g.stackguard0
  0x000d JLS	58	    ;; jumps to 0x3a if SP <= g.stackguard0


  type g struct {
	stack       stack   // 16 bytes
	// stackguard0 is the stack pointer compared in the Go stack growth prologue.
	// It is stack.lo+StackGuard normally, but can be StackPreempt to trigger a preemption.
	stackguard0 uintptr
	stackguard1 uintptr

	// ...omitted dozens of fields...
}
```
  TLS 是一个由 runtime 维护的虚拟寄存器，保存了指向当前 g 的指针，这个 g 的数据结构会跟踪 goroutine 运行时的所有状态值。
  我们可以看到 16(CX) 对应的是 g.stackguard0，是 runtime 维护的一个阈值，该值会被拿来与栈指针(stack-pointer)进行比较以判断一个 goroutine 是否马上要用完当前的栈空间。


1）TEXT指令声明了符号””.add，指令紧接在类似于函数的主体中。TEXT块的最后必须是某种形式的跳转，通常是一个RET(伪)指令。(如果没有，链接器会追加一个跳转到块自身的指令，TEXT块中没有fallthrough) 符号的后面，参数是标志和栈帧的大小.
TEXT    "".add(SB), $0-24    $0-24表示栈帧的大小，描述了函数有0字节的栈帧，并且需要24字节的参数和返回值的空间

2）MOVQ指令
  MOV指令就是将左侧的内容放到右侧的寄存器或者地址中去

3）LEAQ指令：（Load Effective Address） 这个表示将左侧的变量取地址，然后存放在右边的寄存器中
  LEAQ type.int(SB), AX

4）FUNCDATA指令：这个是golang编译器自带的指令
  FUNCDATA    $0, gclocals·54241e171da8af6ae173d69da0236748(SB)
  FUNCDATA    $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
  标注做GC的一些变量，用来给GC收集进行提示

5）CALL指令
  CALL runtime.convT2E(SB)
  这个表示函数调用，将调用runtime包下面convT2E函数

6）MOVL
  把全局数组的地址压栈：MOVL $array(SB), TOS。
  把全局数组的第二个元素压栈：MOVL array+4(SB), TOS

7） FUNCDATA和PCDATA是编译器产生的，用于保存一些给垃圾收集的信息。

8）指令修饰符
  DUPOK:允许一个二进制文件里有多个实例

  NOSPLIT: FOR TEXT，routine或者routine的子函数，必须把栈的空间的头填满，用来保护栈分隔

  RODATA:FOR DATA/GLOBL，把数据放在只读段

  NOPTR: FOR DATA/GLOBL，数据没有指针，不需要被垃圾收集扫描

  WRAPPER: FOR TEXT，wrapper function，不需要被以禁用recover计数

  NEEDCTXT:FOR TEXT，闭包

9）NOSPLIT表示不用写入参数的大小

10）ADDQ 加法操作

11) retq 从函数返回
  最后的 RET 伪指令告诉 Go 汇编器插入一些指令，这些指令是对应的目标平台中的调用规约所要求的，从子过程中返回时所需要的指令。
  一般情况下这样的指令会使在 0(SP) 寄存器中保存的函数返回地址被 pop 出栈，并跳回到该地址。
