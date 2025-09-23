package submittedresourcerequeststatuses

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubmittedResourceRequestStatusProperties struct {
	FaultCode         *string                `json:"faultCode,omitempty"`
	ProvisioningState *RequestState          `json:"provisioningState,omitempty"`
	RequestSubmitTime *string                `json:"requestSubmitTime,omitempty"`
	RequestedResource *GroupQuotaRequestBase `json:"requestedResource,omitempty"`
}

func (o *SubmittedResourceRequestStatusProperties) GetRequestSubmitTimeAsTime() (*time.Time, error) {
	if o.RequestSubmitTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RequestSubmitTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SubmittedResourceRequestStatusProperties) SetRequestSubmitTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RequestSubmitTime = &formatted
}
