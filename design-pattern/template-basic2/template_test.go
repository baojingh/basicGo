package template_basic1

import "testing"

/**
  Author     : He Bao Jing
  Date       : 5/10/2024 9:27 AM
  Description:
*/

func TestTemplateMethod(t *testing.T) {
	sms := &Sms{}
	otp := Otp{iOtp: sms}
	otp.genAndSendOTP(4)

	email := &Email{}
	otp2 := Otp{iOtp: email}
	otp2.genAndSendOTP(4)

}
