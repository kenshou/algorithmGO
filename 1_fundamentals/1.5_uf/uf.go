package __5_uf

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type UF struct {
	id    []int
	count int
}

type UFInterface interface {
	init(n int)
	union(p int, q int)
	find(p int) int
	connected(p int, q int) bool
	countSet() int
}

func (u *UF) init(n int) {
	u.id = make([]int, n)
	for i, _ := range u.id {
		u.id[i] = i
	}
	u.count = n
}

func (u *UF) countSet() int {
	return u.count
}

func initByFile(puf *UFInterface, fileName string) {
	uf := *puf
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(file)
	for i, line := range strings.Split(fileStr, "\n") {
		if i == 0 {
			var n int
			fmt.Sscanf(line, "%d", &n)
			log.Println("nums = ", n)
			uf.init(n)
		} else {
			var p, q int
			fmt.Sscanf(line, "%d %d", &p, &q)
			if uf.find(p) != uf.find(q) {
				uf.union(p, q)
				log.Println("union ", p, ",", q, " : ", uf)
			}
		}
	}
	log.Println(uf.countSet(), " countSet")
}
