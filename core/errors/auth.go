package errors

import "github.com/gofiber/fiber/v2"

func AuthInputDataError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"errorCode": "1001001",
		"detail":    "입력하신 정보가 올바르지 않습니다",
		"error":     err,
	})
}

func AuthNotFoundUser(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"errorCode": "1001002",
		"detail":    "사용자 조회를 실패하였습니다",
		"error":     err,
	})
}

func AuthInvalidPassword(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"errorCode": "1001003",
		"detail":    "패스워드가 올바르지 않습니다",
		"error":     err,
	})
}

func AuthCanNotGenToken(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"errorCode": "1001004",
		"detail":    "토큰 발급을 실패하였습니다",
		"error":     err,
	})
}

func AuthCompletedPhoneVerified(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"errorCode": "1001005",
		"detail":    "이미 인증된 휴대폰 번호입니다",
		"error":     err,
	})
}

func AuthNotFoundPhone(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"errorCode": "1001006",
		"detail":    "유효하지 않은 휴대폰 정보 입니다",
		"error":     err,
	})
}

func AuthNotMatchedVerifiedCode(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
		"errorCode": "1001007",
		"detail":    "휴대폰 인증키가 일치하지 않습니다",
		"error":     err,
	})
}
