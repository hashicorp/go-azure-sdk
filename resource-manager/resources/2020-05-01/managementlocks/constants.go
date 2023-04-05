package managementlocks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LockLevel string

const (
	LockLevelCanNotDelete LockLevel = "CanNotDelete"
	LockLevelNotSpecified LockLevel = "NotSpecified"
	LockLevelReadOnly     LockLevel = "ReadOnly"
)

func PossibleValuesForLockLevel() []string {
	return []string{
		string(LockLevelCanNotDelete),
		string(LockLevelNotSpecified),
		string(LockLevelReadOnly),
	}
}
