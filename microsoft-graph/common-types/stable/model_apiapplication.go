package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiApplication struct {
	AcceptMappedClaims          *bool                       `json:"acceptMappedClaims,omitempty"`
	KnownClientApplications     *[]string                   `json:"knownClientApplications,omitempty"`
	ODataType                   *string                     `json:"@odata.type,omitempty"`
	Oauth2PermissionScopes      *[]PermissionScope          `json:"oauth2PermissionScopes,omitempty"`
	PreAuthorizedApplications   *[]PreAuthorizedApplication `json:"preAuthorizedApplications,omitempty"`
	RequestedAccessTokenVersion *int64                      `json:"requestedAccessTokenVersion,omitempty"`
}
