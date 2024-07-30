package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRuleActions struct {
	AssignCategories      *[]string                         `json:"assignCategories,omitempty"`
	CopyToFolder          *string                           `json:"copyToFolder,omitempty"`
	Delete                *bool                             `json:"delete,omitempty"`
	ForwardAsAttachmentTo *[]Recipient                      `json:"forwardAsAttachmentTo,omitempty"`
	ForwardTo             *[]Recipient                      `json:"forwardTo,omitempty"`
	MarkAsRead            *bool                             `json:"markAsRead,omitempty"`
	MarkImportance        *MessageRuleActionsMarkImportance `json:"markImportance,omitempty"`
	MoveToFolder          *string                           `json:"moveToFolder,omitempty"`
	ODataType             *string                           `json:"@odata.type,omitempty"`
	PermanentDelete       *bool                             `json:"permanentDelete,omitempty"`
	RedirectTo            *[]Recipient                      `json:"redirectTo,omitempty"`
	StopProcessingRules   *bool                             `json:"stopProcessingRules,omitempty"`
}
