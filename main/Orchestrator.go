package main

import "sync"

type LoadBalancer struct {
	regServers Backend[]
	algo       loadBalancingAlgo
}

// Backend the application servers
type Backend struct {
	Host      string `json:"host"`
	Port      string `json:"port"`
	IsHealthy string
	LastPing  int64
	Mutex     sync.RWMutex
}

// Config Load balancer configuration
type Config struct {
	Port     string    `json:"port"`
	Backends []Backend `json:"backends"`
}
