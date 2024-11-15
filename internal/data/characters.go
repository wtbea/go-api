package data

type Character struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Age             string `json:"age"`
	HorrorGenre     string `json:"horrorGenre"`
	FirstAppearance string `json:"firstAppearance"`
}
