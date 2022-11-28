package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayProperties struct {
	ApiMetadataProperties *GatewayApiMetadataProperties `json:"apiMetadataProperties"`
	CorsProperties        *GatewayCorsProperties        `json:"corsProperties"`
	HTTPSOnly             *bool                         `json:"httpsOnly,omitempty"`
	Instances             *[]GatewayInstance            `json:"instances,omitempty"`
	OperatorProperties    *GatewayOperatorProperties    `json:"operatorProperties"`
	ProvisioningState     *GatewayProvisioningState     `json:"provisioningState,omitempty"`
	Public                *bool                         `json:"public,omitempty"`
	ResourceRequests      *GatewayResourceRequests      `json:"resourceRequests"`
	SsoProperties         *SsoProperties                `json:"ssoProperties"`
	Url                   *string                       `json:"url,omitempty"`
}
