package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedOrganization struct {
	CreatedBy        *string                     `json:"createdBy,omitempty"`
	CreatedDateTime  *string                     `json:"createdDateTime,omitempty"`
	Description      *string                     `json:"description,omitempty"`
	DisplayName      *string                     `json:"displayName,omitempty"`
	ExternalSponsors *[]DirectoryObject          `json:"externalSponsors,omitempty"`
	Id               *string                     `json:"id,omitempty"`
	IdentitySources  *[]IdentitySource           `json:"identitySources,omitempty"`
	InternalSponsors *[]DirectoryObject          `json:"internalSponsors,omitempty"`
	ModifiedBy       *string                     `json:"modifiedBy,omitempty"`
	ModifiedDateTime *string                     `json:"modifiedDateTime,omitempty"`
	ODataType        *string                     `json:"@odata.type,omitempty"`
	State            *ConnectedOrganizationState `json:"state,omitempty"`
}
