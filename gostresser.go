package main

import "strconv"
import "os"
import "time"
import "net"
import "fmt"
//Author: Sami Yessou
// Gostresser - Easy and Fast TCP/UDP connection stresser using Goroutines workers
// it's my first golang project so don't expect any pro dev code, just for fun and testing

func main() {


// Checking if Args are present
// if not display usage
 if len(os.Args) > 1 {
        hostname := os.Args[1]
        port := os.Args[2]
        proto := os.Args[3]
        seconds_args  := os.Args[4]
        num_workers_args := os.Args[5]

        socket := hostname+":"+port
        i63, _ := strconv.ParseInt(seconds_args, 10, 32)
        seconds := int(i63)
        i64, _ := strconv.ParseInt(num_workers_args, 10, 32)
        num_workers := int(i64)
        fmt.Println("Start Load stressing: ", socket)
        for i := 0; i < num_workers; i++ {
                go loadtest(socket,proto)
                fmt.Println("Worker ",i)
                                 }
        time.Sleep(time.Second * time.Duration(seconds))
        } else {
        fmt.Println("Usage example:\n ./gostresser <hostname> <port> <protocol> <seconds> <num_workers>\n ./gostresser example.com 80 tcp 120 15")
        }
}


func loadtest(socket string,proto string){
    conn, _ := net.Dial(proto, socket)
    for range time.Tick(time.Millisecond * 100) {
  // OPTIONS * HTTP/1.1
        fmt.Fprintf(conn, "GET /index.php HTTP/1.1" + "\n")
        fmt.Fprintf(conn, "OPTIONS * HTTP/1.1" + "\n")
        fmt.Print(".")
    //fmt.Println(".")
    // Add multiple requests and non specific to HTTP
        }

}
