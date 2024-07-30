package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessagePolicyViolation struct {
	DlpAction         *ChatMessagePolicyViolationDlpAction      `json:"dlpAction,omitempty"`
	JustificationText *string                                   `json:"justificationText,omitempty"`
	ODataType         *string                                   `json:"@odata.type,omitempty"`
	PolicyTip         *ChatMessagePolicyViolationPolicyTip      `json:"policyTip,omitempty"`
	UserAction        *ChatMessagePolicyViolationUserAction     `json:"userAction,omitempty"`
	VerdictDetails    *ChatMessagePolicyViolationVerdictDetails `json:"verdictDetails,omitempty"`
}
