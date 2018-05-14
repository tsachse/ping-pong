# Ping-Pong mit Gopherjs

Eine Spielerei mit GopherJS

- gin
- gopherjs
- gopherjs-vue

## Server

Der Server start auf Port 3000

`````
go run main.go
`````

## GopherJS App

http://localhost:3000/public/

Die App befindet sich im Unterverzeichnis public

`````
cd public
gopherjs build app.go -o app.js --tags debug -v -w
`````
