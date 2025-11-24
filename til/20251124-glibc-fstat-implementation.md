# 20251124 - glibc의 fstat 구현

[Why glibc's fstat() is slow](https://lwn.net/Articles/944214/) 글을 정리한 것이다.

1.  glibc의 `fstat()` 은 `fstatat()` 의 호출로 구현되어 있다.
    `pathname` 을 `""`, `AT_EMPTY_PATH` flag를 지정하여
    `dirfd` 상에서 직접 동작하는 방식이다.
    즉 실제 fstat system call을 호출하지 않는다.

2.  이 방식은 overhead이다. 즉 `fstatat()`의 wrapper 방식을 사용하는 것보다, 단순한 system call을 호출하는 게 더 빠르다.
    이로 인한 성능 저하가 발생한다는 것을 발견하였다.

3.  Linux kernel 상에서 patch를 적용하였다.
    -   처음에는 `""` 와 `AT_EMPTY_PATH` 조합의 인수를 받으면 fstat으로 전환하는 패치를 하였으나,
        해당 코드에서의 empty string 확인 부분으로 인한 성능 문제는 여전히 발생했다.
    -   나중 patch는 정확히는 모르겠지만 ad-hoc과 같은 방식으로 보인다

궁금해서 현재 소스 코드를 확인해 봤는데, 여전히 fstatat을 사용하는 것으로 보였다 ([GitHub mirror](https://github.com/bminor/glibc/blob/0f7b73f2ed70e783cd02ab77503645b03ee1d332/io/fstat.c#L29)).

