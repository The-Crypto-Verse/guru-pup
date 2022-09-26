package lib

import (
	"os"
	"strings"
)

// AtomicHub collection name
const COLLECTION_NAME = "cryptopuppyz"

// environment variables
var (
	TOKEN            = os.Getenv("TOKEN")
	DETA_PROJECT_KEY = os.Getenv("DETA_PROJECT_KEY")
	GUILDS           = strings.Split(os.Getenv("GUILDS"), ",")
)
