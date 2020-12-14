package __5_uf

type QuickUnionUF struct {
	UF
}

func (this *QuickUnionUF) find(p int) int {
	pID := this.id[p]
	for pID != this.id[pID] {
		pID = this.id[pID]
	}
	return pID
}

func (this *QuickUnionUF) union(p int, q int) {
	rootP := this.find(p)
	rootQ := this.find(q)
	if rootP == rootQ {
		return
	}
	this.id[rootP] = rootQ
	this.count--
}

func (this *QuickUnionUF) connected(p int, q int) bool {
	return this.find(p) == this.find(q)
}
