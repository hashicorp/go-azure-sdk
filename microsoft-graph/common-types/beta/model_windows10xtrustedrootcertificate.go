package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XTrustedRootCertificate struct {
	Assignments            *[]DeviceManagementResourceAccessProfileAssignment `json:"assignments,omitempty"`
	CertFileName           *string                                            `json:"certFileName,omitempty"`
	CreationDateTime       *string                                            `json:"creationDateTime,omitempty"`
	Description            *string                                            `json:"description,omitempty"`
	DestinationStore       *Windows10XTrustedRootCertificateDestinationStore  `json:"destinationStore,omitempty"`
	DisplayName            *string                                            `json:"displayName,omitempty"`
	Id                     *string                                            `json:"id,omitempty"`
	LastModifiedDateTime   *string                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType              *string                                            `json:"@odata.type,omitempty"`
	RoleScopeTagIds        *[]string                                          `json:"roleScopeTagIds,omitempty"`
	TrustedRootCertificate *string                                            `json:"trustedRootCertificate,omitempty"`
	Version                *int64                                             `json:"version,omitempty"`
}
