package certificateordersdiagnostics

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorType string

const (
	DetectorTypeAnalysis         DetectorType = "Analysis"
	DetectorTypeCategoryOverview DetectorType = "CategoryOverview"
	DetectorTypeDetector         DetectorType = "Detector"
)

func PossibleValuesForDetectorType() []string {
	return []string{
		string(DetectorTypeAnalysis),
		string(DetectorTypeCategoryOverview),
		string(DetectorTypeDetector),
	}
}

func parseDetectorType(input string) (*DetectorType, error) {
	vals := map[string]DetectorType{
		"analysis":         DetectorTypeAnalysis,
		"categoryoverview": DetectorTypeCategoryOverview,
		"detector":         DetectorTypeDetector,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectorType(input)
	return &out, nil
}

type InsightStatus string

const (
	InsightStatusCritical InsightStatus = "Critical"
	InsightStatusInfo     InsightStatus = "Info"
	InsightStatusNone     InsightStatus = "None"
	InsightStatusSuccess  InsightStatus = "Success"
	InsightStatusWarning  InsightStatus = "Warning"
)

func PossibleValuesForInsightStatus() []string {
	return []string{
		string(InsightStatusCritical),
		string(InsightStatusInfo),
		string(InsightStatusNone),
		string(InsightStatusSuccess),
		string(InsightStatusWarning),
	}
}

func parseInsightStatus(input string) (*InsightStatus, error) {
	vals := map[string]InsightStatus{
		"critical": InsightStatusCritical,
		"info":     InsightStatusInfo,
		"none":     InsightStatusNone,
		"success":  InsightStatusSuccess,
		"warning":  InsightStatusWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InsightStatus(input)
	return &out, nil
}

type RenderingType string

const (
	RenderingTypeAppInsight               RenderingType = "AppInsight"
	RenderingTypeAppInsightEnablement     RenderingType = "AppInsightEnablement"
	RenderingTypeCard                     RenderingType = "Card"
	RenderingTypeChangeAnalysisOnboarding RenderingType = "ChangeAnalysisOnboarding"
	RenderingTypeChangeSets               RenderingType = "ChangeSets"
	RenderingTypeChangesView              RenderingType = "ChangesView"
	RenderingTypeDataSummary              RenderingType = "DataSummary"
	RenderingTypeDependencyGraph          RenderingType = "DependencyGraph"
	RenderingTypeDetector                 RenderingType = "Detector"
	RenderingTypeDownTime                 RenderingType = "DownTime"
	RenderingTypeDropDown                 RenderingType = "DropDown"
	RenderingTypeDynamicInsight           RenderingType = "DynamicInsight"
	RenderingTypeEmail                    RenderingType = "Email"
	RenderingTypeForm                     RenderingType = "Form"
	RenderingTypeGuage                    RenderingType = "Guage"
	RenderingTypeInsights                 RenderingType = "Insights"
	RenderingTypeMarkdown                 RenderingType = "Markdown"
	RenderingTypeNoGraph                  RenderingType = "NoGraph"
	RenderingTypePieChart                 RenderingType = "PieChart"
	RenderingTypeSearchComponent          RenderingType = "SearchComponent"
	RenderingTypeSolution                 RenderingType = "Solution"
	RenderingTypeSummaryCard              RenderingType = "SummaryCard"
	RenderingTypeTable                    RenderingType = "Table"
	RenderingTypeTimeSeries               RenderingType = "TimeSeries"
	RenderingTypeTimeSeriesPerInstance    RenderingType = "TimeSeriesPerInstance"
)

func PossibleValuesForRenderingType() []string {
	return []string{
		string(RenderingTypeAppInsight),
		string(RenderingTypeAppInsightEnablement),
		string(RenderingTypeCard),
		string(RenderingTypeChangeAnalysisOnboarding),
		string(RenderingTypeChangeSets),
		string(RenderingTypeChangesView),
		string(RenderingTypeDataSummary),
		string(RenderingTypeDependencyGraph),
		string(RenderingTypeDetector),
		string(RenderingTypeDownTime),
		string(RenderingTypeDropDown),
		string(RenderingTypeDynamicInsight),
		string(RenderingTypeEmail),
		string(RenderingTypeForm),
		string(RenderingTypeGuage),
		string(RenderingTypeInsights),
		string(RenderingTypeMarkdown),
		string(RenderingTypeNoGraph),
		string(RenderingTypePieChart),
		string(RenderingTypeSearchComponent),
		string(RenderingTypeSolution),
		string(RenderingTypeSummaryCard),
		string(RenderingTypeTable),
		string(RenderingTypeTimeSeries),
		string(RenderingTypeTimeSeriesPerInstance),
	}
}

func parseRenderingType(input string) (*RenderingType, error) {
	vals := map[string]RenderingType{
		"appinsight":               RenderingTypeAppInsight,
		"appinsightenablement":     RenderingTypeAppInsightEnablement,
		"card":                     RenderingTypeCard,
		"changeanalysisonboarding": RenderingTypeChangeAnalysisOnboarding,
		"changesets":               RenderingTypeChangeSets,
		"changesview":              RenderingTypeChangesView,
		"datasummary":              RenderingTypeDataSummary,
		"dependencygraph":          RenderingTypeDependencyGraph,
		"detector":                 RenderingTypeDetector,
		"downtime":                 RenderingTypeDownTime,
		"dropdown":                 RenderingTypeDropDown,
		"dynamicinsight":           RenderingTypeDynamicInsight,
		"email":                    RenderingTypeEmail,
		"form":                     RenderingTypeForm,
		"guage":                    RenderingTypeGuage,
		"insights":                 RenderingTypeInsights,
		"markdown":                 RenderingTypeMarkdown,
		"nograph":                  RenderingTypeNoGraph,
		"piechart":                 RenderingTypePieChart,
		"searchcomponent":          RenderingTypeSearchComponent,
		"solution":                 RenderingTypeSolution,
		"summarycard":              RenderingTypeSummaryCard,
		"table":                    RenderingTypeTable,
		"timeseries":               RenderingTypeTimeSeries,
		"timeseriesperinstance":    RenderingTypeTimeSeriesPerInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RenderingType(input)
	return &out, nil
}
