package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternallyAccessibleGcpStorageBucketFinding struct {
	Accessibility       *ExternallyAccessibleGcpStorageBucketFindingAccessibility       `json:"accessibility,omitempty"`
	CreatedDateTime     *string                                                         `json:"createdDateTime,omitempty"`
	EncryptionManagedBy *ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy `json:"encryptionManagedBy,omitempty"`
	Id                  *string                                                         `json:"id,omitempty"`
	ODataType           *string                                                         `json:"@odata.type,omitempty"`
	StorageBucket       *AuthorizationSystemResource                                    `json:"storageBucket,omitempty"`
}
