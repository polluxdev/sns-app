package interfaces

import (
	"github.com/polluxdev/sns-app/app/domain"
)

type PostRepository struct {
	SQLHandler SQLHandler
}

func (pr *PostRepository) FindAll() (posts domain.Posts, err error) {
	const query = `
	SELECT
		id,
		user_id,
		title
	FROM
		posts
	`

	rows, err := pr.SQLHandler.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id int8
		var userID int8
		var title string
		if err = rows.Scan(&id, &userID, &title); err != nil {
			return
		}
		post := domain.Post{
			ID:     id,
			Title:  title,
			UserID: userID,
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return
	}

	return
}

func (pr *PostRepository) Save(p domain.Post) (id int64, err error) {
	tx, err := pr.SQLHandler.Begin()
	if err != nil {
		return
	}

	const query = `
		INSERT INTO
			posts(user_id, body)
		VALUES
			(?, ?)
	`

	result, err := tx.Exec(query, p.Title, p.UserID)
	if err != nil {
		_ = tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	id, err = result.LastInsertId()
	if err != nil {
		return id, nil
	}

	return
}

func (pr *PostRepository) DeleteByID(postID int) (err error) {
	const query = `
		DELETE
		FROM
			posts
		WHERE
			id = ?
	`

	_, err = pr.SQLHandler.Exec(query, postID)
	if err != nil {
		return
	}

	return
}
