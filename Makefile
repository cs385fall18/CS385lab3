
webserver:

	sudo apt-get install golang -y
	#export GOPATH=$HOME/My_app/go
	docker build -t webserver .
	go build my_app.go
	go run my_app.go
	#docker run -d --name webserver -p 80:8080 webserver:latest
