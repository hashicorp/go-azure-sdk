package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Permission struct {
	ExpirationDateTime    *string                  `json:"expirationDateTime,omitempty"`
	GrantedTo             *IdentitySet             `json:"grantedTo,omitempty"`
	GrantedToIdentities   *[]IdentitySet           `json:"grantedToIdentities,omitempty"`
	GrantedToIdentitiesV2 *[]SharePointIdentitySet `json:"grantedToIdentitiesV2,omitempty"`
	GrantedToV2           *SharePointIdentitySet   `json:"grantedToV2,omitempty"`
	HasPassword           *bool                    `json:"hasPassword,omitempty"`
	Id                    *string                  `json:"id,omitempty"`
	InheritedFrom         *ItemReference           `json:"inheritedFrom,omitempty"`
	Invitation            *SharingInvitation       `json:"invitation,omitempty"`
	Link                  *SharingLink             `json:"link,omitempty"`
	ODataType             *string                  `json:"@odata.type,omitempty"`
	Roles                 *[]string                `json:"roles,omitempty"`
	ShareId               *string                  `json:"shareId,omitempty"`
}
