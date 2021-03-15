package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	model "myapp/pkg/model"
)

type Message struct {
	db           *sql.DB
	messageTable *model.MessageTable
}

func NewMessage(db *sql.DB) *Message {
	return &Message{db, &model.MessageTable{}}
}

func (m *Message) HandleGet(c echo.Context) error {
	messageRecords, err := m.messageTable.SelectAll(c.Request().Context(), m.db)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, messageRecords)
}

func (m *Message) HandlePost(c echo.Context) error {
	messageRecord := &model.MessageRecord{}
	if err := c.Bind(messageRecord); err != nil {
		return err
	}

	err := m.messageTable.Insert(c.Request().Context(), m.db, messageRecord)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, messageRecord)
}

func (m *Message) HandlePut(c echo.Context) error {
	messageRecord := &model.MessageRecord{}
	if err := c.Bind(messageRecord); err != nil {
		return err
	}

	err := m.messageTable.Update(c.Request().Context(), m.db, messageRecord)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, messageRecord)
}

func (m *Message) HandleDelete(c echo.Context) error {
	messageRecord := &model.MessageRecord{}
	if err := c.Bind(messageRecord); err != nil {
		return err
	}

	err := m.messageTable.Delete(c.Request().Context(), m.db, messageRecord)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, messageRecord)
}
