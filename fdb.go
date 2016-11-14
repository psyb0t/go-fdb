package fdb

var illegal_chars = "^[a-zA-Z0-9-_:]+$"
var key_val_file_ext = ".fdbk"

type Collection struct {
    Path string
    Keys []string
    KeyFiles []string
    File2Key map[string]string
    Key2File map[string]string
}
