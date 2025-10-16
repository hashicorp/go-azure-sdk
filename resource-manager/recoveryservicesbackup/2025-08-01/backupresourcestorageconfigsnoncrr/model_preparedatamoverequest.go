package backupresourcestorageconfigsnoncrr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrepareDataMoveRequest struct {
	DataMoveLevel         DataMoveLevel `json:"dataMoveLevel"`
	IgnoreMoved           *bool         `json:"ignoreMoved,omitempty"`
	SourceContainerArmIds *[]string     `json:"sourceContainerArmIds,omitempty"`
	TargetRegion          string        `json:"targetRegion"`
	TargetResourceId      string        `json:"targetResourceId"`
}
