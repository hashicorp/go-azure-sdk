package monitoringmetricconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricConfiguration struct {
	CounterSets     []MetricCounterSet `json:"counterSets"`
	MdmAccount      *string            `json:"mdmAccount,omitempty"`
	MetricNameSpace *string            `json:"metricNameSpace,omitempty"`
	ResourceId      string             `json:"resourceId"`
}
