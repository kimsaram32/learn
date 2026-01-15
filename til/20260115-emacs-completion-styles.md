# 20260115 - Emacs completion styles

## Completion styles

Emacs의 default completion에서,
completion style은 completion의 동작 방식을 지정하는 것이다.

그냥 completion algorithm이라고 생각하면 될 것 같음

- 전체 completion style 들은 `completion-styles-alist` 에 저장된다.
- customization 시엔 위 변수에 존재하는 completion style의 이름 중 하나를 사용한다.

## Default completion의 동작 방식

특정 completion context 상에서,
default completion은 completion style을 다음과 같이 사용한다.

우선 사용할 전체 completion style의 목록을 만든다.
이는 다음 변수들이 순서대로 병합된다.

1. `completion-category-overrides` 에서 현재 category에 해당하는 styles
2. `completion-category-defaults` 에서 현재 category에 해당하는 styles
3. `completion-styles` 에 존재하는 styles

이후 만들어진 목록의 completion style을 하나씩 돌면서 매칭을 시도한다.
- 만약 매칭된 candidate가 존재한다면, 그것을 completion의 결과값으로 사용한다.
- 그렇지 않으면 다음 순서의 completion style을 시도한다.

따라서 completion 목록에서 앞선 값이 뒤의 것을 override 한다.
즉 category 별 override 한 값이 기본 `completion-styles` 를 덮어쓴다.

## Default completion 더 알아본 것들

### Metadata

현재 completion에 대한 metadata를 지정할 수 있다.

metadata에는 여러 필드가 있다. 그 중 내가 재밌게 본 것들
- category 필드는 style 등의 override 에 활용되는, completion의 context 정보를 담는다.
  completion을 수행하는 대상이 무엇인가? 를 지정한다. (e.g. file, org-heading, symbol...)
- annotation 필드는 candidate 옆에 추가로 표시할 정보를 지정한다.

metadata는 다음과 같은 사용처가 있다.
- category 필드를 기반으로 completion style를 override 하는 등의 customization을 진행한다.
- completion UI를 띄울 때 metadata에 지정한 것들을 사용한다 (annotation 달기, grouping...)

metadata의 지정 방식에는 다음과 같은 것들이 있다.
- programmed completion에서 metadata request에 대한 반환값 지정
- `completion-extra-properties` 에 값을 let-bound로 할당하여, 현재 completion에 한한 metadata 설정

### Default completion을 사용하는 3rd party package

Ivy, Helm과 같은 프레임워크들은 default completion을 override 하는 방식이다.
그러나 default completion 위에서 동작하는 패키지들도 있다.

- Vertico, Icomplete: completion UI 제공
- Consult: `completing-read`를 기반으로 search & navigation 진행해주는 framework...
  default completion을 Ivy, Helm과 같은 포지션으로 사용하기 위한 것이다.
- Marginalia: completion candidate에 대한 custom annotation을 지정할 수 있게 해 준다.

## Todo

- completion 기능의 역사를 알아야 아직 이해하지 못한 부분들을 파악할 수 있을 것이다.
- API가 왜 이렇게 설계되었지? 같은 부분들은 초기 버전을 보고 그것이 어떻게 개선되었는지를 알아야 이해할 수 있으니

## Thoughts

- Emacs의 빌트인 기능은 단순한 DS로 높은 확장성을 가지는 것들이 많다.
- 나는 Helm과 Ivy 같이 default behavior를 "overwrite" 하는 기분을 들게 하는 것들은 선호하지 않는 것 같다.
  빌트인을 기반으로 확장해 나가는 것을 더 좋아하는 듯
