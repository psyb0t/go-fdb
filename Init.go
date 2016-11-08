package fdb

import (
    "strconv"
    "io/ioutil"
)

func (c *Collection) Init() error {
    index_files, err := ioutil.ReadDir(c.Path + "/" + index_to_key_dir)

    if err != nil {
        return err
    }

    c.KeyToIndex = make(map[*string]*int)
    c.IndexToKey = make(map[*int]*string)

    for _, file := range index_files {
        fname := file.Name()

        index, err := strconv.Atoi(fname)

        if err != nil {
            return err
        }

        key, err := c.KeyByIndex(index)

        if err != nil {
            return err
        }

        c.Indexes = append(c.Indexes, index)
        c.Keys = append(c.Keys, key)

        c.KeyToIndex[&key] = &index
        c.IndexToKey[&index] = &key
    }

    return nil
}
