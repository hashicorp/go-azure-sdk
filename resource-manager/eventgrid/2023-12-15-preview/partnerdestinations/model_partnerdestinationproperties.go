package partnerdestinations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerDestinationProperties struct {
	ActivationState                 *PartnerDestinationActivationState   `json:"activationState,omitempty"`
	EndpointBaseURL                 *string                              `json:"endpointBaseUrl,omitempty"`
	EndpointServiceContext          *string                              `json:"endpointServiceContext,omitempty"`
	ExpirationTimeIfNotActivatedUtc *string                              `json:"expirationTimeIfNotActivatedUtc,omitempty"`
	MessageForActivation            *string                              `json:"messageForActivation,omitempty"`
	PartnerRegistrationImmutableId  *string                              `json:"partnerRegistrationImmutableId,omitempty"`
	ProvisioningState               *PartnerDestinationProvisioningState `json:"provisioningState,omitempty"`
}

func (o *PartnerDestinationProperties) GetExpirationTimeIfNotActivatedUtcAsTime() (*time.Time, error) {
	if o.ExpirationTimeIfNotActivatedUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationTimeIfNotActivatedUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *PartnerDestinationProperties) SetExpirationTimeIfNotActivatedUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationTimeIfNotActivatedUtc = &formatted
}
