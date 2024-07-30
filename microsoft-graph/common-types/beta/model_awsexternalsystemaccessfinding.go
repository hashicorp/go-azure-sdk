package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsExternalSystemAccessFinding struct {
	AccessMethods        *AwsExternalSystemAccessFindingAccessMethods `json:"accessMethods,omitempty"`
	AffectedSystem       *AuthorizationSystem                         `json:"affectedSystem,omitempty"`
	CreatedDateTime      *string                                      `json:"createdDateTime,omitempty"`
	Id                   *string                                      `json:"id,omitempty"`
	ODataType            *string                                      `json:"@odata.type,omitempty"`
	SystemWithAccess     *AuthorizationSystemInfo                     `json:"systemWithAccess,omitempty"`
	TrustedIdentityCount *int64                                       `json:"trustedIdentityCount,omitempty"`
	TrustsAllIdentities  *bool                                        `json:"trustsAllIdentities,omitempty"`
}
