#### Go方法

##### Go语言中struct的方法

方法就是函数，对象调用的函数就是方法，同样包括值传递和指针传递

##### Go方法接收值
```go
type user struct {
	name  string
	email string
}
//传递值，changeEmail方法没有用
func (u user) print() {
	fmt.Println(u.email)
}
func (u user) changeEmail(email string) {
	u.email = email
}
func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.print()
	//bill@email.com
	bill.changeEmail("new address")
	bill.print()
	//bill@email.com
}
```
两个函数都是传递user的值的副本，因此修改没用。

##### Go方法接收指针
```go
type user struct {
	name string
    email string
}

// 函数接收user类型的指针
func (u *user) printPointer() {
	fmt.Println(u.email)
}
func (u *user) changeEmailPointer(email string) {
	u.email = email
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	bill.printPointer() 
	//bill@email.com
	bill.changeEmailPointer("new address") 
	bill.printPointer()
	//new address
}
```
注意这里传递的就是变量`bill`的指针了。
细心的同学可能发现`bill.printPointer()`这个调用有问题。因为`printPointer`接收的是指针。
其实这样写更准确：`(&bill).printPointer() 
`这里是go编译器帮我们内部处理的，这样能更加方便我们书写方法。

##### &user的写法
大家可能猜到，同理还有另外一种写法,上述的2份代码只改动一处：
`bill := &user{"Bill", "bill@email.com"}`

上述的2份代码运行的结果都不会发生变化，传入指针或者传入值，都是由函数方法决定的。

