package usecases

import "strings"


func (uc *UserUseCase) nameValidate(name string) bool {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "  ", " ")

	
	if name == "" {
		return false
	}
	
	if len(name) < 3 || len(name) > 255 {
		return false
	}

	return true
}