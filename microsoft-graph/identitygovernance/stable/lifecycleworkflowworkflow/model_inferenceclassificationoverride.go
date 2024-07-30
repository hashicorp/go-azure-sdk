package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceClassificationOverride struct {
	ClassifyAs         *InferenceClassificationOverrideClassifyAs `json:"classifyAs,omitempty"`
	Id                 *string                                    `json:"id,omitempty"`
	ODataType          *string                                    `json:"@odata.type,omitempty"`
	SenderEmailAddress *EmailAddress                              `json:"senderEmailAddress,omitempty"`
}
