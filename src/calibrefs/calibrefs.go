package calibrefs

import (
	"fmt"
	"net"

	uuid "github.com/nu7hatch/gouuid"
	"github.com/rminnich/go9p"
)

const (
	CFSEntryNIL     uint8 = 0x0
	CFSEntryAuthors       = 0x1
	CFSEntryTags          = 0x2
	CFSEntryFormats       = 0x4
	CFSEntryTitles        = 0x8
)

type CalibreFS struct {
	namesByEntry  map[uint8]string
	entriesByName map[string]uint8
	go9p.Srv
}

func NewCalibreFS() *CalibreFS {

	calibrefs := new(CalibreFS)
	calibrefs.namesByEntry = map[uint8]string{
		CFSEntryAuthors: "Authors",
		CFSEntryTags:    "Tags",
		CFSEntryFormats: "Formats",
		CFSEntryTitles:  "Titles",
	}
	calibrefs.entriesByName = map[string]uint8{}
	for entry, name := range calibrefs.namesByEntry {
		calibrefs.entriesByName[name] = entry
	}

	return calibrefs
}

func (calibrefs *CalibreFS) StartNetListener(ntype, addr string) error {
	l, err := net.Listen(ntype, addr)
	if err != nil {
		return fmt.Errorf("Unable to start NetListener: %s", err.Error())
	}

	for {
		c, err := l.Accept()
		if err != nil {
			return fmt.Errorf("Unable to handle incoming connection: %s", err.Error())
		}

		calibrefs.NewConn(c)
	}
}

func (calibrefs *CalibreFS) NameByEntry(entry uint8) (entryName *string, err error) {
	name, ok := calibrefs.namesByEntry[entry]
	if !ok {
		return nil, fmt.Errorf("Entry id %d not found!", entry)
	}

	return &name, nil
}

func (calibrefs *CalibreFS) EntryByName(name string) (entry *uint8, err error) {
	_entry, ok := calibrefs.entriesByName[name]
	if !ok {
		return nil, fmt.Errorf("Name %s not found!", name)
	}

	return &_entry, nil
}

type CFSFid struct {
	Id      uuid.UUID
	Entries uint8
}

func newCFSFid(entries uint8) (u *CFSFid, err error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	cfsFid := new(CFSFid)
	cfsFid.Id = *uuid
	cfsFid.Entries = entries
	return cfsFid, nil
}

func rootCFSFid() (cfsFid *CFSFid, err error) {
	return newCFSFid(go9p.QTDIR)
}

func (cfsFid *CFSFid) hasEntry(entry uint8) bool {
	return entry&cfsFid.Entries == 1
}

func rootQid() *go9p.Qid {
	var qid go9p.Qid
	qid.Path = 0
	qid.Version = 0
	qid.Type = go9p.QTDIR

	return &qid
}
