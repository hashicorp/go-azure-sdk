package triggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerSubscriptionOperationStatus struct {
	Status      *EventSubscriptionStatus `json:"status,omitempty"`
	TriggerName *string                  `json:"triggerName,omitempty"`
}
