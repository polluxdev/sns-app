package usecases

import "github.com/polluxdev/sns-app/app/domain"

type PostRepository interface {
	FindAll() (domain.Posts, error)
}
