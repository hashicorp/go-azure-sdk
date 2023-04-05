package budgets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BudgetOperatorType string

const (
	BudgetOperatorTypeIn BudgetOperatorType = "In"
)

func PossibleValuesForBudgetOperatorType() []string {
	return []string{
		string(BudgetOperatorTypeIn),
	}
}

type CategoryType string

const (
	CategoryTypeCost CategoryType = "Cost"
)

func PossibleValuesForCategoryType() []string {
	return []string{
		string(CategoryTypeCost),
	}
}

type CultureCode string

const (
	CultureCodeCsNegativecz CultureCode = "cs-cz"
	CultureCodeDaNegativedk CultureCode = "da-dk"
	CultureCodeDeNegativede CultureCode = "de-de"
	CultureCodeEnNegativegb CultureCode = "en-gb"
	CultureCodeEnNegativeus CultureCode = "en-us"
	CultureCodeEsNegativees CultureCode = "es-es"
	CultureCodeFrNegativefr CultureCode = "fr-fr"
	CultureCodeHuNegativehu CultureCode = "hu-hu"
	CultureCodeItNegativeit CultureCode = "it-it"
	CultureCodeJaNegativejp CultureCode = "ja-jp"
	CultureCodeKoNegativekr CultureCode = "ko-kr"
	CultureCodeNbNegativeno CultureCode = "nb-no"
	CultureCodeNlNegativenl CultureCode = "nl-nl"
	CultureCodePlNegativepl CultureCode = "pl-pl"
	CultureCodePtNegativebr CultureCode = "pt-br"
	CultureCodePtNegativept CultureCode = "pt-pt"
	CultureCodeRuNegativeru CultureCode = "ru-ru"
	CultureCodeSvNegativese CultureCode = "sv-se"
	CultureCodeTrNegativetr CultureCode = "tr-tr"
	CultureCodeZhNegativecn CultureCode = "zh-cn"
	CultureCodeZhNegativetw CultureCode = "zh-tw"
)

func PossibleValuesForCultureCode() []string {
	return []string{
		string(CultureCodeCsNegativecz),
		string(CultureCodeDaNegativedk),
		string(CultureCodeDeNegativede),
		string(CultureCodeEnNegativegb),
		string(CultureCodeEnNegativeus),
		string(CultureCodeEsNegativees),
		string(CultureCodeFrNegativefr),
		string(CultureCodeHuNegativehu),
		string(CultureCodeItNegativeit),
		string(CultureCodeJaNegativejp),
		string(CultureCodeKoNegativekr),
		string(CultureCodeNbNegativeno),
		string(CultureCodeNlNegativenl),
		string(CultureCodePlNegativepl),
		string(CultureCodePtNegativebr),
		string(CultureCodePtNegativept),
		string(CultureCodeRuNegativeru),
		string(CultureCodeSvNegativese),
		string(CultureCodeTrNegativetr),
		string(CultureCodeZhNegativecn),
		string(CultureCodeZhNegativetw),
	}
}

type OperatorType string

const (
	OperatorTypeEqualTo              OperatorType = "EqualTo"
	OperatorTypeGreaterThan          OperatorType = "GreaterThan"
	OperatorTypeGreaterThanOrEqualTo OperatorType = "GreaterThanOrEqualTo"
)

func PossibleValuesForOperatorType() []string {
	return []string{
		string(OperatorTypeEqualTo),
		string(OperatorTypeGreaterThan),
		string(OperatorTypeGreaterThanOrEqualTo),
	}
}

type ThresholdType string

const (
	ThresholdTypeActual     ThresholdType = "Actual"
	ThresholdTypeForecasted ThresholdType = "Forecasted"
)

func PossibleValuesForThresholdType() []string {
	return []string{
		string(ThresholdTypeActual),
		string(ThresholdTypeForecasted),
	}
}

type TimeGrainType string

const (
	TimeGrainTypeAnnually       TimeGrainType = "Annually"
	TimeGrainTypeBillingAnnual  TimeGrainType = "BillingAnnual"
	TimeGrainTypeBillingMonth   TimeGrainType = "BillingMonth"
	TimeGrainTypeBillingQuarter TimeGrainType = "BillingQuarter"
	TimeGrainTypeMonthly        TimeGrainType = "Monthly"
	TimeGrainTypeQuarterly      TimeGrainType = "Quarterly"
)

func PossibleValuesForTimeGrainType() []string {
	return []string{
		string(TimeGrainTypeAnnually),
		string(TimeGrainTypeBillingAnnual),
		string(TimeGrainTypeBillingMonth),
		string(TimeGrainTypeBillingQuarter),
		string(TimeGrainTypeMonthly),
		string(TimeGrainTypeQuarterly),
	}
}
