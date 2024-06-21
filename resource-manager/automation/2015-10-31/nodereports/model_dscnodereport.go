package nodereports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeReport struct {
	ConfigurationVersion *string               `json:"configurationVersion,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
	Errors               *[]DscReportError     `json:"errors,omitempty"`
	HostName             *string               `json:"hostName,omitempty"`
	IPV4Addresses        *[]string             `json:"iPV4Addresses,omitempty"`
	IPV6Addresses        *[]string             `json:"iPV6Addresses,omitempty"`
	Id                   *string               `json:"id,omitempty"`
	LastModifiedTime     *string               `json:"lastModifiedTime,omitempty"`
	MetaConfiguration    *DscMetaConfiguration `json:"metaConfiguration,omitempty"`
	NumberOfResources    *int64                `json:"numberOfResources,omitempty"`
	RawErrors            *string               `json:"rawErrors,omitempty"`
	RebootRequested      *string               `json:"rebootRequested,omitempty"`
	RefreshMode          *string               `json:"refreshMode,omitempty"`
	ReportFormatVersion  *string               `json:"reportFormatVersion,omitempty"`
	ReportId             *string               `json:"reportId,omitempty"`
	Resources            *[]DscReportResource  `json:"resources,omitempty"`
	StartTime            *string               `json:"startTime,omitempty"`
	Status               *string               `json:"status,omitempty"`
	Type                 *string               `json:"type,omitempty"`
}
