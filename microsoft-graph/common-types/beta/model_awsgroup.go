package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsGroup struct {
	AuthorizationSystem *AuthorizationSystem               `json:"authorizationSystem,omitempty"`
	DisplayName         *string                            `json:"displayName,omitempty"`
	ExternalId          *string                            `json:"externalId,omitempty"`
	Id                  *string                            `json:"id,omitempty"`
	ODataType           *string                            `json:"@odata.type,omitempty"`
	Source              *AuthorizationSystemIdentitySource `json:"source,omitempty"`
}
