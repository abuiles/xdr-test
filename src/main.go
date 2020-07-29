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
	log.Println(base)
	decoded, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(decoded)

	a := AccountEntryExtensionV2{}
	err = a.UnmarshalBinary(decoded)
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range a.SignerSponsoringIDs {
		if s == nil {
			log.Printf("Signer %d is not sponsored", i)
		} else {
			log.Printf("Signer %d is sponsored by %d", i, *s.Ed25519)
		}
	}

	if a.SponsoringId == nil {
		log.Println("entry is not sponsored")
	} else {
		log.Println("entry is sponsored")
	}

	b, err := a.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(base64.StdEncoding.EncodeToString(b))
}
