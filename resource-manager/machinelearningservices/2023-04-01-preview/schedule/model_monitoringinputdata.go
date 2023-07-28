package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringInputData struct {
	Asset                    *interface{}               `json:"asset,omitempty"`
	DataContext              MonitoringInputDataContext `json:"dataContext"`
	PreprocessingComponentId *string                    `json:"preprocessingComponentId,omitempty"`
	TargetColumnName         *string                    `json:"targetColumnName,omitempty"`
}
