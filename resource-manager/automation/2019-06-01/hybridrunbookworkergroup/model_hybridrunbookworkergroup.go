package hybridrunbookworkergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridRunbookWorkerGroup struct {
	Credential           *RunAsCredentialAssociationProperty `json:"credential"`
	GroupType            *GroupTypeEnum                      `json:"groupType,omitempty"`
	HybridRunbookWorkers *[]HybridRunbookWorker              `json:"hybridRunbookWorkers,omitempty"`
	Id                   *string                             `json:"id,omitempty"`
	Name                 *string                             `json:"name,omitempty"`
}
