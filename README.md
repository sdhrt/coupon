# Coupon Code System

A full-stack coupon management and redemption system built with **Go**, **React + ShadCN**, and **PostgreSQL**.


## Setup Instructions

1. Use Docker
```bash
docker compose up --build
```

2. Self host

To self host the coupon management system, to run the backend you can simply run
```
make run
```
To list the flags that can be passed
```
make help
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
