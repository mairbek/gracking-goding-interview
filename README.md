Modify `.env` file and run `docker-compose run app sh`.

Inside the container run `glide install --strip-vendor` to load deps.

Inside the container run `go run main.go`.
