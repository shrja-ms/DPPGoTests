package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"

	generated "github.com/shrja-ms/TestGoSample"
)

func TestBackupRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/Backup.json")
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

func TestValidateForBackupRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/ValidateForBackup.json")
	validateforbackuprequest := generated.ValidateForBackupRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &validateforbackuprequest)

	//Validate the incoming request is proper BackupRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the BackupRequest to verify same object got created
	validateforbackuprequeststring, err2 := json.Marshal(validateforbackuprequest)

	// Deserializing again to recreate the backuprequest.
	validateforbackuprequestRegenerated := generated.ValidateForBackupRequest{}
	err3 := json.Unmarshal([]byte(validateforbackuprequeststring), &validateforbackuprequestRegenerated)

	//Verify both the Incoming Backup Request and regenerated BackupRequest are same.
	if !reflect.DeepEqual(validateforbackuprequestRegenerated, validateforbackuprequest) {
		t.Errorf("BackupRequestMarshalled() = %v, BackupRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(validateforbackuprequeststring), string(testFile1), err2, err3)
	}
}

func TestCommitOrRollbackBackupRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/CommitOrRollbackBackup.json")
	commitorrollbackbackuprequest := generated.CommitOrRollbackBackupRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &commitorrollbackbackuprequest)

	//Validate the incoming request is proper BackupRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the BackupRequest to verify same object got created
	commitorrollbackbackuprequeststring, err2 := json.Marshal(commitorrollbackbackuprequest)

	// Deserializing again to recreate the backuprequest.
	commitorrollbackbackuprequestRegenerated := generated.CommitOrRollbackBackupRequest{}
	err3 := json.Unmarshal([]byte(commitorrollbackbackuprequeststring), &commitorrollbackbackuprequestRegenerated)

	//Verify both the Incoming Backup Request and regenerated BackupRequest are same.
	if !reflect.DeepEqual(commitorrollbackbackuprequestRegenerated, commitorrollbackbackuprequest) {
		t.Errorf("BackupRequestMarshalled() = %v, BackupRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(commitorrollbackbackuprequeststring), string(testFile1), err2, err3)
	}
}
