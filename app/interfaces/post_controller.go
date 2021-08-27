package interfaces

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/polluxdev/sns-app/app/usecases"
)

type PostController struct {
	PostInteractor usecases.PostInteractor
	Logger         usecases.Logger
}

func NewPostController(sqlHandler SQLHandler, logger usecases.Logger) *PostController {
	return &PostController{
		PostInteractor: usecases.PostInteractor{
			PostRepository: &PostRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

func (pc *PostController) Index(c echo.Context) error {
	pc.Logger.LogAccess("%s %s %s\n", c.Request().RemoteAddr, c.Request().Method, c.Request().URL)

	_, err := pc.PostInteractor.Index()
	if err != nil {
		pc.Logger.LogError("%s", err)
		c.String(http.StatusBadRequest, "ERROR")
	}

	return c.String(http.StatusOK, "Post List")
}
