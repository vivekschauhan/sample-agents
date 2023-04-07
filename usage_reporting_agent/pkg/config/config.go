package config

import (
	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
)

// AgentConfig - represents the config for agent
type AgentConfig struct {
	CentralCfg corecfg.CentralConfig `config:"central"`
	GatewayCfg *GatewayConfig        `config:"gateway-section"`
}

// GatewayConfig - represents the config for gateway
type GatewayConfig struct {
	corecfg.IConfigValidator
	corecfg.IResourceConfigCallback
	LogFile        string `config:"logFile"`
	ProcessOnInput bool   `config:"processOnInput"`
	ConfigKey1     string `config:"config_key_1"`
	ConfigKey2     string `config:"config_key_2"`
	ConfigKey3     string `config:"config_key_3"`
}

// ValidateCfg - Validates the gateway config
func (c *GatewayConfig) ValidateCfg() (err error) {
	// if c.LogFile == "" {
	// 	return errors.New("Invalid gateway configuration: logFile is not configured")
	// }
	// if c.ConfigKey1 == "" {
	// 	return errors.New("Invalid gateway configuration: config_key_1 is not configured")
	// }

	return
}

// ApplyResources - Applies the apply API Server resource to the agent config
func (c *GatewayConfig) ApplyResources(agentResource *v1.ResourceInstance) error {
	// config, ok := agentResource.Spec["key"]
	// if !ok {
	// 	return errors.New("resource did not have a config map")
	// }

	// if value, ok := config.(map[string]interface{})["somekey"]; ok {
	// 	c.ConfigKey2 = value.(string)
	// }

	return nil
}
