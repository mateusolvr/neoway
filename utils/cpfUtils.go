package utils

import (
	"regexp"
	"strconv"
)

// CPFUtils : Valida e higieniza o número de CPF
func CPFUtils(cpf string) (string, bool) {

	clean := regexp.MustCompile(`[^\d]`)
	cleanCPF := clean.ReplaceAllString(cpf, "")

	// Retorn falso se o CPF não possuir 11 dígitos
	if len(cleanCPF) != 11 {
		return cleanCPF, false
	}

	// Conferência do primeiro dígito verificador
	sumFirstDigit := 0
	for j := 0; j < 9; j++ {
		digit, _ := strconv.Atoi(cleanCPF[j : j+1])
		sumFirstDigit = sumFirstDigit + digit*(10-j)
	}

	firstVerifyingDigit, _ := strconv.Atoi(cleanCPF[9:10])
	firstCheck := ((sumFirstDigit * 10) % 11) == firstVerifyingDigit

	if firstCheck == false {
		return cleanCPF, false
	}

	// Conferência do segundo dígito verificador
	sumSecDigit := 0
	for j := 0; j < 10; j++ {
		digit, _ := strconv.Atoi(cleanCPF[j : j+1])
		sumSecDigit = sumSecDigit + digit*(11-j)
	}

	secVerifyingDigit, _ := strconv.Atoi(cleanCPF[10:11])
	secCheck := ((sumSecDigit * 10) % 11) == secVerifyingDigit

	if secCheck == false {
		return cleanCPF, false
	}

	// Conferência dígitos todos iguais
	actualDigit := cleanCPF[0:1]
	actualCheck := true

	for i := 1; i < len(cleanCPF); i++ {
		// Compara cada um dos dígitos do CPF
		currentDigit := cleanCPF[i : i+1]
		currentCheck := currentDigit == actualDigit
		actualCheck = currentCheck && actualCheck
		actualDigit = currentDigit
	}

	if actualCheck == true {
		return cleanCPF, false
	}

	return cleanCPF, true

}