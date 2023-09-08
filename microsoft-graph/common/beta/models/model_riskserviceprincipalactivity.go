package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskServicePrincipalActivity struct {
	Detail         *RiskServicePrincipalActivityDetail `json:"detail,omitempty"`
	ODataType      *string                             `json:"@odata.type,omitempty"`
	RiskEventTypes *[]string                           `json:"riskEventTypes,omitempty"`
}
