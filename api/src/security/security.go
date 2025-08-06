package security

import "golang.org/x/crypto/bcrypt"

//Receive a string and hash it
func Hash(passwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
}

// Validate passwds
func ValidyPasswd(hashedPasswd string, clearPasswd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(clearPasswd))

}
