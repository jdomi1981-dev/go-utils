package utils

import (
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// HasWhiteSpaces indicate if a string has only white spaces
func HasWhiteSpaces(value string) bool {
	var invalidStr bool
	if len(strings.TrimSpace(value)) == 0 {
		invalidStr = true
	}
	return invalidStr
}

// SearchValue search a value inside a string
func SearchValue(evaluatedValue string, line string, separator string) bool {
	var vFind bool
	splitProperty := strings.SplitSeq(line, separator)

	for line := range splitProperty {
		if strings.EqualFold(evaluatedValue, line) {
			vFind = true
		}
	}
	return vFind
}

// GetStringsArray make a substring by separator and returns a array of strings
func GetStringsArray(line string, separator string) []string {
	splitProperty := strings.Split(line, separator)
	return splitProperty
}

// GetCurrentTimeMillis provides the current time in milliseconds.
func GetCurrentTimeMillis(timeProvider func() time.Time) int64 {
	return timeProvider().UnixNano() / 1e6
}

// NowPostgres returns a datetime string in 2006-01-02 15:04:05 format
func NowPostgres() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

func RemoveNonDigits(s string) string {
	re := regexp.MustCompile(`[^0-9]+`)
	return re.ReplaceAllString(s, "")
}

func CleanText(s string) string {
	// 1. Proteger ñ y Ñ antes de normalizar
	s = strings.ReplaceAll(s, "ñ", "__enie__")
	s = strings.ReplaceAll(s, "Ñ", "__ENIE__")

	// 2. Normalizar a NFD para separar letras y acentos
	s = norm.NFD.String(s)

	// 3. Eliminar marcas de acento (Mn)
	var b strings.Builder
	for _, r := range s {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}
	s = b.String()

	// 4. Restaurar ñ y Ñ
	s = strings.ReplaceAll(s, "__enie__", "ñ")
	s = strings.ReplaceAll(s, "__ENIE__", "Ñ")

	// 5. Eliminar todo lo que no sea letra o espacio
	re := regexp.MustCompile(`[^a-zA-ZñÑ\s]+`)
	s = re.ReplaceAllString(s, "")

	// 6. Limpiar espacios múltiples
	reSpaces := regexp.MustCompile(`\s+`)
	s = strings.TrimSpace(reSpaces.ReplaceAllString(s, " "))

	return s
}

func ValidEmail(email string) bool {
	// Expresión regular simple y práctica (RFC 5322 simplificada)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ValidDocumentId(input string) bool {
	// 1. Eliminar puntos como separadores de miles
	s := strings.ReplaceAll(input, ".", "")
	s = strings.TrimSpace(s)

	/*
		Formato general:
		- Cédula/DNI: 6–10 dígitos, opcional guion o dígito verificador
		- Brasil RG: 7–9 dígitos + dígito opcional
		- Uruguay CI: 7–8 dígitos + dígito verificador
		- México CURP: 18 caracteres alfanuméricos
		- Pasaportes LATAM & Caribe: 6–9 caracteres alfanuméricos, letras iniciales opcionales
	*/
	pattern := `(?i)^(` +
		`[0-9]{6,10}` + // cédula/DNI simple
		`|[0-9]{7,9}-?[0-9Xx]` + // Brasil / Uruguay con dígito verificador
		`|[A-Z]{1,2}[0-9]{6,9}` + // pasaportes latinoamericanos/caribeños
		`|[A-Z]{4}[0-9]{6}[A-Z0-9]{8}` + // México CURP (18 caracteres)
		`)$`

	re := regexp.MustCompile(pattern)
	return re.MatchString(s)
}
