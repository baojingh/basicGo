
```console
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative  \
helloworld/helloworld.proto
```


```dockerfile
protoc --go_out=plugins=grpc:. *.proto
```

