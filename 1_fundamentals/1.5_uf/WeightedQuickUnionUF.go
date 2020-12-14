package __5_uf

type WeightedQuickUnionUF struct {
	QuickUnionUF
	size []int
}

func (this *WeightedQuickUnionUF) init(n int) {
	this.QuickUnionUF.init(n)
	this.size = make([]int, n)
	for i, _ := range this.size {
		this.size[i] = 1
	}
}

func (this *WeightedQuickUnionUF) union(p int, q int) {
	rootP := this.find(p)
	rootQ := this.find(q)
	if rootP == rootQ {
		return
	}
	sizeP := this.size[rootP]
	sizeQ := this.size[rootQ]
	if sizeP > sizeQ {
		this.id[rootQ] = rootP
		this.size[rootP] += sizeQ
	} else {
		this.id[rootP] = rootQ
		this.size[rootQ] += sizeP
	}

	this.count--
}
