#builtin_types

`builtin 包为Go的预声明标识符提供了文档`

```go
$ go doc builtin
package builtin // import "builtin"

Package builtin provides documentation for Go's predeclared identifiers. The
items documented here are not actually in package builtin but their
descriptions here allow godoc to present documentation for the language's
special identifiers.

const true = 0 == 0 ...
const iota = 0
func close(c chan<- Type)
func delete(m map[Type]Type1, key Type)
func panic(v interface{})
func print(args ...Type)
func println(args ...Type)
func recover() interface{}
func cap(v Type) int
func copy(dst, src []Type) int
func len(v Type) int
type ComplexType complex64
    func complex(r, i FloatType) ComplexType
type FloatType float32
    func imag(c ComplexType) FloatType
    func real(c ComplexType) FloatType
type IntegerType int
type Type int
    var nil Type
    func append(slice []Type, elems ...Type) []Type
    func make(t Type, size ...IntegerType) Type
    func new(Type) *Type
type Type1 int
type bool bool
type byte = uint8
type complex128 complex128
type complex64 complex64
type error interface{ ... }
type float32 float32
type float64 float64
type int int
    func cap(v Type) int
    func copy(dst, src []Type) int
    func len(v Type) int
type int16 int16
type int32 int32
type int64 int64
type int8 int8
type rune = int32
type string string
type uint uint
type uint16 uint16
type uint32 uint32
type uint64 uint64
type uint8 uint8
type uintptr uintptr

```
## [[int]]
## [[int8]]
## [[int16]]
## [[int32]]
## [[int64]]
## [[uint]]
## [[uint8]]
## [[uint16]]
## [[uint32]]
## [[uint64]]
## [[uintptr]]
## [[float32]]
## [[float64]]
## [[complex128]]
## [[complex64]]
## [[bool]]
## [[byte]]
## [[rune]]
## [[string]]
## [[error]]
## [[Type]]