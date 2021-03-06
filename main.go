package main

import (
	"fmt"
	"net/http"
	"os"
	"uploadServer/controllers"
	//"uploadServer/database"
)

func SetupRoutes()	{

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/",fs))
	http.HandleFunc("/", controllers.Init)
	http.HandleFunc("/upload", controllers.UploadFile)//upload file
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/logout", controllers.LogOut)
	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/main", controllers.MainPage)
	http.HandleFunc("/room", controllers.Room)
	http.HandleFunc("/redirect", controllers.Redirect)
	http.HandleFunc("/ws", controllers.WS)
	//http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	//http.ListenAndServe(":8080", nil)
}
func main() {
	fmt.Println("File upload example")
	SetupRoutes()
}


