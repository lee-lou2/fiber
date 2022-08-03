package handler

import (
	"car/cache"
	"car/conf"
	"car/core/errors"
	"car/core/template"
	"car/core/utils"
	"car/database"
	"car/model"
	errorDetail "errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func CheckPasswordHash(password, hash string) error {
	// 패스워드 확인
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func getUserByEmail(email string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Email: email}).Find(&user).Error; err != nil {
		if errorDetail.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func Login(c *fiber.Ctx) error {
	type requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input requestBody

	// 데이터 변환
	if err := c.BodyParser(&input); err != nil {
		return errors.AuthInputDataError(c, err)
	}
	email := input.Email
	pass := input.Password

	// 사용자 조회
	user, err := getUserByEmail(email)
	if err != nil || user.ID == 0 {
		return errors.AuthNotFoundUser(c, err)
	}

	// 패스워드 확인
	if err := CheckPasswordHash(pass, user.Password); err != nil {
		return errors.AuthInvalidPassword(c, nil)
	}

	// 토큰 발급
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	accessToken, err := token.SignedString([]byte(conf.Config("JWT_SECRET_KEY")))
	if err != nil {
		return errors.AuthCanNotGenToken(c, err)
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"token":   accessToken,
	})
}

func AuthVerifiedEmail(c *fiber.Ctx) error {
	/*
		사용자 이메일 검증
	*/
	var user model.User

	// 데이터 조회
	userUuid := c.Query("uuid")
	userCode := c.Query("verifiedCode")

	// 사용자가 존재하는지 확인
	db := database.DB
	db.Where("uuid = ?", userUuid).First(&user)
	if user.ID == 0 {
		return errors.UserNotFoundUUID(c, nil)
	}

	// 검증 여부
	if user.IsVerified {
		return errors.UserExistVerified(c, nil)
	}

	cacheVerifiedCode, isSuccess := cache.MemoryCache.Get(
		"VERIFIED_EMAIL_CODE_" + strconv.Itoa(int(user.ID)),
	)
	if !isSuccess {
		return c.JSON(fiber.Map{"isVerified": false})
	}

	if userCode != cacheVerifiedCode.(string) {
		return errors.UserNotMatchedVerifiedCode(c, nil)
	}

	// 인증 완료
	user.IsVerified = true
	db.Save(&user)
	return c.JSON(fiber.Map{"isVerified": true})
}

func AuthReVerifiedEmail(c *fiber.Ctx) error {
	/*
		로그인 했지만 아직 이메일 인증이 안된 경우
	*/
	// 이메일 주소 조회
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

	// 검증 여부
	if user.IsVerified {
		return errors.UserExistVerified(c, nil)
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
		"[차빼주세요] 이메일 재검증",
		template.AuthVerifiedEmailTemplate(
			user.UUID,
			verifiedCode,
		),
	)

	return c.JSON(fiber.Map{"email": user.Email})
}

func AuthVerifiedPhone(c *fiber.Ctx) error {
	type RequestBody struct {
		Phone string `json:"phone"`
	}
	var requestBody RequestBody

	// 데이터 파싱
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.UserInputDataError(c, err)
	}

	var phone model.Phone

	// 로그인된 사용자 불러오기

	// 휴대폰 정보가 있으면 가져오고 아니면 만들기
	db := database.DB
	db.FirstOrCreate(&phone, model.Phone{Number: requestBody.Phone})

	// 이미 인증된 경우 오류 반환
	if phone.IsVerified {
		return errors.AuthCompletedPhoneVerified(c, nil)
	}

	// UUID 저장
	if phone.UUID == "" {
		phone.UUID = generateUuid()
		db.Save(&phone)
	}

	// 인증키 저장
	verifiedCode := utils.RandStringBytes(6, utils.IntBytes)
	cache.MemoryCache.Set(
		"VERIFIED_PHONE_CODE_"+strconv.Itoa(int(phone.ID)),
		verifiedCode,
		5*time.Minute,
	)

	// 문자 발송
	NotifySendSMS(
		phone.Number,
		template.AuthVerifiedPhoneTemplate(
			verifiedCode,
		),
	)

	return c.JSON(fiber.Map{"phoneUuid": phone.UUID})
}

func AuthVerifiedPhoneCode(c *fiber.Ctx) error {
	/*
		휴대폰 인증
	*/
	type RequestBody struct {
		UUID string `json:"uuid"`
		Code string `json:"code"`
	}
	var requestBody RequestBody

	// 데이터 파싱
	if err := c.BodyParser(&requestBody); err != nil {
		return errors.UserInputDataError(c, err)
	}

	var phone model.Phone

	// 휴대폰 정보가 존재하는지 확인
	db := database.DB
	db.Where("uuid = ?", requestBody.UUID).First(&phone)
	if phone.ID == 0 {
		return errors.AuthNotFoundPhone(c, nil)
	}
	// 이미 인증된 경우 오류 반환
	if phone.IsVerified {
		return errors.AuthCompletedPhoneVerified(c, nil)
	}
	cacheVerifiedCode, isSuccess := cache.MemoryCache.Get(
		"VERIFIED_PHONE_CODE_" + strconv.Itoa(int(phone.ID)),
	)
	if !isSuccess {
		return c.JSON(fiber.Map{"isVerified": false})
	}

	if requestBody.Code != cacheVerifiedCode.(string) {
		return errors.AuthNotMatchedVerifiedCode(c, nil)
	}
	return c.JSON(fiber.Map{"isVerified": true})
}
