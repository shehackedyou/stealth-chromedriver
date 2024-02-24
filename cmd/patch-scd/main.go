package main

import (
	"fmt"
	"log"

	chromedriver "github.com/Davincible/go-undetected-chromedriver"
)

func main() {
	fmt.Println("stealth chromedriver patcher")
	fmt.Println("============================")

	log.Printf("DOWNLOADING chrome driver and extracting it from zip archive...")

	if bPath, err := chromedriver.Download(); err == nil {
		log.Printf("succeeded in downloading and unzipping: %v", bPath)
		log.Printf("Its now downloaded, now we need to patch the fucker with the NEW methodology...")
		log.Printf("and it works... but it left the thingy in /tmp/ we should obvio patch then mv it\n")
		log.Printf("PATCHING to obstuficate the cdc javascript objects...")
		b, err := chromedriver.Patch(bPath)
		if err != nil {
			log.Printf("failed to patch the binary: %v", err)
		} else {
			log.Printf("successfully patched binary then moved to: %v", b.DataDir)
		}
	} else {
		log.Printf("failed to download and unzip: %v", err)
	}
}
