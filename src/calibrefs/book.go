package calibrefs

const (
	PDF  int = iota
	MOBI     = iota
	EPUB     = iota
	PRC      = iota
	TXT      = iota
)

type Format int

type BookEntry struct {
	Title  string
	Author string
	Tags   []*string
	Format []Format
}
