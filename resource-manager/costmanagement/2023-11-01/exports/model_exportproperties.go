package exports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportProperties struct {
	Definition          ExportDefinition           `json:"definition"`
	DeliveryInfo        ExportDeliveryInfo         `json:"deliveryInfo"`
	Format              *FormatType                `json:"format,omitempty"`
	NextRunTimeEstimate *string                    `json:"nextRunTimeEstimate,omitempty"`
	PartitionData       *bool                      `json:"partitionData,omitempty"`
	RunHistory          *ExportExecutionListResult `json:"runHistory,omitempty"`
	Schedule            *ExportSchedule            `json:"schedule,omitempty"`
}
