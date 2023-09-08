package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAutomation struct {
	CreatedBy            *EmailIdentity              `json:"createdBy,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	Description          *string                     `json:"description,omitempty"`
	DisplayName          *string                     `json:"displayName,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	LastModifiedBy       *EmailIdentity              `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	LastRunDateTime      *string                     `json:"lastRunDateTime,omitempty"`
	NextRunDateTime      *string                     `json:"nextRunDateTime,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	Runs                 *[]SimulationAutomationRun  `json:"runs,omitempty"`
	Status               *SimulationAutomationStatus `json:"status,omitempty"`
}
