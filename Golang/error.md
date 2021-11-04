#error
#错误

### type error

```go
type error interface {
    Error() string
}
```

* 内建error接口类型是约定用于表示错误信息，[[nil]]值表示无错误。

Builtin Types
