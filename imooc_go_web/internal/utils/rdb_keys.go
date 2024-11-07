package utils

import "fmt"

func GetAuthKey(sessionID string) string {
	authKey := fmt.Sprintf("session_auth:#{sessionID}")
	return authKey
}

func GetSessionKey(userID string) string {
	sessionKey := fmt.Sprintf("session_key:#{userID}")
	return sessionKey
}
