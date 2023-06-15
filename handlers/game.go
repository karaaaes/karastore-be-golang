package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"belajar-golang-bookstore2/config"
	"belajar-golang-bookstore2/models"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	db, err := config.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, cover_game, game_name, years, play_to_complete, created_at, modified_at FROM games")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get games"})
		fmt.Println(err)
		return
	}
	defer rows.Close()

	games := []models.Game{}
	for rows.Next() {
		var game models.Game
		err := rows.Scan(&game.ID, &game.Cover_Game, &game.Game_Name, &game.Years, &game.Play_To_Complete, &game.Created_At, &game.Modified_At)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan game"})
			return
		}
		games = append(games, game)
	}

	c.JSON(http.StatusOK, games)
}

func GetHeadlineGames(c *gin.Context) {
	db, err := config.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, cover_game, game_name, description, years, genre, play_to_complete, type, created_at, modified_at FROM games WHERE type ='headline' ORDER BY id DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get games"})
		fmt.Println(err)
		return
	}
	defer rows.Close()

	games := []models.Game{}
	for rows.Next() {
		var game models.Game
		var description sql.NullString
		err := rows.Scan(&game.ID, &game.Cover_Game, &game.Game_Name, &description, &game.Years, &game.Genre, &game.Play_To_Complete, &game.Type, &game.Created_At, &game.Modified_At)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan game"})
			fmt.Println(err)
			return
		}

		if description.Valid {
			game.Description = description.String
		} else {
			game.Description = "" // atau bisa menggunakan nilai default lain jika ingin
		}

		games = append(games, game)
	}

	c.JSON(http.StatusOK, games)
}

func GetPopularGames(c *gin.Context) {
	db, err := config.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,squared_cover_game,game_name,years,genre,play_to_complete,type,total_comment,total_view,created_at,modified_At FROM games WHERE squared_cover_game != NULL OR squared_cover_game != '' ORDER BY total_comment DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get games"})
		fmt.Println(err)
		return
	}
	defer rows.Close()

	games := []models.PopularGame{}
	for rows.Next() {
		var game models.PopularGame
		err := rows.Scan(&game.ID, &game.Squared_Cover_Game, &game.Game_Name, &game.Years, &game.Genre, &game.Play_To_Complete, &game.Type, &game.Total_Comment, &game.Total_View, &game.Created_At, &game.Modified_At)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan game"})
			fmt.Println(err)
			return
		}

		games = append(games, game)
	}

	c.JSON(http.StatusOK, games)
}

func GetTrendingGames(c *gin.Context) {
	db, err := config.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,squared_cover_game,game_name,years,genre,play_to_complete,type,total_comment,total_view,created_at,modified_At FROM games WHERE squared_cover_game != NULL OR squared_cover_game != '' ORDER BY total_view DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get games"})
		fmt.Println(err)
		return
	}
	defer rows.Close()

	games := []models.PopularGame{}
	for rows.Next() {
		var game models.PopularGame
		err := rows.Scan(&game.ID, &game.Squared_Cover_Game, &game.Game_Name, &game.Years, &game.Genre, &game.Play_To_Complete, &game.Type, &game.Total_Comment, &game.Total_View, &game.Created_At, &game.Modified_At)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan game"})
			fmt.Println(err)
			return
		}

		games = append(games, game)
	}

	c.JSON(http.StatusOK, games)
}

func GetGameDetails(c *gin.Context) {
	db, err := config.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	gameID := c.Param("id")

	var game models.DetailGame
	err = db.QueryRow("SELECT id, squared_cover_game, game_name, description, years, genre, price, play_to_complete, type, total_comment, total_view FROM games WHERE id = ?", gameID).Scan(&game.ID, &game.Squared_Cover_Game, &game.Game_Name, &game.Description, &game.Years, &game.Genre, &game.Price, &game.Play_To_Complete, &game.Type, &game.Total_Comment, &game.Total_View)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get game"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, game)
}

func PostTransactions(c *gin.Context) {
	db, err := config.NewDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var tgame models.TransactionsDetails
	if err := c.ShouldBindJSON(&tgame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})

		fmt.Println(err)
		return
	}

	result, err := db.Exec("INSERT INTO transactions_games (transaction_id, game_id, price, full_name, address, email, phone_number, payment_status, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		tgame.Transaction_ID, tgame.Game_ID, tgame.Price, tgame.Full_Name, tgame.Address, tgame.Email, tgame.Phone_Number, tgame.Payment_Status, tgame.Created_At)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		fmt.Println(err)
		return
	}

	transactionID, _ := result.LastInsertId()
	tgame.ID = int(transactionID)

	c.JSON(http.StatusCreated, tgame)
}
