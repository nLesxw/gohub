package bootstrap

import (
	"gohub/pkg/config"
	"gohub/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {

	logger.InitLogger(
		config.GetStirng("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetStirng("log.type"),
		config.GetStirng("log.level"),
	)
}