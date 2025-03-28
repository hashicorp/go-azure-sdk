package policytokens

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalEvaluationEndpointInvocationResult struct {
	Claims     *interface{}            `json:"claims,omitempty"`
	Expiration *string                 `json:"expiration,omitempty"`
	Message    *string                 `json:"message,omitempty"`
	PolicyInfo *PolicyLogInfo          `json:"policyInfo,omitempty"`
	Result     *ExternalEndpointResult `json:"result,omitempty"`
	RetryAfter *string                 `json:"retryAfter,omitempty"`
}

func (o *ExternalEvaluationEndpointInvocationResult) GetExpirationAsTime() (*time.Time, error) {
	if o.Expiration == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Expiration, "2006-01-02T15:04:05Z07:00")
}

func (o *ExternalEvaluationEndpointInvocationResult) SetExpirationAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Expiration = &formatted
}

func (o *ExternalEvaluationEndpointInvocationResult) GetRetryAfterAsTime() (*time.Time, error) {
	if o.RetryAfter == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RetryAfter, "2006-01-02T15:04:05Z07:00")
}

func (o *ExternalEvaluationEndpointInvocationResult) SetRetryAfterAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RetryAfter = &formatted
}
