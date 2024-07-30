package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerInstance struct {
	CustomExtensionId     *string                               `json:"customExtensionId,omitempty"`
	ExternalCorrelationId *string                               `json:"externalCorrelationId,omitempty"`
	ODataType             *string                               `json:"@odata.type,omitempty"`
	Stage                 *CustomExtensionHandlerInstanceStage  `json:"stage,omitempty"`
	Status                *CustomExtensionHandlerInstanceStatus `json:"status,omitempty"`
}
