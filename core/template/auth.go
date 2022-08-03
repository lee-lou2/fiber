package template

import "fmt"

func AuthVerifiedEmailTemplate(userUuid, verifiedCode string) string {
	return fmt.Sprintf(`아래 링크 클릭시 회원 가입이 완료됩니다<br>
<a href='http://localhost:3000/v1/auth/verified/email?uuid=%s&verifiedCode=%s'>인증하기</a>
`, userUuid, verifiedCode,
	)
}

func AuthChangePasswordEmailTemplate(userUuid, verifiedCode string) string {
	return fmt.Sprintf(`패스워드를 변경해주세요<br>
<a href='http://localhost:3000/v1/auth/verified/email?uuid=%s&verifiedCode=%s'>패스워드 변경하기</a>
`, userUuid, verifiedCode,
	)
}

func AuthVerifiedPhoneTemplate(verifiedCode string) string {
	return fmt.Sprintf(`[%s] 차빼주세요에서 전송한 인증 번호입니다`, verifiedCode)
}
