package routes

import (
	"chatrooms/gosrc/controllers"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func initPostsRoutes(gp fiber.Router, pc *controllers.PostsController) {
	gp.Post("/:roomId/posts", verifyAuth, postPosts(pc))
	gp.Get("/:roomId/posts", verifyAuth, getPosts(pc))
	gp.Use("/:roomId/ws", verifyAuth, wsPostsUpgrade, wsPosts(pc))
}

func postPosts(pc *controllers.PostsController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body struct {
			Content string `json:"content"`
		}
		err := c.BodyParser(&body)
		if err != nil {
			// TODO handle bad request
			return err
		}

		roomId := c.Params("roomId")
		userId := c.Locals("id").(string)

		if err := pc.CreatePost(userId, roomId, body.Content); err != nil {
			// TODO handle bad request
			return err
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getPosts(pc *controllers.PostsController) fiber.Handler {
	return func(c *fiber.Ctx) error {
		roomId := c.Params("roomId")

		posts, err := pc.ListPosts(roomId)
		if err != nil {
			// TODO handle bad request
			return err
		}

		return c.JSON(posts)
	}
}

func wsPostsUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

func wsPosts(pc *controllers.PostsController) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		roomId := c.Params("roomId")

		err := pc.SubscribeMessages(roomId, func(data []byte) error {
			return c.WriteMessage(websocket.TextMessage, data)
		})

		if err != nil {
			c.Close()
		}
	})
}
