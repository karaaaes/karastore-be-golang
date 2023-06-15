package models

type TransactionsDetails struct {
	ID             int     `json:"id"`
	Transaction_ID string  `json:"transaction_id"`
	Game_ID        int     `json:"game_id"`
	Price          float32 `json:"price"`
	Full_Name      string  `json:"full_name"`
	Address        string  `json:"address"`
	Email          string  `json:"email"`
	Phone_Number   string  `json:"phone_number"`
	Payment_Status string  `json:"payment_status"`
	Created_At     string  `json:"created_at"`
}
