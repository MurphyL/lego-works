package iam

import (
	"sort"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/MurphyL/lego-works/pkg/iam/internal/login"
)

func GetHashParts(acc login.Account) []byte {
	parts := []string{acc.GetLoginUsername(), acc.GetLoginPassword()}
	sort.Strings(parts)
	return []byte(strings.Join(parts, "-"))
}

func GetHashPassword(acc login.Account) string {
	combined := GetHashParts(acc)
	cipher, _ := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	return string(cipher)
}

func CompareHashPassword(hash string, acc login.Account) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), GetHashParts(acc)) == nil
}
