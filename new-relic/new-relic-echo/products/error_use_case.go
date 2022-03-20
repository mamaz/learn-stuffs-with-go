package products

import "log"

func (puc ProductUC) MakeError() int {
	return 1
}

func (puc ProductUC) MakeNullPtr() int {
	var aMap []int = nil
	return aMap[1]
}

func (puc ProductUC) MakeFatalError() int {
	log.Fatal("Mati lo!")
	return 1
}
