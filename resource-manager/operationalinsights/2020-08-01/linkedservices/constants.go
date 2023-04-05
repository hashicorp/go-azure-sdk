package linkedservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedServiceEntityStatus string

const (
	LinkedServiceEntityStatusDeleting            LinkedServiceEntityStatus = "Deleting"
	LinkedServiceEntityStatusProvisioningAccount LinkedServiceEntityStatus = "ProvisioningAccount"
	LinkedServiceEntityStatusSucceeded           LinkedServiceEntityStatus = "Succeeded"
	LinkedServiceEntityStatusUpdating            LinkedServiceEntityStatus = "Updating"
)

func PossibleValuesForLinkedServiceEntityStatus() []string {
	return []string{
		string(LinkedServiceEntityStatusDeleting),
		string(LinkedServiceEntityStatusProvisioningAccount),
		string(LinkedServiceEntityStatusSucceeded),
		string(LinkedServiceEntityStatusUpdating),
	}
}
