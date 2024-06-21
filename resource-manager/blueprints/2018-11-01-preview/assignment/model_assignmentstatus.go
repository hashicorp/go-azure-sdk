package assignment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentStatus struct {
	LastModified     *string   `json:"lastModified,omitempty"`
	ManagedResources *[]string `json:"managedResources,omitempty"`
	TimeCreated      *string   `json:"timeCreated,omitempty"`
}
