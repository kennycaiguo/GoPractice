#### Go镶嵌类型

##### Go镶嵌
```go
package main

import (
    "fmt"
)

type user struct {
    name  string
    email string
}

func (u *user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n",
    u.name,
    u.email)
}

type admin struct {
    user
    level string
    name  string
    email string
}

func main() {
    ad := admin{
        user: user{
            name:  "common user",
            email: "common user@email.com",
        },
        level: "super",
    }
    ad.user.notify()
    //Sending user email to common user<common user@email.com>

    ad.notify()
    //Sending user email to common user<common user@email.com>
}
```