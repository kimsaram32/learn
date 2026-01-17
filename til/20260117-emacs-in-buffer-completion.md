# 20260117 - Emacs in-buffer completion

## `completion-at-point`

- 현재 buffer의 point를 기준으로 completion candidate 목록을 만들어낸 뒤 completion을 진행한다.
- 내부적으로 `completion-in-region` 을 호출하며, completion UI는 여기서 그려주는 듯

## 3rd-party package: company-mode에서 corfu로 옮기기

built-in completion을 개선하는 3rd party 패키지들이 존재한다.
기존에는 company-mode를 사용했었으나, 다음과 같은 이유로 넘어가기로 했다.

- company-mode는 자체적인 frontend와 backend 기능을 가지고 동작한다. 나는 Emacs의 standard API를 사용하는 것을 선호하는 편이라는 것을 최근에 알게 되었다. 때문에 나와 맞지 않는다는 생각이 들었다.
- 내가 잘 못 찾아서 그런 것일 수도 있으나 문서화가 부족해서 customization에 어려움을 겪었다. 직접 소스를 뜯기보다는 그냥 다른 패키지를 시도하는 것이 더 낫다고 판단했다.

옮겨간 패키지는 corfu 이며, `completion-at-point` 를 기반으로 작동하고 UI만 바꿔주는 형태이다.

## Autocomplete delay 관련

company-mode를 사용할 땐 autocompletion의 delay를 0으로 두었었다. 즉 타이핑 즉시 completion을 시작하는 형태였다.

corfu의 문서를 읽던 도중, 이 설정을 권장하지 않는다는 내용을 보게 되었다. 불필요한 경우에도 completion을 위한 연산이 요구되기에 오버헤드가 될 수 있다고 한다.

따라서 corfu에서는 autocompletion에만 의존하지 않고 텍스트 트리거 (현재 `.` 또는 `-` 입력 시 트리거되도록 설정함) 등의 기능을 더 사용해보기로 했다.
