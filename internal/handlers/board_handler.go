package handler

import (
	model "example/todo-api/internal/models"
	service "example/todo-api/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	userScope *service.UserScopeService
}

func NewBoardController(uService *service.UserScopeService) *BoardController {
	return &BoardController{
		userScope: uService,
	}
}

func (tc *BoardController) getUserId(c *gin.Context) uint {
	return c.GetUint("userId")
}

func (tc *BoardController) forUser(c *gin.Context) *service.BoardService {
	userId := tc.getUserId(c)
	return tc.userScope.ForUserOwner(userId).Board()
}

// GetBoards godoc
//
//	@Summary		Get all boards
//	@Description	get boards
//	@Tags			Boards
//	@Produce		json
//	@Success		200	{array}		model.BoardModel
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/boards [get]
//
// @Security BearerAuth
func (tc *BoardController) GetBoards(c *gin.Context) {

	boards, err := tc.forUser(c).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": boards,
	})
}

// CreateBoards godoc
//
//	@Summary		Create board
//	@Description	Create board
//	@Tags			Boards
//	@Param			board	body	model.BoardCreateRequest	true	"Board data"
//	@Produce		json
//	@accept			json
//	@Success		200	{object}	model.BoardModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/boards [post]
//
// @Security BearerAuth
func (tc *BoardController) CreateBoards(c *gin.Context) {
	var board model.BoardCreateRequest

	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	createdBoard, err := tc.forUser(c).Create(&board)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": createdBoard,
	})
}

// GetBoardById godoc
//
//	@Summary		Get board by id
//	@Description	get board by id
//	@Param			id	path	int	true	"Board ID"
//	@Tags			Boards
//	@Produce		json
//	@Success		302	{object}	model.BoardModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/boards/:id [get]
//
// @Security BearerAuth
func (tc *BoardController) GetBoardById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return

	}

	board, err := tc.forUser(c).GetById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": board,
	})
}

// GetBoardByTitle godoc
//
//	@Summary		Get board by title
//	@Description	get board by title
//	@Param			title	path	string	true	"Board title"
//	@Tags			Boards
//	@Produce		json
//	@Success		200	{object}	model.BoardModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/boards/title/:title [get]
//
// @Security BearerAuth
func (tc *BoardController) GetBoardByTitle(c *gin.Context) {
	title := c.Param("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Title"})
		return

	}

	board, err := tc.forUser(c).GetByTitle(title)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": board,
	})
}

// UpdateBoard godoc
//
//	@Summary		Update board
//	@Description	update board
//	@Tags			Boards
//	@Param			board	body	model.BoardModel	true	"Board data"
//	@Produce		json
//	@accept			json
//	@Success		200	{object}	model.BoardModel
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/boards [put]
//
// @Security BearerAuth
func (tc *BoardController) UpdateBoard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	board, err := tc.forUser(c).GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := tc.forUser(c).Update(board); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": board,
	})
}

// DeleteBoard godoc
//
//	@Summary		Delete board by id
//	@Description	delete board by id
//	@Param			id	path	int	true	"Board ID"
//	@Tags			Boards
//	@Produce		json
//	@Success		200	{object}	model.BooleanResponse
//	@Failure		400	{object}	model.HTTPError	"Invalid request"
//	@Failure		500	{object}	model.HTTPError	"Internal server error"
//	@Router			/boards/:id [delete]
//
// @Security BearerAuth
func (tc *BoardController) DeleteBoard(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := tc.forUser(c).Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": true,
	})

}

func (tc *BoardController) AddUsers(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req struct {
		UserIDs []uint `json:"user_ids" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board, err := tc.forUser(c).AddUsers(uint(id), req.UserIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": board})
}
