package models

type Product struct {
    ID            int64    `meddler:"id,pk"`
    Price         int      `meddler:"price"`
    Name          string   `meddler:"name"`
    Description   string   `meddler:"description"`
}