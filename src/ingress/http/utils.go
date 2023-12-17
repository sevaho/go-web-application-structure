package http

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/sevaho/gowas/src/logger"
)

func RecoverWebsocketHandler(conn *websocket.Conn) {
	if err := recover(); err != nil {
		logger.Logger.Error().Msgf("Something went wrong while establishing a websocket connection. %v", err)
		conn.WriteJSON(fiber.Map{"customError": "error occurred"})
	}
}
