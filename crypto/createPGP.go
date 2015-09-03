package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func main() {
	var e *openpgp.Entity
	e, err := openpgp.NewEntity("itis", "test", "itis@itis3.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add more identities here if you wish

	// Sign all the identities
	for _, id := range e.Identities {
		err := id.SelfSignature.SignUserId(id.UserId.Id, e.PrimaryKey, e.PrivateKey, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	public, err := armor.Encode(os.Stdout, openpgp.PublicKeyType, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.Serialize(public)
	public.Close()
	fmt.Printf("\n")

	private, err := armor.Encode(os.Stdout, openpgp.PrivateKeyType, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.SerializePrivate(private, nil)
	private.Close()
	fmt.Printf("\n")
}
