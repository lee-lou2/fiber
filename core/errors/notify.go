package errors

import "github.com/gofiber/fiber/v2"

func NotifyCanNotSendEmail(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"errorCode": "1002001",
		"detail":    "이메일 발송을 실패하였습니다",
		"error":     err,
	})
}
