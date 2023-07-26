"# omdb" 

## Description
This is OMDB Wrapper Service.

Stack Tech: Golang, Redis, gRPC

## Clone/Download Repo
    git clone https://github.com/kenanya/omdb.git

## How to Start
cd omdb<br/>
go run cmd\server\main.go

## Consume Service
We can use <a href="https://appimage.github.io/BloomRPC/">BloomRPC</a> to test consuming this service. After you download and install the BloomRPC, you have to import the protobuf file at omdb/api/v1/omdb-service.proto. As the default, you will get the initial random value as sample request when the protobuf file has been imported. You can use the initial value or yours to test the service.

## The API documentation
Below are the sample requests and expected responses for each API:
### 1. Get movie by ID
#### Sample Request
{
"id": "tt4853102"
}

#### Sample Response
{
"actors": [
"Kevin Conroy",
"Mark Hamill",
"Tara Strong"
],
"id": "tt4853102",
"title": "Batman: The Killing Joke",
"year": "2016",
"rated": "R",
"genre": "Animation, Action, Crime",
"plot": "As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.",
"director": "Sam Liu",
"language": "English",
"country": "United States",
"type": "movie",
"poster_url": "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
}


### 2. Search Movie
#### Sample Request
{
"query": "batman",
"type": "movie",
"page": 10
}

#### Sample Response
{
"movies": [
{
"id": "tt5479456",
"title": "Batman V Arrow",
"year": "2016",
"type": "movie",
"poster_url": "https://m.media-amazon.com/images/M/MV5BY2UzYzQwNzctMzc2MC00MDk2LTljZjgtMGE3NDdiYmY5MzFiXkEyXkFqcGdeQXVyNDk4MzA4Mjk@._V1_SX300.jpg"
},
{
"id": "tt1006834",
"title": "Beyond Batman: From Jack to the Joker",
"year": "2005",
"type": "movie",
"poster_url": "https://m.media-amazon.com/images/M/MV5BNzA4ZjZkODYtM2Y4NS00NDBjLWE5NjMtYzg2MzUwYjkzOTFjXkEyXkFqcGdeQXVyMjQ0NzE0MQ@@._V1_SX300.jpg"
},
{
"id": "tt1074939",
"title": "Batman Begins: Behind the Mask",
"year": "2005",
"type": "movie",
"poster_url": "N/A"
},
],
"total_results": "439"
}



