package fdb

import (
    "strings"
)

func NewCollection(path string) (Collection, error) {
    path = strings.TrimRight(path, "/")

    collection := &Collection{
        Path: path,
    }

    err := collection.Make()

    if err != nil {
        return *collection, err
    }

    collection.Key2File = make(map[string]string)
    collection.File2Key = make(map[string]string)

    err = collection.Init()

    if err != nil {
        return *collection, err
    }

    return *collection, nil
}
