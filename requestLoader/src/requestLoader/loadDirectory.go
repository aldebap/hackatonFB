////////////////////////////////////////////////////////////////////////////////
//	loadDirectory.go  -  May/06/2017  -  aldebaran perseke
//
//	Polling the load directory for request files
////////////////////////////////////////////////////////////////////////////////

package requestLoader

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const poolingDelay = 30

////////////////////////////////////////////////////////////////////////////////
//	Polling the load directory for request files
////////////////////////////////////////////////////////////////////////////////

func LoadDirectoryPolling(_loadDirectory string) {

	for {
		//	get the list of files in load directory
		fileList, err := ioutil.ReadDir(_loadDirectory)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[error] can't get file list from directory %s\n", _loadDirectory)
			return
		}

		//	iterate to the list of files
		for _, file := range fileList {

			//	load only regular files
			if true == file.Mode().IsRegular() {
				LoadRequestFile(_loadDirectory + "\\" + file.Name())
			}
		}

		time.Sleep(time.Duration(poolingDelay) * time.Second)
	}
}
