package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}

type BuyProduct struct {
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

// слоистая архитектура
// 1)handler/controller(обработка запроса)
// 2)service(бизнес логика)
// 3)storage/repository(данные или БД)
