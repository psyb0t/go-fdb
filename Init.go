package fdb

import (
    "log"
    "regexp"
    "io/ioutil"
)

func (c *Collection) Init() error {
    log.Print("Initializing FDB RAM stuff...")

    index_files, err := ioutil.ReadDir(c.Path)

    if err != nil {
        return err
    }

    re := regexp.MustCompile("\\d{19}~(.*?)" + key_val_file_ext)
    for _, file := range index_files {
        keyfile := file.Name()

        key := re.ReplaceAllString(keyfile, "$1")

        c.KeyInMem(key, keyfile)
    }

    log.Print("FDB RAM stuff initialized!")

    return nil
}
