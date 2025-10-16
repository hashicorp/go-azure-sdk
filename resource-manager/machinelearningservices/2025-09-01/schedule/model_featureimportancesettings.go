package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureImportanceSettings struct {
	Mode         *FeatureImportanceMode `json:"mode,omitempty"`
	TargetColumn *string                `json:"targetColumn,omitempty"`
}
