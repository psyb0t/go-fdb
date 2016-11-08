package fdb

var index_to_key_dir = ".i2k"
var key_to_index_dir = ".k2i"
var last_index_file = ".index_at"
var key_val_file_ext = ".fdbk"

type Collection struct {
    Path string
    Keys []string
    Indexes []int
    KeyToIndex map[*string]*int
    IndexToKey map[*int]*string
}
