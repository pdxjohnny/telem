package crypto

var ConfigOptions = map[string]interface{}{
	"name": map[string]interface{}{
		"value": "First Last",
		"help":  "Name for the key pair",
	},
	"comment": map[string]interface{}{
		"value": "",
		"help":  "Comment for the key pair",
	},
	"email": map[string]interface{}{
		"value": "user@example.com",
		"help":  "Email for the key pair",
	},
	"stdout": map[string]interface{}{
		"value": false,
		"help":  "Print keys to stdout",
	},
	"path": map[string]interface{}{
		"value": "keys/crypto/",
		"help":  "Where generated keys are writen to",
	},
	"gen": map[string]interface{}{
		"value": false,
		"help":  "Generate a key pair",
	},
}
