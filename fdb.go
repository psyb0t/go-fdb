package fdb

var reserved_cnames = []string{"keys"}
var keyno_file = ".keys"
var keyfile_ext = ".fdbk"

type Key struct {
    Index int
    Name string
    Path string
}

type Collection struct {
    Path string
    Keys []string
    IndexAt int
}
