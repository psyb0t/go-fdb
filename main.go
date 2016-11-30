package main

import (
    "os"
    "fmt"
    "strings"
    "log"
    "io/ioutil"

    "github.com/valyala/fasthttp"
    "github.com/psyb0t/go-fdb/fdb"
)

var databases = map[string]map[string]*fdb.Collection{}

func init() {
    InitConfig()
    InitDatabases()
}

func InitDatabases() {
    if _, err := os.Stat(config.DatabasesPath); os.IsNotExist(err) {
        if err := os.MkdirAll(config.DatabasesPath, 0644); err != nil {
            log.Fatal("Could not create databases dir")
        }
    }

    dbdirs, _ := ioutil.ReadDir(config.DatabasesPath)
    for _, dbdir := range(dbdirs) {
        coldirs, _ := ioutil.ReadDir(config.DatabasesPath + dbdir.Name())
        for _, coldir := range(coldirs) {
            collection, err := fdb.NewCollection(coldir.Name())

            if err != nil {
                continue
            }

            databases[dbdir.Name()] = make(map[string]*fdb.Collection)
            databases[dbdir.Name()][coldir.Name()] = collection
        }
    }
}

func DoFDB(ctx *fasthttp.RequestCtx, params []string) {
    //if len(params) > 0 {
    //    dbname := params[0]
    //}
    //
    //if len(params) > 1 {
    //    collection := params[1]
    //}
    //
    //if len(params) > 2 {
    //    action := params[2]
    //}
    //
    //if len(params) > 3 {
    //    action_params := params[3]
    //}
}

func Perform(ctx *fasthttp.RequestCtx) {
    fmt.Println(databases)

    log.Print(fmt.Sprintf("%s [%s] %s", strings.Split(
        ctx.RemoteAddr().String(), ":")[0],
        string(ctx.Method()), string(ctx.Path())))

    ctx.SetContentType("application/json")

    params := strings.Split(strings.ToLower(strings.Trim(
        string(ctx.Path()), "/")), "/")

    DoFDB(ctx, params)
}

func main() {
    if err := fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", config.ListenHost,
      config.ListenPort), Perform); err != nil {
        log.Fatalf("Server error: %s", err)
    }
}
