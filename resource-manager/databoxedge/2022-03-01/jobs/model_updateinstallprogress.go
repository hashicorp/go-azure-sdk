package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateInstallProgress struct {
	NumberOfUpdatesInstalled *int64 `json:"numberOfUpdatesInstalled,omitempty"`
	NumberOfUpdatesToInstall *int64 `json:"numberOfUpdatesToInstall,omitempty"`
	PercentComplete          *int64 `json:"percentComplete,omitempty"`
}
