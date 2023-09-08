package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptError struct {
	Code                                       *DeviceComplianceScriptErrorCode                                       `json:"code,omitempty"`
	DeviceComplianceScriptRulesValidationError *DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError `json:"deviceComplianceScriptRulesValidationError,omitempty"`
	Message                                    *string                                                                `json:"message,omitempty"`
	ODataType                                  *string                                                                `json:"@odata.type,omitempty"`
}
