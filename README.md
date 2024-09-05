# go-pointer

## INSTALL

```go
go get -u github.com/hidori/go-pointer@latest
```

## USAGE

```go
 // returns pointer of value 1
 fmt.Println(pointer.Of(1))

 val := 1

 // returns 1
 fmt.Println(pointer.ValueOrDefault(&val, 3))

 // returns 3
 fmt.Println(pointer.ValueOrDefault((*int)(nil), 3))

 // returns 1
 fmt.Println(pointer.ValueOrEmpty(&val))

 // returns 0
 fmt.Println(pointer.ValueOrEmpty((*int)(nil)))
```
