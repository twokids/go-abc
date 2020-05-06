package animal

import (
	"fmt"
	"go-abc/basic/basic_03_import/life"
)

func init() {
	fmt.Println("this is animal init")
}

func Study() {
	fmt.Println("animal can study")
	life.Live()
}
