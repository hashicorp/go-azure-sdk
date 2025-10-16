package incidentbookmarks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPEntityProperties struct {
	AdditionalData     *interface{}          `json:"additionalData,omitempty"`
	Address            *string               `json:"address,omitempty"`
	FriendlyName       *string               `json:"friendlyName,omitempty"`
	Location           *GeoLocation          `json:"location,omitempty"`
	ThreatIntelligence *[]ThreatIntelligence `json:"threatIntelligence,omitempty"`
}
