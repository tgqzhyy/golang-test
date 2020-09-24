package api

import "github.com/hashicorp/consul/agent/router"

func init() {
	group1 := router.New("api", "/v1")
	group1.POST
}
