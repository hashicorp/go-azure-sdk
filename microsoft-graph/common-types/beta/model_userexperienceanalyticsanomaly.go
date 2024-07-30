package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomaly struct {
	AnomalyFirstOccurrenceDateTime  *string                                    `json:"anomalyFirstOccurrenceDateTime,omitempty"`
	AnomalyId                       *string                                    `json:"anomalyId,omitempty"`
	AnomalyLatestOccurrenceDateTime *string                                    `json:"anomalyLatestOccurrenceDateTime,omitempty"`
	AnomalyName                     *string                                    `json:"anomalyName,omitempty"`
	AnomalyType                     *UserExperienceAnalyticsAnomalyAnomalyType `json:"anomalyType,omitempty"`
	AssetName                       *string                                    `json:"assetName,omitempty"`
	AssetPublisher                  *string                                    `json:"assetPublisher,omitempty"`
	AssetVersion                    *string                                    `json:"assetVersion,omitempty"`
	DetectionModelId                *string                                    `json:"detectionModelId,omitempty"`
	DeviceImpactedCount             *int64                                     `json:"deviceImpactedCount,omitempty"`
	Id                              *string                                    `json:"id,omitempty"`
	IssueId                         *string                                    `json:"issueId,omitempty"`
	ODataType                       *string                                    `json:"@odata.type,omitempty"`
	Severity                        *UserExperienceAnalyticsAnomalySeverity    `json:"severity,omitempty"`
	State                           *UserExperienceAnalyticsAnomalyState       `json:"state,omitempty"`
}
