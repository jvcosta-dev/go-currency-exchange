# GO Currency Exchange

This project leverages Go to intelligently serve currency data from over 160 different countries, using the public Exchange Rate API while minimizing the number of requests made through efficient routines.

## Table of Contents

- [Project Info](#project-info)
- [Running the Application](#running-the-application)
- [Examples](#examples)
  - [Convert Currency](#convert-curency)
  - [Get Rate from Currency](#get-rate-from-currency)
  - [Get All Rates](#get-all-rates)

## Project Info

- **GO**: 1.22.2

## Running the Application

First get your API_KEY on https://app.exchangerate-api.com/dashboard
Now create your `.env.local` file in internal/config dir and setup your `API_KEY`

Now run the main application

```bash
go run cmd/web/main.go
```

## Examples

### Convert Currency

Covert one currency to another and return its value

- **Request**: `GET`
- **Endpoint**: `convert/{from}/{to}/{amount}`
- **Path Params**: `from: "BRL", to: "USD", amount: 1`

- **Response Example**

```json
{
  "from": "BRL",
  "to": "USD",
  "amount": 1,
  "result": 0.172,
  "rate": 5.8055,
  "last_update": "Tue, 05 Nov 2024 00:00:01 +0000"
}
```

### Get Rate from Currency

Return currency rate based on USD

- **Request**: `GET`
- **Endpoint**: `rate/{base}`
- **Path Params**: `base: "BRL"`

- **Response Example**

```json
{
  "base": "BRL",
  "rate": 5.8055,
  "last_update": "Tue, 05 Nov 2024 00:00:01 +0000"
}
```

### Get All Rates

Return all currency rates based on USD

- **Request**: `GET`
- **Endpoint**: `rates`

- **Response Example**

```json
{
  "rates": {
    "AED": 3.6725,
    "AFN": 66.873,
    "ALL": 90.4668,
    "AMD": 387.0439,
    "ANG": 1.79,
    "AOA": 920.9918,
    "ARS": 993.33,
    "AUD": 1.5177,
    "AWG": 1.79,
    "AZN": 1.6993,
    "BAM": 1.7972,
    "BBD": 2,
    "BDT": 119.4957,
    "BGN": 1.7968,
    "BHD": 0.376,
    "BIF": 2910.6207,
    "BMD": 1,
    "BND": 1.3187,
    "BOB": 6.9214,
    "BRL": 5.8055,
    [...]
}
```
