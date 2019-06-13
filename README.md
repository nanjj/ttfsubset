# ttfsubset

Install:

```
go get github.com/nanjj/ttfsubset
```

Output ttf subset to file:

```
ttfsubset NotoSansCJKsc-Regular.ttf -o mcuts.ttf -r "姓名分数日期时间："
```

Output ttf subset in golang bytes format to stdout:

```
ttfsubset NotoSansCJKsc-Regular.ttf -r "姓名分数日期时间："
```
