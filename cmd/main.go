package main

import "fmt"

func returnsError(value int) error {
    return fmt.Errorf("This is an error with value %v", value)
}

type Foo struct { }

func (f *Foo) thisIsOnFoo() error {
    return fmt.Errorf("This is an error from Foo")
}


func main(){
    err := returnsError(5)
    foo := Foo{}

    foo.thisIsOnFoo();
}
