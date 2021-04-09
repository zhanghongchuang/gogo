package src

import "fmt"

type Driver struct {

}

func (This *Driver) Connection()  {
	fmt.Print("connection")
}