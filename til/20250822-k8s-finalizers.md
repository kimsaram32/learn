# 20250822 - K8s Finalizers

finailzer는 object가 삭제되기 전 수행해야 할 작업을 표시하는 tag이다.
그 자체로는 아무 기능이 없으며, 연관된 controller가 작업을 핸들링한다.

## 삭제 프로세스

API server가 object를 삭제하라는 명령을 받아도 바로 삭제를 진행하지 않는다.
우선 삭제 요청이 들어왔다는 것을 알리기 위해 `.metadata.deletionTimestamp`를 추가하고,
finalizer가 비워질 때까지 기다린다.

controller는 `.metadata.deletionTimestamp`를 감지하여 관련 액션을 수행하고,
그 액션에 해당하는 finalizer를 지운다.

## Sources

- <https://kubernetes.io/blog/2021/05/14/using-finalizers-to-control-deletion/>
