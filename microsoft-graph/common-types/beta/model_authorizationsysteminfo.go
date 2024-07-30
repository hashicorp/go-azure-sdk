package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemInfo struct {
	AuthorizationSystemType *AuthorizationSystemInfoAuthorizationSystemType `json:"authorizationSystemType,omitempty"`
	DisplayName             *string                                         `json:"displayName,omitempty"`
	Id                      *string                                         `json:"id,omitempty"`
	ODataType               *string                                         `json:"@odata.type,omitempty"`
}
