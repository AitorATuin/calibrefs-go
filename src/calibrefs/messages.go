package calibrefs

import "github.com/rminnich/go9p"

func (calibrefs *CalibreFS) Attach(req *go9p.SrvReq) {
	Logger.Debugf("Attach message")
	var err error
	if req.Afid != nil {
		req.RespondError(go9p.Enoauth)
		return
	}

	req.Fid.Aux, err = rootCFSFid()
	if err != nil {
		Logger.Errorf("Uname to create rootCFSFid")
		req.RespondError(err)
	}
	Logger.Debugf("Created rootCFSFid")
	req.RespondRattach(rootQid())
}

func (c *CalibreFS) Walk(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Open(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Create(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Read(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Write(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Clunk(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Remove(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Stat(req *go9p.SrvReq) {
	panic("not implemented")
}

func (c *CalibreFS) Wstat(req *go9p.SrvReq) {
	panic("not implemented")
}
