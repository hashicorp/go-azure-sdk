package alerts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertCategory string

const (
	AlertCategoryBilling AlertCategory = "Billing"
	AlertCategoryCost    AlertCategory = "Cost"
	AlertCategorySystem  AlertCategory = "System"
	AlertCategoryUsage   AlertCategory = "Usage"
)

func PossibleValuesForAlertCategory() []string {
	return []string{
		string(AlertCategoryBilling),
		string(AlertCategoryCost),
		string(AlertCategorySystem),
		string(AlertCategoryUsage),
	}
}

type AlertCriteria string

const (
	AlertCriteriaCostThresholdExceeded          AlertCriteria = "CostThresholdExceeded"
	AlertCriteriaCreditThresholdApproaching     AlertCriteria = "CreditThresholdApproaching"
	AlertCriteriaCreditThresholdReached         AlertCriteria = "CreditThresholdReached"
	AlertCriteriaCrossCloudCollectionError      AlertCriteria = "CrossCloudCollectionError"
	AlertCriteriaCrossCloudNewDataAvailable     AlertCriteria = "CrossCloudNewDataAvailable"
	AlertCriteriaForecastCostThresholdExceeded  AlertCriteria = "ForecastCostThresholdExceeded"
	AlertCriteriaForecastUsageThresholdExceeded AlertCriteria = "ForecastUsageThresholdExceeded"
	AlertCriteriaGeneralThresholdError          AlertCriteria = "GeneralThresholdError"
	AlertCriteriaInvoiceDueDateApproaching      AlertCriteria = "InvoiceDueDateApproaching"
	AlertCriteriaInvoiceDueDateReached          AlertCriteria = "InvoiceDueDateReached"
	AlertCriteriaMultiCurrency                  AlertCriteria = "MultiCurrency"
	AlertCriteriaQuotaThresholdApproaching      AlertCriteria = "QuotaThresholdApproaching"
	AlertCriteriaQuotaThresholdReached          AlertCriteria = "QuotaThresholdReached"
	AlertCriteriaUsageThresholdExceeded         AlertCriteria = "UsageThresholdExceeded"
)

func PossibleValuesForAlertCriteria() []string {
	return []string{
		string(AlertCriteriaCostThresholdExceeded),
		string(AlertCriteriaCreditThresholdApproaching),
		string(AlertCriteriaCreditThresholdReached),
		string(AlertCriteriaCrossCloudCollectionError),
		string(AlertCriteriaCrossCloudNewDataAvailable),
		string(AlertCriteriaForecastCostThresholdExceeded),
		string(AlertCriteriaForecastUsageThresholdExceeded),
		string(AlertCriteriaGeneralThresholdError),
		string(AlertCriteriaInvoiceDueDateApproaching),
		string(AlertCriteriaInvoiceDueDateReached),
		string(AlertCriteriaMultiCurrency),
		string(AlertCriteriaQuotaThresholdApproaching),
		string(AlertCriteriaQuotaThresholdReached),
		string(AlertCriteriaUsageThresholdExceeded),
	}
}

type AlertOperator string

const (
	AlertOperatorEqualTo              AlertOperator = "EqualTo"
	AlertOperatorGreaterThan          AlertOperator = "GreaterThan"
	AlertOperatorGreaterThanOrEqualTo AlertOperator = "GreaterThanOrEqualTo"
	AlertOperatorLessThan             AlertOperator = "LessThan"
	AlertOperatorLessThanOrEqualTo    AlertOperator = "LessThanOrEqualTo"
	AlertOperatorNone                 AlertOperator = "None"
)

func PossibleValuesForAlertOperator() []string {
	return []string{
		string(AlertOperatorEqualTo),
		string(AlertOperatorGreaterThan),
		string(AlertOperatorGreaterThanOrEqualTo),
		string(AlertOperatorLessThan),
		string(AlertOperatorLessThanOrEqualTo),
		string(AlertOperatorNone),
	}
}

type AlertSource string

const (
	AlertSourcePreset AlertSource = "Preset"
	AlertSourceUser   AlertSource = "User"
)

func PossibleValuesForAlertSource() []string {
	return []string{
		string(AlertSourcePreset),
		string(AlertSourceUser),
	}
}

type AlertStatus string

const (
	AlertStatusActive     AlertStatus = "Active"
	AlertStatusDismissed  AlertStatus = "Dismissed"
	AlertStatusNone       AlertStatus = "None"
	AlertStatusOverridden AlertStatus = "Overridden"
	AlertStatusResolved   AlertStatus = "Resolved"
)

func PossibleValuesForAlertStatus() []string {
	return []string{
		string(AlertStatusActive),
		string(AlertStatusDismissed),
		string(AlertStatusNone),
		string(AlertStatusOverridden),
		string(AlertStatusResolved),
	}
}

type AlertTimeGrainType string

const (
	AlertTimeGrainTypeAnnually       AlertTimeGrainType = "Annually"
	AlertTimeGrainTypeBillingAnnual  AlertTimeGrainType = "BillingAnnual"
	AlertTimeGrainTypeBillingMonth   AlertTimeGrainType = "BillingMonth"
	AlertTimeGrainTypeBillingQuarter AlertTimeGrainType = "BillingQuarter"
	AlertTimeGrainTypeMonthly        AlertTimeGrainType = "Monthly"
	AlertTimeGrainTypeNone           AlertTimeGrainType = "None"
	AlertTimeGrainTypeQuarterly      AlertTimeGrainType = "Quarterly"
)

func PossibleValuesForAlertTimeGrainType() []string {
	return []string{
		string(AlertTimeGrainTypeAnnually),
		string(AlertTimeGrainTypeBillingAnnual),
		string(AlertTimeGrainTypeBillingMonth),
		string(AlertTimeGrainTypeBillingQuarter),
		string(AlertTimeGrainTypeMonthly),
		string(AlertTimeGrainTypeNone),
		string(AlertTimeGrainTypeQuarterly),
	}
}

type AlertType string

const (
	AlertTypeBudget         AlertType = "Budget"
	AlertTypeBudgetForecast AlertType = "BudgetForecast"
	AlertTypeCredit         AlertType = "Credit"
	AlertTypeGeneral        AlertType = "General"
	AlertTypeInvoice        AlertType = "Invoice"
	AlertTypeQuota          AlertType = "Quota"
	AlertTypeXCloud         AlertType = "xCloud"
)

func PossibleValuesForAlertType() []string {
	return []string{
		string(AlertTypeBudget),
		string(AlertTypeBudgetForecast),
		string(AlertTypeCredit),
		string(AlertTypeGeneral),
		string(AlertTypeInvoice),
		string(AlertTypeQuota),
		string(AlertTypeXCloud),
	}
}

type ExternalCloudProviderType string

const (
	ExternalCloudProviderTypeExternalBillingAccounts ExternalCloudProviderType = "externalBillingAccounts"
	ExternalCloudProviderTypeExternalSubscriptions   ExternalCloudProviderType = "externalSubscriptions"
)

func PossibleValuesForExternalCloudProviderType() []string {
	return []string{
		string(ExternalCloudProviderTypeExternalBillingAccounts),
		string(ExternalCloudProviderTypeExternalSubscriptions),
	}
}
