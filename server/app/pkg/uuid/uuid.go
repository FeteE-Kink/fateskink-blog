package uuid

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateUUID() string {
	timestamp := time.Now().UnixNano()
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		timestamp>>32,
		timestamp&0xFFFF,
		uint16(randomBytes[0])<<8|uint16(randomBytes[1])&0x0FFF|0x4000,
		uint16(randomBytes[2]&0x3F|0x80)<<8|uint16(randomBytes[3]),
		randomBytes[4:])
}
