package models

type DetailGame struct {
	ID                 int     `json:"id"`
	Squared_Cover_Game string  `json:"squared_cover_game"`
	Game_Name          string  `json:"title"`
	Description        string  `json:"description"`
	Years              int     `json:"years"`
	Genre              string  `json:"genre"`
	Price              float32 `json:"price"`
	Play_To_Complete   int     `json:"play_to_complete"`
	Type               string  `json:"type"`
	Total_Comment      string  `json:"total_comment"`
	Total_View         string  `json:"total_view"`
}
