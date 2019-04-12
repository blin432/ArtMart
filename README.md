# Artmart

Full stack E-commerce application to buy hand-made crafts

## [Click here to visit live app](ec2-18-191-209-102.us-east-2.compute.amazonaws.com)

By Faris Huskovic and Benjamin Lin


## Tech Stack Used

React.js for the client-side/front-end

Golang for the server-side/back-end

Postgres DB for persistance

## Dependencies

inside /frontend install the node dependencies

	npm install
	
outside front end install the Go dependencies

	go get github.com/go-pg/pg
	go get github.com/joho/godotenv
	go get github.com/gorilla/securecookie
	go get golang.org/x/crypto/bcrypt
	
## Running

from root open two terminal windows.

- start the client

		cd frontend && npm start
	
- start the server

		cd ../ && go run main.go
	

## Building

First create a production build of the front-end

	cd frontend && npm run build

Next go back up 1 dir and build the binary

	cd ../ && go build .
	
Now just run the binary

	./server
	
## [Click here to visit live app](ec2-18-191-209-102.us-east-2.compute.amazonaws.com)
