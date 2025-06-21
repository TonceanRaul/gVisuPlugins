package main

import "github.com/TonceanRaul/gVisu/api/pluginapi"

// GrayscalePlugin implements the pluginapi.FilterPlugin interface (assumed)
type GrayscalePlugin struct{}

func (GrayscalePlugin) GetFilters() map[string]string {
	return map[string]string{
		"Grayscale": "format=gray",
	}
}

// Exported symbol for dynamic plugin loading (must be a var, not const or func)
var Plugin GrayscalePlugin
