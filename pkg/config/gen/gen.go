package main

import (
	cfg "github.com/conductorone/baton-ms365/pkg/config"
	"github.com/conductorone/baton-sdk/pkg/config"
)

func main() {
	config.Generate("ms365", cfg.Config)
}
