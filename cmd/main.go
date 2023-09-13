package main

import {
    "fmt"
    "errors"
}

func example() error {
    return fmt.Errorf("here is an error with a string");
}

func otherExample() error {
    return errors.New("here is an error, but with errors");
}

func exampleNoError() error {
    return nil;
}

func exampleWithData(should bool) (*Thing, error) {
    if should {
        return &Thing{}, nil;
    }
    return nil, fmt.Errorf("nice try, guy");
}

func main() {
    err := example();
    if err != nil {
        fmt.Println(err);
    }

    _, err = exampleWithData(true);
    if err != nil {
        fmt.Println(err);
    }
}
