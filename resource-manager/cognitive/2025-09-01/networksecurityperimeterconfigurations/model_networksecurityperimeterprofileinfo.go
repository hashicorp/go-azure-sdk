package networksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterProfileInfo struct {
	AccessRules               *[]NetworkSecurityPerimeterAccessRule `json:"accessRules,omitempty"`
	AccessRulesVersion        *int64                                `json:"accessRulesVersion,omitempty"`
	DiagnosticSettingsVersion *int64                                `json:"diagnosticSettingsVersion,omitempty"`
	EnabledLogCategories      *[]string                             `json:"enabledLogCategories,omitempty"`
	Name                      *string                               `json:"name,omitempty"`
}
