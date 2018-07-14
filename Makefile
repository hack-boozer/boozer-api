
build:
	go build
	./boozer-api

deploy:
	git checkout -b deploy
	dep ensure
	GOOS=linux GOARCH=amd64 go build
	git add -f boozer-api
	git commit -m 'deploy to heroku'
	git push -f heroku master
