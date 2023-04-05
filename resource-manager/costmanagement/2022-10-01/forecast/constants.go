package forecast

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type ForecastOperatorType string

const (
	ForecastOperatorTypeIn ForecastOperatorType = "In"
)

func PossibleValuesForForecastOperatorType() []string {
	return []string{
		string(ForecastOperatorTypeIn),
	}
}

type ForecastTimeframe string

const (
	ForecastTimeframeCustom ForecastTimeframe = "Custom"
)

func PossibleValuesForForecastTimeframe() []string {
	return []string{
		string(ForecastTimeframeCustom),
	}
}

type ForecastType string

const (
	ForecastTypeActualCost    ForecastType = "ActualCost"
	ForecastTypeAmortizedCost ForecastType = "AmortizedCost"
	ForecastTypeUsage         ForecastType = "Usage"
)

func PossibleValuesForForecastType() []string {
	return []string{
		string(ForecastTypeActualCost),
		string(ForecastTypeAmortizedCost),
		string(ForecastTypeUsage),
	}
}

type FunctionName string

const (
	FunctionNameCost          FunctionName = "Cost"
	FunctionNameCostUSD       FunctionName = "CostUSD"
	FunctionNamePreTaxCost    FunctionName = "PreTaxCost"
	FunctionNamePreTaxCostUSD FunctionName = "PreTaxCostUSD"
)

func PossibleValuesForFunctionName() []string {
	return []string{
		string(FunctionNameCost),
		string(FunctionNameCostUSD),
		string(FunctionNamePreTaxCost),
		string(FunctionNamePreTaxCostUSD),
	}
}

type FunctionType string

const (
	FunctionTypeSum FunctionType = "Sum"
)

func PossibleValuesForFunctionType() []string {
	return []string{
		string(FunctionTypeSum),
	}
}

type GranularityType string

const (
	GranularityTypeDaily GranularityType = "Daily"
)

func PossibleValuesForGranularityType() []string {
	return []string{
		string(GranularityTypeDaily),
	}
}
