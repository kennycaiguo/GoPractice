#### Go函数参数传递和引用类型

##### Go语言的类型

目前学到的

基本类型：
string int bool 

引用类型：
切片,映射(map),接口(interface),函数(func)


##### 函数参数的传递

go中函数参数分两种，值传递和指针传递
值传递：就是传递的变量的拷贝
指针传递：就是传递的变量的指针

已知：
```go
type person struct{
    name string
    age int
}
```
值传递：
```go
func(p person){
    ...
}
```
指针传递：
```go
func(p *person){
    ...
}
```

###### struct作为函数参数的值传递
看一个关于函数参数值传递的例子：
```go
type person struct{
    name string
    age int
}

func main(){
    jim := person{"jim",10} //变量jim为类型person
    fmt.Println(jim) //{jim 10}
    changename(jim)  //传进函数的是值的拷贝
    fmt.Println(jim) //{jim 10}
}

func changename(p person){
    p.name = "another" //这里是对拷贝赋值，没有效果
}
```
这里的changename接受的参数是值传递，就是说形参变量`p`是变量`jim`的副本。

###### struct作为函数参数的指针传递

我们换一个写法试试，改用指针传递，就能修改变量：
```go
type person struct{
    name string
    age int
}

func main(){
    jim := person{"jim",10}
    fmt.Println(jim) //{jim 10}
    changename(&jim)  //传进函数的是变量jim的指针
    fmt.Println(jim) //{another 10}
}

func changename(p *person){
    p.name = "another" //这里是指针指向的变量赋值，所以有效果
    /*
    细心的同学可能发现p是一个指针变量，为什么能直接加.name?
    此处是go语言编译器自动转换了写法
    准确来写应该是：
    (*p).name = "another"
    */
}
```

写到这里大家可能会觉得struct应该是一个基本类型，可能有的同学不太明白，函数参数的值传递，指针传递和变量的基本类型，引用类型有什么关系。

我再举一个例子，依然是值传递和指针传递，但是这次参数选一个引用类型：切片

###### 切片作为函数参数的值传递
```go
func main(){
    persons := []string{"jim","tom","panda"}
    fmt.Println(persons) //[jim tom panda]
    changeperson(persons)  //传进函数的是切片的拷贝
    fmt.Println(persons) //[another tom panda]
}

func changeperson(p []string){
    p[0] = "another" //这里是对拷贝赋值，但仍有效果
}
```
为什么切片作为值的拷贝传递，修改拷贝的值仍然有效果？
因为之前讲过，切片是引用类型，切片本身可以理解为包含3个字段的struct：指向底层数组的指针，长度，容量。因此即使在函数`changeperson`中变量`p`是值的拷贝，切片作为引用类型，变量`p`仍然包含了指向底层数组的指针,因此能够修改底层数组成功。

###### js中一个引用类型的例子
我再举一个javascript的例子进行比较一下(不熟悉js的同学可忽略)

```javascript
let obj={
    name:"old name"
}

let changename= function(object){
    object.name="new name"
}

console.log(obj) //{ name: 'old name' }
changename(obj)
console.log(obj) //{ name: 'new name' }
```
js中`Object`是引用类型，这里obj传递的实际上仍然是值的复制(js中没有显式的指针的概念)，但是在这个副本中仍然包含了引用类型这个对象在内存中的指针，因此可以修改内存中的值。

所以在js中这么写是无效的：
```javascript
let obj={
    name:"old name"
}

let changename= function(object){
    object.name="new name"
    object=null //是没有用的，object只是一个值的复制
}

console.log(obj) //{ name: 'old name' }
changename(obj)
console.log(obj) //{ name: 'new name' }
```

###### 切片作为指针传递

```go
func main(){
    persons := []string{"jim","tom","panda"}
    fmt.Println(persons) //[jim tom panda]
    changeperson(&persons)  //传进函数的是变量的指针
    fmt.Println(persons) //[another tom panda]
}

func changeperson(p *[]string){
    (*p)[0] = "another" //这里直接取到指向指针地址对应的变量，并赋值
}
```
可见对于引用变量，起码对于切片而言，传递值和传递指针，变化并不是那么明显。