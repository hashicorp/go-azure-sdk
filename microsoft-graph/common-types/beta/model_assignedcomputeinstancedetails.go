package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedComputeInstanceDetails struct {
	AccessedStorageBuckets  *[]AuthorizationSystemResource `json:"accessedStorageBuckets,omitempty"`
	AssignedComputeInstance *AuthorizationSystemResource   `json:"assignedComputeInstance,omitempty"`
	Id                      *string                        `json:"id,omitempty"`
	ODataType               *string                        `json:"@odata.type,omitempty"`
}
