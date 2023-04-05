package quotas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceState string

const (
	ResourceStateCanceled  ResourceState = "Canceled"
	ResourceStateDeleted   ResourceState = "Deleted"
	ResourceStateFailed    ResourceState = "Failed"
	ResourceStateSucceeded ResourceState = "Succeeded"
)

func PossibleValuesForResourceState() []string {
	return []string{
		string(ResourceStateCanceled),
		string(ResourceStateDeleted),
		string(ResourceStateFailed),
		string(ResourceStateSucceeded),
	}
}
