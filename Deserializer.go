package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	generated "github.com/shrja-ms/TestGoSample"
)

func main() {
	gomega.RegisterFailHandler(ginkgo.Fail)
	fmt.Println("!... Hello World ...!")
	fmt.Println(json.Marshal(validate()))
}

func validate() *generated.BackupRequest {

	testFile1, err := ioutil.ReadFile("./Backup.json")
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
	backuprequest := generated.BackupRequest{}
	err = json.Unmarshal([]byte(testFile1), &backuprequest)
	gomega.Expect(err).ShouldNot(gomega.HaveOccurred())

	return &backuprequest
}
