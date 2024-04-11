# go-pointer

## INSTALL

```go
go get github.com/hidori/go-pointer@latest
```

## USAGE

## Of()

```go
fmt.Println(pointer.Of(1))
// returns pointer of value 1
```

## ValueOrDefault()

```go
val := 1
fmt.Println(pointer.ValueOrDefault(&val), 3)
// 1

fmt.Println(pointer.ValueOrDefault((int*)(nil)), 3)
// 3
```

## ValueOrEmpty()

```go
val := 1
fmt.Println(pointer.ValueOrEmpty(&val))
// 1

fmt.Println(pointer.ValueOrEmpty((int*)(nil)))
// 0
```
