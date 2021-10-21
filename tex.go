package main

import ("fmt"; "os")

var VERSION = 0;

func file_exists(path string) bool {
    result := false;
    if _, er := os.Stat(path); er == nil {
        result = true;
    }
    return result;
}

func main() {
    fmt.Printf("TEXporter v%v\n", VERSION);

    if !file_exists("in") {
        os.Mkdir("in", os.ModeDir)
    }
    if !file_exists("out") {
        os.Mkdir("out", os.ModeDir)
    }
    if !file_exists("config") {
        os.Mkdir("config", os.ModeDir)
    }

    paths, _ := os.ReadDir("in");

    for it_index, it := range paths {
        fmt.Println(it_index, it.Name(), it.IsDir());
    }
}
