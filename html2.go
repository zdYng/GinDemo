package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
	"fmt"
	"net/http"
)

func formatAsDate(t time.Time)string{
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d",year,month,day)
}

func main(){
	router := gin.Default()
	router.Delims("{[{","}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("templates/raw.tmpl")
	router.GET("/raw", func(context *gin.Context) {
		context.HTML(http.StatusOK,"raw.tmpl",map[string]interface{}{
			"now": time.Date(2019,9,1,0,0,0,0,time.UTC),
		})
	})
	//html := template.Must(template.ParseFiles("file1","file2"))
	//router.SetHTMLTemplate(html)
	router.Run(":8080")
}

