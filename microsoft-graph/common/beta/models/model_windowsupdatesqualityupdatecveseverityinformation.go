package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCveSeverityInformation struct {
	ExploitedCves *[]WindowsUpdatesCveInformation                               `json:"exploitedCves,omitempty"`
	MaxSeverity   *WindowsUpdatesQualityUpdateCveSeverityInformationMaxSeverity `json:"maxSeverity,omitempty"`
	ODataType     *string                                                       `json:"@odata.type,omitempty"`
}
