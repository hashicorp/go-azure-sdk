package datamove

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerDataMoveRequest struct {
	CorrelationId         string        `json:"correlationId"`
	DataMoveLevel         DataMoveLevel `json:"dataMoveLevel"`
	PauseGC               *bool         `json:"pauseGC,omitempty"`
	SourceContainerArmIds *[]string     `json:"sourceContainerArmIds,omitempty"`
	SourceRegion          string        `json:"sourceRegion"`
	SourceResourceId      string        `json:"sourceResourceId"`
}
