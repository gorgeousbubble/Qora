package app

type PackFile struct {
	SrcFiles  []string `json:"src_files" yaml:"src_files" xml:"src_files"`
	DestFile  string   `json:"dest_file" yaml:"dest_file" xml:"dest_file"`
	Algorithm string   `json:"algorithm" yaml:"algorithm" xml:"algorithm"`
}

type UnpackFile struct {
	SrcFile  string `json:"src_file" yaml:"src_file" xml:"src_file"`
	DestFile string `json:"dest_file" yaml:"dest_file" xml:"dest_file"`
}
