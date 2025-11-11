# 20251111 - System call 작동 방식

로우레벨 알못이지만 알게 된 내용 적어보기

## 인수 전달 방식

- 인수는 레지스터에 지정된다.
  각 아키텍처별로 사용되는 레지스터는 `syscall(2)` man page에서 확인할 수 있다.
- 추가 데이터는 포인터로 넘긴다.
- 사실 system call specific 적인 부분은 아닌 것 같기도 함

## 구현 방식

system call의 발동을 구현하는 방식에는 여러 가지가 있다.

- 인터럽트 방식
  - 인터럽트 핸들러 기반으로 구현한다.
  - e.g. x86 - `int 0x80`
- instruction 방식
  - architecture specific한 instruction이 kernel mode의 빠른 전환을 해준다.
  - e.g. arm `SVC`, x86 `SYSENTER`, `SYSEXIT`, `SYSCALL` (newer)

인터럽트 방식은 성능 문제로 인해 이제는 obsolete이고, 현재는 instruction 방식이 사용된다고 한다.

## arm64 구조에서의 system call 방식 알아보기

macOS에서 Ubuntu VM을 돌려서, arm64 구조에서의 system call 방식을 뜯어보았다.

### libc shared library 뜯어보기: objdump

objdump를 이용해 libc를 disassemble 한다.
```sh
objdump -d /lib/aarch64-linux-gnu/libc.so.6 | less
```

### `SVC` instruction 찾아보기

arm64에서 system call을 구현하는 instruction은 `SVC`이므로 해당 instruction을 검색해보았다.

찾다 보니 이런 부분이 나왔다
```
00000000000b41e0 <setitimer@@GLIBC_2.17>:
   b41e0:       a9bf7bfd        stp     x29, x30, [sp, #-16]!
   b41e4:       2a0003e0        mov     w0, w0
   b41e8:       d2800ce8        mov     x8, #0x67                       // #103
   b41ec:       910003fd        mov     x29, sp
   b41f0:       d4000001        svc     #0x0
   ... (생략)
```

### 찾은 부분 해석하기

1. standard library의 `setitimer` 함수에 해당하는 부분이다.
2. x8 레지스터에 들어간 103이 syscall number에 해당하는 것으로 보여서,
  [arm64 system call table](https://www.chromium.org/chromium-os/developer-library/reference/linux-constants/syscalls/#arm64-64-bit) 에서 찾아보았다.
3. 103에 해당하는 system call은 `setitimer`이다.
4. 즉 이 함수는 `setitimer` system call의 wrapper function 임을 알 수 있었다.
