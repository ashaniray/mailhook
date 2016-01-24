package main

import (
	"log"
	"net/http"
	"path"
)

var assetMap = map[string]string{
	".css":   "text/css",
	".js":    "application/javascript",
	".woff2": "font/opentype",
	".otf":   "font/opentype",
	".svg":   "image/svg+xml",
	".woff":  "font/opentype",
	".eot":   "font/opentype",
	".ttf":   "font/opentype",
}

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	asset := r.URL.Path[1:]
	ext := path.Ext(r.URL.Path)

	data, err := Asset(asset)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	mt, ok := assetMap[ext]

	if ok {
		w.Header().Set("Content-Type", mt)
	} else {
		w.Header().Set("Content-Type", "text/plain")
	}

	w.Write(data)
}
