package pay

import (
	"github.com/smartwalle/alipay/v3"
)

var Client *alipay.Client

const (
	kAppId = "9021000142652708"
	//私钥
	kPrivateKey = "MIIEowIBAAKCAQEAoWFzAsuV2DzI6hmnB9WgMj3GwFGkWxOffhDrTA5lyf6j7seJbHabh42MgbHEFeVpStqKC3CikvmnwYYEvPADsc0A7DQVwCPBzLqfeYGgC2laPi69rOG+hr7y0U+/p8w0XBsxTtjsjpoxRvfT7lEcpHXjZbNJyscUDwf7pgrSySzZJsw11QNXp1N2wnVwg+MIZfZUsilcKDz4fEd1TqhDPS+GyoHZTX+jfwOeqvZH9z06A7Uxys4akIBg6vHC0fhNj7UwXnz45vHEnJAeMRS7ZHxaikKUE4S0BGpKeuX6g4OIkGcHZqCVOkMKJiMhMAbwkAQklorXDLyBV506ZC0QswIDAQABAoIBAGeNhWREyJkhn/Z/kWt4i7vRYihj+ueqNsyJBMMf9fbgm8aLaUYc8X3QuVk7fUxjkeXDR5NBqkBPTHLkkUO/a0968V2fqllJWIELO2uXXuATsSF3kqRNkpkaC+t9lciRQwPbOw+SKHD0xiv4uJgSmiN5tfm7St2AUwG2KKcKCOmnBYgd1ou+i3c+67ShEsfhJ9zYTwkuxXUpEhEiOTImEN060eHG/rCpt5N9mDqeZexr2ZzRTErp4oDQMfwvvuxTs39mq+DS31C+2v6r0rgqcIFciiQQkF3ZM43gW0g0x4lItIeXBRKYh5YL5eXCO6BnOQ3xL99y38pWqsSbOnJrjwECgYEA0qA48uAUbsVxmaQum7rl4xLnJTN1gFbZGkFRgDXhiKe45BgG1GyPxgnSWNIRnPqMq/ugiRwP+YTiWvL32PUpZWS/+/OJV1m6mJrX5RINtcweBa3AXfRIdaNB3opXvOwIzUTaWrWAgTWeCsnYBSjzzFsqgp1Byj9Z3rFUjL1csDMCgYEAxCVo1VWoah3JngeSSkA56EKahjo5jkZqAMH4BOK5KNut4SI0ykgsYUhVDRXnrY6fN51ZHnBhdKdZ2ZmoQq3172+LmJwfVTEPoPOY/EnPT3/USwhJOrD2PYuIruY8oePQ3yYmeo6fmvUeSUvFXp8jAkeMtJReDPWSkv8spqDznYECgYBSds61MMxvLjiy+dgRutQk2pLLOuGHVPl9UROSygW6VkiKbWnHI5YK6G+FvQGOX35SG5uX8vfCLqfdCgHF3P7PJuoPwCMGoyfudbmPg0kA46DhgkytvGXeQQQaGDoPyq4LiCihmSxt6kstWCeOpaEGYq87IkzO3YzXJsC+takjhwKBgQCWSGdjFQc9juccOobedwknKGjGsTPKC69KN1O1QpVFddNqrE+wMM44Fzh7Fy50LtNUxC01AjvZKiPQckeWFz1Yn6lgWzYyiq2Dz0CHBHJfYfyhJI8e9dkk3JT0FJmeHDO71bojqsUk7+utku19CwbCx1lKPmc1HK6aDAMFBErDgQKBgETsxQy0L2jJ+G3lSLUHya0rBuyow23cQsvM+P3eHzs3JcsSXrJ3bAuX4dY5XYGJ7isdsPzvIKEOXfrWIgbXwJUd50cHhxILL/ixZaUAjny5ON2BpJfaMJ5P9oYs/Ucbq7J08+mZmb4nRq+um/ALqMB5f7jMrsgFnUZL6rT1/RSq"
	// TODO 设置回调地址域名
	kServerDomain = "http://hhyz5x.natappfree.cc"
)

func InitAlipay() {
	var err error
	Client, err = alipay.New(kAppId, kPrivateKey, false) // false表示使用沙盒环境
	if err != nil {
		panic(err)
	}

	// 加载证书
	if err = Client.LoadAppCertPublicKeyFromFile("service/pay/appPublicCert.crt"); err != nil {
		panic(err)
	}
	if err = Client.LoadAliPayRootCertFromFile("service/pay/alipayRootCert.crt"); err != nil {
		panic(err)
	}
	if err = Client.LoadAlipayCertPublicKeyFromFile("service/pay/alipayPublicCert.crt"); err != nil {
		panic(err)
	}
}

func GetServerDomain() string {
	return kServerDomain
}
