package nodereports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscReportResource struct {
	DependsOn         *[]DscReportResourceNavigation `json:"dependsOn,omitempty"`
	DurationInSeconds *float64                       `json:"durationInSeconds,omitempty"`
	Error             *string                        `json:"error,omitempty"`
	ModuleName        *string                        `json:"moduleName,omitempty"`
	ModuleVersion     *string                        `json:"moduleVersion,omitempty"`
	ResourceId        *string                        `json:"resourceId,omitempty"`
	ResourceName      *string                        `json:"resourceName,omitempty"`
	SourceInfo        *string                        `json:"sourceInfo,omitempty"`
	StartDate         *string                        `json:"startDate,omitempty"`
	Status            *string                        `json:"status,omitempty"`
}
