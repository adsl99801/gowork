// parse includes/config.txt into a map
//
// format of includes/config.txt:
//   key=value
//
// access by:
//  value := config["key"]
//
package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var config map[string]string
var confFile string

//sample o
func sample() {
	confFile = "config.txt"
	config = ReadConfig(confFile)

}

//ReadConfig o
func ReadConfig(filenameFullpath string) map[string]string {
	prg := "ReadConfig()"

	var options map[string]string
	options = make(map[string]string)

	file, err := os.Open(filenameFullpath)
	if err != nil {
		log.Printf("%s: os.Open(): %s\n", prg, err)
		return options
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "=") == true {
			re, err := regexp.Compile(`([^=]+)=(.*)`)
			if err != nil {
				log.Printf("%s: regexp.Compile(): error=%s", prg, err)
				return options
			}
			configOption := re.FindStringSubmatch(line)[1]
			configValue := re.FindStringSubmatch(line)[2]
			options[configOption] = configValue
			log.Printf("%s: out[]: %s ... config_option=%s, config_value=%s\n", prg, line, configOption, configValue)
		}
	}
	log.Printf("%s: options[]: %+v\n", prg, options)

	if err := scanner.Err(); err != nil {
		log.Printf("%s: scanner.Err(): %s\n", prg, err)
		return options
	}
	return options
}
