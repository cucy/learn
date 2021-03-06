 package main

 import (
         "fmt"
         "os"
         "time"
         "net/http"
 )

 func main() {

         start := time.Now()

         url := "http://www.golang.org"

         result, err := http.Get(url)

         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer result.Body.Close()

         elapsed := time.Since(start).Seconds()

         fmt.Printf("http.Get to %s took %v seconds \n", url, elapsed)

 }
