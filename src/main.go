package main

import (
	"encoding/base64"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("node", "index.js")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	base := string(out)
	// var account AccountEntryExtensionV2

	log.Println(fmtTest)
	decoded, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		log.Fatal(err)
	}

	a := AccountEntryExtensionV2{}
	err = a.UnmarshalBinary(decoded)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(a)
}
