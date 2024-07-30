package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternallyAccessibleAwsStorageBucketFinding struct {
	Accessibility      *ExternallyAccessibleAwsStorageBucketFindingAccessibility `json:"accessibility,omitempty"`
	AccountsWithAccess *AccountsWithAccess                                       `json:"accountsWithAccess,omitempty"`
	CreatedDateTime    *string                                                   `json:"createdDateTime,omitempty"`
	Id                 *string                                                   `json:"id,omitempty"`
	ODataType          *string                                                   `json:"@odata.type,omitempty"`
	StorageBucket      *AuthorizationSystemResource                              `json:"storageBucket,omitempty"`
}
