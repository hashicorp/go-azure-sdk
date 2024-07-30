package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpEvaluatePoliciesRequest struct {
	EvaluationInput  *DlpEvaluationInput `json:"evaluationInput,omitempty"`
	NotificationInfo *DlpNotification    `json:"notificationInfo,omitempty"`
	ODataType        *string             `json:"@odata.type,omitempty"`
	Target           *string             `json:"target,omitempty"`
}
