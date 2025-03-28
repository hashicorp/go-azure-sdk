package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GraphQuery struct {
	BaseQuery  string `json:"baseQuery"`
	Legend     string `json:"legend"`
	MetricName string `json:"metricName"`
}
