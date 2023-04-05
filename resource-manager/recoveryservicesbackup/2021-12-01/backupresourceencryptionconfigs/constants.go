package backupresourceencryptionconfigs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionAtRestType string

const (
	EncryptionAtRestTypeCustomerManaged  EncryptionAtRestType = "CustomerManaged"
	EncryptionAtRestTypeInvalid          EncryptionAtRestType = "Invalid"
	EncryptionAtRestTypeMicrosoftManaged EncryptionAtRestType = "MicrosoftManaged"
)

func PossibleValuesForEncryptionAtRestType() []string {
	return []string{
		string(EncryptionAtRestTypeCustomerManaged),
		string(EncryptionAtRestTypeInvalid),
		string(EncryptionAtRestTypeMicrosoftManaged),
	}
}

type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateDisabled InfrastructureEncryptionState = "Disabled"
	InfrastructureEncryptionStateEnabled  InfrastructureEncryptionState = "Enabled"
	InfrastructureEncryptionStateInvalid  InfrastructureEncryptionState = "Invalid"
)

func PossibleValuesForInfrastructureEncryptionState() []string {
	return []string{
		string(InfrastructureEncryptionStateDisabled),
		string(InfrastructureEncryptionStateEnabled),
		string(InfrastructureEncryptionStateInvalid),
	}
}

type LastUpdateStatus string

const (
	LastUpdateStatusFailed              LastUpdateStatus = "Failed"
	LastUpdateStatusFirstInitialization LastUpdateStatus = "FirstInitialization"
	LastUpdateStatusInitialized         LastUpdateStatus = "Initialized"
	LastUpdateStatusInvalid             LastUpdateStatus = "Invalid"
	LastUpdateStatusNotEnabled          LastUpdateStatus = "NotEnabled"
	LastUpdateStatusPartiallyFailed     LastUpdateStatus = "PartiallyFailed"
	LastUpdateStatusPartiallySucceeded  LastUpdateStatus = "PartiallySucceeded"
	LastUpdateStatusSucceeded           LastUpdateStatus = "Succeeded"
)

func PossibleValuesForLastUpdateStatus() []string {
	return []string{
		string(LastUpdateStatusFailed),
		string(LastUpdateStatusFirstInitialization),
		string(LastUpdateStatusInitialized),
		string(LastUpdateStatusInvalid),
		string(LastUpdateStatusNotEnabled),
		string(LastUpdateStatusPartiallyFailed),
		string(LastUpdateStatusPartiallySucceeded),
		string(LastUpdateStatusSucceeded),
	}
}
