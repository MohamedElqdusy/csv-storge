# README

## System Requirements
( Golang:1.11, Redis)
#### Note: This project builds outside Gopath as it uses Go-modules
# Run
`$ docker-compose up`
###### Note: changing csv_file_path is from Dockerfile 

# Run locally
## 1- Install all the dependencies:
Load environment variables
`$ source c.env`

Build
`$ go build`
## 2- Run the service:
`$ ./csv-storage csv_file_path`

##### End points:
- `http://localhost:1321/` Welcome
- `http://localhost:1321/promotions/:id` GET an object with an id

