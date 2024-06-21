package activitylogs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventData struct {
	Authorization        *SenderAuthorization `json:"authorization,omitempty"`
	Caller               *string              `json:"caller,omitempty"`
	Category             *LocalizableString   `json:"category,omitempty"`
	Claims               *map[string]string   `json:"claims,omitempty"`
	CorrelationId        *string              `json:"correlationId,omitempty"`
	Description          *string              `json:"description,omitempty"`
	EventDataId          *string              `json:"eventDataId,omitempty"`
	EventName            *LocalizableString   `json:"eventName,omitempty"`
	EventTimestamp       *string              `json:"eventTimestamp,omitempty"`
	HTTPRequest          *HTTPRequestInfo     `json:"httpRequest,omitempty"`
	Id                   *string              `json:"id,omitempty"`
	Level                *EventLevel          `json:"level,omitempty"`
	OperationId          *string              `json:"operationId,omitempty"`
	OperationName        *LocalizableString   `json:"operationName,omitempty"`
	Properties           *map[string]string   `json:"properties,omitempty"`
	ResourceGroupName    *string              `json:"resourceGroupName,omitempty"`
	ResourceId           *string              `json:"resourceId,omitempty"`
	ResourceProviderName *LocalizableString   `json:"resourceProviderName,omitempty"`
	ResourceType         *LocalizableString   `json:"resourceType,omitempty"`
	Status               *LocalizableString   `json:"status,omitempty"`
	SubStatus            *LocalizableString   `json:"subStatus,omitempty"`
	SubmissionTimestamp  *string              `json:"submissionTimestamp,omitempty"`
	SubscriptionId       *string              `json:"subscriptionId,omitempty"`
	TenantId             *string              `json:"tenantId,omitempty"`
}
