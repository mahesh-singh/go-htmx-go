package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	client string
	store  Store
}

func NewApiServer() *ApiServer {
	env := GetEnv()
	fmt.Print(env)
	return &ApiServer{client: env.Listen, store: NewDbInstance()}
}

func (api *ApiServer) Home(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}

func (api *ApiServer) Contact(c *gin.Context) {
	person := Person{Name: "Mahesh Singh", Address: "This is my address", Email: "myemail@mydomin.com"}
	c.HTML(http.StatusOK, "contact.html", person)
}

func (api *ApiServer) ContactEdit(c *gin.Context) {
	//id := c.Params("id")
	person := Person{Name: "Mahesh Singh", Address: "This is my address", Email: "myemail@mydomin.com"}

	c.HTML(http.StatusOK, "contact-edit.html", person)
}

func (api *ApiServer) ContactSave(c *gin.Context) {
	err := c.Request.ParseForm()
}

func (api *ApiServer) Start() {
	//api.store.preStart()

	//gin default router
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", api.Home)

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/contact", api.Contact)

	r.GET("/contact/1/edit", api.ContactEdit)

	err := r.Run(api.client)
	if err != nil {
		panic(err)
	}
}
