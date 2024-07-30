package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsIdentityAccessManagementKeyAgeFinding struct {
	AccessKey             *AwsAccessKey                                   `json:"accessKey,omitempty"`
	ActionSummary         *ActionSummary                                  `json:"actionSummary,omitempty"`
	AwsAccessKeyDetails   *AwsAccessKeyDetails                            `json:"awsAccessKeyDetails,omitempty"`
	CreatedDateTime       *string                                         `json:"createdDateTime,omitempty"`
	Id                    *string                                         `json:"id,omitempty"`
	ODataType             *string                                         `json:"@odata.type,omitempty"`
	PermissionsCreepIndex *PermissionsCreepIndex                          `json:"permissionsCreepIndex,omitempty"`
	Status                *AwsIdentityAccessManagementKeyAgeFindingStatus `json:"status,omitempty"`
}
