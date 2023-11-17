package usecases

import (
	"regexp"
	"strings"
)

// Tambem preciso formatar o nome real do usario para que ele seja salvo no banco de dados formatado
// nao acho que isso seja um problema, pois o nome real so sera salvo se atender todos os criterios de validacao
func nameInvalid(name string) bool {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "  ", " ")
	// regex para substituir mais de um espaço por um espaço só
	name = regexp.MustCompile(`/( )+/g`).ReplaceAllString(name, " ")


	if name == "" {
		return true
	}
	
	if len(name) < 4 || len(name) > 255 {
		return true
	}

	return false
}