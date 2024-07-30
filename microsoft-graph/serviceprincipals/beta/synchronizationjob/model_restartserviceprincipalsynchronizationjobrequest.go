package synchronizationjob

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestartServicePrincipalSynchronizationJobRequest struct {
	Criteria *SynchronizationJobRestartCriteria `json:"criteria,omitempty"`
}
