package products

import (
	"fmt"
	"log"
)

func (puc ProductUC) MakeError() error {
	return fmt.Errorf("test error")
}

func (puc ProductUC) MakeNullPtr() int {
	var aMap []int = nil
	return aMap[1]
}

// NewRelic can not catch intentional exit (os.Exit(1))
func (puc ProductUC) MakeFatalError() int {
	log.Fatal("Mati lo!")
	return 1
}
