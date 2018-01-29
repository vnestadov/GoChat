package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"golang.org/x/net/websocket"
	"io"
)

var dirPath string

func IndexPage(w http.ResponseWriter, req *http.Request, filename string) {
    
    fp, err := os.Open(dirPath + "/" + filename)
    if err != nil{
    	log.Println("Could not open file", err.Error())
    	w.Write([]byte("500 internal server error"))
    	return
    }
    
    defer fp.Close()

    _, err = io.Copy(w, fp)
    if err != nil {
    	log.Println("Could not send file contents", err.Error())
    	w.Write([]byte("500 internal server error"))
    	return
    }
}
func EchoServer (ws *websocket.Conn) {
	log.Println("websocket connected:" + ws.RemoteAddr().String())
	defer log.Println("websocket disconnected:" + ws.RemoteAddr().String())
	_, err := io.Copy(ws,ws)
	if err != nil{
		log.Println("Copy error: " + err.Error())
	}
}

func main() {
	if len(os.Args) < 2{
		log.Fatal("Usage: chatExample <dir>")
	}

	dirPath = os.Args[1]

	fmt.Println("Starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		IndexPage(w,req,"index.html")
	})

	http.HandleFunc("/script.js",func(w http.ResponseWriter, req *http.Request){ 
		IndexPage(w,req,"script.js")
	})

	http.Handle("/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
		}
}