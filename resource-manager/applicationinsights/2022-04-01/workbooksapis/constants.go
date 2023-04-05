package workbooksapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CategoryType string

const (
	CategoryTypePerformance CategoryType = "performance"
	CategoryTypeRetention   CategoryType = "retention"
	CategoryTypeTSG         CategoryType = "TSG"
	CategoryTypeWorkbook    CategoryType = "workbook"
)

func PossibleValuesForCategoryType() []string {
	return []string{
		string(CategoryTypePerformance),
		string(CategoryTypeRetention),
		string(CategoryTypeTSG),
		string(CategoryTypeWorkbook),
	}
}

type WorkbookSharedTypeKind string

const (
	WorkbookSharedTypeKindShared WorkbookSharedTypeKind = "shared"
)

func PossibleValuesForWorkbookSharedTypeKind() []string {
	return []string{
		string(WorkbookSharedTypeKindShared),
	}
}

type WorkbookUpdateSharedTypeKind string

const (
	WorkbookUpdateSharedTypeKindShared WorkbookUpdateSharedTypeKind = "shared"
)

func PossibleValuesForWorkbookUpdateSharedTypeKind() []string {
	return []string{
		string(WorkbookUpdateSharedTypeKindShared),
	}
}
