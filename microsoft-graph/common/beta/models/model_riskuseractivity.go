package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskUserActivity struct {
	Detail         *RiskUserActivityDetail       `json:"detail,omitempty"`
	EventTypes     *[]RiskUserActivityEventTypes `json:"eventTypes,omitempty"`
	ODataType      *string                       `json:"@odata.type,omitempty"`
	RiskEventTypes *[]string                     `json:"riskEventTypes,omitempty"`
}
