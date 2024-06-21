package views

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ViewProperties struct {
	Accumulated *AccumulatedType        `json:"accumulated,omitempty"`
	Chart       *ChartType              `json:"chart,omitempty"`
	CreatedOn   *string                 `json:"createdOn,omitempty"`
	Currency    *string                 `json:"currency,omitempty"`
	DateRange   *string                 `json:"dateRange,omitempty"`
	DisplayName *string                 `json:"displayName,omitempty"`
	Kpis        *[]KpiProperties        `json:"kpis,omitempty"`
	Metric      *MetricType             `json:"metric,omitempty"`
	ModifiedOn  *string                 `json:"modifiedOn,omitempty"`
	Pivots      *[]PivotProperties      `json:"pivots,omitempty"`
	Query       *ReportConfigDefinition `json:"query,omitempty"`
	Scope       *string                 `json:"scope,omitempty"`
}
