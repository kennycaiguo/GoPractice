let obj={
    name:"old name"
}

let changename= function(object){
    object.name="new name"
}

console.log(obj) //{ name: 'old name' }
changename(obj)
console.log(obj) //{ name: 'new name' }