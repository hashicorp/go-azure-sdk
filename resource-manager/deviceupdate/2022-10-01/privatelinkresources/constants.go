package privatelinkresources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupIdProvisioningState string

const (
	GroupIdProvisioningStateCanceled  GroupIdProvisioningState = "Canceled"
	GroupIdProvisioningStateFailed    GroupIdProvisioningState = "Failed"
	GroupIdProvisioningStateSucceeded GroupIdProvisioningState = "Succeeded"
)

func PossibleValuesForGroupIdProvisioningState() []string {
	return []string{
		string(GroupIdProvisioningStateCanceled),
		string(GroupIdProvisioningStateFailed),
		string(GroupIdProvisioningStateSucceeded),
	}
}
