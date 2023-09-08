package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutInstance struct {
	CustomExtensionId     *string                               `json:"customExtensionId,omitempty"`
	Detail                *string                               `json:"detail,omitempty"`
	ExternalCorrelationId *string                               `json:"externalCorrelationId,omitempty"`
	Id                    *string                               `json:"id,omitempty"`
	ODataType             *string                               `json:"@odata.type,omitempty"`
	Status                *CustomExtensionCalloutInstanceStatus `json:"status,omitempty"`
}
