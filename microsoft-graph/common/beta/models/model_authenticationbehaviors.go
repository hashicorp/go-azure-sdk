package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationBehaviors struct {
	ODataType                     *string `json:"@odata.type,omitempty"`
	RemoveUnverifiedEmailClaim    *bool   `json:"removeUnverifiedEmailClaim,omitempty"`
	RequireClientServicePrincipal *bool   `json:"requireClientServicePrincipal,omitempty"`
}
