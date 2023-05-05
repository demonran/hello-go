package main

import (
	"fmt"
	"hello-go/common/db"
	"log"
	"os/exec"
	"time"
)

type BaseModel interface {
	getId() int
	getName() string
}

type Greet func(name string) string

func (Greet) say() string {
	return "say"
}

type User struct {
	Name string `json:"name" yaml:"name"`
}

func (user User) getId() int {
	return 0
}

func english(name string) string {
	return name
}

func (user *User) setName(name string) {

	fmt.Printf("123%v", user)
	user.Name = name
}

func main() {
	user := User{"zhangsan"}
	fmt.Printf(db.DB)
	//user.setName("lisi")
	////user := new(User)
	//field, _ := reflect.TypeOf(user).FieldByName("Name")
	//fmt.Println(field.Tag.Get("yaml"))

	//reflectValue := reflect.ValueOf(user)
	//fmt.Println(reflectValue.CanSet())
	////fmt.Println(reflectValue.Elem().CanSet())
	//
	//var str = "str"
	//ptr := &str
	//fmt.Println(ptr)
	//fmt.Println(user)
	//fmt.Println(new(string))

	fmt.Printf("%#v", user)

	pipeline := make(chan int)
	go func() {
		for {
			fmt.Println("准备发送数据")
			pipeline <- 100
		}

	}()

	go func() {
		num := <-pipeline
		fmt.Printf("接受数据: %d", num)

	}()
	time.Sleep(time.Second)

	cmd := exec.Command("ls", "-al", "/")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	fmt.Printf("out: %s", string(out))
}

func init() {
	db.DB = "1234"
}
