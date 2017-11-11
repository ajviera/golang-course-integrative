# The real challenge in Golang begins

<p align="center">
  <img src="https://github.com/ajviera/golang-course-integrative/blob/master/gopher-star-wars.png"/>
</p>

## Golang workshop integrative by Mercadolibre

_Use the following commands to start the exercises_

`Exercise`
```sh
cd src/integrative
go run integrative.go

 curl -i -H \
 "Accept: application/json" -H \
 "Content-Type: application/json" \
 "http://localhost:8080//categories/:id/prices"
```

`Test Integrative`
```sh
go test -coverprofile=test_reports/cover.out
go tool cover -html=test_reports/cover.out -o test_reports/coverage.html
go test -bench=.
```

`Test Maths`
```sh
cd ../maths
go test -coverprofile=test_maths_reports/cover.out
go tool cover -html=test_maths_reports/cover.out -o test_maths_reports/coverage.html
go test -bench=.
```