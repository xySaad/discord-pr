package github

import (
	"crypto/hmac"
	"crypto/sha256"
	"discord-pr/bot"
	"discord-pr/config"
	"discord-pr/mux"
	"encoding/hex"
	"net/http"
)

func GithubSignature(w http.ResponseWriter, r *mux.Request, bot *bot.Bot) bool {
	h := r.Header.Get
	// Compute HMAC SHA256 of body
	mac := hmac.New(sha256.New, []byte(config.WEBHOOK_SECRET))
	mac.Write(r.Body())
	expectedSignature := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	// Compare with GitHub signature
	if !hmac.Equal([]byte(expectedSignature), []byte(h("X-Hub-Signature-256"))) {
		http.Error(w, "invalid signature", http.StatusUnauthorized)
		return false
	}
	return true
}
