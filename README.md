Alura Go Projects
=================

## Get the repo in your machine


```bash
Clone the repo in your machine
$ git clone https://github.com/mahauni/Alura-go.git

And enter the folder
$ cd ./Alura-go
```


## GO REST API

# For you to run the frontend aplication you will need to have the repo in you local machine

```bash
Enter the frontend folder
$ cd ./RestAPI-Go/frontend/

Install all dependencies
$ npm install
$ npm update
$ npm run

And then seach in your machine http://localhost:3000
```

# If you only want to run the back-end you will need to have docker in your machine running
```bash
$ cd ./RestAPI-Go/

Initialize the database via docker
$ docker-compose up

After all is up, run the api
$ go run main.go

And the the output of the request will all go to the http://localhost:8000

The headers to the browser you will have to look in the routers.go file
```
