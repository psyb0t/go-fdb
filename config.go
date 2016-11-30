package main

import (
    "os"
    "log"
    "io/ioutil"
    "path/filepath"
    "encoding/json"
)

var config_file = "/etc/fdb/config"
var databases_path = "/etc/fdb/databases/"
var log_path = "/var/log/fdb/fdb.log"
var listen_host = "127.0.0.1"
var listen_port = 60676

type Config struct {
    DatabasesPath string
    LogPath string
    ListenHost string
    ListenPort int
}

var config = &Config {
    DatabasesPath: databases_path,
    LogPath: log_path,
    ListenHost: listen_host,
    ListenPort: listen_port,
}

func InitConfig() {
    if _, err := os.Stat(config_file); os.IsNotExist(err) {
        config_dir := filepath.Dir(config_file)

        if _, err := os.Stat(config_dir); os.IsNotExist(err) {
            if err := os.MkdirAll(config_dir, 0644); err != nil {
                log.Fatal("Could not create config dir")
            }
        }

        json_config, err := json.MarshalIndent(config, "", "  ")

        if err != nil {
            log.Fatal("Could not serialize config data")
        }

        err = ioutil.WriteFile(config_file, json_config, 0644)

        if err != nil {
            log.Fatal("Could not write default config file")
        }

        return
    }

    json_data, err := ioutil.ReadFile(config_file)

    if err != nil {
        log.Fatal("Could not read from config file")
    }

    err = json.Unmarshal(json_data, config)

    if err != nil {
        log.Fatal("Could not unserialize config data")
    }

    return
}
