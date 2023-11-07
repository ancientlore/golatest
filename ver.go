package main

type file struct {
	Filename string `json:"filename"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Version  string `json:"version"`
	SHA256   []byte `json:"sha256"`
	Size     int    `json:"size"`
	Kind     string `json:"kind"`
}

type release struct {
	Version string `json:"version"`
	Stable  bool   `json:"stable"`
	Files   []file `json:"files"`
}
