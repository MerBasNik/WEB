package repository

import (
	"fmt"
	"strings"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
)

type ChatItemPostgres struct {
	db *sqlx.DB
}

func NewChatItemPostgres(db *sqlx.DB) *ChatItemPostgres {
	return &ChatItemPostgres{db: db}
}

func (r *ChatItemPostgres) GetUsers(userId, chatId int) ([]int, error) {
	var usersID []int
	query := fmt.Sprintf("SELECT tl.user_id FROM %s tl WHERE tl.chatlists_id = $1", usersChatListsTable)
	if err := r.db.Select(&usersID, query, chatId); err != nil {
		return usersID, err
	}

	return usersID, nil
}

func (r *ChatItemPostgres) Create(chatId int, item chat.ChatItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (username, description, chatlists_id) values ($1, $2, $3) RETURNING id", chatItemsTable)

	row := tx.QueryRow(createItemQuery, item.Username, item.Description, item.Chatlist_id)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ChatItemPostgres) GetAll(userId, listId int) ([]chat.ChatItem, error) {
	var items []chat.ChatItem
	query := fmt.Sprintf(`SELECT ti.id, ti.username, ti.description, ti.chatlists_id FROM %s ti 
	INNER JOIN %s ul on ul.chatlists_id = li.chatlists_id WHERE ti.chatlists_id = $1 AND ul.user_id = $2`,
	chatItemsTable, usersChatListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ChatItemPostgres) GetById(userId, itemId int) (chat.ChatItem, error) {
	var item chat.ChatItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti 
	INNER JOIN %s ul on ul.chatlists_id = ti.chatlists_id WHERE ti.id = $1 AND ul.user_id = $2`,
	chatItemsTable, usersChatListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *ChatItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s ul 
	WHERE ti.chatlists_id = ul.chatlists_id AND ul.user_id = $1 AND ti.id = $2`,
	chatItemsTable, usersChatListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *ChatItemPostgres) Update(userId, itemId int, input chat.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s ul WHERE ti.chatlists_id = ul.chatlists_id
		AND ul.user_id = $%d AND ti.id = $%d`,
		chatItemsTable, setQuery, usersChatListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
