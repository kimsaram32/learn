# 20251222 - System call wrapper functions (recap / short)

standard library는 system call에 대한 wrapper function을 제공한다.
내부 구현은 실제 대응되는 system call을 호출하여 이루어질 수 있지만,
그렇지 않은 경우도 있다 (e.g. VDSO - todo).

application program은 이러한 wrapper function을 호출하는 형태로 system call을 간접적으로 사용한다.

application -> stdlib -> syscall의 흐름을 가진다.

Linux Kernel Development ch. 5
