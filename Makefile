
webserver:

	sudo apt-get install golang -y
	#export GOPATH=$HOME/My_app/go
	docker build -t webserver .
	docker tag webserver:latest  gcr.io/gifted-decker-214823/echohost:latest

	go build my_app.go
	#go run my_app.go
	#docker run -d --name webserver -p 80:8080 webserver:latest
clean:
	docker system prune -a -f
	rm -R my_app
