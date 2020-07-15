package main

import (
	"log"
	"net/http"
	"os"
)

//func main() {
//	//调用其他依赖文件
//	libfunc()
//	fmt.Println("MAIN FUNC IS RUNNING~")
//}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    ":8000",
		Handler: http.FileServer(http.Dir(cwd)),
	}
	log.Fatal(srv.ListenAndServe())
}
