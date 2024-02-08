package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	// Define flags
	length := flag.Int("l", 6, "Length of mOTP output (default 6 characters)")
	window := flag.Int64("w", 0, "Number of counter values before and after current one to show (for testing time-skew)")

	// Parse the flags
	flag.Parse()

	// Remaining arguments after flags are parsed
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("Missing required arguments: secret and pin")
		return
	}

	secret := args[0]
	pin := args[1]

	epochTime := time.Now().Unix()
	counter := epochTime / 10

	for i := -*window; i <= *window; i++ {
		data := fmt.Sprintf("%d%s%s", counter+i, secret, pin)

		hash := md5.Sum([]byte(data))

		code := hex.EncodeToString(hash[:])

		if *length < len(code) {
			code = code[:*length]
		}
        fmt.Println(code)
	}

}
