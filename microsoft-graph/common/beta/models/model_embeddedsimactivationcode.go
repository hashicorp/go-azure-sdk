package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCode struct {
	IntegratedCircuitCardIdentifier *string `json:"integratedCircuitCardIdentifier,omitempty"`
	MatchingIdentifier              *string `json:"matchingIdentifier,omitempty"`
	ODataType                       *string `json:"@odata.type,omitempty"`
	SmdpPlusServerAddress           *string `json:"smdpPlusServerAddress,omitempty"`
}
