package repository

import (
	"fmt"
	"strings"

	chat "github.com/MerBasNik/rndmCoffee"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ChatListPostgres struct {
	db *sqlx.DB
}

func NewChatListPostgres(db *sqlx.DB) *ChatListPostgres {
	return &ChatListPostgres{db: db}
}

func (r *ChatListPostgres) Create(userId chat.UsersForChat) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", chatListsTable)
	row := tx.QueryRow(createListQuery, "Новая встреча")
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, chatlists_id, chatName) VALUES ($1, $2, $3)", usersChatListsTable)
	_, err = tx.Exec(createUsersListQuery, userId.FirstUserId, id, "")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	_, err = tx.Exec(createUsersListQuery, userId.SecondUserId, id, "")
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *ChatListPostgres) RenameChat(userId, chatId int, chat chat.UpdateChat) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if chat.ChatName != nil {
		setValues = append(setValues, fmt.Sprintf("chatName=$%d", argId))
		args = append(args, chat.ChatName)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s WHERE tl.user_id=$%d AND tl.chatlists_id=$%d",
	usersChatListsTable, setQuery, argId, argId+1)
	args = append(args, userId, chatId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ChatListPostgres) GetAll(userId int) ([]chat.ChatList, error) {
	var chats 		  []chat.ChatList
	var list_chatName []string

	query := fmt.Sprintf("SELECT tl.id, tl.title FROM %s tl INNER JOIN %s ul on tl.id = ul.chatlists_id WHERE ul.user_id = $1",
		chatListsTable, usersChatListsTable)
	err := r.db.Select(&chats, query, userId)
	if err != nil {
		return chats, err
	}

	query = fmt.Sprintf("SELECT tl.chatName FROM %s tl WHERE tl.user_id = $1", usersChatListsTable)
	if err := r.db.Select(&list_chatName, query, userId); err != nil {
		return chats, err
	}

	for i := 0; i < len(list_chatName); i++ {
		if list_chatName[i] != "" {
			chats[i].Title = list_chatName[i]
		}
	}

	return chats, nil
}

func (r *ChatListPostgres) GetById(userId, chatId int) (chat.ChatList, error) {
	var chat 	 chat.ChatList
	var chatName string
	query := fmt.Sprintf(`SELECT tl.id, tl.title FROM %s tl INNER JOIN %s ul on tl.id = ul.chatlists_id WHERE ul.user_id = $1 AND ul.chatlists_id = $2`,
		chatListsTable, usersChatListsTable)
	err := r.db.Get(&chat, query, userId, chatId)
	if err != nil {
		return chat, err
	}

	query = fmt.Sprintf(`SELECT tl.chatName FROM %s tl WHERE tl.user_id = $1 AND tl.chatlists_id = $2`,
	usersChatListsTable)
	if err := r.db.Get(&chatName, query, userId, chatId); err != nil {
		return chat, err
	}

	if chatName != "" {
		chat.Title = string(chatName)
	}

	return chat, nil
}

func (r *ChatListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.chatlists_id AND ul.user_id=$1 AND ul.chatlists_id=$2",
		chatListsTable, usersChatListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *ChatListPostgres) Update(userId, listId int, input chat.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.chatlists_id AND ul.chatlists_id=$%d AND ul.user_id=$%d",
		chatListsTable, setQuery, usersChatListsTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ChatListPostgres) FindByTime(userId int, input chat.FindUserInput) (int, error) {
	var id int
	tx, err := r.db.Begin()
	if err != nil {
		return id, err
	}

	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id, start_day, end_day, start_time, end_time) VALUES ($1, $2, $3, $4, $5)", findUsersTable)
	_, err = tx.Exec(createListQuery, userId, input.StartDay, input.EndDay, input.StartTime, input.EndTime)
	if err != nil {
		tx.Rollback()
		return id, err
	}
	tx.Commit()

	query := fmt.Sprintf(`SELECT tl.user_id FROM %s tl WHERE
	(tl.start_day <= $1) AND ($2 <= tl.end_day) AND 
	(tl.start_time <= $3) AND ($4 <= tl.end_time) AND tl.user_id!=$5`, findUsersTable)
	err = r.db.Get(&id, query, input.EndDay, input.StartDay, input.EndTime, input.StartTime, userId)

	return id, err
}

func (r *ChatListPostgres) FindByHobby(userId1, userId2 int) ([]chat.UserHobby, error) {
	var lists []chat.UserHobby

	query := fmt.Sprintf(`SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id
	WHERE ul.user_id=$1 INTERSECT SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id
	WHERE ul.user_id=$2`,
	userHobbyTable, usersHobbyListsTable, userHobbyTable, usersHobbyListsTable)
	err := r.db.Select(&lists, query, userId1, userId2)

	return lists, err
}

func (r *ChatListPostgres) DeleteFindUsers(userId chat.UsersForChat) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.user_id = $1", findUsersTable)
	if _, err := r.db.Exec(query, userId.FirstUserId); err != nil {
		return err
	}
	if _, err := r.db.Exec(query, userId.SecondUserId); err != nil {
		return err
	}

	return nil
}