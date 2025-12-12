# 20251212 - emacs-mac Installation Troubleshooting

Emacs의 macOS용 port 중 하나인 [jdtsmith/emacs-mac](https://github.com/jdtsmith/emacs-mac/tree/a9f6592fe275774e6c14e3058a4f3488f9914e5e)
을 설치하면서 겪은 트러블슈팅 과정을 정리하였다.

## libgccjit 없음: Homebrew와 관련한 path 설정

`./configure` 도중 libgccjit가 없다는 에러 메시지가 발생했다.
libgccjit를 Homebrew로 설치했지만, Homebrew에서 해당 라이브러리를 설치한 경로가
컴파일 및 링킹 시 플래그로 등록되어 있지 않기 때문이었다.

따라서 다음을 `.zshrc` 에 추가하여 해결하였다.

```sh
export BREW="/opt/homebrew"
export LIBRARY_PATH="$LIBRARY_PATH:$BREW/lib/gcc/current"
export LDFLAGS="-L$BREW/lib"
export CPPFLAGS="-I$BREW/include"
```

## tree-sitter 버전 충돌 문제

`make` 도중 다음과 같은 에러 메시지가 발생했다.

```
treesit.c:749:21: error: call to undeclared function 'ts_language_version'; ISO C99 and later do not support implicit function declarations [-Wimplicit-function-declaration]
  749 |                             make_fixnum (ts_language_version (lang)));
      |                                          ^
treesit.c:749:21: note: did you mean 'ts_language_abi_version'?
/opt/homebrew/include/tree_sitter/api.h:1249:10: note: 'ts_language_abi_version' declared here
 1249 | uint32_t ts_language_abi_version(const TSLanguage *self);
      |          ^
treesit.c:820:27: error: call to undeclared function 'ts_language_version'; ISO C99 and later do not support implicit function declarations [-Wimplicit-function-declaration]
  820 |       uint32_t version =  ts_language_version (ts_language);
```

### 시도: Homebrew `--HEAD` 설치 (실패)

`tree-sitter` 의 버전 차이로 발생하는 문제로 보였기에,
`brew intall --HEAD tree-sitter` 로 직접 빌드하여 설치해 보았지만 실패하였다.

### 해결: 0.25 버전 설치

[upstream Emacs에서의 treesit.c](https://github.com/emacs-mirror/emacs/blob/10d022741cf854817ec241c0d1c7b68b3f748d68/src/treesit.c#L730)
코드를 살펴보니, `ts_language_version` 이 tree-sitter 0.26 버전에서 제거되었고,
해당 버전에서는 `ts_language_abi_version` 을 사용하도록 되어 있었다.

그러나 [포팅된 Emacs 소스 코드에서의 treesit.c](https://github.com/jdtsmith/emacs-mac/blob/a9f6592fe275774e6c14e3058a4f3488f9914e5e/src/treesit.c) 에서는
여전히 `ts_language_version` 을 사용하고 있었으며,
내 환경에서 tree-sitter는 0.26 버전이 설치되어 있었다.

따라서 0.25 버전의 tree-sitter를 설치함으로써 해결하였다.

### 요약

![](./20251212-versions.png)

### GitHub Issue

해결을 하고 나니 바로 직전에 [GitHub Issue](https://github.com/jdtsmith/emacs-mac/issues/121) 가 열렸었음을 확인하였다.

## Pure Storage Overflow 문제

startup 시 다음과 같은 경고 메시지가 발생했다.

```
Warning Warning!!! Pure space overflow !!!Warning Warning
```

[info 문서의 Pure Storage 부분](https://www.gnu.org/software/emacs/manual/html_node/elisp/Pure-Storage.html) 에서 `src/puresize.h` 의 `SYSTEM_PURESIZE_EXTRA` 를 늘려 해결하라고 작성되어 있었다.
해당 지침을 따라서 해결하였다.

```c
#ifndef SYSTEM_PURESIZE_EXTRA
#define SYSTEM_PURESIZE_EXTRA 100000
#endif
```

[linuxquestions](https://www.linuxquestions.org/questions/slackware-14/gnu-emacs-27-1-pure-space-overflow-warning-on-startup-screen-4175688386/)

## Meta key 바인딩 변경됨

1. 이전 빌드에선 Mac Option key -> Emacs Meta key였기에,
   사용의 편의를 위해 Karabiner를 통해 Command key를 Option key로 매핑해놨었다.
2. 여기 빌드에서는 Mac Command key -> Emacs Meta key인 것으로 보였다.
   따라서 기존 키바인딩을 제거하여 이전과 같이 사용할 수 있도록 하였다.
