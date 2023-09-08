package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2KeyRestrictions struct {
	AaGuids         *[]string                            `json:"aaGuids,omitempty"`
	EnforcementType *Fido2KeyRestrictionsEnforcementType `json:"enforcementType,omitempty"`
	IsEnforced      *bool                                `json:"isEnforced,omitempty"`
	ODataType       *string                              `json:"@odata.type,omitempty"`
}
