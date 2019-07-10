package main

import (
	"flag"
	"net/http"
	"log"
	"fmt"
)

var httpAddr = flag.String("http", ":8080", "Listen address")

func main() {
	flag.Parse()
	http.Handle("assets/", http.StripPrefix("assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pusher, ok := w.(http.Pusher)
		if ok {
			// Push is supported. Try pushing rather than
			// waiting for the browser request these static assets.
			if err := pusher.Push("assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
			if err := pusher.Push("assets/style.css", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		fmt.Fprintf(w, indexHTML)
	})
	log.Fatal(http.ListenAndServeTLS(*httpAddr, "testdata/cert.pem", "testdata/key.pem", nil))
}

const indexHTML = `<html>
<head>
	<title>Hello World</title>
	<script src="assets/app.js"></script>
	<link rel="stylesheet" href="assets/style.css"">
</head>
<body>
Hello, gopher!
</body>
</html>
`

//var html = template.Must(template.New("https").Parse(`
//<html>
//<head>
//	<title>Https Test</title>
//<script src="/assets/app.js"></script>
//</head>
//<body>
//	<h1 style="color:red;">Welcome,Ginner!</h1>
//</body>
//</html>
//`))
//var httpAddr = flag.String("http",":8080","Listen address")
//
//func main(){
//	//flag.Parse()
//	//http.Handler("")
//	r := gin.Default()
//	r.Static("assets","assets")
//	r.SetHTMLTemplate(html)
//
//	r.GET("/", func(context *gin.Context) {
//		if pusher := context.Writer.Pusher();pusher!=nil{
//			if err:=pusher.Push("assets/app.js",nil);err!=nil{
//				log.Printf("Failed to push:%v",err)
//			}
//		}
//		context.HTML(200,"https",gin.H{
//			"status":"success",
//		})
//	})
//	r.RunTLS(":8080","testdata/cert.pem","testdata/key.pem")
//}
