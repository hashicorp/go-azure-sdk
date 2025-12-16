package openapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *AlertCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertCategory(input string) (*AlertCategory, error) {
	vals := map[string]AlertCategory{
		"billing": AlertCategoryBilling,
		"cost":    AlertCategoryCost,
		"system":  AlertCategorySystem,
		"usage":   AlertCategoryUsage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertCategory(input)
	return &out, nil
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

func (s *AlertCriteria) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertCriteria(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertCriteria(input string) (*AlertCriteria, error) {
	vals := map[string]AlertCriteria{
		"costthresholdexceeded":          AlertCriteriaCostThresholdExceeded,
		"creditthresholdapproaching":     AlertCriteriaCreditThresholdApproaching,
		"creditthresholdreached":         AlertCriteriaCreditThresholdReached,
		"crosscloudcollectionerror":      AlertCriteriaCrossCloudCollectionError,
		"crosscloudnewdataavailable":     AlertCriteriaCrossCloudNewDataAvailable,
		"forecastcostthresholdexceeded":  AlertCriteriaForecastCostThresholdExceeded,
		"forecastusagethresholdexceeded": AlertCriteriaForecastUsageThresholdExceeded,
		"generalthresholderror":          AlertCriteriaGeneralThresholdError,
		"invoiceduedateapproaching":      AlertCriteriaInvoiceDueDateApproaching,
		"invoiceduedatereached":          AlertCriteriaInvoiceDueDateReached,
		"multicurrency":                  AlertCriteriaMultiCurrency,
		"quotathresholdapproaching":      AlertCriteriaQuotaThresholdApproaching,
		"quotathresholdreached":          AlertCriteriaQuotaThresholdReached,
		"usagethresholdexceeded":         AlertCriteriaUsageThresholdExceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertCriteria(input)
	return &out, nil
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

func (s *AlertOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertOperator(input string) (*AlertOperator, error) {
	vals := map[string]AlertOperator{
		"equalto":              AlertOperatorEqualTo,
		"greaterthan":          AlertOperatorGreaterThan,
		"greaterthanorequalto": AlertOperatorGreaterThanOrEqualTo,
		"lessthan":             AlertOperatorLessThan,
		"lessthanorequalto":    AlertOperatorLessThanOrEqualTo,
		"none":                 AlertOperatorNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertOperator(input)
	return &out, nil
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

func (s *AlertSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertSource(input string) (*AlertSource, error) {
	vals := map[string]AlertSource{
		"preset": AlertSourcePreset,
		"user":   AlertSourceUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertSource(input)
	return &out, nil
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

func (s *AlertStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertStatus(input string) (*AlertStatus, error) {
	vals := map[string]AlertStatus{
		"active":     AlertStatusActive,
		"dismissed":  AlertStatusDismissed,
		"none":       AlertStatusNone,
		"overridden": AlertStatusOverridden,
		"resolved":   AlertStatusResolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertStatus(input)
	return &out, nil
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

func (s *AlertTimeGrainType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertTimeGrainType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertTimeGrainType(input string) (*AlertTimeGrainType, error) {
	vals := map[string]AlertTimeGrainType{
		"annually":       AlertTimeGrainTypeAnnually,
		"billingannual":  AlertTimeGrainTypeBillingAnnual,
		"billingmonth":   AlertTimeGrainTypeBillingMonth,
		"billingquarter": AlertTimeGrainTypeBillingQuarter,
		"monthly":        AlertTimeGrainTypeMonthly,
		"none":           AlertTimeGrainTypeNone,
		"quarterly":      AlertTimeGrainTypeQuarterly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertTimeGrainType(input)
	return &out, nil
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

func (s *AlertType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertType(input string) (*AlertType, error) {
	vals := map[string]AlertType{
		"budget":         AlertTypeBudget,
		"budgetforecast": AlertTypeBudgetForecast,
		"credit":         AlertTypeCredit,
		"general":        AlertTypeGeneral,
		"invoice":        AlertTypeInvoice,
		"quota":          AlertTypeQuota,
		"xcloud":         AlertTypeXCloud,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertType(input)
	return &out, nil
}

type BenefitKind string

const (
	BenefitKindIncludedQuantity BenefitKind = "IncludedQuantity"
	BenefitKindReservation      BenefitKind = "Reservation"
	BenefitKindSavingsPlan      BenefitKind = "SavingsPlan"
)

func PossibleValuesForBenefitKind() []string {
	return []string{
		string(BenefitKindIncludedQuantity),
		string(BenefitKindReservation),
		string(BenefitKindSavingsPlan),
	}
}

func (s *BenefitKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBenefitKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBenefitKind(input string) (*BenefitKind, error) {
	vals := map[string]BenefitKind{
		"includedquantity": BenefitKindIncludedQuantity,
		"reservation":      BenefitKindReservation,
		"savingsplan":      BenefitKindSavingsPlan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BenefitKind(input)
	return &out, nil
}

type BenefitUtilizationSummaryReportSchema string

const (
	BenefitUtilizationSummaryReportSchemaAvgUtilizationPercentage BenefitUtilizationSummaryReportSchema = "AvgUtilizationPercentage"
	BenefitUtilizationSummaryReportSchemaBenefitId                BenefitUtilizationSummaryReportSchema = "BenefitId"
	BenefitUtilizationSummaryReportSchemaBenefitOrderId           BenefitUtilizationSummaryReportSchema = "BenefitOrderId"
	BenefitUtilizationSummaryReportSchemaBenefitType              BenefitUtilizationSummaryReportSchema = "BenefitType"
	BenefitUtilizationSummaryReportSchemaKind                     BenefitUtilizationSummaryReportSchema = "Kind"
	BenefitUtilizationSummaryReportSchemaMaxUtilizationPercentage BenefitUtilizationSummaryReportSchema = "MaxUtilizationPercentage"
	BenefitUtilizationSummaryReportSchemaMinUtilizationPercentage BenefitUtilizationSummaryReportSchema = "MinUtilizationPercentage"
	BenefitUtilizationSummaryReportSchemaUsageDate                BenefitUtilizationSummaryReportSchema = "UsageDate"
	BenefitUtilizationSummaryReportSchemaUtilizedPercentage       BenefitUtilizationSummaryReportSchema = "UtilizedPercentage"
)

func PossibleValuesForBenefitUtilizationSummaryReportSchema() []string {
	return []string{
		string(BenefitUtilizationSummaryReportSchemaAvgUtilizationPercentage),
		string(BenefitUtilizationSummaryReportSchemaBenefitId),
		string(BenefitUtilizationSummaryReportSchemaBenefitOrderId),
		string(BenefitUtilizationSummaryReportSchemaBenefitType),
		string(BenefitUtilizationSummaryReportSchemaKind),
		string(BenefitUtilizationSummaryReportSchemaMaxUtilizationPercentage),
		string(BenefitUtilizationSummaryReportSchemaMinUtilizationPercentage),
		string(BenefitUtilizationSummaryReportSchemaUsageDate),
		string(BenefitUtilizationSummaryReportSchemaUtilizedPercentage),
	}
}

func (s *BenefitUtilizationSummaryReportSchema) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBenefitUtilizationSummaryReportSchema(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBenefitUtilizationSummaryReportSchema(input string) (*BenefitUtilizationSummaryReportSchema, error) {
	vals := map[string]BenefitUtilizationSummaryReportSchema{
		"avgutilizationpercentage": BenefitUtilizationSummaryReportSchemaAvgUtilizationPercentage,
		"benefitid":                BenefitUtilizationSummaryReportSchemaBenefitId,
		"benefitorderid":           BenefitUtilizationSummaryReportSchemaBenefitOrderId,
		"benefittype":              BenefitUtilizationSummaryReportSchemaBenefitType,
		"kind":                     BenefitUtilizationSummaryReportSchemaKind,
		"maxutilizationpercentage": BenefitUtilizationSummaryReportSchemaMaxUtilizationPercentage,
		"minutilizationpercentage": BenefitUtilizationSummaryReportSchemaMinUtilizationPercentage,
		"usagedate":                BenefitUtilizationSummaryReportSchemaUsageDate,
		"utilizedpercentage":       BenefitUtilizationSummaryReportSchemaUtilizedPercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BenefitUtilizationSummaryReportSchema(input)
	return &out, nil
}

type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

func PossibleValuesForCheckNameAvailabilityReason() []string {
	return []string{
		string(CheckNameAvailabilityReasonAlreadyExists),
		string(CheckNameAvailabilityReasonInvalid),
	}
}

func (s *CheckNameAvailabilityReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCheckNameAvailabilityReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCheckNameAvailabilityReason(input string) (*CheckNameAvailabilityReason, error) {
	vals := map[string]CheckNameAvailabilityReason{
		"alreadyexists": CheckNameAvailabilityReasonAlreadyExists,
		"invalid":       CheckNameAvailabilityReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CheckNameAvailabilityReason(input)
	return &out, nil
}

type CostDetailsDataFormat string

const (
	CostDetailsDataFormatCsv CostDetailsDataFormat = "Csv"
)

func PossibleValuesForCostDetailsDataFormat() []string {
	return []string{
		string(CostDetailsDataFormatCsv),
	}
}

func (s *CostDetailsDataFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostDetailsDataFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostDetailsDataFormat(input string) (*CostDetailsDataFormat, error) {
	vals := map[string]CostDetailsDataFormat{
		"csv": CostDetailsDataFormatCsv,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostDetailsDataFormat(input)
	return &out, nil
}

type CostDetailsMetricType string

const (
	CostDetailsMetricTypeActualCost    CostDetailsMetricType = "ActualCost"
	CostDetailsMetricTypeAmortizedCost CostDetailsMetricType = "AmortizedCost"
)

func PossibleValuesForCostDetailsMetricType() []string {
	return []string{
		string(CostDetailsMetricTypeActualCost),
		string(CostDetailsMetricTypeAmortizedCost),
	}
}

func (s *CostDetailsMetricType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostDetailsMetricType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostDetailsMetricType(input string) (*CostDetailsMetricType, error) {
	vals := map[string]CostDetailsMetricType{
		"actualcost":    CostDetailsMetricTypeActualCost,
		"amortizedcost": CostDetailsMetricTypeAmortizedCost,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostDetailsMetricType(input)
	return &out, nil
}

type CostDetailsStatusType string

const (
	CostDetailsStatusTypeCompleted   CostDetailsStatusType = "Completed"
	CostDetailsStatusTypeFailed      CostDetailsStatusType = "Failed"
	CostDetailsStatusTypeNoDataFound CostDetailsStatusType = "NoDataFound"
)

func PossibleValuesForCostDetailsStatusType() []string {
	return []string{
		string(CostDetailsStatusTypeCompleted),
		string(CostDetailsStatusTypeFailed),
		string(CostDetailsStatusTypeNoDataFound),
	}
}

func (s *CostDetailsStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCostDetailsStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCostDetailsStatusType(input string) (*CostDetailsStatusType, error) {
	vals := map[string]CostDetailsStatusType{
		"completed":   CostDetailsStatusTypeCompleted,
		"failed":      CostDetailsStatusTypeFailed,
		"nodatafound": CostDetailsStatusTypeNoDataFound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CostDetailsStatusType(input)
	return &out, nil
}

type ExportType string

const (
	ExportTypeActualCost                 ExportType = "ActualCost"
	ExportTypeAmortizedCost              ExportType = "AmortizedCost"
	ExportTypeFocusCost                  ExportType = "FocusCost"
	ExportTypePriceSheet                 ExportType = "PriceSheet"
	ExportTypeReservationDetails         ExportType = "ReservationDetails"
	ExportTypeReservationRecommendations ExportType = "ReservationRecommendations"
	ExportTypeReservationTransactions    ExportType = "ReservationTransactions"
	ExportTypeUsage                      ExportType = "Usage"
)

func PossibleValuesForExportType() []string {
	return []string{
		string(ExportTypeActualCost),
		string(ExportTypeAmortizedCost),
		string(ExportTypeFocusCost),
		string(ExportTypePriceSheet),
		string(ExportTypeReservationDetails),
		string(ExportTypeReservationRecommendations),
		string(ExportTypeReservationTransactions),
		string(ExportTypeUsage),
	}
}

func (s *ExportType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExportType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExportType(input string) (*ExportType, error) {
	vals := map[string]ExportType{
		"actualcost":                 ExportTypeActualCost,
		"amortizedcost":              ExportTypeAmortizedCost,
		"focuscost":                  ExportTypeFocusCost,
		"pricesheet":                 ExportTypePriceSheet,
		"reservationdetails":         ExportTypeReservationDetails,
		"reservationrecommendations": ExportTypeReservationRecommendations,
		"reservationtransactions":    ExportTypeReservationTransactions,
		"usage":                      ExportTypeUsage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExportType(input)
	return &out, nil
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

func (s *ExternalCloudProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalCloudProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalCloudProviderType(input string) (*ExternalCloudProviderType, error) {
	vals := map[string]ExternalCloudProviderType{
		"externalbillingaccounts": ExternalCloudProviderTypeExternalBillingAccounts,
		"externalsubscriptions":   ExternalCloudProviderTypeExternalSubscriptions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalCloudProviderType(input)
	return &out, nil
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

func (s *ForecastOperatorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseForecastOperatorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseForecastOperatorType(input string) (*ForecastOperatorType, error) {
	vals := map[string]ForecastOperatorType{
		"in": ForecastOperatorTypeIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastOperatorType(input)
	return &out, nil
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

func (s *ForecastTimeframe) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseForecastTimeframe(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseForecastTimeframe(input string) (*ForecastTimeframe, error) {
	vals := map[string]ForecastTimeframe{
		"custom": ForecastTimeframeCustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastTimeframe(input)
	return &out, nil
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

func (s *ForecastType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseForecastType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseForecastType(input string) (*ForecastType, error) {
	vals := map[string]ForecastType{
		"actualcost":    ForecastTypeActualCost,
		"amortizedcost": ForecastTypeAmortizedCost,
		"usage":         ForecastTypeUsage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ForecastType(input)
	return &out, nil
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

func (s *FunctionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFunctionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFunctionName(input string) (*FunctionName, error) {
	vals := map[string]FunctionName{
		"cost":          FunctionNameCost,
		"costusd":       FunctionNameCostUSD,
		"pretaxcost":    FunctionNamePreTaxCost,
		"pretaxcostusd": FunctionNamePreTaxCostUSD,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FunctionName(input)
	return &out, nil
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

func (s *FunctionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFunctionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFunctionType(input string) (*FunctionType, error) {
	vals := map[string]FunctionType{
		"sum": FunctionTypeSum,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FunctionType(input)
	return &out, nil
}

type GenerateDetailedCostReportMetricType string

const (
	GenerateDetailedCostReportMetricTypeActualCost    GenerateDetailedCostReportMetricType = "ActualCost"
	GenerateDetailedCostReportMetricTypeAmortizedCost GenerateDetailedCostReportMetricType = "AmortizedCost"
)

func PossibleValuesForGenerateDetailedCostReportMetricType() []string {
	return []string{
		string(GenerateDetailedCostReportMetricTypeActualCost),
		string(GenerateDetailedCostReportMetricTypeAmortizedCost),
	}
}

func (s *GenerateDetailedCostReportMetricType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGenerateDetailedCostReportMetricType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGenerateDetailedCostReportMetricType(input string) (*GenerateDetailedCostReportMetricType, error) {
	vals := map[string]GenerateDetailedCostReportMetricType{
		"actualcost":    GenerateDetailedCostReportMetricTypeActualCost,
		"amortizedcost": GenerateDetailedCostReportMetricTypeAmortizedCost,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GenerateDetailedCostReportMetricType(input)
	return &out, nil
}

type Grain string

const (
	GrainDaily   Grain = "Daily"
	GrainHourly  Grain = "Hourly"
	GrainMonthly Grain = "Monthly"
)

func PossibleValuesForGrain() []string {
	return []string{
		string(GrainDaily),
		string(GrainHourly),
		string(GrainMonthly),
	}
}

func (s *Grain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGrain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGrain(input string) (*Grain, error) {
	vals := map[string]Grain{
		"daily":   GrainDaily,
		"hourly":  GrainHourly,
		"monthly": GrainMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Grain(input)
	return &out, nil
}

type GrainParameter string

const (
	GrainParameterDaily   GrainParameter = "Daily"
	GrainParameterHourly  GrainParameter = "Hourly"
	GrainParameterMonthly GrainParameter = "Monthly"
)

func PossibleValuesForGrainParameter() []string {
	return []string{
		string(GrainParameterDaily),
		string(GrainParameterHourly),
		string(GrainParameterMonthly),
	}
}

func (s *GrainParameter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGrainParameter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGrainParameter(input string) (*GrainParameter, error) {
	vals := map[string]GrainParameter{
		"daily":   GrainParameterDaily,
		"hourly":  GrainParameterHourly,
		"monthly": GrainParameterMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GrainParameter(input)
	return &out, nil
}

type GranularityType string

const (
	GranularityTypeDaily   GranularityType = "Daily"
	GranularityTypeMonthly GranularityType = "Monthly"
)

func PossibleValuesForGranularityType() []string {
	return []string{
		string(GranularityTypeDaily),
		string(GranularityTypeMonthly),
	}
}

func (s *GranularityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGranularityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGranularityType(input string) (*GranularityType, error) {
	vals := map[string]GranularityType{
		"daily":   GranularityTypeDaily,
		"monthly": GranularityTypeMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GranularityType(input)
	return &out, nil
}

type LookBackPeriod string

const (
	LookBackPeriodLastSevenDays     LookBackPeriod = "Last7Days"
	LookBackPeriodLastSixZeroDays   LookBackPeriod = "Last60Days"
	LookBackPeriodLastThreeZeroDays LookBackPeriod = "Last30Days"
)

func PossibleValuesForLookBackPeriod() []string {
	return []string{
		string(LookBackPeriodLastSevenDays),
		string(LookBackPeriodLastSixZeroDays),
		string(LookBackPeriodLastThreeZeroDays),
	}
}

func (s *LookBackPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLookBackPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLookBackPeriod(input string) (*LookBackPeriod, error) {
	vals := map[string]LookBackPeriod{
		"last7days":  LookBackPeriodLastSevenDays,
		"last60days": LookBackPeriodLastSixZeroDays,
		"last30days": LookBackPeriodLastThreeZeroDays,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LookBackPeriod(input)
	return &out, nil
}

type OperationStatusType string

const (
	OperationStatusTypeCompleted OperationStatusType = "Completed"
	OperationStatusTypeFailed    OperationStatusType = "Failed"
	OperationStatusTypeRunning   OperationStatusType = "Running"
)

func PossibleValuesForOperationStatusType() []string {
	return []string{
		string(OperationStatusTypeCompleted),
		string(OperationStatusTypeFailed),
		string(OperationStatusTypeRunning),
	}
}

func (s *OperationStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationStatusType(input string) (*OperationStatusType, error) {
	vals := map[string]OperationStatusType{
		"completed": OperationStatusTypeCompleted,
		"failed":    OperationStatusTypeFailed,
		"running":   OperationStatusTypeRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatusType(input)
	return &out, nil
}

type QueryColumnType string

const (
	QueryColumnTypeDimension QueryColumnType = "Dimension"
	QueryColumnTypeTagKey    QueryColumnType = "TagKey"
)

func PossibleValuesForQueryColumnType() []string {
	return []string{
		string(QueryColumnTypeDimension),
		string(QueryColumnTypeTagKey),
	}
}

func (s *QueryColumnType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQueryColumnType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQueryColumnType(input string) (*QueryColumnType, error) {
	vals := map[string]QueryColumnType{
		"dimension": QueryColumnTypeDimension,
		"tagkey":    QueryColumnTypeTagKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QueryColumnType(input)
	return &out, nil
}

type QueryOperatorType string

const (
	QueryOperatorTypeIn QueryOperatorType = "In"
)

func PossibleValuesForQueryOperatorType() []string {
	return []string{
		string(QueryOperatorTypeIn),
	}
}

func (s *QueryOperatorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseQueryOperatorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseQueryOperatorType(input string) (*QueryOperatorType, error) {
	vals := map[string]QueryOperatorType{
		"in": QueryOperatorTypeIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := QueryOperatorType(input)
	return &out, nil
}

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
	ReasonValid         Reason = "Valid"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
		string(ReasonInvalid),
		string(ReasonValid),
	}
}

func (s *Reason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReason(input string) (*Reason, error) {
	vals := map[string]Reason{
		"alreadyexists": ReasonAlreadyExists,
		"invalid":       ReasonInvalid,
		"valid":         ReasonValid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Reason(input)
	return &out, nil
}

type ReservationReportSchema string

const (
	ReservationReportSchemaInstanceFlexibilityGroup ReservationReportSchema = "InstanceFlexibilityGroup"
	ReservationReportSchemaInstanceFlexibilityRatio ReservationReportSchema = "InstanceFlexibilityRatio"
	ReservationReportSchemaInstanceId               ReservationReportSchema = "InstanceId"
	ReservationReportSchemaKind                     ReservationReportSchema = "Kind"
	ReservationReportSchemaReservationId            ReservationReportSchema = "ReservationId"
	ReservationReportSchemaReservationOrderId       ReservationReportSchema = "ReservationOrderId"
	ReservationReportSchemaReservedHours            ReservationReportSchema = "ReservedHours"
	ReservationReportSchemaSkuName                  ReservationReportSchema = "SkuName"
	ReservationReportSchemaTotalReservedQuantity    ReservationReportSchema = "TotalReservedQuantity"
	ReservationReportSchemaUsageDate                ReservationReportSchema = "UsageDate"
	ReservationReportSchemaUsedHours                ReservationReportSchema = "UsedHours"
)

func PossibleValuesForReservationReportSchema() []string {
	return []string{
		string(ReservationReportSchemaInstanceFlexibilityGroup),
		string(ReservationReportSchemaInstanceFlexibilityRatio),
		string(ReservationReportSchemaInstanceId),
		string(ReservationReportSchemaKind),
		string(ReservationReportSchemaReservationId),
		string(ReservationReportSchemaReservationOrderId),
		string(ReservationReportSchemaReservedHours),
		string(ReservationReportSchemaSkuName),
		string(ReservationReportSchemaTotalReservedQuantity),
		string(ReservationReportSchemaUsageDate),
		string(ReservationReportSchemaUsedHours),
	}
}

func (s *ReservationReportSchema) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReservationReportSchema(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReservationReportSchema(input string) (*ReservationReportSchema, error) {
	vals := map[string]ReservationReportSchema{
		"instanceflexibilitygroup": ReservationReportSchemaInstanceFlexibilityGroup,
		"instanceflexibilityratio": ReservationReportSchemaInstanceFlexibilityRatio,
		"instanceid":               ReservationReportSchemaInstanceId,
		"kind":                     ReservationReportSchemaKind,
		"reservationid":            ReservationReportSchemaReservationId,
		"reservationorderid":       ReservationReportSchemaReservationOrderId,
		"reservedhours":            ReservationReportSchemaReservedHours,
		"skuname":                  ReservationReportSchemaSkuName,
		"totalreservedquantity":    ReservationReportSchemaTotalReservedQuantity,
		"usagedate":                ReservationReportSchemaUsageDate,
		"usedhours":                ReservationReportSchemaUsedHours,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReservationReportSchema(input)
	return &out, nil
}

type Scope string

const (
	ScopeShared Scope = "Shared"
	ScopeSingle Scope = "Single"
)

func PossibleValuesForScope() []string {
	return []string{
		string(ScopeShared),
		string(ScopeSingle),
	}
}

func (s *Scope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScope(input string) (*Scope, error) {
	vals := map[string]Scope{
		"shared": ScopeShared,
		"single": ScopeSingle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Scope(input)
	return &out, nil
}

type Term string

const (
	TermPOneY   Term = "P1Y"
	TermPThreeY Term = "P3Y"
)

func PossibleValuesForTerm() []string {
	return []string{
		string(TermPOneY),
		string(TermPThreeY),
	}
}

func (s *Term) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTerm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTerm(input string) (*Term, error) {
	vals := map[string]Term{
		"p1y": TermPOneY,
		"p3y": TermPThreeY,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Term(input)
	return &out, nil
}

type TimeframeType string

const (
	TimeframeTypeBillingMonthToDate  TimeframeType = "BillingMonthToDate"
	TimeframeTypeCustom              TimeframeType = "Custom"
	TimeframeTypeMonthToDate         TimeframeType = "MonthToDate"
	TimeframeTypeTheCurrentMonth     TimeframeType = "TheCurrentMonth"
	TimeframeTypeTheLastBillingMonth TimeframeType = "TheLastBillingMonth"
	TimeframeTypeTheLastMonth        TimeframeType = "TheLastMonth"
	TimeframeTypeWeekToDate          TimeframeType = "WeekToDate"
)

func PossibleValuesForTimeframeType() []string {
	return []string{
		string(TimeframeTypeBillingMonthToDate),
		string(TimeframeTypeCustom),
		string(TimeframeTypeMonthToDate),
		string(TimeframeTypeTheCurrentMonth),
		string(TimeframeTypeTheLastBillingMonth),
		string(TimeframeTypeTheLastMonth),
		string(TimeframeTypeWeekToDate),
	}
}

func (s *TimeframeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeframeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeframeType(input string) (*TimeframeType, error) {
	vals := map[string]TimeframeType{
		"billingmonthtodate":  TimeframeTypeBillingMonthToDate,
		"custom":              TimeframeTypeCustom,
		"monthtodate":         TimeframeTypeMonthToDate,
		"thecurrentmonth":     TimeframeTypeTheCurrentMonth,
		"thelastbillingmonth": TimeframeTypeTheLastBillingMonth,
		"thelastmonth":        TimeframeTypeTheLastMonth,
		"weektodate":          TimeframeTypeWeekToDate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeframeType(input)
	return &out, nil
}
