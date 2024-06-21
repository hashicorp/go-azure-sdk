package quotarequests

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaRequestProperties struct {
	Error             *ServiceErrorDetail `json:"error,omitempty"`
	Message           *string             `json:"message,omitempty"`
	ProvisioningState *QuotaRequestState  `json:"provisioningState,omitempty"`
	RequestSubmitTime *string             `json:"requestSubmitTime,omitempty"`
	Value             *[]SubRequest       `json:"value,omitempty"`
}
