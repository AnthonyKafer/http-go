package main

import (
	"context"
	"fmt"
	nameWriter "go-server/NameWriter"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/write-my-name/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/write-my-name/"):]

		res := nameWriter.WriteName(
			name,
			r.URL.Query().Get("style"),
			r.URL.Query().Get("color"),
		)

		fmt.Fprint(w, res)
	})

	mux.HandleFunc("/art/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/art/"):]

		http.ServeFile(w, r, "assets/"+name+".txt")
	})

	mux.HandleFunc("/get-fonts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, nameWriter.GetFonts())
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("METHOD:", r.Method)
		fmt.Println("PATH:", r.URL.Path)

		http.Error(w, "not found: "+r.URL.Path, http.StatusNotFound)
	})

	adapter := httpadapter.NewV2(mux)

	lambda.Start(func(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		return adapter.ProxyWithContext(ctx, req)
	})
}
