package data

import (
	"database/sql"

	"follow-along.whathebea.com/internal/validator"
)

func ValidateCharacter(v *validator.Validator, character *Character) {
	v.Check(character.Name != "", "name", "must be provided")
	v.Check(character.Age != "", "age", "must be provided")
	v.Check(character.HorrorGenre != "", "horrorGenre", "must be provided")
	v.Check(character.FirstAppearance != "", "firstAppearance", "must be provided")
}

type Character struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Age             string `json:"age"`
	HorrorGenre     string `json:"horrorGenre"`
	FirstAppearance string `json:"firstAppearance"`
}

type CharacterModel struct {
	DB *sql.DB
}

func (c CharacterModel) Insert(character *Character) error {
	return nil
}

func (c CharacterModel) Get(id int) (*Character, error) {
	return nil, nil
}

func (c CharacterModel) Update(character *Character) error {
	return nil
}

func (c CharacterModel) Delete(id int) error {
	return nil
}
