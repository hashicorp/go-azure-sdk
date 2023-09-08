package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProtectionRoot struct {
	ODataType                      *string                          `json:"@odata.type,omitempty"`
	RiskDetections                 *[]RiskDetection                 `json:"riskDetections,omitempty"`
	RiskyServicePrincipals         *[]RiskyServicePrincipal         `json:"riskyServicePrincipals,omitempty"`
	RiskyUsers                     *[]RiskyUser                     `json:"riskyUsers,omitempty"`
	ServicePrincipalRiskDetections *[]ServicePrincipalRiskDetection `json:"servicePrincipalRiskDetections,omitempty"`
}
