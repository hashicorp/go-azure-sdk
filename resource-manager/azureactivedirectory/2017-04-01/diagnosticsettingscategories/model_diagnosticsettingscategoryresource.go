package diagnosticsettingscategories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticSettingsCategoryResource struct {
	Id         *string                     `json:"id,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *DiagnosticSettingsCategory `json:"properties,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}
