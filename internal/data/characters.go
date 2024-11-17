package data

import (
	"database/sql"
	"errors"

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
	query := `INSERT INTO characters (name, age, horror_genre, first_appearance) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id`

	args := []any{character.Name, character.Age, character.HorrorGenre, character.FirstAppearance}

	return c.DB.QueryRow(query, args...).Scan(&character.ID)
}

func (c CharacterModel) Get(id int64) (*Character, error) {
	query := `SELECT id, name, age, horror_genre, first_appearance FROM characters WHERE id = $1`

	row := c.DB.QueryRow(query, id)

	var character Character

	err := row.Scan(&character.ID, &character.Name, &character.Age, &character.HorrorGenre, &character.FirstAppearance)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &character, nil
}

func (c CharacterModel) Update(character *Character) error {
	return nil
}

func (c CharacterModel) Delete(id int) error {
	return nil
}
