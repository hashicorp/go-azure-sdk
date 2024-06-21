package serveradvancedthreatprotectionsettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionProperties struct {
	CreationTime *string                       `json:"creationTime,omitempty"`
	State        AdvancedThreatProtectionState `json:"state"`
}
