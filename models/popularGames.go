package models

type PopularGame struct {
	ID                 int    `json:"id"`
	Squared_Cover_Game string `json:"squared_cover_game"`
	Game_Name          string `json:"title"`
	Years              int    `json:"years"`
	Genre              string `json:"genre"`
	Play_To_Complete   int    `json:"play_to_complete"`
	Type               string `json:"type"`
	Total_Comment      string `json:"total_comment"`
	Total_View         string `json:"total_view"`
	Created_At         string `json:"created_at"`
	Modified_At        string `json:"modified_at"`
}
