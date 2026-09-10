package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/go")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			log.Printf("读到的内容:%s \n", string(buf[:n]))
		}
		if err == io.EOF {
			log.Println("读取完毕")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("读到的内容:%s \n", string(buf[:n]))
	}
}
