package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecurityState struct {
	DestinationServiceIp   *string `json:"destinationServiceIp,omitempty"`
	DestinationServiceName *string `json:"destinationServiceName,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	RiskScore              *string `json:"riskScore,omitempty"`
}
