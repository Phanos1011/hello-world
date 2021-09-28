/*package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//var mydata string[];
	// запрос в формате http
	//response, request := http.Get("http://localhost:8000/_/api/")
	response, request := http.Get("http://localhost:8000/_/api/v3/key_values")

	if request != nil {
		log.Fatal("failed to read file:",request) //Fatal is equivalent to Print() followed by a call to os.Exit(1)
	}
	defer response.Body.Close()  //закрываем файл defer до выхода из функции мейн



	// копируем инфо в нормальный вывод

	if request != nil {
		n, request := io.Copy(os.Stdout, response.Body)
		log.Fatal("failed to copy file",request)
	}

	fmt.Println("number of bytes:",n)
}
*/



package main
import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"

)

type PageContent struct{
	Info [] Data `json:"data"`
	Included [] string `json:"included"`
	Links LinkContent `json:"links"`
}

type LinkContent struct{
	Next string `json:"next"`
	Self string `json:"self"`
}

type Data struct {
	Id string `json:"id"`
	TypeData string `json:"type"`
	Attributes AttributesContent `json:"attributes"`
}
type AttributesContent struct{
	Value ValueContent `json:"value"`
	U32 int `json:"u32"`
}
type ValueContent struct{
	TypeValue TypeContent `json:"type"`
}
type TypeContent struct{
	Value int `json:"value"`
	Name string `json:"name"`
}

type Item struct {
	Pages [] PageContent
}
func main() {
	MyLink:="http://localhost:8000/_/api/v3/key_values"
	length:=1
	counter:=0
	Datan:=PageContent{}


	for  ; length!=0 ;  {
		    if counter!=0 {
				fmt.Println("Data=", Datan.Info[0])
			}
			response, err := http.Get(MyLink)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()



			dataInBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("failed to read json file, error: %v", err)
				return
			}

			error:= json.Unmarshal([]byte(dataInBytes), &Datan)
			if error != nil {
				fmt.Println(error)
				return
			}
            counter++
			fmt.Println("Self of page ",counter,"=",Datan.Links.Self)
			fmt.Println("Next of page ",(counter+1),"=",Datan.Links.Next)

			length=len(Datan.Info)
			MyLink="http://localhost:8000/_/api"+Datan.Links.Next
	}



}
/*
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// запрос в формате http
	response, request := http.Get("http://localhost:8000/_/api/v3/key_values")
	if request != nil {
		log.Fatal("failed to read file:",request) //Fatal is equivalent to Print() followed by a call to os.Exit(1)
	}
	defer response.Body.Close()  //закрываем файл defer до выхода из функции мейн

	// копируем инфо в нормальный вывод
	n, request := io.Copy(os.Stdout, response.Body)
	if request != nil {
		log.Fatal("failed to copy file",request)
	}

	fmt.Println("number of bytes:", n)
} */
