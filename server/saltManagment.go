package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// yes it does say a$$word, as that is the optimal password
//
// This salt is simple and only requires 1 hash
func normalSalt(assword string) (salt string) {
	// gets the margin
	beginning := int(assword[0]) % len(assword)
	end := int(assword[len(assword)-1]) % len(assword)
	hasher := sha256.New()
	hasher.Write([]byte(assword[beginning:end]))
	return hex.EncodeToString(hasher.Sum(nil))
}
// hashes the user
func preHash(luser string) string {
	hasher := sha256.New()
	hasher.Write([]byte(luser))
	luser = hex.EncodeToString(hasher.Sum(nil))
	return luser
}

func getSalt(luser string) (salt string) {
	luserSalt, err := readJson("salty.json")
	if err != nil {
		return normalSalt(luser)
	}
	salt, exists := luserSalt[luser]
	if !exists {
		return normalSalt(luser)
	} else {
		return salt
	}
}

// adds the new user (password properly salted) to the luserInfo (userInfo)
func makeNewUser(salter map[string]string, luserInfo map[string]string, luser string, hashedPass string) (newSalter map[string]string, newLuserInfo map[string]string, err error) {
	// 80% of user checks will be done outside of this function before sending. but this will just check if it is inside the current architecture. 
	// auth flow. get client hash (however made) salt hash, then check if salted hash matches existing hash
	// salt has
	temp_salt := getSalt(hashedPass)
	tempPass := hashedPass + temp_salt
	hasher := sha256.New()
	hasher.Write([]byte(tempPass))
	tempPass = hex.EncodeToString(hasher.Sum(nil))
	// add hash to the users and salt configs
	if _,ok := luserInfo[luser]; !ok {
		luserInfo[luser] = tempPass
	} else {
		return salter, luserInfo, errors.New("Invalid username or password")
	}
	return salter, luserInfo, nil
}

