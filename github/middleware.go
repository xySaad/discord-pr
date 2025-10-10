package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"discord-pr/config"
	"encoding/hex"
	"net/http"

	"github.com/xySaad/gocord"
	"github.com/xySaad/trail"
)

func GithubSignature(c *trail.Context[*gocord.Bot]) bool {
	// Compute HMAC SHA256 of body
	mac := hmac.New(sha256.New, []byte(config.WEBHOOK_SECRET))
	mac.Write(c.BodyNoErr())
	expectedSignature := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	// Compare with GitHub signature
	if !hmac.Equal([]byte(expectedSignature), []byte(c.Header("X-Hub-Signature-256"))) {
		http.Error(c.Response, "invalid signature", http.StatusUnauthorized)
		return false
	}
	return true
}
