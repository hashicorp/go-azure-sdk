package policytokens

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyTokenResponse struct {
	ChangeReference *string                                       `json:"changeReference,omitempty"`
	Expiration      *string                                       `json:"expiration,omitempty"`
	Message         *string                                       `json:"message,omitempty"`
	Result          *PolicyTokenResult                            `json:"result,omitempty"`
	Results         *[]ExternalEvaluationEndpointInvocationResult `json:"results,omitempty"`
	RetryAfter      *string                                       `json:"retryAfter,omitempty"`
	Token           *string                                       `json:"token,omitempty"`
	TokenId         *string                                       `json:"tokenId,omitempty"`
}

func (o *PolicyTokenResponse) GetExpirationAsTime() (*time.Time, error) {
	if o.Expiration == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Expiration, "2006-01-02T15:04:05Z07:00")
}

func (o *PolicyTokenResponse) SetExpirationAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Expiration = &formatted
}

func (o *PolicyTokenResponse) GetRetryAfterAsTime() (*time.Time, error) {
	if o.RetryAfter == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RetryAfter, "2006-01-02T15:04:05Z07:00")
}

func (o *PolicyTokenResponse) SetRetryAfterAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RetryAfter = &formatted
}
