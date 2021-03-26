package main

import (
	"fmt"
	"io/ioutil"
)

func init() {
	bannerPaths := []string{"./banner.txt", "../doc/banner.txt", "./doc/banner.txt", "./misc/doc/banner.txt"}
	for _, v := range bannerPaths {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
			break
		}
	}
}

func main() {

}
