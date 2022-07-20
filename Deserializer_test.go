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

func TestRestoreRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/Restore.json")
	restorerequest := generated.RestoreRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &restorerequest)

	//Validate the incoming request is proper RestoreRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the RestoreRequest to verify same object got created
	restorerequeststring, err2 := json.Marshal(restorerequest)

	// Deserializing again to recreate the restorerequest.
	restorerequestRegenerated := generated.RestoreRequest{}
	err3 := json.Unmarshal([]byte(restorerequeststring), &restorerequestRegenerated)

	//Verify both the Incoming Restore Request and regenerated RestoreRequest are same.
	if !reflect.DeepEqual(restorerequestRegenerated, restorerequest) {
		t.Errorf("RestoreRequestMarshalled() = %v, RestoreRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(restorerequeststring), string(testFile1), err2, err3)
	}
}

func TestValidateForRestoreRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/ValidateForRestore.json")
	validateforrestorerequest := generated.ValidateForRestoreRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &validateforrestorerequest)

	//Validate the incoming request is proper RestoreRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the RestoreRequest to verify same object got created
	validateforrestorerequeststring, err2 := json.Marshal(validateforrestorerequest)

	// Deserializing again to recreate the restorerequest.
	validateforrestorerequestRegenerated := generated.ValidateForRestoreRequest{}
	err3 := json.Unmarshal([]byte(validateforrestorerequeststring), &validateforrestorerequestRegenerated)

	//Verify both the Incoming Restore Request and regenerated RestoreRequest are same.
	if !reflect.DeepEqual(validateforrestorerequestRegenerated, validateforrestorerequest) {
		t.Errorf("RestoreRequestMarshalled() = %v, RestoreRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(validateforrestorerequeststring), string(testFile1), err2, err3)
	}
}

func TestCommitOrRollbackRestoreRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/CommitOrRollbackRestore.json")
	commitorrollbackrestorerequest := generated.CommitOrRollbackRestoreRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &commitorrollbackrestorerequest)

	//Validate the incoming request is proper RestoreRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the RestoreRequest to verify same object got created
	commitorrollbackrestorerequeststring, err2 := json.Marshal(commitorrollbackrestorerequest)

	// Deserializing again to recreate the restorerequest.
	commitorrollbackrestorerequestRegenerated := generated.CommitOrRollbackRestoreRequest{}
	err3 := json.Unmarshal([]byte(commitorrollbackrestorerequeststring), &commitorrollbackrestorerequestRegenerated)

	//Verify both the Incoming Restore Request and regenerated RestoreRequest are same.
	if !reflect.DeepEqual(commitorrollbackrestorerequestRegenerated, commitorrollbackrestorerequest) {
		t.Errorf("RestoreRequestMarshalled() = %v, RestoreRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(commitorrollbackrestorerequeststring), string(testFile1), err2, err3)
	}
}

func TestStartProtectionRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/StartProtection.json")
	startprotectionrequest := generated.StartProtectionRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &startprotectionrequest)

	//Validate the incoming request is proper StartProtectionRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the StartProtectionRequest to verify same object got created
	startprotectionrequeststring, err2 := json.Marshal(startprotectionrequest)

	// Deserializing again to recreate the startprotectionrequest.
	startprotectionrequestRegenerated := generated.StartProtectionRequest{}
	err3 := json.Unmarshal([]byte(startprotectionrequeststring), &startprotectionrequestRegenerated)

	//Verify both the Incoming StartProtection Request and regenerated StartProtectionRequest are same.
	if !reflect.DeepEqual(startprotectionrequestRegenerated, startprotectionrequest) {
		t.Errorf("StartProtectionRequestMarshalled() = %v, StartProtectionRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(startprotectionrequeststring), string(testFile1), err2, err3)
	}
}

func TestValidateForProtectionRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/ValidateForProtection.json")
	validateforprotectionrequest := generated.ValidateForProtectionRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &validateforprotectionrequest)

	//Validate the incoming request is proper StartProtectionRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the StartProtectionRequest to verify same object got created
	validateforprotectionrequeststring, err2 := json.Marshal(validateforprotectionrequest)

	// Deserializing again to recreate the startprotectionrequest.
	validateforprotectionrequestRegenerated := generated.ValidateForProtectionRequest{}
	err3 := json.Unmarshal([]byte(validateforprotectionrequeststring), &validateforprotectionrequestRegenerated)

	//Verify both the Incoming StartProtection Request and regenerated StartProtectionRequest are same.
	if !reflect.DeepEqual(validateforprotectionrequestRegenerated, validateforprotectionrequest) {
		t.Errorf("StartProtectionRequestMarshalled() = %v, StartProtectionRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(validateforprotectionrequeststring), string(testFile1), err2, err3)
	}
}

func TestStopProtectionRequest(t *testing.T) {
	//Reading the input to plugin
	testFile1, err := ioutil.ReadFile("./Requests/StopProtection.json")
	stopprotectionrequest := generated.StopProtectionRequest{}

	//Deserializing using contract in the SDK
	err = json.Unmarshal([]byte(testFile1), &stopprotectionrequest)

	//Validate the incoming request is proper StartProtectionRequest
	if err != nil {
		t.Errorf("Deserialization error = %v", err)
		return
	}

	//Stringify the StartProtectionRequest to verify same object got created
	stopprotectionrequeststring, err2 := json.Marshal(stopprotectionrequest)

	// Deserializing again to recreate the startprotectionrequest.
	stopprotectionrequestRegenerated := generated.StopProtectionRequest{}
	err3 := json.Unmarshal([]byte(stopprotectionrequeststring), &stopprotectionrequestRegenerated)

	//Verify both the Incoming StartProtection Request and regenerated StartProtectionRequest are same.
	if !reflect.DeepEqual(stopprotectionrequestRegenerated, stopprotectionrequest) {
		t.Errorf("StartProtectionRequestMarshalled() = %v, StartProtectionRequestFile =  %v, ErrorInSerializing = %v, ErrorInDeserializing = %v", string(stopprotectionrequeststring), string(testFile1), err2, err3)
	}
}
