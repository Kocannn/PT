package helper

import "golang.org/x/crypto/bcrypt"

func HassPass(pass string) (string, error) {
	p := []byte(pass)
	hashedPassword, err := bcrypt.GenerateFromPassword(p, 12)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


func ComparePass(hass, pass []byte) (bool, error){

  err := bcrypt.CompareHashAndPassword(hass, pass)
  if err != nil {
    return false, err
  }
  return true, nil
}