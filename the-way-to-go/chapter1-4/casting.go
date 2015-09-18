package main

import (
  "fmt"
  "math"
)

func main() {
  var n int16 = 34
  var m int32
  // compiler error: cannot use n (type int16) as type int32 in assignment
  // m = n
  m = int32(n)

  fmt.Println(Unit8FromInt(123))
  fmt.Println(IntFromFloat64(21.2213343443))

  fmt.Printf("32 bit int is: %d\n", m)
  fmt.Printf("16 bit int is: %d\n", n)
}

func Unit8FromInt(n int) (uint8, error) {
  if 0 <= n && n <= math.MaxUint8 {// conversion is safe
    return uint8(n), nil
  }
  return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
  if math.MinInt32 <= x && x <= math.MaxInt32 {// x lies in the integer range
    whole, fraction := math.Modf(x)
    if fraction >= 0.5 {
      whole++
    }
    return int(whole)
  }
  panic(fmt.Sprintf("%g is out of the int32 range", x))//数字超出要转换的类型的取值范围就会引发panic
}

//位左移常见实现存储单位的用例
type ByteSize float64
const (
  _ = iota // 通过复制给空白标识符来忽略初始值
  Kb ByteSize = 1<<(10*iota)
  Mb
  Gb
  Tb
  Pb
  Eb
  Zb
  Yb
)

//在通讯中使用位左移表示标识的用例
type BitFlag int
const (
  Active BitFlag = 1 << iota // 1 << 0 == 1
  Send // 1 << 1 == 2
  Receive // 1 << 2 == 4
)

flag := Active | Send // == 3
