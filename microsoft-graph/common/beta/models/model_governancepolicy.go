package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernancePolicy struct {
	DecisionMakerCriteria *[]GovernanceCriteria         `json:"decisionMakerCriteria,omitempty"`
	NotificationPolicy    *GovernanceNotificationPolicy `json:"notificationPolicy,omitempty"`
	ODataType             *string                       `json:"@odata.type,omitempty"`
}
