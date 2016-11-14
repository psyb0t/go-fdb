package fdb

import (
    "errors"
    "regexp"
    "fmt"
)

func (c *Collection) Set(key string, value string) error {
    if c.KeyExists(key) {
        return errors.New("Key already exists")
    }

    regexpr, err := regexp.Compile(illegal_chars)

    if err != nil {
        return err
    }

    if !regexpr.MatchString(key) {
        return errors.New("Invalid characters in key name")
    }

    err = c.AddKey(key, value)
    fmt.Println(c.Keys[len(c.Keys)-1])

    if err != nil {
        return err
    }

    return nil
}
