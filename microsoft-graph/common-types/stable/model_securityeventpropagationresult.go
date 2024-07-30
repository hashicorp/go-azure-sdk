package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventPropagationResult struct {
	Location          *string                               `json:"location,omitempty"`
	ODataType         *string                               `json:"@odata.type,omitempty"`
	ServiceName       *string                               `json:"serviceName,omitempty"`
	Status            *SecurityEventPropagationResultStatus `json:"status,omitempty"`
	StatusInformation *string                               `json:"statusInformation,omitempty"`
}
