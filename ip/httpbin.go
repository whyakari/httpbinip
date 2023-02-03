// pacote pegar ip
package httpbin

import (
	"io/ioutil"
	"net/http"
)

// Função de pegar o ip 
func GetIp() (string, error) {
    resp, err := http.Get("http://httpbin.org/ip")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    return string(body), nil

}

