# FileDB - Even a cheap ass VPS can hold and serve 3million DB entries

## Usage

```go
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

## Testing alongside redis

```
I have a website which contains videos and video categories(see [go-content-master](https://github.com/psyb0t/go-content-master/))

DB has
213637 Videos
5350 Categories

Init time to load required stuffz:
FDB: 2sec
Redis: 11sec

RAM:
FDB: 677MB RAM
Redis: 2227MB RAM

Disk space:
FDB: 897MB
Redis: 1018MB

ApacheBuffer percentage of the requests served within a certain time (ms)
Launched with 1000 concurrent requests and 20000 total request number
--------------------------------------
FDB          | REDIS
             |
 50%     56  |  50%     41
 66%     86  |  66%     63
 75%    223  |  75%     89
 80%    262  |  80%    224
 90%    715  |  90%    350
 95%   1058  |  95%   1022
 98%   1395  |  98%   1257
 99%   2331  |  99%   1497
100%  11565  | 100%   5047
```

I think that with a little bit of caching and more goroutines the latency numbers will go down.

No problems on single element get / set / update.

Ranges are messier.
