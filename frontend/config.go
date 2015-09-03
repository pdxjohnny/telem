package frontend

var ConfigOptions = map[string]interface{}{
	"address": map[string]interface{}{
		"value": "0.0.0.0",
		"help":  "Address  to bind to",
	},
	"port": map[string]interface{}{
		"value": 25000,
		"help":  "Port to listen on",
	},
	"cert": map[string]interface{}{
		"value": "keys/frontend/cert.pem",
		"help":  "Cert file to use for webserver",
	},
	"key": map[string]interface{}{
		"value": "keys/frontend/key.pem",
		"help":  "Key file to use for webserver",
	},
	"upload": map[string]interface{}{
		"value": "tmp/frontend/uploaded/",
		"help":  "Where uploaded files are stored",
	},
}
