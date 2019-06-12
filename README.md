# ttfsubset
Install:

```
go get github.com/nanjj/ttfsubset
```

Usage:
- Subset a smaller ttf file:
    ```
    ttfsubset NotoSansCJKsc-Regular.ttf -o mcuts.ttf -r "姓名分数日期时间："
    ```
- Print golang bytes to std output:
    ```
    ttfsubset NotoSansCJKsc-Regular.ttf -r "姓名分数日期时间："
    ```
