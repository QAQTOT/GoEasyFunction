package cipher

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func HmacSHA256Encode(key, message string) string {

	// 创建一个新的HMAC对象，使用SHA256哈希算法
	h := hmac.New(sha256.New, []byte(key))

	// 写入要签名的数据
	h.Write([]byte(message))

	// 计算HMAC-SHA256摘要（即哈希值）
	hashed := h.Sum(nil)

	// 将哈希值转换为Base64编码以便于显示或存储
	base64Hash := base64.StdEncoding.EncodeToString(hashed)

	return base64Hash
}
