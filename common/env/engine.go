package env

import (
	"os"

	"apple/common/enum"
)

var engine string

func init() {
	engine = os.Getenv("ENGINE")
}

func GetEngineNum() string {
	return engine
}

func IsEngineNum1() bool {
	return engine == enum.ENGINE_NAME_1
}
