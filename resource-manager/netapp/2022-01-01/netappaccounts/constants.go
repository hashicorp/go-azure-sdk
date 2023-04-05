package netappaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryStatus string

const (
	ActiveDirectoryStatusCreated  ActiveDirectoryStatus = "Created"
	ActiveDirectoryStatusDeleted  ActiveDirectoryStatus = "Deleted"
	ActiveDirectoryStatusError    ActiveDirectoryStatus = "Error"
	ActiveDirectoryStatusInUse    ActiveDirectoryStatus = "InUse"
	ActiveDirectoryStatusUpdating ActiveDirectoryStatus = "Updating"
)

func PossibleValuesForActiveDirectoryStatus() []string {
	return []string{
		string(ActiveDirectoryStatusCreated),
		string(ActiveDirectoryStatusDeleted),
		string(ActiveDirectoryStatusError),
		string(ActiveDirectoryStatusInUse),
		string(ActiveDirectoryStatusUpdating),
	}
}
