package main

import(
    "fmt"
    // 自己包的引入为gopath+其余path
    "go_code/goproject01/chapter03/package1"
)

func main(){
    // fmt.Println(package1.name_package1) 
    // 不能被导入且访问
    fmt.Println(package1.Name_package1)
}