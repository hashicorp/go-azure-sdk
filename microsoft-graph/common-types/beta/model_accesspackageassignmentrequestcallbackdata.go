package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestCallbackData struct {
	CustomExtensionStageInstanceDetail *string                                          `json:"customExtensionStageInstanceDetail,omitempty"`
	CustomExtensionStageInstanceId     *string                                          `json:"customExtensionStageInstanceId,omitempty"`
	ODataType                          *string                                          `json:"@odata.type,omitempty"`
	Stage                              *AccessPackageAssignmentRequestCallbackDataStage `json:"stage,omitempty"`
	State                              *string                                          `json:"state,omitempty"`
}
