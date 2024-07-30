package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationEnforcedRestrictionsSessionControl struct {
	IsEnabled *bool   `json:"isEnabled,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
