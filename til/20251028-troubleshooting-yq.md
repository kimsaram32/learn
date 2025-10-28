# 20251028 - yq troubleshooting

[GitHub commit](https://github.com/team-xquare/xquare-onpremise-project-gitops-repo/commit/ceef1319510aa5bc627a55d3eebf0bf1fc7fce4e)

## 배경

- GitOps repo에서 config 파일 업데이트를 GitHub Action의 workflow_dispatch 기능으로 처리한다.
- 이때 전달하는 payload에는 수행할 액션과 업데이트 할 subresource를 지정한다.
- workflow는 payload에서의 요구를 바탕으로, manifest 파일을 `yq`로 업데이트하는 방식이다.

## 문제

- apply 액션을 addon subresource에 적용하는 부분이 작동하지 않았다.

## 1. `//` operator의 syntax error 문제

## Problem

로그 확인 결과, 다음과 같은 오류가 발생했음을 알 수 있었다.

```
Error: '//' expects 2 args but there is 1
```

## 원인 분석

처음엔 `add_or_update_sub_resource` 함수에 문제가 있다고 생각했다.  
여기가 더 복잡했기 때문이다.

```bash
add_or_update_sub_resource () {
  local resource_type="$1"
  yq eval "with(.${resource_type}; (first(.name == \"$NAME\") // (. += { \"name\": \"$NAME\" } | first(.name == \"$NAME\"))) *= ${SPEC})" -i $FILE
}
```

그러나 application을 subresource로 하여 apply 액션을 다시 건드려 봤는데,  
이 경우엔 잘 작동하였다.

`add_or_update_sub_resource` 함수는 addon과 application 두 subresource에서 공유된다.  
즉 이 함수엔 문제가 없고, addon 부분에만 사용되는 코드에서 문제가 있다는 뜻이다.  
해당 부분에서 yq를 사용하는 건 이 줄뿐이었다.

```bash
yq eval ".addons //= []" -i $FILE
```

이를 로컬에서 돌려보니, 같은 오류 메시지가 발생했다.  
즉 문법상의 syntax error임을 알 수 있었다.

## Fix

다음과 같이 syntax에 맞도록 수정하였다.

```bash
yq eval ".addons // (.addons = [])" -i $FILE
```

## 2. 새로운 addons 필드 생성 문제

1.을 고치고 나자 발생한 문제이다.

## Problem

다음과 같은 오류 메시지가 발생했다.

```
Error: cannot index array with 'addons' (strconv.ParseInt: parsing "addons": invalid syntax)
```

## 원인 분석 과정

[관련 GitHub Issue](https://github.com/mikefarah/yq/issues/1076)

현재 개체가 array일 때, 객체식으로 필드명을 통해 접근하면 발생하는 오류이다.  
addons 접근은 top level 개체에 대해서 이루어진다.  
즉 top level 개체가 배열이어서 발생하는 오류라는 뜻이었다.

그러나 레포에 올라와 있는 manifest는 명백히 객체였다.

```
applications:
  - ...
addons:
  - ...
```

그럼에도 불구하고 코드 실행 시엔 top level을 array로 인식하였다.

좀 더 분석해보기 위해 로컬에서 정확히 같은 변수를 집어넣고 코드를 돌렸다.
이때는 잘 작동하였다.

이처럼 reproducing이 잘 되지 않아, 원인을 찾는 데 시간이 걸렸다.

## 원인

탐색 끝에 원인을 찾았다. 이는 다음과 같았다.

1. addons 배열이 없을 때, 다음이 실행된다.
   ```bash
   yq eval ".addons // (.addons = [])" -i ...
   ```
2. 나의 의도는 addons 필드가 추가된 전체 개체가 manifest에 저장되는 것이었다.  
   그러나 위 코드는 그렇게 실행되지 않는다.  
   업데이트 후 manifest에는 빈 배열 `[]` 만이 저장된다.
3. 이후 실행되는 yq 코드는 empty array 가 담긴 파일 상에서 작동하게 된다.  
   결국 여기에 객체 접근 `.addons` 를 사용하게 되어 오류가 발생한다.

## 해결

다음과 같은 코드로 수정하여 해결했다.  
yq 실행 후 파일에는 addons 필드의 parent, 즉 top level 개체가 담기도록 수정했다.

```bash
yq eval "(.addons // (.addons = [])) | parent" -i $FILE
```

## Thoughts

## yq의 이해 부족

두 버그 모두 같은 코드를 수정하면서 발생했다.  
이 코드는 전부 스스로 작성한 것이다.  
따라서 근본적인 원인은 yq를 제대로 이해하지 않은 상태에서 코드를 작성했기 때문이라고 볼 수 있다.

## 2번째 버그 원인 분석은 파일이 static 하다고 가정했기에 오래 걸렸다

2번째 버그를 분석할 때, 파일의 형식이 잘못되어서라는 것 까지는 비교적 수월하게 파악했다.  
그러나 그 파일에 대한 문제를 찾는 것부터 막혔다.

- 나는 문제를 job 실행 전 repo 상에 올라와 있는 파일에서만 찾았다.
- 실제 문제는 job 상의 이전 `yq` 실행에서 파일을 잘못 건드리면서 발생했다.

결론적으로 내가 분석 시 놓치고 있었던 부분은 입력 파일의 실행 중 업데이트였다.  
정적인 부분만 신경을 썼고, 동적인 부분은 놓치고 있었다.
