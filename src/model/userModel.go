package model

import db "../database"

type User struct{
    Id            uint64    `json:"id"`
    Username      string    `json:"username"`
    Password      string    `json:"password"`
    Name          string    `json:"name"`
    City          string    `json:"city"`
}

type Users []User

func Insert(user User) {
    conn := db.Connect()
    res, err := conn.Prepare("INSERT INTO Employee(name, city) VALUE(?,?)")
    if err != nil{
        panic(err)
    }

    res.Exec(user.Name, user.City)

    defer conn.Close()
}

func FetchAll() Users{
    conn := db.Connect()
    res, err := conn.Query("SELECT * FROM Employee")
    if err != nil{
        panic(err)
    }

    users := Users{}
    for res.Next(){
        var id uint64
        var name string
        var city string

        err = res.Scan(&id, &name, &city)
        if err != nil{
            panic(err)
        }

        user := User{id, "", "", name, city}
        users = append(users, user)
    }

    defer conn.Close()
    return users
}
