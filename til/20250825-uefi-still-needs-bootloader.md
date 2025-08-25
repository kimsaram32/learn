# 20250825 - Linux - UEFI 시스템에서도 bootloader 사용이 필요한 이유

BIOS와 다르게 UEFI에선 bootloader를 사용하지 않고 자체적으로 커널을 불러올 수 있다.
그럼에도 불구하고 여전히 bootloader를 사용하는 것이 좋다.

기능적으로 봤을 때 UEFI에서 커널을 직접 불러오는 것은 bootloader를 사용하는 것보다 부족한 부분이 많다.
- EFI partition을 자주 바꾸는 것은 불안정할 수 있다. (BIOS 같은 경우 역사적으로 그랬다)
- kernel parameter 설정 등이 제한될 수 있다.

또한 UEFI가 커널까지 불러오게 된다면 너무 많은 책임을 지게 된다고 생각한다.
커널을 불러온다는 기능은 bootloader에서 추상화해서 담당할 필요가 있다.

## Sources

<https://utcc.utoronto.ca/~cks/space/blog/linux/WhyBootloaderOnUEFI>
