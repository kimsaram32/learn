# 20251120 - x64 system call internals

x64 architecture에서의 system call을 알아보았다.
저번에 arm64 환경에서 했던 것과 비슷하게 진행하였다.

## objdump로 dump 하기

`ldd` 로 glibc의 경로를 알아낸 후, `objdump` 로 dump를 생성하였다.

## system call instruction 살펴보기

1. x64 상에서 system call을 날릴 땐 주로 SYSCALL instruction을 사용한다.
이를 확인하기 위해 dump 상에서 `syscall` 로 검색해보니 많이 찾아볼 수 있었다.

2. register를 통해 인수를 전달하는 것도 확인할 수 있었다.
syscall num을 저장하는 register는 %eax 또는 %rax를 사용하는 것을 확인할 수 있었다.
%rax에 syscall num을 지정한다고 배웠지만,
%eax가 %rax의 rightmost 32 bit에 해당되어서 둘 다 쓰는 것으로 보인다 [Stack Overflow](https://stackoverflow.com/questions/44972293/how-is-rax-different-from-eax).
