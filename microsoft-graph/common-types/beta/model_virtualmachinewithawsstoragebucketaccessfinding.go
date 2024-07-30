package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineWithAwsStorageBucketAccessFinding struct {
	AccessibleCount       *int64                       `json:"accessibleCount,omitempty"`
	BucketCount           *int64                       `json:"bucketCount,omitempty"`
	CreatedDateTime       *string                      `json:"createdDateTime,omitempty"`
	Ec2Instance           *AuthorizationSystemResource `json:"ec2Instance,omitempty"`
	Id                    *string                      `json:"id,omitempty"`
	ODataType             *string                      `json:"@odata.type,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex       `json:"permissionsCreepIndex,omitempty"`
	Role                  *AwsRole                     `json:"role,omitempty"`
}
