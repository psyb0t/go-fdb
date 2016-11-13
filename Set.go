package fdb

import (
    "fmt"
    "errors"
    "regexp"
)

func (c *Collection) Set(key string, value string) error {
    fmt.Println("Checking", key)

    if c.KeyExists(key) {
        fmt.Println("Key already exists", c.Key2File[key])
        fmt.Println("Key already exists", key)
        return errors.New("Key already exists")
    }

    regexpr, err := regexp.Compile("^[a-zA-Z0-9-_:]+$")

    if err != nil {
        fmt.Println("Regex fail", key)
        return err
    }

    if !regexpr.MatchString(key) {
        fmt.Println("Match fail", key)
        return errors.New("Invalid characters in key name")
    }

    fmt.Println("Writing", key)

    err = c.AddKey(key, value)

    if err != nil {
        return err
    }

    return nil
}
