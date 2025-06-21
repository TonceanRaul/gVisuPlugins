package main

type GrayscalePlugin struct{}

func (GrayscalePlugin) GetFilters() map[string]string {
	return map[string]string{
		"Grayscale": "format=gray",
	}
}

// Exported symbol so plugin can be loaded
var Plugin GrayscalePlugin
