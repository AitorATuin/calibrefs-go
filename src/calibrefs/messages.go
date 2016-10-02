package calibrefs

import "github.com/rminnich/go9p"

func (calibrefs *CalibreFS) Attach(req *go9p.SrvReq) {
	var err error
	req.Fid.Aux, err = rootCFSFid()
	if err != nil {
		req.RespondError(err)
	}
	req.RespondRattach(rootQid())
}
