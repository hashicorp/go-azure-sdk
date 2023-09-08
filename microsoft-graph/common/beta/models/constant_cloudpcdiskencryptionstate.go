package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDiskEncryptionState string

const (
	CloudPCDiskEncryptionStateencryptedUsingCustomerManagedKey CloudPCDiskEncryptionState = "EncryptedUsingCustomerManagedKey"
	CloudPCDiskEncryptionStateencryptedUsingPlatformManagedKey CloudPCDiskEncryptionState = "EncryptedUsingPlatformManagedKey"
	CloudPCDiskEncryptionStatenotAvailable                     CloudPCDiskEncryptionState = "NotAvailable"
	CloudPCDiskEncryptionStatenotEncrypted                     CloudPCDiskEncryptionState = "NotEncrypted"
)

func PossibleValuesForCloudPCDiskEncryptionState() []string {
	return []string{
		string(CloudPCDiskEncryptionStateencryptedUsingCustomerManagedKey),
		string(CloudPCDiskEncryptionStateencryptedUsingPlatformManagedKey),
		string(CloudPCDiskEncryptionStatenotAvailable),
		string(CloudPCDiskEncryptionStatenotEncrypted),
	}
}

func parseCloudPCDiskEncryptionState(input string) (*CloudPCDiskEncryptionState, error) {
	vals := map[string]CloudPCDiskEncryptionState{
		"encryptedusingcustomermanagedkey": CloudPCDiskEncryptionStateencryptedUsingCustomerManagedKey,
		"encryptedusingplatformmanagedkey": CloudPCDiskEncryptionStateencryptedUsingPlatformManagedKey,
		"notavailable":                     CloudPCDiskEncryptionStatenotAvailable,
		"notencrypted":                     CloudPCDiskEncryptionStatenotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDiskEncryptionState(input)
	return &out, nil
}
