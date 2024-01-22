package billingsubscriptions

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

type BillingSubscriptionStatus string

const (
	BillingSubscriptionStatusActive    BillingSubscriptionStatus = "Active"
	BillingSubscriptionStatusAutoRenew BillingSubscriptionStatus = "AutoRenew"
	BillingSubscriptionStatusCancelled BillingSubscriptionStatus = "Cancelled"
	BillingSubscriptionStatusDeleted   BillingSubscriptionStatus = "Deleted"
	BillingSubscriptionStatusDisabled  BillingSubscriptionStatus = "Disabled"
	BillingSubscriptionStatusExpired   BillingSubscriptionStatus = "Expired"
	BillingSubscriptionStatusExpiring  BillingSubscriptionStatus = "Expiring"
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

type SubscriptionEnrollmentAccountStatus string

const (
	SubscriptionEnrollmentAccountStatusActive         SubscriptionEnrollmentAccountStatus = "Active"
	SubscriptionEnrollmentAccountStatusCancelled      SubscriptionEnrollmentAccountStatus = "Cancelled"
	SubscriptionEnrollmentAccountStatusDeleted        SubscriptionEnrollmentAccountStatus = "Deleted"
	SubscriptionEnrollmentAccountStatusExpired        SubscriptionEnrollmentAccountStatus = "Expired"
	SubscriptionEnrollmentAccountStatusTransferredOut SubscriptionEnrollmentAccountStatus = "TransferredOut"
	SubscriptionEnrollmentAccountStatusTransferring   SubscriptionEnrollmentAccountStatus = "Transferring"
)

func PossibleValuesForSubscriptionEnrollmentAccountStatus() []string {
	return []string{
		string(SubscriptionEnrollmentAccountStatusActive),
		string(SubscriptionEnrollmentAccountStatusCancelled),
		string(SubscriptionEnrollmentAccountStatusDeleted),
		string(SubscriptionEnrollmentAccountStatusExpired),
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
	SubscriptionTransferValidationErrorCodeProductInactive                     SubscriptionTransferValidationErrorCode = "ProductInactive"
	SubscriptionTransferValidationErrorCodeProductNotFound                     SubscriptionTransferValidationErrorCode = "ProductNotFound"
	SubscriptionTransferValidationErrorCodeProductTypeNotSupported             SubscriptionTransferValidationErrorCode = "ProductTypeNotSupported"
	SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue         SubscriptionTransferValidationErrorCode = "SourceBillingProfilePastDue"
	SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive        SubscriptionTransferValidationErrorCode = "SourceInvoiceSectionInactive"
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
		string(SubscriptionTransferValidationErrorCodeProductInactive),
		string(SubscriptionTransferValidationErrorCodeProductNotFound),
		string(SubscriptionTransferValidationErrorCodeProductTypeNotSupported),
		string(SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue),
		string(SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive),
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
		"productinactive":                     SubscriptionTransferValidationErrorCodeProductInactive,
		"productnotfound":                     SubscriptionTransferValidationErrorCodeProductNotFound,
		"producttypenotsupported":             SubscriptionTransferValidationErrorCodeProductTypeNotSupported,
		"sourcebillingprofilepastdue":         SubscriptionTransferValidationErrorCodeSourceBillingProfilePastDue,
		"sourceinvoicesectioninactive":        SubscriptionTransferValidationErrorCodeSourceInvoiceSectionInactive,
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
