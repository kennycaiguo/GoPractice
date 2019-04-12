let obj={
    name:"old name"
}

let changename= function(object){
    object.name="new name"
    object=null //是没有用的
}

console.log(obj) //{ name: 'old name' }
changename(obj)
console.log(obj) //{ name: 'new name' }