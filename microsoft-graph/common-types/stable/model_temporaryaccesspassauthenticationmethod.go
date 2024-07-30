package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemporaryAccessPassAuthenticationMethod struct {
	CreatedDateTime       *string `json:"createdDateTime,omitempty"`
	Id                    *string `json:"id,omitempty"`
	IsUsable              *bool   `json:"isUsable,omitempty"`
	IsUsableOnce          *bool   `json:"isUsableOnce,omitempty"`
	LifetimeInMinutes     *int64  `json:"lifetimeInMinutes,omitempty"`
	MethodUsabilityReason *string `json:"methodUsabilityReason,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	StartDateTime         *string `json:"startDateTime,omitempty"`
	TemporaryAccessPass   *string `json:"temporaryAccessPass,omitempty"`
}
