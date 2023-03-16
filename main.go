package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	heredotcom "github.com/krylphi/travels-naming/internal/api/here.com"
	"github.com/krylphi/travels-naming/internal/domain/naming"
	"github.com/krylphi/travels-naming/internal/util"
)

func main() {

	parameters := os.Args[1:]
	if len(parameters) == 0 {
		log.Fatalln("expecting filename as an argument")
	}

	filename := parameters[0]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer util.CloseOrLog(file)

	apiKey, err := heredotcom.GetApiKey()
	if err != nil {
		log.Fatal(err)
	}
	api := heredotcom.NewAPI(apiKey)
	// could not test on actual API because for some reason issued valid token returns 403 error on request,
	// although no mentions of this behaviour in documentation
	// (apart from RESTRICTED tagged apis, but revgeocode endpoint do not have this tag)
	// thus I have to use mock api, which imitates real api for testing.
	// api = heredotcom.MockAPI{}
	processor := naming.NewProcessor(api)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		params := strings.Split(line, ",")
		if len(params) != 3 {
			continue
		}
		res, err := processor.Process(params[0], params[1], params[2])
		if err != nil {
			log.Printf("error processing data: time: %s, lat: %s, lon: %s, err: %s\n",
				params[0], params[1], params[2], err.Error(),
			)
		}
		log.Printf("input: [time: %s, lat: %s, lon: %s] output: %s \n", params[0], params[1], params[2], res)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
