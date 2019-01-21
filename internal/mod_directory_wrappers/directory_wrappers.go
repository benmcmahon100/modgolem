package directory_wrappers

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type Source struct {
	host 		string
	nicename	string
}

type Configuration struct {
	mod_directories	[]Source
}

func main() {
	configuration := Configuration{}

	err := gonfig.GetConf("directories.conf", &configuration)


	if err != nil {
		fmt.Printf(err.Error())
	}


}
