package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/nomad/plugins"

	"github.com/ubiquity-community/nomad-driver-nspawn/nspawn"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of the nspawn driver plugin
func factory(log hclog.Logger) interface{} {
	return nspawn.NewNspawnDriver(log)
}
