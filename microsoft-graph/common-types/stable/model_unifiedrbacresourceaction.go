package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRbacResourceAction struct {
	ActionVerb                      *string `json:"actionVerb,omitempty"`
	AuthenticationContextId         *string `json:"authenticationContextId,omitempty"`
	Description                     *string `json:"description,omitempty"`
	Id                              *string `json:"id,omitempty"`
	IsAuthenticationContextSettable *bool   `json:"isAuthenticationContextSettable,omitempty"`
	Name                            *string `json:"name,omitempty"`
	ODataType                       *string `json:"@odata.type,omitempty"`
	ResourceScopeId                 *string `json:"resourceScopeId,omitempty"`
}
