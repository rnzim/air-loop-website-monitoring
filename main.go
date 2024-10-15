package main
import(
  "fmt"
  "net/http"
  "strings"
  "time"
  "os"
)

func main(){
  
  sites := []string{"htps://www.google.com.br/pygame"}
  // var requestCount uint64
  menuMain()
  opt := choose()
  if opt == 1{
    for{
       time.Sleep(time.Second)
       monitor(sites)
    }
  }
  
}

func menuMain(){
  fmt.Println("Monitor De Sites 2000")
  fmt.Println("-----------------------")
  fmt.Println("1 - Monitorar")
  fmt.Println("2 - Logs")
  fmt.Println("3 - Custom Mon")
  fmt.Println("4 - Sair")
}
func menuCustom(){
  fmt.Println("1 - Digite a url")
  fmt.Println("2 - Logs")
  fmt.Println("3 - Sair")
}
func choose() uint8 {
  var opt uint8
  fmt.Scan(&opt)
  return opt
}
func monitor(sites []string){
  for i, site := range sites{
    domain:=       strings.Split(site,".")
    result,err := http.Get(site)
    if err != nil{
      fmt.Println("Erro Ao Fazer Requisição",err)
    os.Exit(-1)
    }else{
      if result.StatusCode == 200{
        
        fmt.Println(i,"->",domain[1],"\033[32mOnline\033[0m")
      }else{
        fmt.Println(i,"->",domain[1],"\033[31mOffline\033[0m",result.StatusCode)
      }
     // requestCount++
    }
    
  }
/*func saveLog() {
    file, err := os.OpenFile("logs.txt", os.O_CREATE||os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Erro ao Criar Arquivo: ", err)
        return
    }
    defer file.Close()

    _, err = file.WriteString("testekkk")
    if err != nil {
        fmt.Println("Erro ao Escrever no Arquivo: ", err)
    }
}

*/
 
