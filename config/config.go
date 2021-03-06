package config

import (
	stdcfg "github.com/Microkubes/microservice-tools/config"
	"github.com/Microkubes/microservice-tools/gateway"
)

type ServiceConfig struct {
	// Service holds the confgiuration for connecting and registering the service with the API Gateway
	Service *gateway.MicroserviceConfig `json:"service"`
	// SecurityConfig holds the security configuration
	stdcfg.SecurityConfig `json:"security,omitempty"`
	// DBConfig holds the database connection configuration
	stdcfg.DBConfig `json:"database"`
	// GatewayURL is the URL of the API Gateway
	GatewayURL string `json:"gatewayUrl"`
	// GatewayAdminURL is the administration URL of the API Gateway. Used for purposes of registration of a
	// microservice with the API gateway.
	GatewayAdminURL string `json:"gatewayAdminUrl"`
	// ContainerManager is the platform for managing containerized services
	// Can be swarm or kubernetes
	ContainerManager string `json:"containerManager,omitempty"`
	//Version is version of the service
	Version string `json:"version"`
	// RabbitMQ holds information about the rabbitmq server
	RabbitMQ map[string]string `json:"rabbitmq"`
}

func (svc *ServiceConfig) ToStandardConfig() *stdcfg.ServiceConfig {
	return &stdcfg.ServiceConfig{
		Service:          svc.Service,
		SecurityConfig:   svc.SecurityConfig,
		DBConfig:         svc.DBConfig,
		GatewayURL:       svc.GatewayURL,
		GatewayAdminURL:  svc.GatewayAdminURL,
		ContainerManager: svc.ContainerManager,
		Version:          svc.Version,
	}
}
