package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NspProfile struct {
	AccessRules               *[]NspProfileAccessRule `json:"accessRules,omitempty"`
	AccessRulesVersion        *string                 `json:"accessRulesVersion,omitempty"`
	DiagnosticSettingsVersion *string                 `json:"diagnosticSettingsVersion,omitempty"`
	EnabledLogCategories      *[]string               `json:"enabledLogCategories,omitempty"`
	Name                      *string                 `json:"name,omitempty"`
}
