package signs

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"mobileSign/util/https"
)


var (
	pub *rsa.PublicKey
)

type PublicKeyRes struct {
	Status string `json:"status"`
	Data   struct {
		Pub string `json:"pub"`
	} `json:"data"`
}

func CheckToken(payload, signature string) bool {
	pl, _ := base64.StdEncoding.DecodeString(payload)
	tmp, _ := base64.StdEncoding.DecodeString(signature)

	return sign(pub, pl, tmp)
}

func sign(pubkey *rsa.PublicKey, payload, signature []byte) (ok bool) {

	h := sha256.New()
	h.Write(payload)
	hash := h.Sum(nil)

	err := rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hash, signature)
	if err != nil {
		refreshPub()
		err := rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hash, signature)
		return err == nil
	}
	return true
}

func refreshPub() {
	res, _ := https.SendGet("https://wx.redrock.team/magicloop/keycenter/public")
	key := PublicKeyRes{}
	_ = json.Unmarshal(res, &key)
	block, _ := pem.Decode([]byte(key.Data.Pub))
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub = pubInterface.(*rsa.PublicKey)
}

func init() {
	refreshPub()
}

