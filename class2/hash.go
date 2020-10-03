package class2

import (
	"encoding/hex"
	"crypto/sha256"
)
//hash takes string, returns string and error

func Hash(text string)(string,error){
	
	hash:=sha256.New()
	_, err := hash.Write([]byte(text)) // _ means ignoring a variable
	
	if err != nil {
		return "", err;
	}
	hashedBytes := hash.Sum(nil)

	return hex.EncodeToString(hashedBytes), nil;
}