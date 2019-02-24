package Configuration

import (
	"net/http"
	"strconv"
	"strings"
)

type Server struct{
	*Configuration
}

func initializeDefaultConfiguration() *Configuration{
	var defaultConfiguration Configuration

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

func (server *Server) StartServer(configuration *Configuration) {
	server.Configuration = initializeDefaultConfiguration()

	address := strings.Join([]string{server.HostName, ":", strconv.Itoa(server.Port)}, "")

	http.ListenAndServe(address, nil)
}

func (server *Server) AddPage(pattern string, function func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, function)
}

