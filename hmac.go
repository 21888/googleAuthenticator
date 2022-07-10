
package googleAuthenticator

import (
	"crypto/hmac"
	"crypto/sha1"
)

func HmacSha1(key, data []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}
