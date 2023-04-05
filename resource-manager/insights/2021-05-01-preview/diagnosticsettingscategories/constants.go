package diagnosticsettingscategories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CategoryType string

const (
	CategoryTypeLogs    CategoryType = "Logs"
	CategoryTypeMetrics CategoryType = "Metrics"
)

func PossibleValuesForCategoryType() []string {
	return []string{
		string(CategoryTypeLogs),
		string(CategoryTypeMetrics),
	}
}
