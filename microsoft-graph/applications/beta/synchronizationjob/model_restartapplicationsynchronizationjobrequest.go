package synchronizationjob

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestartApplicationSynchronizationJobRequest struct {
	Criteria *SynchronizationJobRestartCriteria `json:"criteria,omitempty"`
}
