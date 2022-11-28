package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "chonjiang.com"

func GetEncrypt(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum(data))
}
