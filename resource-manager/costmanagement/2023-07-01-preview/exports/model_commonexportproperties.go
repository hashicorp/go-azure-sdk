package exports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonExportProperties struct {
	CompressionMode         *CompressionModeType       `json:"compressionMode,omitempty"`
	DataOverwriteBehavior   *DataOverwriteBehaviorType `json:"dataOverwriteBehavior,omitempty"`
	Definition              ExportDefinition           `json:"definition"`
	DeliveryInfo            ExportDeliveryInfo         `json:"deliveryInfo"`
	ExportDescription       *string                    `json:"exportDescription,omitempty"`
	Format                  *FormatType                `json:"format,omitempty"`
	NextRunTimeEstimate     *string                    `json:"nextRunTimeEstimate,omitempty"`
	PartitionData           *bool                      `json:"partitionData,omitempty"`
	RunHistory              *ExportExecutionListResult `json:"runHistory,omitempty"`
	SystemSuspensionContext *ExportSuspensionContext   `json:"systemSuspensionContext,omitempty"`
}
