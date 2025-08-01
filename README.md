# Coupon Code System

A full-stack coupon management and redemption system built with **Go**, **React + ShadCN**, and **PostgreSQL**.


## Setup Instructions

1. Use Docker
Configure the environment files and flags in the docker compose file and the Dockerfile
```bash
docker compose up --build
```

2. Self host

To self host the coupon management system, you need a postgres server.
Then you can pass the dsn or connection string to the executable doing
```
go build -o ./build/main ./cmd/api
./build/main -db-dsn postgres://user:password@host/database?sslmode=false # replace user,password,host,database
```

To list the flags that can be passed (port, dsn, jwt_secret)
```
./build/main -h
```

To run the frontend, go to the frontend folder and run
```
npm run dev
# pnpm run dev
```

## Features

-  JWT-based authentication
-  Public and private coupons
-  Redemption tracking
-  Secure, scalable, and testable backend
-  Designed with modular, clean architecture

## API endpoints

- /register
- /login
- /coupon/get
- /coupon/create
- /coupon/redeem
- /coupon/redemptions

## Database Schema (ER Diagram)

<center>
<img src="./assets/erd.svg" height="600px" />
</center>
