package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gogearbox/gearbox"
)

func main() {
	app := gearbox.New()
	// app.Use(middlewareCors)
	app.Post("/api/upload", handlerUpload)
	app.Start(":8090")
}

func middlewareCors(ctx gearbox.Context) {
	ctx.Set("Access-Control-Allow-Origin", "*")
	ctx.Next()
}

func handlerUpload(ctx gearbox.Context) {
	h, e := ctx.Context().FormFile("image")
	if e != nil {
		ctx.SendString(e.Error())
		return
	}

	f, e := h.Open()
	if e != nil {
		ctx.SendString(e.Error())
		return
	}
	defer f.Close()

	fileName := fmt.Sprintf("%d.png", time.Now().UnixNano())
	n, e := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if e != nil {
		ctx.SendString(e.Error())
		return
	}
	defer n.Close()
	_, e = io.Copy(n, f)
	if e != nil {
		ctx.SendString(e.Error())
		return
	}
	ctx.SendString("uploaded as " + fileName)
}
