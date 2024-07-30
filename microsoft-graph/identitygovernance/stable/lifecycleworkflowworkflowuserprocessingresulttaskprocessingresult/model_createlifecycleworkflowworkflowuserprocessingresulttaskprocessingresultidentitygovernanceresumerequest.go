package lifecycleworkflowworkflowuserprocessingresulttaskprocessingresult

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateLifecycleWorkflowWorkflowUserProcessingResultTaskProcessingResultIdentityGovernanceResumeRequest struct {
	Data   *IdentityGovernanceCustomTaskExtensionCallbackData `json:"data,omitempty"`
	Source *string                                            `json:"source,omitempty"`
	Type   *string                                            `json:"type,omitempty"`
}
