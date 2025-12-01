# 20251201 - Linux file-based API flow

일반적으로 Linux에서 file 기반 API를 사용할 때의 플로우이다.
전체 흐름만 정리하였다.

1. C library functions (wrapper functions)
   - system call에 대한 wrapper function 들이다.
   - 상위에게 system call 호출을 추상화해서 제공한다.
   - 이는 application이 system call 호출 구현이 아닌
     이들의 동작 방식으로 생각하게 해 준다.
   
2. system call
   - userland에서 file 기반 API를 사용하기 위한 entrypoint로 사용한다.
   - VFS 상에서의 API를 사용한다.
   
3. VFS
   - 여러 파일 시스템을 통합하여 하나의 logical file system으로 제공한다.
   
4. logical file system
   - underlying file system은 device 위에서의 file system의 역할을 수행한다.
   - 또한 VFS 상에서의 연산을 구현하여 제공한다.

5. block devices
   - driver 위에서 작동하여 device들을 추상화한다.
     block 단위의 random data access를 제공한다.

6. storage drivers
   - 물리적인 하드웨어 장치에 대한 interface라고 생각함...

리눅스의 다른 컴포넌트와 마찬가지로...
abstraction layer를 통한 계층적 구조를 이룬다.

## Sources

- Linux Kernel Tree
