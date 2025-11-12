# 20251112 - 실습: fork()와 exec()

fork()와 exec()의 기본 사용법

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>

int main(void)
{
pid_t pid;

pid = fork(); // (1)

switch (pid)
{
  case -1: // (2)
    perror("fork() failed");
    exit(EXIT_FAILURE);
    break;
  case 0: // (3)
    printf("hello from child %d\n", getpid());
    execl("/usr/bin/echo", "/usr/bin/echo", "Hello", "World", (char*) NULL); // (5)
    exit(EXIT_FAILURE);
    break;
  default: // (4)
    printf("parent %d\n", getpid());
    wait(NULL); // (6)
    printf("exiting %d\n", getpid());
    exit(EXIT_SUCCESS);
    break;
}
}
```

1) fork() 이후 자식 프로세스는 현재 프로그램을 그대로 복제받아 실행된다.
   즉 fork() 호출 이후 코드는 두 프로세스에서 동일하게 실행된다.
   때문에 fork()는 부모 프로세스와 자식 프로세스별로 다른 값을 반환한다.
   이를 통해 둘을 구분해줄 수 있다.
2) 프로세스 생성이 실패했을 경우 -1이 리턴된다. 자식 프로세스는 생성되지 않았으므로 동작이 없고요...
3) 자식 프로세스에선 0을 리턴한다.
4) 부모 프로세스에선 생성된 프로세스의 pid를 리턴한다.
5) execl()은 exec* family의 한 종류이다.
   여기선 그냥 실행된 프로그램 + 인수 목록을 써준다고 생각하자
   - `/usr/bin/echo` 프로그램을 실행하는데,
     인수 목록을 알맞게 넣어주어야 한다.
   - 컨벤션 상 첫 번째 인수는 프로그램의 이름이다.
     또한 마지막 인수는 null pointer가 되어야 한다.
6) wait()을 통해 자식 프로세스가 종료될 때까지 기다린다.

## Sources

The Linux Programming Interface 24.2
