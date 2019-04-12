#### 多态

##### 多态的意义
重载式多态，也叫编译时多态。也就是说这种多态再编译时已经确定好了。重载大家都知道，方法名相同而参数列表不同的一组方法就是重载。在调用这种重载的方法时，通过传入不同的参数最后得到不同的结果。

但是这里是有歧义的，有的人觉得不应该把重载也算作多态。因为很多人对多态的理解是：程序中定义的引用变量所指向的具体类型和通过该引用变量发出的方法调用在编程时并不确定，而是在程序运行期间才确定，这种情况叫做多态。 这个定义中描述的就是我们的第二种多态—重写式多态。并且，重载式多态并不是面向对象编程特有的，而多态却是面向对象三大特性之一

我觉得大家也没有必要在定义上去深究这些，我的理解是：同一个行为具有多个不同表现形式或形态的能力就是多态，所以我认为重载也是一种多态，如果你不同意这种观点，我也接受。

重写式多态，也叫运行时多态。这种多态通过动态绑定（dynamic binding）技术来实现，是指在执行期间判断所引用对象的实际类型，根据其实际的类型调用其相应的方法。也就是说，只有程序运行起来，你才知道调用的是哪个子类的方法。 这种多态通过函数的重写以及向上转型来实现，我们上面代码中的例子就是一个完整的重写式多态。我们接下来讲的所有多态都是重写式多态，因为它才是面向对象编程中真正的多态。

##### 接口实现多态的例子
```go
package main

import (
    "fmt"
)

type person interface {
    sayInfo()
}

type user struct {
    name  string
    email string
}

type admin struct {
    name  string
    email string
}

func (u *user) sayInfo() {
    fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)
}

func (a *admin) sayInfo() {
fmt.Printf("Sending admin email to %s<%s>\n",a.name,a.email)
}

func main() {
    bill := user{"Bill", "bill@email.com"}
    sendInfo(&bill)

    lisa := admin{"Lisa", "lisa@email.com"}
    sendInfo(&lisa)
}

func sendInfo(p person) {
    p.sayInfo()
}
```
可以看到admin和user有着相同的方法`sayInfo`，可以通过接口封装在同一个函数中。这样只有在实际调用时，调用一个函数才会知道会执行哪一个方法。