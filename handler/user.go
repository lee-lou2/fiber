package handler

import (
	"strconv"
	"time"

	"api/cache"
	"api/core/errors"
	"api/core/template"
	"api/core/utils"
	"api/database"
	"api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	// 패스워드 해싱
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validatePassword(c *fiber.Ctx, password string) error {
	// 패스워드 유효성 검사
	if password == "" {
		return errors.UserPasswordValidationError(c, nil)
	}
	return nil
}

func UserCreate(c *fiber.Ctx) error {
	/*
		사용자 생성
	*/
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var requestBody RequestBody
	var user model.User

	// 데이터 파싱
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.UserInputDataError(c, err)
	}

	// 존재 여부 확인
	db := database.DB
	db.Where("email = ?", requestBody.Email).First(&user)
	if user.ID != 0 {
		return errors.UserExistUser(c, nil)
	}

	// 패스워드 유효성 검사
	if err := validatePassword(c, requestBody.Password); err != nil {
		return err
	}

	// 패스워드 해싱
	hash, err := hashPassword(requestBody.Password)
	if err != nil {
		return errors.UserFailedHashingPassword(c, err)
	}

	// 데이터 생성
	user.Email = requestBody.Email
	user.Password = hash
	user.UUID = generateUuid()
	if err := db.Create(&user).Error; err != nil {
		return errors.UserFailedCreateUser(c, err)
	}

	// 인증키 저장
	verifiedCode := utils.RandStringBytes(8, utils.LetterBytes)
	cache.MemoryCache.Set(
		"VERIFIED_EMAIL_CODE_"+strconv.Itoa(int(user.ID)),
		verifiedCode,
		5*time.Minute,
	)

	// 이메일 발송
	NotifySendEmail(
		user.Email,
		"회원 가입 완료",
		template.AuthVerifiedEmailTemplate(
			user.UUID,
			verifiedCode,
		),
	)

	return c.JSON(fiber.Map{"email": user.Email})
}

func generateUuid() string {
	/*
		UUID 생성
	*/
	var user model.User

	db := database.DB
	newUuid := uuid.NewString()
	if db.Where(&model.User{UUID: newUuid}).Find(&user); user.ID != 0 {
		newUuid = generateUuid()
	}
	return newUuid
}

func UserSendChangePasswordEmail(c *fiber.Ctx) error {
	/*
		패스워드를 잊어버린 경우 해당 이메일로 패스워드 변경 메일 전송
	*/
	type RequestBody struct {
		Email string `json:"email"`
	}
	var requestBody RequestBody
	var user model.User

	// 데이터 파싱
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.UserInputDataError(c, err)
	}

	// 사용자가 존재하는지 확인
	db := database.DB
	db.Where("email = ?", requestBody.Email).First(&user)
	if user.ID == 0 {
		return errors.UserNotFoundEmail(c, nil)
	}

	// 인증키 저장
	verifiedCode := utils.RandStringBytes(8, utils.LetterBytes)
	cache.MemoryCache.Set(
		"CHANGE_PASSWORD_CODE_"+strconv.Itoa(int(user.ID)),
		verifiedCode,
		5*time.Minute,
	)

	// 이메일 발송
	NotifySendEmail(
		user.Email,
		"[차빼주세요] 패스워드 변경",
		template.AuthChangePasswordEmailTemplate(
			user.UUID,
			verifiedCode,
		),
	)

	return c.JSON(fiber.Map{"email": user.Email})
}

func UserChangePassword(c *fiber.Ctx) error {
	/*
		패스워드 변경
	*/
	type RequestBody struct {
		UUID     string `json:"uuid"`
		Password string `json:"password"`
		Code     string `json:"code"`
	}
	var requestBody RequestBody
	var user model.User

	// 데이터 파싱
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.UserInputDataError(c, err)
	}

	// 사용자가 존재하는지 확인
	db := database.DB
	db.Where("uuid = ?", requestBody.UUID).First(&user)
	if user.ID == 0 {
		return errors.UserNotFoundUUID(c, nil)
	}

	// 코드 유효성 검사
	cacheCode, _ := cache.MemoryCache.Get(
		"CHANGE_PASSWORD_CODE_" + strconv.Itoa(int(user.ID)),
	)
	if requestBody.Code != cacheCode.(string) {
		return errors.UserNotMatchedChangePasswordCode(c, nil)
	}

	// 패스워드 유효성 검사
	if err := validatePassword(c, requestBody.Password); err != nil {
		return err
	}

	// 패스워드 해싱
	hash, err := hashPassword(requestBody.Password)
	if err != nil {
		return errors.UserFailedHashingPassword(c, err)
	}

	// 데이터 저장
	user.Password = hash
	db.Save(&user)

	return c.JSON(fiber.Map{"isCompleted": true})
}
