package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAlertRuleDefinitionTemplate struct {
	DefaultSeverity *ManagedTenantsAlertRuleDefinitionTemplateDefaultSeverity `json:"defaultSeverity,omitempty"`
	ODataType       *string                                                   `json:"@odata.type,omitempty"`
}
