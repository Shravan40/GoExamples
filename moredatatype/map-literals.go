package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex {
    "Bell Labs": Vertex{
        40.406324, -134.3524123,
    },
    "Google": Vertex{
        34.452512,-12.23512,
    },
}

func main(){
    fmt.Println(m)
}