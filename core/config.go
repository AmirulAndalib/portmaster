package core

import (
	"flag"

	"github.com/safing/portbase/config"
	"github.com/safing/portbase/log"
)

var (
	devMode        config.BoolOption
	defaultDevMode bool
)

func init() {
	flag.BoolVar(&defaultDevMode, "devmode", false, "force development mode")
}

func logFlagOverrides() {
	if defaultDevMode {
		log.Warning("core: core/devMode default config is being forced by -devmode flag")
	}
}

func registerConfig() error {
	err := config.Register(&config.Option{
		Name:           "Development Mode",
		Key:            "core/devMode",
		Description:    "In Development Mode security restrictions are lifted/softened to enable easier access to Portmaster for debugging and testing purposes. This is potentially very insecure, only activate if you know what you are doing.",
		ExpertiseLevel: config.ExpertiseLevelDeveloper,
		OptType:        config.OptTypeBool,
		DefaultValue:   defaultDevMode,
	})
	if err != nil {
		return err
	}

	return nil
}