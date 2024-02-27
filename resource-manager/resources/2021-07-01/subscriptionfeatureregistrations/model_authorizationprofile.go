package subscriptionfeatureregistrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationProfile struct {
	ApprovedTime      *string `json:"approvedTime,omitempty"`
	Approver          *string `json:"approver,omitempty"`
	RequestedTime     *string `json:"requestedTime,omitempty"`
	Requester         *string `json:"requester,omitempty"`
	RequesterObjectId *string `json:"requesterObjectId,omitempty"`
}

func (o *AuthorizationProfile) GetApprovedTimeAsTime() (*time.Time, error) {
	if o.ApprovedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ApprovedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AuthorizationProfile) GetRequestedTimeAsTime() (*time.Time, error) {
	if o.RequestedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RequestedTime, "2006-01-02T15:04:05Z07:00")
}
