package subscriptionfeatureregistrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationProfile struct {
	ApprovedTime      *string `json:"approvedTime,omitempty"`
	Approver          *string `json:"approver,omitempty"`
	RequestedTime     *string `json:"requestedTime,omitempty"`
	Requester         *string `json:"requester,omitempty"`
	RequesterObjectId *string `json:"requesterObjectId,omitempty"`
}
