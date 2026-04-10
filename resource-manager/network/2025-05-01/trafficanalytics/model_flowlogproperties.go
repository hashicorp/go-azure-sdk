package trafficanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlowLogProperties struct {
	Enabled                  bool                             `json:"enabled"`
	EnabledFilteringCriteria *string                          `json:"enabledFilteringCriteria,omitempty"`
	Format                   *CommonFlowLogFormatParameters   `json:"format,omitempty"`
	RecordTypes              *string                          `json:"recordTypes,omitempty"`
	RetentionPolicy          *CommonRetentionPolicyParameters `json:"retentionPolicy,omitempty"`
	StorageId                string                           `json:"storageId"`
}
