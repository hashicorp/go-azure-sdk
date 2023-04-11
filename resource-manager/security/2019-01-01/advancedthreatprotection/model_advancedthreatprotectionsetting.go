package advancedthreatprotection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionSetting struct {
	Id         *string                             `json:"id,omitempty"`
	Name       *string                             `json:"name,omitempty"`
	Properties *AdvancedThreatProtectionProperties `json:"properties,omitempty"`
	Type       *string                             `json:"type,omitempty"`
}
