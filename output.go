package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

func stdout(wg *sync.WaitGroup, filtered, verbose bool, ss []status) {
	for _, s := range ss {
		if filtered && s.status != doesNotExistStatus {
			fmt.Println(fmt.Sprintf("[%d] %s", s.status, s.domain))
		} else if !filtered {
			fmt.Println(fmt.Sprintf("[%d] %s", s.status, s.domain))
		}
	}
	wg.Done()
}

func writeFile(wg *sync.WaitGroup, filepath string, ss []status) {
	file, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		return
	}
	writer := bufio.NewWriter(file)
	for _, s := range ss {
		_, err := writer.WriteString(fmt.Sprintf("[%d] %s\n", s.status, s.domain))
		if err != nil {
			log.Println(err)
			return
		}
	}
	writer.Flush()
	wg.Done()
}
