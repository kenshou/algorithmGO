package __5_uf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestUF_initByFile(t *testing.T) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	fileName := dir + "/../1_fundamentals/1.5_uf/tinyUF.txt"
	log.Println("fileName: " + fileName)
	//var uf UFInterface= &QuickFindUF{}
	//var uf UFInterface= &QuickUnionUF{}
	var uf UFInterface = &WeightedQuickUnionUF{}
	initByFile(&uf, fileName)
	log.Println(uf)

}
