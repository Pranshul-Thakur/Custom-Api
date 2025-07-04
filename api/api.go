package api

type CoinbalanceParams struct { // Parameters that the API will take
	username string
}

type CoinbalanceResponse struct { // Responses the succesfull status code and account balance
	code    int // Response code
	balance int64
}

type Error struct { // Response returned when error occurs
	code    int // Response code
	message string
}
