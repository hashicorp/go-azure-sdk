package advancedthreatprotectionsettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerThreatProtectionProperties struct {
	CreationTime *string               `json:"creationTime,omitempty"`
	State        ThreatProtectionState `json:"state"`
}
