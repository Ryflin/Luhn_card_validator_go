package client

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)




// same as crazy_salt but uses slightly different means to accomplish it making in truly one directional
// 
// not performant probably should not use
func crazySalt2(assword string) (salt string) {
	var step = 3
	var cheek = []string{}
	for i := 0; i < len(assword); i+= step {
		cheek = append(cheek, assword[i:i+step])
	}

	hasher := sha256.New()
	for i := 0; i < len(cheek); i++ {
		hasher.Write([]byte(cheek[i]))
		cheeky_word := hex.EncodeToString(hasher.Sum(nil))
		hasher.Reset()
		for j := 0; j < len(cheek[i]); j++ {
			index := int(cheeky_word[j])
			index = index % len(cheeky_word)
			cheek[i] = strings.Join([]string{cheeky_word[0:index],cheeky_word[index:]}, string(cheek[i][j]) )
		}
		hasher.Write([]byte(cheek[i]))
		cheek[i] = hex.EncodeToString(hasher.Sum(nil))
	}
	cheeky := strings.Join(cheek, "")
	hasher.Write([]byte(cheeky))
	return hex.EncodeToString(hasher.Sum(nil))
}
