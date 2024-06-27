package billingsubscription

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoRenew string

const (
	AutoRenewOff AutoRenew = "Off"
	AutoRenewOn  AutoRenew = "On"
)

func PossibleValuesForAutoRenew() []string {
	return []string{
		string(AutoRenewOff),
		string(AutoRenewOn),
	}
}

func (s *AutoRenew) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoRenew(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoRenew(input string) (*AutoRenew, error) {
	vals := map[string]AutoRenew{
		"off": AutoRenewOff,
		"on":  AutoRenewOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoRenew(input)
	return &out, nil
}

type BillingSubscriptionOperationStatus string

const (
	BillingSubscriptionOperationStatusLockedForUpdate BillingSubscriptionOperationStatus = "LockedForUpdate"
	BillingSubscriptionOperationStatusNone            BillingSubscriptionOperationStatus = "None"
	BillingSubscriptionOperationStatusOther           BillingSubscriptionOperationStatus = "Other"
)

func PossibleValuesForBillingSubscriptionOperationStatus() []string {
	return []string{
		string(BillingSubscriptionOperationStatusLockedForUpdate),
		string(BillingSubscriptionOperationStatusNone),
		string(BillingSubscriptionOperationStatusOther),
	}
}

func (s *BillingSubscriptionOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingSubscriptionOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingSubscriptionOperationStatus(input string) (*BillingSubscriptionOperationStatus, error) {
	vals := map[string]BillingSubscriptionOperationStatus{
		"lockedforupdate": BillingSubscriptionOperationStatusLockedForUpdate,
		"none":            BillingSubscriptionOperationStatusNone,
		"other":           BillingSubscriptionOperationStatusOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingSubscriptionOperationStatus(input)
	return &out, nil
}

type BillingSubscriptionStatus string

const (
	BillingSubscriptionStatusActive    BillingSubscriptionStatus = "Active"
	BillingSubscriptionStatusAutoRenew BillingSubscriptionStatus = "AutoRenew"
	BillingSubscriptionStatusCancelled BillingSubscriptionStatus = "Cancelled"
	BillingSubscriptionStatusDeleted   BillingSubscriptionStatus = "Deleted"
	BillingSubscriptionStatusDisabled  BillingSubscriptionStatus = "Disabled"
	BillingSubscriptionStatusExpired   BillingSubscriptionStatus = "Expired"
	BillingSubscriptionStatusExpiring  BillingSubscriptionStatus = "Expiring"
	BillingSubscriptionStatusFailed    BillingSubscriptionStatus = "Failed"
	BillingSubscriptionStatusOther     BillingSubscriptionStatus = "Other"
	BillingSubscriptionStatusSuspended BillingSubscriptionStatus = "Suspended"
	BillingSubscriptionStatusUnknown   BillingSubscriptionStatus = "Unknown"
	BillingSubscriptionStatusWarned    BillingSubscriptionStatus = "Warned"
)

func PossibleValuesForBillingSubscriptionStatus() []string {
	return []string{
		string(BillingSubscriptionStatusActive),
		string(BillingSubscriptionStatusAutoRenew),
		string(BillingSubscriptionStatusCancelled),
		string(BillingSubscriptionStatusDeleted),
		string(BillingSubscriptionStatusDisabled),
		string(BillingSubscriptionStatusExpired),
		string(BillingSubscriptionStatusExpiring),
		string(BillingSubscriptionStatusFailed),
		string(BillingSubscriptionStatusOther),
		string(BillingSubscriptionStatusSuspended),
		string(BillingSubscriptionStatusUnknown),
		string(BillingSubscriptionStatusWarned),
	}
}

func (s *BillingSubscriptionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBillingSubscriptionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBillingSubscriptionStatus(input string) (*BillingSubscriptionStatus, error) {
	vals := map[string]BillingSubscriptionStatus{
		"active":    BillingSubscriptionStatusActive,
		"autorenew": BillingSubscriptionStatusAutoRenew,
		"cancelled": BillingSubscriptionStatusCancelled,
		"deleted":   BillingSubscriptionStatusDeleted,
		"disabled":  BillingSubscriptionStatusDisabled,
		"expired":   BillingSubscriptionStatusExpired,
		"expiring":  BillingSubscriptionStatusExpiring,
		"failed":    BillingSubscriptionStatusFailed,
		"other":     BillingSubscriptionStatusOther,
		"suspended": BillingSubscriptionStatusSuspended,
		"unknown":   BillingSubscriptionStatusUnknown,
		"warned":    BillingSubscriptionStatusWarned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BillingSubscriptionStatus(input)
	return &out, nil
}

type Cancellation string

const (
	CancellationAllowed    Cancellation = "Allowed"
	CancellationNotAllowed Cancellation = "NotAllowed"
)

func PossibleValuesForCancellation() []string {
	return []string{
		string(CancellationAllowed),
		string(CancellationNotAllowed),
	}
}

func (s *Cancellation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCancellation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCancellation(input string) (*Cancellation, error) {
	vals := map[string]Cancellation{
		"allowed":    CancellationAllowed,
		"notallowed": CancellationNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Cancellation(input)
	return &out, nil
}

type CancellationReason string

const (
	CancellationReasonCompromise CancellationReason = "Compromise"
	CancellationReasonDispute    CancellationReason = "Dispute"
	CancellationReasonOther      CancellationReason = "Other"
)

func PossibleValuesForCancellationReason() []string {
	return []string{
		string(CancellationReasonCompromise),
		string(CancellationReasonDispute),
		string(CancellationReasonOther),
	}
}

func (s *CancellationReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCancellationReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCancellationReason(input string) (*CancellationReason, error) {
	vals := map[string]CancellationReason{
		"compromise": CancellationReasonCompromise,
		"dispute":    CancellationReasonDispute,
		"other":      CancellationReasonOther,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CancellationReason(input)
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

type SubscriptionEnrollmentAccountStatus string

const (
	SubscriptionEnrollmentAccountStatusActive         SubscriptionEnrollmentAccountStatus = "Active"
	SubscriptionEnrollmentAccountStatusCancelled      SubscriptionEnrollmentAccountStatus = "Cancelled"
	SubscriptionEnrollmentAccountStatusDeleted        SubscriptionEnrollmentAccountStatus = "Deleted"
	SubscriptionEnrollmentAccountStatusExpired        SubscriptionEnrollmentAccountStatus = "Expired"
	SubscriptionEnrollmentAccountStatusInactive       SubscriptionEnrollmentAccountStatus = "Inactive"
	SubscriptionEnrollmentAccountStatusTransferredOut SubscriptionEnrollmentAccountStatus = "TransferredOut"
	SubscriptionEnrollmentAccountStatusTransferring   SubscriptionEnrollmentAccountStatus = "Transferring"
)

func PossibleValuesForSubscriptionEnrollmentAccountStatus() []string {
	return []string{
		string(SubscriptionEnrollmentAccountStatusActive),
		string(SubscriptionEnrollmentAccountStatusCancelled),
		string(SubscriptionEnrollmentAccountStatusDeleted),
		string(SubscriptionEnrollmentAccountStatusExpired),
		string(SubscriptionEnrollmentAccountStatusInactive),
		string(SubscriptionEnrollmentAccountStatusTransferredOut),
		string(SubscriptionEnrollmentAccountStatusTransferring),
	}
}

func (s *SubscriptionEnrollmentAccountStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionEnrollmentAccountStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionEnrollmentAccountStatus(input string) (*SubscriptionEnrollmentAccountStatus, error) {
	vals := map[string]SubscriptionEnrollmentAccountStatus{
		"active":         SubscriptionEnrollmentAccountStatusActive,
		"cancelled":      SubscriptionEnrollmentAccountStatusCancelled,
		"deleted":        SubscriptionEnrollmentAccountStatusDeleted,
		"expired":        SubscriptionEnrollmentAccountStatusExpired,
		"inactive":       SubscriptionEnrollmentAccountStatusInactive,
		"transferredout": SubscriptionEnrollmentAccountStatusTransferredOut,
		"transferring":   SubscriptionEnrollmentAccountStatusTransferring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionEnrollmentAccountStatus(input)
	return &out, nil
}

type SubscriptionStatusReason string

const (
	SubscriptionStatusReasonCancelled            SubscriptionStatusReason = "Cancelled"
	SubscriptionStatusReasonExpired              SubscriptionStatusReason = "Expired"
	SubscriptionStatusReasonNone                 SubscriptionStatusReason = "None"
	SubscriptionStatusReasonOther                SubscriptionStatusReason = "Other"
	SubscriptionStatusReasonPastDue              SubscriptionStatusReason = "PastDue"
	SubscriptionStatusReasonPolicyViolation      SubscriptionStatusReason = "PolicyViolation"
	SubscriptionStatusReasonSpendingLimitReached SubscriptionStatusReason = "SpendingLimitReached"
	SubscriptionStatusReasonSuspiciousActivity   SubscriptionStatusReason = "SuspiciousActivity"
	SubscriptionStatusReasonTransferred          SubscriptionStatusReason = "Transferred"
)

func PossibleValuesForSubscriptionStatusReason() []string {
	return []string{
		string(SubscriptionStatusReasonCancelled),
		string(SubscriptionStatusReasonExpired),
		string(SubscriptionStatusReasonNone),
		string(SubscriptionStatusReasonOther),
		string(SubscriptionStatusReasonPastDue),
		string(SubscriptionStatusReasonPolicyViolation),
		string(SubscriptionStatusReasonSpendingLimitReached),
		string(SubscriptionStatusReasonSuspiciousActivity),
		string(SubscriptionStatusReasonTransferred),
	}
}

func (s *SubscriptionStatusReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionStatusReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionStatusReason(input string) (*SubscriptionStatusReason, error) {
	vals := map[string]SubscriptionStatusReason{
		"cancelled":            SubscriptionStatusReasonCancelled,
		"expired":              SubscriptionStatusReasonExpired,
		"none":                 SubscriptionStatusReasonNone,
		"other":                SubscriptionStatusReasonOther,
		"pastdue":              SubscriptionStatusReasonPastDue,
		"policyviolation":      SubscriptionStatusReasonPolicyViolation,
		"spendinglimitreached": SubscriptionStatusReasonSpendingLimitReached,
		"suspiciousactivity":   SubscriptionStatusReasonSuspiciousActivity,
		"transferred":          SubscriptionStatusReasonTransferred,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionStatusReason(input)
	return &out, nil
}

type SubscriptionTransferValidationErrorCode string

const (
	SubscriptionTransferValidationErrorCodeAccountIsLocked                     SubscriptionTransferValidationErrorCode = "AccountIsLocked"
	SubscriptionTransferValidationErrorCodeAssetHasCap                         SubscriptionTransferValidationErrorCode = "AssetHasCap"
	SubscriptionTransferValidationErrorCodeAssetNotActive                      SubscriptionTransferValidationErrorCode = "AssetNotActive"
	SubscriptionTransferValidationErrorCodeBillingAccountInactive              SubscriptionTransferValidationErrorCode = "BillingAccountInactive"
	SubscriptionTransferValidationErrorCodeBillingProfilePastDue               SubscriptionTransferValidationErrorCode = "BillingProfilePastDue"
	SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed       SubscriptionTransferValidationErrorCode = "CrossBillingAccountNotAllowed"
	SubscriptionTransferValidationErrorCodeDestinationBillingProfileInactive   SubscriptionTransferValidationErrorCode = "DestinationBillingProfileInactive"
	SubscriptionTransferValidationErrorCodeDestinationBillingProfileNotFound   SubscriptionTransferValidationErrorCode = "DestinationBillingProfileNotFound"
	SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue    SubscriptionTransferValidationErrorCode = "DestinationBillingProfilePastDue"
	SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionInactive   SubscriptionTransferValidationErrorCode = "DestinationInvoiceSectionInactive"
	SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionNotFound   SubscriptionTransferValidationErrorCode = "DestinationInvoiceSectionNotFound"
	SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination SubscriptionTransferValidationErrorCode = "InsufficientPermissionOnDestination"
	SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource      SubscriptionTransferValidationErrorCode = "InsufficientPermissionOnSource"
	SubscriptionTransferValidationErrorCodeInvalidDestination                  SubscriptionTransferValidationErrorCode = "InvalidDestination"
	SubscriptionTransferValidationErrorCodeInvalidSource                       SubscriptionTransferValidationErrorCode = "InvalidSource"
	SubscriptionTransferValidationErrorCodeInvoiceSectionIsRestricted          SubscriptionTransferValidationErrorCode = "InvoiceSectionIsRestricted"
	SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination  SubscriptionTransferValidationErrorCode = "MarketplaceNotEnabledOnDestination"
	SubscriptionTransferValidationErrorCodeNoActiveAzurePlan                   SubscriptionTransferValidationErrorCode = "NoActiveAzurePlan"
	SubscriptionTransferValidationErrorCodeNone                                SubscriptionTransferValidationErrorCode = "None"
	SubscriptionTransferValidationErrorCodeOther                               SubscriptionTransferValidationErrorCode = "Other"
	SubscriptionTransferValidationErrorCodeProductInactive                     SubscriptionTransferValidationErrorCode = "ProductInactive"
	SubscriptionTransferValidationErrorCodeProductNotFound                     SubscriptionTransferValidationErrorCode = "ProductNotFound"
	SubscriptionTransferValidationErrorCodeProductTypeNotSupported             SubscriptionTransferValidationErrorCode = "ProductTypeNotSupported"
	SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue         SubscriptionTransferValidationErrorCode = "SourceBillingProfilePastDue"
	SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive        SubscriptionTransferValidationErrorCode = "SourceInvoiceSectionInactive"
	SubscriptionTransferValidationErrorCodeSubscriptionHasReservations         SubscriptionTransferValidationErrorCode = "SubscriptionHasReservations"
	SubscriptionTransferValidationErrorCodeSubscriptionNotActive               SubscriptionTransferValidationErrorCode = "SubscriptionNotActive"
	SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported        SubscriptionTransferValidationErrorCode = "SubscriptionTypeNotSupported"
)

func PossibleValuesForSubscriptionTransferValidationErrorCode() []string {
	return []string{
		string(SubscriptionTransferValidationErrorCodeAccountIsLocked),
		string(SubscriptionTransferValidationErrorCodeAssetHasCap),
		string(SubscriptionTransferValidationErrorCodeAssetNotActive),
		string(SubscriptionTransferValidationErrorCodeBillingAccountInactive),
		string(SubscriptionTransferValidationErrorCodeBillingProfilePastDue),
		string(SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed),
		string(SubscriptionTransferValidationErrorCodeDestinationBillingProfileInactive),
		string(SubscriptionTransferValidationErrorCodeDestinationBillingProfileNotFound),
		string(SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue),
		string(SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionInactive),
		string(SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionNotFound),
		string(SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination),
		string(SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource),
		string(SubscriptionTransferValidationErrorCodeInvalidDestination),
		string(SubscriptionTransferValidationErrorCodeInvalidSource),
		string(SubscriptionTransferValidationErrorCodeInvoiceSectionIsRestricted),
		string(SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination),
		string(SubscriptionTransferValidationErrorCodeNoActiveAzurePlan),
		string(SubscriptionTransferValidationErrorCodeNone),
		string(SubscriptionTransferValidationErrorCodeOther),
		string(SubscriptionTransferValidationErrorCodeProductInactive),
		string(SubscriptionTransferValidationErrorCodeProductNotFound),
		string(SubscriptionTransferValidationErrorCodeProductTypeNotSupported),
		string(SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue),
		string(SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive),
		string(SubscriptionTransferValidationErrorCodeSubscriptionHasReservations),
		string(SubscriptionTransferValidationErrorCodeSubscriptionNotActive),
		string(SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported),
	}
}

func (s *SubscriptionTransferValidationErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscriptionTransferValidationErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscriptionTransferValidationErrorCode(input string) (*SubscriptionTransferValidationErrorCode, error) {
	vals := map[string]SubscriptionTransferValidationErrorCode{
		"accountislocked":                     SubscriptionTransferValidationErrorCodeAccountIsLocked,
		"assethascap":                         SubscriptionTransferValidationErrorCodeAssetHasCap,
		"assetnotactive":                      SubscriptionTransferValidationErrorCodeAssetNotActive,
		"billingaccountinactive":              SubscriptionTransferValidationErrorCodeBillingAccountInactive,
		"billingprofilepastdue":               SubscriptionTransferValidationErrorCodeBillingProfilePastDue,
		"crossbillingaccountnotallowed":       SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed,
		"destinationbillingprofileinactive":   SubscriptionTransferValidationErrorCodeDestinationBillingProfileInactive,
		"destinationbillingprofilenotfound":   SubscriptionTransferValidationErrorCodeDestinationBillingProfileNotFound,
		"destinationbillingprofilepastdue":    SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue,
		"destinationinvoicesectioninactive":   SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionInactive,
		"destinationinvoicesectionnotfound":   SubscriptionTransferValidationErrorCodeDestinationInvoiceSectionNotFound,
		"insufficientpermissionondestination": SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination,
		"insufficientpermissiononsource":      SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource,
		"invaliddestination":                  SubscriptionTransferValidationErrorCodeInvalidDestination,
		"invalidsource":                       SubscriptionTransferValidationErrorCodeInvalidSource,
		"invoicesectionisrestricted":          SubscriptionTransferValidationErrorCodeInvoiceSectionIsRestricted,
		"marketplacenotenabledondestination":  SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination,
		"noactiveazureplan":                   SubscriptionTransferValidationErrorCodeNoActiveAzurePlan,
		"none":                                SubscriptionTransferValidationErrorCodeNone,
		"other":                               SubscriptionTransferValidationErrorCodeOther,
		"productinactive":                     SubscriptionTransferValidationErrorCodeProductInactive,
		"productnotfound":                     SubscriptionTransferValidationErrorCodeProductNotFound,
		"producttypenotsupported":             SubscriptionTransferValidationErrorCodeProductTypeNotSupported,
		"sourcebillingprofilepastdue":         SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue,
		"sourceinvoicesectioninactive":        SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive,
		"subscriptionhasreservations":         SubscriptionTransferValidationErrorCodeSubscriptionHasReservations,
		"subscriptionnotactive":               SubscriptionTransferValidationErrorCodeSubscriptionNotActive,
		"subscriptiontypenotsupported":        SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscriptionTransferValidationErrorCode(input)
	return &out, nil
}
