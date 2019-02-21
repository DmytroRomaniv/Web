package Models

import (
	"../Configuration"
	"net/http"
	"strings"
)

type Server struct{
	*Configuration.Configuration
}

func (server *Server) initializeDefaultConfiguration() *Configuration.Configuration{
	var defaultConfiguration Configuration.Configuration

	if defaultConfiguration.Name == "" {
		defaultConfiguration.Name = "Server"
	}
	if defaultConfiguration.HostName == "" {
		defaultConfiguration.HostName = "localhost"
	}
	if defaultConfiguration.IP == "" {
		defaultConfiguration.IP = "127.0.0.1"
	}
	if defaultConfiguration.Port == 0 {
		defaultConfiguration.Port = 8080
	}

	return &defaultConfiguration
}

func (server *Server) StartServer(configuration *Configuration.Configuration) error{
	address := strings.Join([]string{server.Name, ":", string(server.Port)}, "")

	http.ListenAndServe(address, nil)
	
	return nil
}

func (server *Server) AddPage(pattern string, function func(http.ResponseWriter, *http.Request) )  {
	http.HandleFunc(pattern, function)
}