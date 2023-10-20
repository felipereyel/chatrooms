package routes

import (
	"chatrooms/gosrc/controllers"

	"github.com/gofiber/fiber/v2"
)

func initPostsRoutes(gp fiber.Router, pc *controllers.PostsController) {
	gp.Post("/:roomId/posts", verifyAuth, postPosts(pc))
	gp.Get("/:roomId/posts", verifyAuth, getPosts(pc))
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

		post, err := pc.CreatePost(userId, roomId, body.Content)
		if err != nil {
			// TODO handle bad request
			return err
		}

		return c.JSON(post)
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
