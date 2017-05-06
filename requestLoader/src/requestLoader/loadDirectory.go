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

const processedDirectory = "processed"
const poolingDelay = 30

////////////////////////////////////////////////////////////////////////////////
//	Polling the load directory for request files
////////////////////////////////////////////////////////////////////////////////

func LoadDirectoryPolling(_loadDirectory string) {

	//	check for the processed directory
	_, err := os.Stat(_loadDirectory + "/" + processedDirectory)
	if nil != err {
		if os.IsNotExist(err) {
			os.Mkdir(_loadDirectory+"/"+processedDirectory, 644)
		} else {
			fmt.Fprintf(os.Stderr, "[error] can't get info about directory name %s\n", _loadDirectory)
		}
	}

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
				LoadRequestFile(_loadDirectory + "/" + file.Name())
				os.Rename(_loadDirectory+"/"+file.Name(), _loadDirectory+"/"+processedDirectory+"/"+file.Name())
			}
		}

		//	sleep for a while until check it again
		time.Sleep(time.Duration(poolingDelay) * time.Second)
	}
}
