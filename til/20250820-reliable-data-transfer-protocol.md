# 20250820 - Reliable Data Transfer Protocol (recap)

reliable data transfer protocol의 구현에 있어서 어려운 점은 하위 레이어가 unreliable한 경우가 대부분이라는 것이다.
그 위에서 데이터의 오류 검출이나 전송 누락 등을 처리해야 하므로, 여러 복잡한 기능들이 들어가게 된다.

e.g. TCP가 의존하는 하위 레이어는 IP이다. IP는 best-effort 형식으로만 패킷을 전송하므로 본질적으로 unreliable한 프로토콜이다.
TCP는 이 위에서 호스트들 사이의 reliable한 연결을 제공해주어야 한다.

## Sources

- Computer Networking - a Top-Down Approach
