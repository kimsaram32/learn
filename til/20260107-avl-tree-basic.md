# 20260107 - AVL tree 기본적인 설명

AVL tree의 연산적인 부분 제외하고 기본적인 것만 high-level적으로 설명해보기

## 기본 설명

AVL tree는 self-balancing binary search tree의 한 종류이다.

균형을 유지하는 방식은 left subtree와 right subtree의 높이 차를
2 미만으로 유지하는 것이다 (= balance factor를 [-1, 1] range 사이로 유지).

AVL tree에 대한 연산은 일반 BST와 같이 이루어지되,
modification이 일어나는 연산의 경우 균형이 깨지면 rotation을 수행한다.
이때 "균형이 깨진 방식"에 따라 적용해야 하는 rotation도 달라진다.

## Balance factor

balance factor는 AVL tree에서 균형 유지의 기준점이 되는 값이다.
AVL tree는 모든 node의 BF가 [-1, 1] range에 속하도록 하여 균형을 유지한다.
(todo: 아직 high-level적인 이해만 하고 있어서 "균형을 유지한다" 의 수학적 의미를 정확히 파악하지 못함).

node $N$의 balance factor를 $\mathrm{BF}(N)$이라고 할 때,
$\mathrm{BF}(N) = \mathrm{Height}(\mathrm{LeftSubtree}(N)) - \mathrm{Height}(\mathrm{RightSubTree}(N))$ 이다.

balance factor에 따라 따로 붙는 말이 있다...
- $\mathrm{BF}(N) > 0$이면 $N$은 left-heavy
- $\mathrm{BF}(N) = 0$이면 $N$은 balanced
- $\mathrm{BF}(N) < 0$이면 $N$은 right-heavy

## insertion and deletion

AVL tree에서의 insertion과 deletion operation 연산은 공통적으로 다음과 같이 진행한다.

1. BST와 같은 방식으로 insertion/deletion을 진행한다.
   
2. 삽입된/삭제된 요소의 모든 ancestor를 가장 가까운 것부터 순회하여 (tracing),
   이것의 balance factor가 [-1, 1] range를 벗어났는지 (= "깨졌는지") 확인한다.
   
3. balance factor가 깨진 첫 번째 ancestor에 대하여 rebalancing을 수행한다.

ancestor만 확인하는 이유는 이것들만이 balance factor가 깨졌을 가능성을 갖기 때문이다.
다른 node들은 깨졌을 가능성이 없으므로 확인하지 않는다.

## Sources

- <https://en.wikipedia.org/wiki/AVL_tree>
- 나중에 책 읽으면서 다시 정리하기...
