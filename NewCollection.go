package fdb

import (
    "strings"
    "path/filepath"
)

func NewCollection(path string) (*Collection, error) {
    path = strings.TrimRight(path, "/")

    collection := &Collection{
        Path: path,
    }

    err := collection.Make()

    if err != nil {
        return nil, err
    }

    files, err := filepath.Glob(path + "/*" + keyfile_ext)

    if err != nil {
        return nil, err
    }

    for _, fpath := range files {
        _, fname := filepath.Split(fpath)

        fname = strings.TrimRight(fname, keyfile_ext)
        //split_fname := strings.Split(fname, "~")
        //
        //index, err := strconv.Atoi(split_fname[0])
        //
        //if err != nil {
        //    continue
        //}

        //keyname := split_fname[1]

        collection.Keys = append(collection.Keys, fname)
    }

    collection.UpdateIndexAt(len(collection.Keys))

    return collection, nil
}
