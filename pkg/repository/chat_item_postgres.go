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

func (r *ChatItemPostgres) Create(listId int, item chat.ChatItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", chatItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (chatlists_id, chatitems_id) values ($1, $2)", itemsListsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (r *ChatItemPostgres) GetAll(userId, listId int) ([]chat.ChatItem, error) {
	var items []chat.ChatItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li on li.chatitems_id = ti.id
	INNER JOIN %s ul on ul.chatlists_id = li.chatlists_id WHERE li.chatlists_id = $1 AND ul.user_id = $2`,
	chatItemsTable, itemsListsTable, usersChatListsTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *ChatItemPostgres) GetById(userId, itemId int) (chat.ChatItem, error) {
	var item chat.ChatItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description FROM %s ti INNER JOIN %s li on li.chatitems_id = ti.id
	INNER JOIN %s ul on ul.chatlists_id = li.chatlists_id WHERE ti.id = $1 AND ul.user_id = $2`,
	chatItemsTable, itemsListsTable, usersChatListsTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *ChatItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
	WHERE ti.id = li.chatitems_id AND li.chatlists_id = ul.chatlists_id AND ul.user_id = $1 AND ti.id = $2`,
	chatItemsTable, itemsListsTable, usersChatListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *ChatItemPostgres) Update(userId, itemId int, input chat.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
						WHERE ti.id = li.chatitems_id AND li.chatlists_id = ul.chatlists_id AND ul.user_id = $%d AND ti.id = $%d`,
		chatItemsTable, setQuery, itemsListsTable, usersChatListsTable, argId, argId+1)
	args = append(args, userId, itemId)

	_, err := r.db.Exec(query, args...)
	return err
}
