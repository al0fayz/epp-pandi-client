package examples

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/al0fayz/epp-pandi-client/frames"
)

func DeleteDomainByFile() {
	//handle file
	dir, _ := os.Getwd()
	current_path := filepath.Join(dir, "/examples/example_domain.txt")
	readFile, err := os.Open(current_path)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	//by new line
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		dom := fileScanner.Text()
		domainFrame := frames.DomainDeleteType{}
		domainFrame.SetDomain(dom)
		//login
		client, err := Login()
		if err != nil {
			fmt.Println(err)
		}
		//delete domain
		resDelete, err := client.DeleteDomain(&domainFrame)
		if err != nil {
			fmt.Println(err)
		}
		//response delete
		//print xml
		// fmt.Println(string(resDelete))
		response := frames.Response{}
		if err := xml.Unmarshal(resDelete, &response); err != nil {
			fmt.Println(err)
		}
		code := response.Result[0].Code
		if code == 1000 {
			fmt.Println("Domain deleted!", dom)
		} else {
			fmt.Println("Message : ", response.Result[0].Message)
		}

		//close connection
		client.Conn.Close()
	}

	readFile.Close()
	//end handlefile
}
