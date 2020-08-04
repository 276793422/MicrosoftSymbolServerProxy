package main

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
	"strings"
	"time"
)
var v Config
var m map[string] string

func main() {

	m := make(map[string]string)
	v = LoadConfig("config.json")
	fmt.Println(v)
	fmt.Println(len(v.Path))
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())

	// Routes
	var str string
	for i := 0; i < len(v.Path); i++ {
		if len(v.Path[i].Exec) != 0 {
			fmt.Println("has exec : ", v.Path[i].Exec)
		}
		str = v.Path[i].Path
		if str[0] != '/' {
			str = "/" + str
		}
		if str[len(str) - 1] != '/' {
			str = str + "/"
		}
		fmt.Println("dir :\"" + str + "\"")
		m[str] = v.Path[i].Exec
		e.GET(str, Specify)
		fmt.Println("dir :\"" + str + "*\"")
		m[str] = v.Path[i].Exec
		e.GET(str + "*", Specify)
	}

	e.POST("/zoo.gif", PostFile)

	e.File("/favicon.ico", v.Root +"images/favicon.ico")
	e.File("/index.htm", v.Root + "index.hml")
	e.File("/*", v.Root + "*");
	e.Static("/" , v.Root)
	//	e.GET("/", Param)
	e.GET("/download/symbols/*", MicrosoftSymbols)
	//	e.GET("/blob/", Blob)

	// Start server
	str = v.Ip + ":" + v.Port
	fmt.Println("ip:port = ", str)

	var LogPath string
	if v.Log != "" {
		LogPath = v.Log
	} else {
		LogPath = "Log.txt"
	}

	fd, _ := os.OpenFile(
		LogPath,
		os.O_RDWR|os.O_APPEND,
		0666,
	)
	//	e.Logger.SetOutput(fd)
	e.Use(middleware.LoggerF(fd))

	e.Logger.Fatal(e.Start(str))
}

// Handler
func Param(c echo.Context) error {
	req := c.Request()
	str := v.Root + req.RequestURI
	str = strings.Replace(str, "/", "\\", -1)
	return c.String(http.StatusOK,
		"Param : " + time.Now().String() + "\r\n" +
		"QueryString : " + c.QueryString() + "\r\n" +
		"RealIP : " + c.RealIP()+ "\r\n" +
		"Scheme : " + c.Scheme() + "\r\n" +
		"Root : " + v.Root + "\r\n" +
		"url : " + req.URL.Path + "\r\n" +
		"uri : " + req.RequestURI + "\r\n" +
		"File :" + str + "\r\n" +
		"Path : " + c.Path())
}

// Handler
func Specify(c echo.Context) error {
	req := c.Request()
	str := v.Root + req.RequestURI
	str = strings.Replace(str, "/", "\\", -1)
	exec := m[c.Path()]
	return c.String(http.StatusOK,
		"Specify : " + time.Now().String() + "\r\n" +
		"QueryString : " + c.QueryString() + "\r\n" +
		"RealIP : " + c.RealIP()+ "\r\n" +
		"Scheme : " + c.Scheme() + "\r\n" +
		"Root : " + v.Root + "\r\n" +
		"url : " + req.URL.Path + "\r\n" +
		"uri : " + req.RequestURI + "\r\n" +
		"File :" + str + "\r\n" +
		"Path : " + c.Path() + "\r\n" +
		"Exec :" + exec)
}

// Handler
func Blob(c echo.Context) error {
	var by = []byte{0x31, 0x32, 0x33, 0x34, 0x00}
	return c.Blob(http.StatusOK,
		"File : " + time.Now().String() + "\r\n" +
		"Path : " + c.Path() + "\r\n" +
		"QueryString : " + c.QueryString() + "\r\n" +
		"RealIP : " + c.RealIP() + "\r\n" +
		"Scheme : " + c.Scheme(), by)
}

func PostFile(c echo.Context) error {
	return errors.New("")
}

func IsFileExist(str string) bool {
	f, err := os.Open(str)
	if err != nil && os.IsNotExist(err) {
		return false
	} else {
		f.Close()
		return true
	}
}

// Handler
func MicrosoftSymbols(c echo.Context) error {
	req := c.Request()
	str := v.Root + req.RequestURI
	str = strings.Replace(str, "/", "\\", -1)
	/*
	return c.String(http.StatusOK,
		"Param : " + time.Now().String() + "\r\n" +
			"QueryString : " + c.QueryString() + "\r\n" +
			"RealIP : " + c.RealIP()+ "\r\n" +
			"Scheme : " + c.Scheme() + "\r\n" +
			"Root : " + v.Root + "\r\n" +
			"url : " + req.URL.Path + "\r\n" +
			"uri : " + req.RequestURI + "\r\n" +
			"File :" + str + "\r\n" +
			"Path : " + c.Path())
	*/
	if IsFileExist(str) == false {
		//	下载
		p, _ := os.StartProcess("wget64.exe", []string{"wget64.exe", "-r","http://msdl.microsoft.com" + req.RequestURI}, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		p.Wait();
	}
	if IsFileExist(str) == false {
		return c.String(http.StatusNotFound, "Nemesis Error")
	}
	//	下载成功了，就发给下面
	return c.File(str)
}