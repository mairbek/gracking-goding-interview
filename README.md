Quick getting started with golang using docker. You don't have to set up $GOPATH and stuff,
just run the code in the container.

Modify `.env` file and run `docker-compose run app sh`.

Inside the container run `glide install --strip-vendor` to load deps.

Inside the container run `go run main.go`.
