package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostReputationRule struct {
	Description       *string                             `json:"description,omitempty"`
	Name              *string                             `json:"name,omitempty"`
	ODataType         *string                             `json:"@odata.type,omitempty"`
	RelatedDetailsUrl *string                             `json:"relatedDetailsUrl,omitempty"`
	Severity          *SecurityHostReputationRuleSeverity `json:"severity,omitempty"`
}
