package billingprofile

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileStatus string

const (
	BillingProfileStatusActive      BillingProfileStatus = "Active"
	BillingProfileStatusDeleted     BillingProfileStatus = "Deleted"
	BillingProfileStatusDisabled    BillingProfileStatus = "Disabled"
	BillingProfileStatusOther       BillingProfileStatus = "Other"
	BillingProfileStatusUnderReview BillingProfileStatus = "UnderReview"
	BillingProfileStatusWarned      BillingProfileStatus = "Warned"
)

func PossibleValuesForBillingProfileStatus() []string {
	return []string{
		string(BillingProfileStatusActive),
		string(BillingProfileStatusDeleted),
		string(BillingProfileStatusDisabled),
		string(BillingProfileStatusOther),
		string(BillingProfileStatusUnderReview),
		string(BillingProfileStatusWarned),
	}
}

func (s *BillingProfileStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileStatus(input string) (*BillingProfileStatus, error) {
	vals := map[string]BillingProfileStatus{
		"active":      BillingProfileStatusActive,
		"deleted":     BillingProfileStatusDeleted,
		"disabled":    BillingProfileStatusDisabled,
		"other":       BillingProfileStatusOther,
		"underreview": BillingProfileStatusUnderReview,
		"warned":      BillingProfileStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileStatus(input)
	return &out, nil
}

type BillingProfileStatusReasonCode string

const (
	BillingProfileStatusReasonCodeOther                BillingProfileStatusReasonCode = "Other"
	BillingProfileStatusReasonCodePastDue              BillingProfileStatusReasonCode = "PastDue"
	BillingProfileStatusReasonCodeSpendingLimitExpired BillingProfileStatusReasonCode = "SpendingLimitExpired"
	BillingProfileStatusReasonCodeSpendingLimitReached BillingProfileStatusReasonCode = "SpendingLimitReached"
	BillingProfileStatusReasonCodeUnusualActivity      BillingProfileStatusReasonCode = "UnusualActivity"
)

func PossibleValuesForBillingProfileStatusReasonCode() []string {
	return []string{
		string(BillingProfileStatusReasonCodeOther),
		string(BillingProfileStatusReasonCodePastDue),
		string(BillingProfileStatusReasonCodeSpendingLimitExpired),
		string(BillingProfileStatusReasonCodeSpendingLimitReached),
		string(BillingProfileStatusReasonCodeUnusualActivity),
	}
}

func (s *BillingProfileStatusReasonCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingProfileStatusReasonCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingProfileStatusReasonCode(input string) (*BillingProfileStatusReasonCode, error) {
	vals := map[string]BillingProfileStatusReasonCode{
		"other":                BillingProfileStatusReasonCodeOther,
		"pastdue":              BillingProfileStatusReasonCodePastDue,
		"spendinglimitexpired": BillingProfileStatusReasonCodeSpendingLimitExpired,
		"spendinglimitreached": BillingProfileStatusReasonCodeSpendingLimitReached,
		"unusualactivity":      BillingProfileStatusReasonCodeUnusualActivity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingProfileStatusReasonCode(input)
	return &out, nil
}

type BillingRelationshipType string

const (
	BillingRelationshipTypeCSPCustomer      BillingRelationshipType = "CSPCustomer"
	BillingRelationshipTypeCSPPartner       BillingRelationshipType = "CSPPartner"
	BillingRelationshipTypeDirect           BillingRelationshipType = "Direct"
	BillingRelationshipTypeIndirectCustomer BillingRelationshipType = "IndirectCustomer"
	BillingRelationshipTypeIndirectPartner  BillingRelationshipType = "IndirectPartner"
	BillingRelationshipTypeOther            BillingRelationshipType = "Other"
)

func PossibleValuesForBillingRelationshipType() []string {
	return []string{
		string(BillingRelationshipTypeCSPCustomer),
		string(BillingRelationshipTypeCSPPartner),
		string(BillingRelationshipTypeDirect),
		string(BillingRelationshipTypeIndirectCustomer),
		string(BillingRelationshipTypeIndirectPartner),
		string(BillingRelationshipTypeOther),
	}
}

func (s *BillingRelationshipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingRelationshipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingRelationshipType(input string) (*BillingRelationshipType, error) {
	vals := map[string]BillingRelationshipType{
		"cspcustomer":      BillingRelationshipTypeCSPCustomer,
		"csppartner":       BillingRelationshipTypeCSPPartner,
		"direct":           BillingRelationshipTypeDirect,
		"indirectcustomer": BillingRelationshipTypeIndirectCustomer,
		"indirectpartner":  BillingRelationshipTypeIndirectPartner,
		"other":            BillingRelationshipTypeOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingRelationshipType(input)
	return &out, nil
}

type DeleteBillingProfileEligibilityCode string

const (
	DeleteBillingProfileEligibilityCodeActiveBillingSubscriptions DeleteBillingProfileEligibilityCode = "ActiveBillingSubscriptions"
	DeleteBillingProfileEligibilityCodeActiveCreditCard           DeleteBillingProfileEligibilityCode = "ActiveCreditCard"
	DeleteBillingProfileEligibilityCodeActiveCredits              DeleteBillingProfileEligibilityCode = "ActiveCredits"
	DeleteBillingProfileEligibilityCodeLastBillingProfile         DeleteBillingProfileEligibilityCode = "LastBillingProfile"
	DeleteBillingProfileEligibilityCodeNone                       DeleteBillingProfileEligibilityCode = "None"
	DeleteBillingProfileEligibilityCodeNotSupported               DeleteBillingProfileEligibilityCode = "NotSupported"
	DeleteBillingProfileEligibilityCodeOutstandingCharges         DeleteBillingProfileEligibilityCode = "OutstandingCharges"
	DeleteBillingProfileEligibilityCodePendingCharges             DeleteBillingProfileEligibilityCode = "PendingCharges"
	DeleteBillingProfileEligibilityCodeReservedInstances          DeleteBillingProfileEligibilityCode = "ReservedInstances"
)

func PossibleValuesForDeleteBillingProfileEligibilityCode() []string {
	return []string{
		string(DeleteBillingProfileEligibilityCodeActiveBillingSubscriptions),
		string(DeleteBillingProfileEligibilityCodeActiveCreditCard),
		string(DeleteBillingProfileEligibilityCodeActiveCredits),
		string(DeleteBillingProfileEligibilityCodeLastBillingProfile),
		string(DeleteBillingProfileEligibilityCodeNone),
		string(DeleteBillingProfileEligibilityCodeNotSupported),
		string(DeleteBillingProfileEligibilityCodeOutstandingCharges),
		string(DeleteBillingProfileEligibilityCodePendingCharges),
		string(DeleteBillingProfileEligibilityCodeReservedInstances),
	}
}

func (s *DeleteBillingProfileEligibilityCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeleteBillingProfileEligibilityCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeleteBillingProfileEligibilityCode(input string) (*DeleteBillingProfileEligibilityCode, error) {
	vals := map[string]DeleteBillingProfileEligibilityCode{
		"activebillingsubscriptions": DeleteBillingProfileEligibilityCodeActiveBillingSubscriptions,
		"activecreditcard":           DeleteBillingProfileEligibilityCodeActiveCreditCard,
		"activecredits":              DeleteBillingProfileEligibilityCodeActiveCredits,
		"lastbillingprofile":         DeleteBillingProfileEligibilityCodeLastBillingProfile,
		"none":                       DeleteBillingProfileEligibilityCodeNone,
		"notsupported":               DeleteBillingProfileEligibilityCodeNotSupported,
		"outstandingcharges":         DeleteBillingProfileEligibilityCodeOutstandingCharges,
		"pendingcharges":             DeleteBillingProfileEligibilityCodePendingCharges,
		"reservedinstances":          DeleteBillingProfileEligibilityCodeReservedInstances,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeleteBillingProfileEligibilityCode(input)
	return &out, nil
}

type DeleteBillingProfileEligibilityStatus string

const (
	DeleteBillingProfileEligibilityStatusAllowed    DeleteBillingProfileEligibilityStatus = "Allowed"
	DeleteBillingProfileEligibilityStatusNotAllowed DeleteBillingProfileEligibilityStatus = "NotAllowed"
)

func PossibleValuesForDeleteBillingProfileEligibilityStatus() []string {
	return []string{
		string(DeleteBillingProfileEligibilityStatusAllowed),
		string(DeleteBillingProfileEligibilityStatusNotAllowed),
	}
}

func (s *DeleteBillingProfileEligibilityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeleteBillingProfileEligibilityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeleteBillingProfileEligibilityStatus(input string) (*DeleteBillingProfileEligibilityStatus, error) {
	vals := map[string]DeleteBillingProfileEligibilityStatus{
		"allowed":    DeleteBillingProfileEligibilityStatusAllowed,
		"notallowed": DeleteBillingProfileEligibilityStatusNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeleteBillingProfileEligibilityStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNew          ProvisioningState = "New"
	ProvisioningStatePending      ProvisioningState = "Pending"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNew),
		string(ProvisioningStatePending),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":     ProvisioningStateCanceled,
		"failed":       ProvisioningStateFailed,
		"new":          ProvisioningStateNew,
		"pending":      ProvisioningStatePending,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type SpendingLimit string

const (
	SpendingLimitOff SpendingLimit = "Off"
	SpendingLimitOn  SpendingLimit = "On"
)

func PossibleValuesForSpendingLimit() []string {
	return []string{
		string(SpendingLimitOff),
		string(SpendingLimitOn),
	}
}

func (s *SpendingLimit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpendingLimit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpendingLimit(input string) (*SpendingLimit, error) {
	vals := map[string]SpendingLimit{
		"off": SpendingLimitOff,
		"on":  SpendingLimitOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpendingLimit(input)
	return &out, nil
}

type SpendingLimitStatus string

const (
	SpendingLimitStatusActive       SpendingLimitStatus = "Active"
	SpendingLimitStatusExpired      SpendingLimitStatus = "Expired"
	SpendingLimitStatusLimitReached SpendingLimitStatus = "LimitReached"
	SpendingLimitStatusLimitRemoved SpendingLimitStatus = "LimitRemoved"
	SpendingLimitStatusNone         SpendingLimitStatus = "None"
	SpendingLimitStatusOther        SpendingLimitStatus = "Other"
)

func PossibleValuesForSpendingLimitStatus() []string {
	return []string{
		string(SpendingLimitStatusActive),
		string(SpendingLimitStatusExpired),
		string(SpendingLimitStatusLimitReached),
		string(SpendingLimitStatusLimitRemoved),
		string(SpendingLimitStatusNone),
		string(SpendingLimitStatusOther),
	}
}

func (s *SpendingLimitStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpendingLimitStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpendingLimitStatus(input string) (*SpendingLimitStatus, error) {
	vals := map[string]SpendingLimitStatus{
		"active":       SpendingLimitStatusActive,
		"expired":      SpendingLimitStatusExpired,
		"limitreached": SpendingLimitStatusLimitReached,
		"limitremoved": SpendingLimitStatusLimitRemoved,
		"none":         SpendingLimitStatusNone,
		"other":        SpendingLimitStatusOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpendingLimitStatus(input)
	return &out, nil
}

type SpendingLimitType string

const (
	SpendingLimitTypeAcademicSponsorship     SpendingLimitType = "AcademicSponsorship"
	SpendingLimitTypeAzureConsumptionCredit  SpendingLimitType = "AzureConsumptionCredit"
	SpendingLimitTypeAzureForStudents        SpendingLimitType = "AzureForStudents"
	SpendingLimitTypeAzureForStudentsStarter SpendingLimitType = "AzureForStudentsStarter"
	SpendingLimitTypeAzurePassSponsorship    SpendingLimitType = "AzurePassSponsorship"
	SpendingLimitTypeFreeAccount             SpendingLimitType = "FreeAccount"
	SpendingLimitTypeMSDN                    SpendingLimitType = "MSDN"
	SpendingLimitTypeMpnSponsorship          SpendingLimitType = "MpnSponsorship"
	SpendingLimitTypeNonProfitSponsorship    SpendingLimitType = "NonProfitSponsorship"
	SpendingLimitTypeNone                    SpendingLimitType = "None"
	SpendingLimitTypeOther                   SpendingLimitType = "Other"
	SpendingLimitTypeSandbox                 SpendingLimitType = "Sandbox"
	SpendingLimitTypeSponsorship             SpendingLimitType = "Sponsorship"
	SpendingLimitTypeStartupSponsorship      SpendingLimitType = "StartupSponsorship"
	SpendingLimitTypeVisualStudio            SpendingLimitType = "VisualStudio"
)

func PossibleValuesForSpendingLimitType() []string {
	return []string{
		string(SpendingLimitTypeAcademicSponsorship),
		string(SpendingLimitTypeAzureConsumptionCredit),
		string(SpendingLimitTypeAzureForStudents),
		string(SpendingLimitTypeAzureForStudentsStarter),
		string(SpendingLimitTypeAzurePassSponsorship),
		string(SpendingLimitTypeFreeAccount),
		string(SpendingLimitTypeMSDN),
		string(SpendingLimitTypeMpnSponsorship),
		string(SpendingLimitTypeNonProfitSponsorship),
		string(SpendingLimitTypeNone),
		string(SpendingLimitTypeOther),
		string(SpendingLimitTypeSandbox),
		string(SpendingLimitTypeSponsorship),
		string(SpendingLimitTypeStartupSponsorship),
		string(SpendingLimitTypeVisualStudio),
	}
}

func (s *SpendingLimitType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSpendingLimitType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSpendingLimitType(input string) (*SpendingLimitType, error) {
	vals := map[string]SpendingLimitType{
		"academicsponsorship":     SpendingLimitTypeAcademicSponsorship,
		"azureconsumptioncredit":  SpendingLimitTypeAzureConsumptionCredit,
		"azureforstudents":        SpendingLimitTypeAzureForStudents,
		"azureforstudentsstarter": SpendingLimitTypeAzureForStudentsStarter,
		"azurepasssponsorship":    SpendingLimitTypeAzurePassSponsorship,
		"freeaccount":             SpendingLimitTypeFreeAccount,
		"msdn":                    SpendingLimitTypeMSDN,
		"mpnsponsorship":          SpendingLimitTypeMpnSponsorship,
		"nonprofitsponsorship":    SpendingLimitTypeNonProfitSponsorship,
		"none":                    SpendingLimitTypeNone,
		"other":                   SpendingLimitTypeOther,
		"sandbox":                 SpendingLimitTypeSandbox,
		"sponsorship":             SpendingLimitTypeSponsorship,
		"startupsponsorship":      SpendingLimitTypeStartupSponsorship,
		"visualstudio":            SpendingLimitTypeVisualStudio,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SpendingLimitType(input)
	return &out, nil
}
