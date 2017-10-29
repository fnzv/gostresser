package main


import  (
 "math/rand"
 "strconv"
 "os"
 "time"
 "net"
 "fmt"
 vegeta "github.com/tsenart/vegeta/lib"
)
var URL string
//Author: Sami Yessou
// Gostresser - Easy and Fast TCP/UDP connection stresser using Goroutines workers
// it's my first golang project so don't expect any pro dev code, just for fun and testing

func init() {
    rand.Seed(time.Now().UnixNano())
}

func FloatToString(input_num float64) string {
    // to convert a float number to a string
    return strconv.FormatFloat(input_num, 'f', 6, 64)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}


func main() {


// Checking if Args are present
// if not display usage
 if len(os.Args) > 1 {
        hostname := os.Args[1]
        port := os.Args[2]
	if port == "80" {
		URL = "http://"+hostname+"/"
	}
        proto := os.Args[3]
        seconds_args  := os.Args[4]
        num_workers_args := os.Args[5]
//        URL := "http://testcdn2.sami.pw/img.jpg"
        socket := hostname+":"+port
        i63, _ := strconv.ParseInt(seconds_args, 10, 32)
        seconds := int(i63)
        i64, _ := strconv.ParseInt(num_workers_args, 10, 32)
        num_workers := int(i64)
	fmt.Println("URL IS ", URL)
        fmt.Println("Start Load stressing: ", socket)
        for i := 0; i < num_workers; i++ {
                go loadtest(socket,proto)
        //        fmt.Println("Worker ",i)
		vegeta_integration(URL)
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
	fmt.Fprintf(conn, "GET /"+RandStringRunes(30)+ " HTTP/1.1" + "\n")
        fmt.Print(".")
	defer conn.Close()
    //fmt.Println(".")
    // Add multiple requests and non specific to HTTP
        }

}

func vegeta_integration(url string) {
  rate := uint64(100) // per second
  duration := 5 * time.Second
  targeter := vegeta.NewStaticTargeter(vegeta.Target{
    Method: "GET",
    URL:    url,
  })
  attacker := vegeta.NewAttacker()

  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration) {
    metrics.Add(res)
  }
  metrics.Close()

//  fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
  fmt.Printf("\nMax Lantency: %s\n", metrics.Latencies.Total)
  fmt.Printf("\nMean Lantency: %s\n", metrics.Latencies.Mean)
  fmt.Printf("\nMean Byts In : %s\n", FloatToString(metrics.BytesIn.Mean))
  fmt.Printf("\nMean Bytes Out : %s\n", FloatToString(metrics.BytesOut.Mean))

}
