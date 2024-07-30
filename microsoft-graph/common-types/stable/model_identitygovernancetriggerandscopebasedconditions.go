package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTriggerAndScopeBasedConditions struct {
	ODataType *string                                     `json:"@odata.type,omitempty"`
	Scope     *SubjectSet                                 `json:"scope,omitempty"`
	Trigger   *IdentityGovernanceWorkflowExecutionTrigger `json:"trigger,omitempty"`
}
