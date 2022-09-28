package models

type ExerciceProperties struct {
	Distance float32 `json:"distance"`
	Length   int     `json:"lenth"`
	Rep      int     `json:"rep"`
	Series   []int   `json:"series"`
}
