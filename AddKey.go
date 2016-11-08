package fdb

import (
    "os"
    "strconv"
    "io/ioutil"

    "github.com/psyb0t/go-genutils"
)

func (c *Collection) AddKey(key string, value string) error {
    comp_val, err := genutils.Compress([]byte(value))

    if err != nil {
        return err
    }

    err = ioutil.WriteFile(c.KeyValuePath(key), comp_val, 0644)

    if err != nil {
        return err
    }

    last_index, err := c.LastIndex()

    if err != nil {
        return err
    }

    curr_index := last_index + 1

    index_to_key_path := c.Index2KeyPath(curr_index)
    key_to_index_path := c.Key2IndexPath(key)

    _, err = os.Stat(index_to_key_path)

    if err != nil && !os.IsNotExist(err) {
        return err
    }

    _, err = os.Stat(key_to_index_path)

    if err != nil && !os.IsNotExist(err) {
        return err
    }

    err = ioutil.WriteFile(index_to_key_path, []byte(key), 0644)

    if err != nil {
        return err
    }

    err = ioutil.WriteFile(key_to_index_path, []byte(
        string(curr_index)), 0644)

    if err != nil {
        return err
    }

    c.Keys = append(c.Keys, key)
    c.Indexes = append(c.Indexes, curr_index)
    c.KeyToIndex[&key] = &curr_index
    c.IndexToKey[&curr_index] = &key


    last_index_fpath := c.Path + "/" + last_index_file

    err = ioutil.WriteFile(last_index_fpath, []byte(
        strconv.Itoa(curr_index)), 0644)

    if err != nil {
        return err
    }

    return nil
}
