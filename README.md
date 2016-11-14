# FileDB - Even a cheap ass VPS can hold and serve 3million DB entries

## Usage

```golang
package main

import (
    "fmt"
    "log"

    "github.com/psyb0t/go-fdb"
)

func err_handl(err error) {
    if err != nil {
        log.Fatal("DAFUQ? ", err)
    }
}

func main() {
    var err error

    collection_path := "/home/spammer/ham/eggs/THE_SPAM_DOCS"
    collection, err := fdb.NewCollection(collection_path)
    err_handl(err)

    err = collection.Set("SpamKey1", "SPAM1")
    err_handl(err)

    fmt.Println(collection.Keys) // [SpamKey1]

    key_1_val, err := collection.Get("SpamKey1")
    err_handl(err)

    fmt.Println(key_1_val) // SPAM1

    err = collection.Update("SpamKey1", "NOSPAM")
    err_handl(err)

    key_1_val, err = collection.Get("SpamKey1")
    err_handl(err)

    fmt.Println(key_1_val) // NOSPAM

    collection.Set("SpamKey2", "SPAM2")
    collection.Set("SpamKey3", "SPAM3")

    range_get, err := collection.GetRange(0, 3)
    err_handl(err)

    fmt.Println(range_get) // [NOSPAM SPAM2 SPAM3]

    range_get, err = collection.GetReverseRange(0, 3)
    err_handl(err)

    fmt.Println(range_get) // [SPAM3 SPAM2 NOSPAM]

    range_get, err = collection.GetFilteredReverseRange(
        0, 3, "Spe[a-zA-Z]+")
    err_handl(err)

    fmt.Println(range_get) // []

    collection.Set("SpemKey1", "Intentionally written like this")

    range_get, err = collection.GetFilteredReverseRange(
        0, 3, "Spe[a-zA-Z]+")
    err_handl(err)

    fmt.Println(range_get) // [Intentionally written like this]

    fmt.Println(collection.Keys) // [SpamKey1 SpamKey2 SpamKey3 SpemKey1]
}
```
