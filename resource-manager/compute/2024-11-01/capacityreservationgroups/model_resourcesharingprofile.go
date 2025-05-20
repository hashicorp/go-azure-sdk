package capacityreservationgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSharingProfile struct {
	SubscriptionIds *[]SubResource `json:"subscriptionIds,omitempty"`
}
