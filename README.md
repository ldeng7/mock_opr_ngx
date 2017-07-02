```shell
go build -buildmode=c-shared -o libtrochilus.so
rm libtrochilus.h
mv libtrochilus.so ${OPRPATH}/trochilus_common/lib
```
