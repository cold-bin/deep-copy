## 🚀 deep-copy

### 💨 背景
很早之前就像写了，最初诞生这个想法的时候还是寒假，主要是项目里有很多需要深拷贝的地方，于是乎今天就有了时间代码也很简洁，思路也比较容易理解。<br>
在go语言里，是基本类型`int`、`string`、`bool`、`array`等类型的变量的赋值都是值拷贝，像`slice`、`map`、`chan`、`func`等引用数据类型，在变量赋值时，都是将该引用赋给变量，也就是说，引用数据类型的赋值，是浅拷贝。<br>
但是我们可能会有这样一种业务情况，需要将引用数据类型进行深拷贝以求不受原来引用数据的影响。<br>
但是在业务里，我们最粗暴简单地实现深拷贝的方式是常常是以下几种:
- 先使用`json`、`gob`等库序列化之后，再反序列化，这样确实可以实现深拷贝，但问题也是也突出的：实现不优雅And性能也更差And不能定制化
- 还有一种是性能要高出前者两个数量级的方法，我们手动深拷贝到基本类型，遇到引用类型，就创建新的类型来填充，这样做显然也有点缺点就是台臃肿以及不优雅了
- 还有一种是，结合前两者的优点：不用写太多的代码+稍微好点的性能，（一般懒人才会这样


### ⚙ 实现
- go语言里基础类型赋值都是深拷贝，那我们只需要递归遍历到基础类型即可。

### 📢 用法
用法很简单<br><br>
先导包
```shell
 go get -u "github.com/cold-bin/deep-copy"
```
<br>

```go
// @author cold bin
// @date 2022/9/5

package main

import (
	"fmt"
	deep_copy "github.com/cold-bin/deep-copy"
)

type Struct struct {
	Int int
	Map map[int]string
}

func main() {
	s := Struct{
		Int: 1,
		Map: map[int]string{1: "1"},
	}
	newS := deep_copy.Copy(s)
	s.Map[1] = "2"
	fmt.Println(newS)
	fmt.Println(s)
}
```

### 📝 之后
初版实现比较简单，后续再考虑实现一些更高级的操作，如：深拷贝非导出字段、深度拷贝一些`no-copy`字段`，还有对深度拷贝做一个定制化，都在后续之中......