
webserver:

	sudo apt-get install golang -y
	#export GOPATH=$HOME/My_app/go
	go build my_app.go
	docker build -t webserver .
	docker tag webserver:latest  gcr.io/gifted-decker-214823/echohost:latest

deploy: clean
	#go run my_app.go
	#docker run -d --name webserver -p 80:8080 webserver:latest
	gcloud docker -- push gcr.io/gifted-decker-214823/echohost:latest
	kubectl run echohost --image=gcr.io/gifted-decker-214823/echohost:latest
	kubectl expose deployment echohost --port 80 --type LoadBalancer
clean:


	-kubectl delete services echohost
	-kubectl delete deployment echohost
	#docker system prune -a -f
	-docker rm -f `docker ps -aq`
	-rm -R my_app

# make webserver deploy
