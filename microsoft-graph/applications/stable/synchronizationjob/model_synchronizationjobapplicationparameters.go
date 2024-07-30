package synchronizationjob

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobApplicationParameters struct {
	ODataType *string                      `json:"@odata.type,omitempty"`
	RuleId    *string                      `json:"ruleId,omitempty"`
	Subjects  *[]SynchronizationJobSubject `json:"subjects,omitempty"`
}
