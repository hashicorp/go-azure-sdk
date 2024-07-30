package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAcl struct {
	AccessType     *ExternalConnectorsAclAccessType     `json:"accessType,omitempty"`
	IdentitySource *ExternalConnectorsAclIdentitySource `json:"identitySource,omitempty"`
	ODataType      *string                              `json:"@odata.type,omitempty"`
	Type           *ExternalConnectorsAclType           `json:"type,omitempty"`
	Value          *string                              `json:"value,omitempty"`
}
