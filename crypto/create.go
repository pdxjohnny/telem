package crypto

import (
	"bufio"
	"bytes"

	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"

	"github.com/pdxjohnny/telem/log"
)

func NewPair(name, comment, email string) (*openpgp.Entity, error) {
	entity, err := openpgp.NewEntity(name, comment, email, nil)
	if err != nil {
		log.PrintError("telem.crypto.NewPair", err)
		return nil, err
	}

	// Sign all the identities
	for _, id := range entity.Identities {
		err := id.SelfSignature.SignUserId(id.UserId.Id, entity.PrimaryKey, entity.PrivateKey, nil)
		if err != nil {
			log.PrintError("telem.crypto.NewPair", err)
			return nil, err
		}
	}
	return entity, nil
}

func Public(entity *openpgp.Entity) (string, error) {
	var output bytes.Buffer
	writeBuffer := bufio.NewWriter(&output)
	public, err := armor.Encode(writeBuffer, openpgp.PublicKeyType, nil)
	if err != nil {
		log.PrintError("telem.crypto.Public", err)
		return "", err
	}

	entity.Serialize(public)
	public.Close()
	writeBuffer.Flush()
	return output.String(), nil
}

func Private(entity *openpgp.Entity) (string, error) {
	var output bytes.Buffer
	writeBuffer := bufio.NewWriter(&output)
	private, err := armor.Encode(writeBuffer, openpgp.PrivateKeyType, nil)
	if err != nil {
		log.PrintError("telem.crypto.Private", err)
		return "", err
	}

	entity.SerializePrivate(private, nil)
	private.Close()
	writeBuffer.Flush()
	return output.String(), nil
}
