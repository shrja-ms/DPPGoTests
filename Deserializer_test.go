package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	generated "github.com/shrja-ms/TestGoSample"
)

func TestValidateBackupRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./BackupRequest.json")
	backuprequest := generated.BackupRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &backuprequest)

	//Validate the incoming request is proper BackupRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the BackupRequest to verify same object got created
	backuprequeststring, err2 := json.Marshal(backuprequest)

	// Deserializing again to recreate the backuprequest.
	backuprequestRegenerated := generated.BackupRequest{}
	err3 := json.Unmarshal([]byte(backuprequeststring), &backuprequestRegenerated)

	//Verify both the Incoming Backup Request and regenerated BackupRequest are same.
	if !reflect.DeepEqual(backuprequestRegenerated, backuprequest) {
		t.Errorf("BackupRequestMarshalled() = %v, BackupRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(backuprequeststring), string(testFile1), err2, err3)
	}
}
