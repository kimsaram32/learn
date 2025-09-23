# 20250923 - Elisp evaluation 과정

evaluation은 object가 가진 value를 얻는 과정이다.

object의 타입에 따라 evaluation 방법이 달라진다.

## Evaluate되는 object 호칭 방법

evaluate될 object를 가리키는 말로는 'form', 'expression', 's-expression'등이 있다.

expression은 다른 언어에선 텍스트 상에서와 관련 있으나, Lisp에선 object 자체를 나타낸다.
다만 상황에 따라 다른 언어와 비슷하게 object를 가리키는 텍스트 형식 (= read format) 을 가리키기도 한다.

## Self-evaluating objects

list와 symbol을 제외한 타입의 object는 self-evaluate 하는데,
evaluation의 결과 object가 자기 자신이라는 소리이다.

## Symbols

symbol의 경우엔 variable로써 evaluate되어 값을 얻는다.

## Lists

list의 첫 번째 요소를 examine (evaluate가 아니다...) 한 결과를 통해 나머지 과정을 결정한다.

1. function일 경우 - 나머지 요소를 evaluate한 값들을 인수로 하여 function call로 이어진다.
2. macro일 경우 - 나머지 요소를 evaluate하지 않은 채로 expansion이 이루어진다.
3. special form일 경우 - 그 종류에 따라 다르게 evaluation 된다.
4. 나머지 타입의 경우엔 'void-function' 에러를 발생시킨다.

## Thoughts

전공공부는 안하고 이런 거나 하고 있네;;

## Sources

GNU Emacs Lisp Reference Manual
