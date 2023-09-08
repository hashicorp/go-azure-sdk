package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRuleSetting struct {
	ODataType      *string `json:"@odata.type,omitempty"`
	RuleIdentifier *string `json:"ruleIdentifier,omitempty"`
	Setting        *string `json:"setting,omitempty"`
}
