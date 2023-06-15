package models

type Game struct {
	ID               int    `json:"id"`
	Cover_Game       string `json:"cover_game"`
	Game_Name        string `json:"title"`
	Description      string `json:"description"`
	Years            int    `json:"years"`
	Genre            string `json:"genre"`
	Play_To_Complete int    `json:"play_to_complete"`
	Type             string `json:"type"`
	Created_At       string `json:"created_at"`
	Modified_At      string `json:"modified_at"`
}
