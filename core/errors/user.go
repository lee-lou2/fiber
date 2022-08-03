package errors

import "github.com/gofiber/fiber/v2"

func UserInputDataError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errorCode": "1000001",
		"detail":    "입력하신 정보가 올바르지 않습니다",
		"error":     err,
	})
}

func UserExistUser(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"errorCode": "1000002",
		"detail":    "이미 존재하는 사용자입니다",
		"error":     err,
	})
}

func UserPasswordValidationError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"errorCode": "1000003",
		"detail":    "패스워드 양식이 올바르지 않습니다",
		"error":     err,
	})
}

func UserFailedHashingPassword(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"errorCode": "1000004",
		"detail":    "패스워드 해싱을 실패하였습니다",
		"error":     err,
	})
}

func UserFailedCreateUser(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"errorCode": "1000005",
		"detail":    "사용자 생성을 실패하였습니다",
		"error":     err,
	})
}

func UserNotFoundUUID(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"errorCode": "1000006",
		"detail":    "유효하지 않은 UUID 입니다",
		"error":     err,
	})
}

func UserNotFoundEmail(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"errorCode": "1000007",
		"detail":    "유효하지 않은 이메일 주소 입니다",
		"error":     err,
	})
}

func UserNotMatchedVerifiedCode(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"errorCode": "1000008",
		"detail":    "사용자 검증키가 일치하지 않습니다",
		"error":     err,
	})
}

func UserExistVerified(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"errorCode": "1000009",
		"detail":    "이미 검증된 사용자입니다",
		"error":     err,
	})
}

func UserNotMatchedChangePasswordCode(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"errorCode": "1000010",
		"detail":    "패스워드 검증키가 일치하지 않습니다",
		"error":     err,
	})
}
