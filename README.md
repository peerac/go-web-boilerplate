# JTIK-PNL :: PEPETA
### Aplikasi Pengelolaan Pendaftaran Pembimbing Tugas Akhir 

## Requirement
1. Postgres installed on the system

## Setup
1. Clone the project on `https://github.com/JTIK-PNL/pepeta.git`
2. Open the project with your IDE
3. Open terminal, navigate to project folder and run `go mod tidy` to install the dependencies
4. Add `.env` file with the following content:
```dotenv
# Aplication ENV
APP_PORT=

# DATABASE ENV
DB_HOST=
DB_USER=
DB_PASSWORD=
DB_PORT=
DB_NAME=
```

## Running the app
1. Open root project folder with the terminal and run `go run .`
2. Open browser, navigate to `localhost:<APP_PORT>`

## Setup Hot Reload with [air](https://github.com/cosmtrek/air)
1. Install air globally `go install github.com/cosmtrek/air@latest`
2. Run `air init` to generate air.toml file. you can modify the file, refer to the url above.
3. Run `air` to start the hot reload.