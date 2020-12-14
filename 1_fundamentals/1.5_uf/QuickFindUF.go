package __5_uf

type QuickFindUF struct {
	UF
}

func (this *QuickFindUF) find(p int) int {
	return this.id[p]
}

func (this *QuickFindUF) union(p int, q int) {
	pID := this.id[p]
	qID := this.id[q]
	if pID == qID {
		return
	}
	for i, id := range this.id {
		if id == pID {
			this.id[i] = qID
		}
	}
	this.count--
}

func (this *QuickFindUF) connected(p int, q int) bool {
	return this.find(p) == this.find(q)
}
