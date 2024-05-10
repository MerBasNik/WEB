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

func (r *ChatListPostgres) CreateList(input chat.UsersForChat) (int, error) {
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
	for i := 0; i < len(input.UsersId); i++ {
		_, err = tx.Exec(createUsersListQuery, input.UsersId[i], id, "")
		if err != nil {
			tx.Rollback()
			return 0, err
		}
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

func (r *ChatListPostgres) GetAllLists(userId int) ([]chat.ChatList, error) {
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

func (r *ChatListPostgres) GetListById(userId, chatId int) (chat.ChatList, error) {
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
		chat.Title = chatName
	}

	return chat, nil
}

func (r *ChatListPostgres) DeleteList(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.chatlists_id AND ul.user_id=$1 AND ul.chatlists_id=$2",
		chatListsTable, usersChatListsTable)
	_, err := r.db.Exec(query, userId, listId)

	return err
}

// func (r *ChatListPostgres) UpdateList(userId, listId int, input chat.UpdateListInput) error {
// 	setValues := make([]string, 0)
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.Title != nil {
// 		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
// 		args = append(args, *input.Title)
// 		argId++
// 	}

// 	setQuery := strings.Join(setValues, ", ")

// 	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.chatlists_id AND ul.chatlists_id=$%d AND ul.user_id=$%d",
// 		chatListsTable, setQuery, usersChatListsTable, argId, argId+1)
// 	args = append(args, listId, userId)

// 	logrus.Debugf("updateQuery: %s", query)
// 	logrus.Debugf("args: %s", args)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }

func (r *ChatListPostgres) FindByTime(userId int, input chat.FindUserInput) ([]int, []chat.UsersInfo, error) {
	var list_users_id []int
	var users_info []chat.UsersInfo
	tx, err := r.db.Begin()
	if err != nil {
		return list_users_id, users_info, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (user_id, count, start_day, end_day, start_time, end_time) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", findUsersTable)
	row := tx.QueryRow(createListQuery, userId, input.Count, input.StartDay, input.EndDay, input.StartTime, input.EndTime)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return list_users_id, users_info, err
	}
	tx.Commit()

	query := fmt.Sprintf(`SELECT tl.id, tl.user_id FROM %s tl WHERE (tl.start_day <= $1) AND ($2 <= tl.end_day) AND 
	(tl.start_time <= $3) AND ($4 <= tl.end_time) AND tl.user_id!=$5 AND tl.count=$6 LIMIT $7`, findUsersTable)
	if err := r.db.Select(&users_info, query, input.EndDay, input.StartDay, input.EndTime, input.StartTime, 
		userId, input.Count, input.Count-1); err != nil {
		return list_users_id, users_info, err
	}

	if len(users_info) != 0 {
		users_info = append(users_info, chat.UsersInfo{Id: id, UserId: userId})
	}

	for i := 0; i < len(users_info); i++ {
		list_users_id = append(list_users_id, users_info[i].UserId)
	}

	return list_users_id, users_info, err
}

func (r *ChatListPostgres) UpdateFindUsersTable(users_info []chat.UsersInfo, count int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	if len(users_info) == (count) {
		createListQuery := fmt.Sprintf("UPDATE %s tl SET del=$1 WHERE tl.id=$2", findUsersTable)
		for i := 0; i < len(users_info); i++ {
			if _, err := tx.Exec(createListQuery, true, users_info[i].Id); err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
	}

	return nil
}

func (r *ChatListPostgres) FindThreeByHobby(list_users []int) ([]chat.UserHobby, error) {
	var lists []chat.UserHobby
	var prof_id_list []int
	var prof_id int

	query := fmt.Sprintf(`SELECT tl.profile_id FROM %s tl WHERE tl.user_id=$1`, usersProfileListsTable)
	for i := 0; i < 3; i++ {
		if err := r.db.Get(&prof_id, query, list_users[i]); err != nil {
			return lists, err
		}
		prof_id_list = append(prof_id_list, prof_id)
	}

	query = fmt.Sprintf(`SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id 
	WHERE ul.prof_id=$1 INTERSECT SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id 
	WHERE ul.prof_id=$2 INTERSECT SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id 
	WHERE ul.prof_id=$3`,
	userHobbyTable, usersHobbyListsTable, userHobbyTable, usersHobbyListsTable, userHobbyTable, usersHobbyListsTable)
	err := r.db.Select(&lists, query, prof_id_list[0], prof_id_list[1], prof_id_list[2])

	return lists, err
}

func (r *ChatListPostgres) FindTwoByHobby(list_users []int) ([]chat.UserHobby, error) {
	var lists []chat.UserHobby
	var prof_id_list []int
	var prof_id int

	query := fmt.Sprintf(`SELECT tl.profile_id FROM %s tl WHERE tl.user_id=$1`, usersProfileListsTable)
	for i := 0; i < 2; i++ {
		if err := r.db.Get(&prof_id, query, list_users[i]); err != nil {
			return lists, err
		}
		prof_id_list = append(prof_id_list, prof_id)
	}

	query = fmt.Sprintf(`SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id 
	WHERE ul.prof_id=$1 INTERSECT SELECT tl.id, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.userhobby_id 
	WHERE ul.prof_id=$2`,
	userHobbyTable, usersHobbyListsTable, userHobbyTable, usersHobbyListsTable)
	err := r.db.Select(&lists, query, prof_id_list[0], prof_id_list[1])

	return lists, err
}

func (r *ChatListPostgres) DeleteFindUsers(input chat.UsersForChat) error {
	query := fmt.Sprintf("DELETE FROM %s tl WHERE tl.user_id=$1 and tl.del=true", findUsersTable)
	for i := 0; i < len(input.UsersId); i++ {
		if _, err := r.db.Exec(query, input.UsersId[i]); err != nil {
			return err
		}
	}

	return nil
}

func (r *ChatListPostgres) GetUserByListId(listId int) ([]int, error) {
	var list_users_id []int

	query := fmt.Sprintf(`SELECT tl.user_id FROM %s tl WHERE tl.chatlists_id=$1`, usersChatListsTable)
	err := r.db.Select(&list_users_id, query, listId)

	return list_users_id, err
}

func (r *ChatListPostgres) GetUserAvatar(users_id []int) ([]string, error) {
	users_avatar := make([]string, len(users_id))

	query := fmt.Sprintf(`SELECT tl.photo FROM %s tl INNER JOIN %s ul on tl.id = ul.profile_id
	WHERE ul.user_id=$1`, usersProfileTable, usersProfileListsTable)
	for i := 0; i < len(users_id); i++ {
		if err := r.db.Get(&users_avatar[i], query, users_id[i]); err != nil {
			return users_avatar, err
		}
	}

	return users_avatar, nil
}