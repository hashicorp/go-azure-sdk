package diskencryptionsets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskEncryptionSetType string

const (
	DiskEncryptionSetTypeConfidentialVMEncryptedWithCustomerKey      DiskEncryptionSetType = "ConfidentialVmEncryptedWithCustomerKey"
	DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey             DiskEncryptionSetType = "EncryptionAtRestWithCustomerKey"
	DiskEncryptionSetTypeEncryptionAtRestWithPlatformAndCustomerKeys DiskEncryptionSetType = "EncryptionAtRestWithPlatformAndCustomerKeys"
)

func PossibleValuesForDiskEncryptionSetType() []string {
	return []string{
		string(DiskEncryptionSetTypeConfidentialVMEncryptedWithCustomerKey),
		string(DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		string(DiskEncryptionSetTypeEncryptionAtRestWithPlatformAndCustomerKeys),
	}
}
