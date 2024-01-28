# todo-app

Todo list training project using Go backend and react frontend

Go backend is based on [this guide](https://learning-cloud-native-go.github.io/docs/overview/).

# How to run this thing:

In the root folder of this project you can find `docker-compose.yml` file, which you can use to quickly start the project.
Please be advised that compose config is set up to read env values from `.env` which is not included in repository.
The necessary env vars to run backend can be found in `/backend/config/config.go`
## backend migrations
To see help on backend migrations run below command in the 'backend' dir
```bash
go run ./cmd/migrations --help
```
