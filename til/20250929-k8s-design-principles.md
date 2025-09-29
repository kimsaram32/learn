# 20250929 - Kubernetes Design Principles 보고 느낀 것 정리

<https://youtu.be/ZuIQurh_kDk> 보고 느낀 것 정리

## K8s는 내부 / 외부 API를 구분하지 않음

k8s는 내부 API와 외부 API를 구분하지 않는다.
내부 시스템끼리 소통하는 API가 외부에 투명하게 공개되어 있다.

scheduler와 kubelet에서 watch하는 Pod는 내가 kubectl을 통해 보는 Pod와 근본적으로 같다.

### 확장성으로 이어짐

이러한 내부 / 외부 구분 없는 API는 k8s의 확장성으로 이어진다.

e.g. 새로운 object가 필요할 경우
기본적으로 desired state의 definition과 (CRD) 그것을 current state와 sync하는 작업을 해주면 된다 (물론 구현은 쉽지 않다;).
즉 인프라와 사용자의 요구대로 확장할 수 있다.

e.g. 내부와 외부를 구분하지 않으니, 기본 컴포넌트 (scheduler, controller...) 가 마음에 들지 않으면 갈아끼우면 된다.

### Emacs와 닮은 점

Emacs도 이것과 비슷한 특성을 가졌다. 그래서 확장성으로 유명하다

결국 무한히 확장 가능한 시스템을 만들려면 API의 추상화 수준이 낮거나 아예 없는 것이 좋은 듯
