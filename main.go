package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`
			<html>
				<head>
					<title>Chat</title>
				</head>
				<body>
					Let's chat
				</body>
			</html>
		`))
	})

	if err := http.ListenAndServe(":8180", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
