```shell
go build -buildmode=c-shared -o libmock_opr_ngx.so
rm libmock_opr_ngx.h
mv libmock_opr_ngx.so ${OPR_CODE_PATH}/sudiyi_common_opr/sdy/common/util/luaunit/mock/ngx/lib
```
