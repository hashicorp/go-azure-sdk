package lifecycleworkflowworkflowtasktaskprocessingresult

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtensionCallbackData struct {
	ODataType       *string                                                           `json:"@odata.type,omitempty"`
	OperationStatus *IdentityGovernanceCustomTaskExtensionCallbackDataOperationStatus `json:"operationStatus,omitempty"`
}
