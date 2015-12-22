package monito

import (
	"fmt"
	"io"

	"github.com/PI-Victor/monito/pkg"
	"github.com/PI-Victor/monito/pkg/log"
)

// MainService - Main monitoring service core
type MainService struct {
	ConfigFile string
	output     io.Writer
	MainNode   bool
}

//MainService validate the loading of assets.
func (m *MainService) loadService() error {
	configFile, err := util.ReadConfigFile(m.ConfigFile)
	if err != nil {
		return err
	}
	fmt.Println(configFile)
	return nil
}

// Start - Starts the main service
func (m *MainService) Start() {
	if err := m.loadService(); err != nil {
		log.Panic("Failed to start monito... ")
	}

}

//CheckConfig - validates the configuration file
func (m *MainService) CheckConfig() {
	fmt.Printf("Checking configuration, %s", m.ConfigFile)
}

//LoadAssets - load and test configured components for runtime
func (m *MainService) LoadAssets() {

}
