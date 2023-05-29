package main

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/password/validate", validatePasswordHandler)

	r.Run()
}

type PasswordRequest struct {
	Password string `json:"password"`
}

func validatePasswordHandler(c *gin.Context) {
	var req PasswordRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	r := ValidatePassword(req.Password)
	c.JSON(http.StatusOK, gin.H{"passwordStrength": r})
}

func ValidatePassword(in string) string {
	l := len(in)
	hasTwo := hasTwoNumberAndLetter(in)

	if l >= 8 {
		if hasTwo {
			return "EXCELENTE"
		} else {
			return "BUENA"
		}
	} else if hasTwo {
		return "BUENA"
	}

	return "DEBIL"
}

func hasTwoNumberAndLetter(in string) bool {
	numberCount := 0
	letterCount := 0

	for _, c := range in {

		if unicode.IsDigit(c) {
			numberCount++
		} else if unicode.IsLetter(c) {
			letterCount++
		}

		if numberCount >= 2 && letterCount >= 1 {
			return true
		}
	}

	return false
}
