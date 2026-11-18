package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, datastring := range dataset {
		err := dp.Parse(datastring)
		if err != nil {
			log.Printf("parsing error: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("action error: %v", err)
			continue
		}
		fmt.Println(info)
	}
}
