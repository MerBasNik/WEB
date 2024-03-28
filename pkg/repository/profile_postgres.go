package repository

import (
	"fmt"
	"strings"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) CreateProfile(userId int, profile chat.Profile) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, surname, photo, telegram, city) values ($1, $2, $3, $4, $5) RETURNING id", usersProfileTable)
	row := r.db.QueryRow(query, profile.Name, profile.Surname, profile.Photo, profile.Telegram, profile.City)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, profile_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ProfilePostgres) GetProfile(userId, profileId int) (chat.Profile, error) {
	var profile chat.Profile

	query := fmt.Sprintf(`SELECT tl.id, tl.name, tl.surname, tl.photo, tl.telegram, tl.city FROM %s tl INNER JOIN %s ul on tl.id = ul.profile_id WHERE ul.user_id = $1 AND ul.profile_id = $2`,
		usersProfileTable, usersListsTable)
	err := r.db.Get(&profile, query, userId, profileId)

	return profile, err
}

func (r *ProfilePostgres) EditProfile(userId, profileId int, input chat.UpdateProfile) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.Photo != "" {
		setValues = append(setValues, fmt.Sprintf("photo=$%d", argId))
		args = append(args, input.Photo)
		argId++
	}

	if input.City != nil {
		setValues = append(setValues, fmt.Sprintf("city=$%d", argId))
		args = append(args, *input.City)
		argId++
	}

	if input.Telegram != nil {
		setValues = append(setValues, fmt.Sprintf("telegram=$%d", argId))
		args = append(args, *input.Telegram)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.profile_id AND ul.profile_id=$%d AND ul.user_id=$%d",
		usersProfileTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, profileId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (s *ProfilePostgres) CreateHobby(userId int, hobby chat.UserHobby) (int, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s description VALUES $1 RETURNING id", userHobbyTable)
	row := tx.QueryRow(createListQuery, hobby.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, userhobby_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (s *ProfilePostgres) GetAllHobby(userId int) ([]chat.UserHobby, error) {
	var hobbylist []chat.UserHobby

	query := fmt.Sprintf("SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id WHERE ul.user_id = $1",
		userHobbyTable, usersListsTable)
	err := s.db.Select(&hobbylist, query, userId)

	return hobbylist, err
}

func (s *ProfilePostgres) DeleteHobby(userId, hobbyId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.userhobby_id AND ul.user_id=$1 AND ul.userhobby_id=$2",
		userHobbyTable, usersListsTable)
	_, err := s.db.Exec(query, userId, hobbyId)

	return err
}
