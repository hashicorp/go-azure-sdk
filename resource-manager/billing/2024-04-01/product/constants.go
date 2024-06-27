package product

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

type MoveValidationErrorCode string

const (
	MoveValidationErrorCodeBillingAccountInactive              MoveValidationErrorCode = "BillingAccountInactive"
	MoveValidationErrorCodeDestinationBillingProfileInactive   MoveValidationErrorCode = "DestinationBillingProfileInactive"
	MoveValidationErrorCodeDestinationBillingProfileNotFound   MoveValidationErrorCode = "DestinationBillingProfileNotFound"
	MoveValidationErrorCodeDestinationBillingProfilePastDue    MoveValidationErrorCode = "DestinationBillingProfilePastDue"
	MoveValidationErrorCodeDestinationInvoiceSectionInactive   MoveValidationErrorCode = "DestinationInvoiceSectionInactive"
	MoveValidationErrorCodeDestinationInvoiceSectionNotFound   MoveValidationErrorCode = "DestinationInvoiceSectionNotFound"
	MoveValidationErrorCodeInsufficientPermissionOnDestination MoveValidationErrorCode = "InsufficientPermissionOnDestination"
	MoveValidationErrorCodeInsufficientPermissionOnSource      MoveValidationErrorCode = "InsufficientPermissionOnSource"
	MoveValidationErrorCodeInvalidDestination                  MoveValidationErrorCode = "InvalidDestination"
	MoveValidationErrorCodeInvalidSource                       MoveValidationErrorCode = "InvalidSource"
	MoveValidationErrorCodeMarketplaceNotEnabledOnDestination  MoveValidationErrorCode = "MarketplaceNotEnabledOnDestination"
	MoveValidationErrorCodeOther                               MoveValidationErrorCode = "Other"
	MoveValidationErrorCodeProductInactive                     MoveValidationErrorCode = "ProductInactive"
	MoveValidationErrorCodeProductNotFound                     MoveValidationErrorCode = "ProductNotFound"
	MoveValidationErrorCodeProductTypeNotSupported             MoveValidationErrorCode = "ProductTypeNotSupported"
	MoveValidationErrorCodeSourceBillingProfilePastDue         MoveValidationErrorCode = "SourceBillingProfilePastDue"
	MoveValidationErrorCodeSourceInvoiceSectionInactive        MoveValidationErrorCode = "SourceInvoiceSectionInactive"
)

func PossibleValuesForMoveValidationErrorCode() []string {
	return []string{
		string(MoveValidationErrorCodeBillingAccountInactive),
		string(MoveValidationErrorCodeDestinationBillingProfileInactive),
		string(MoveValidationErrorCodeDestinationBillingProfileNotFound),
		string(MoveValidationErrorCodeDestinationBillingProfilePastDue),
		string(MoveValidationErrorCodeDestinationInvoiceSectionInactive),
		string(MoveValidationErrorCodeDestinationInvoiceSectionNotFound),
		string(MoveValidationErrorCodeInsufficientPermissionOnDestination),
		string(MoveValidationErrorCodeInsufficientPermissionOnSource),
		string(MoveValidationErrorCodeInvalidDestination),
		string(MoveValidationErrorCodeInvalidSource),
		string(MoveValidationErrorCodeMarketplaceNotEnabledOnDestination),
		string(MoveValidationErrorCodeOther),
		string(MoveValidationErrorCodeProductInactive),
		string(MoveValidationErrorCodeProductNotFound),
		string(MoveValidationErrorCodeProductTypeNotSupported),
		string(MoveValidationErrorCodeSourceBillingProfilePastDue),
		string(MoveValidationErrorCodeSourceInvoiceSectionInactive),
	}
}

func (s *MoveValidationErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMoveValidationErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMoveValidationErrorCode(input string) (*MoveValidationErrorCode, error) {
	vals := map[string]MoveValidationErrorCode{
		"billingaccountinactive":              MoveValidationErrorCodeBillingAccountInactive,
		"destinationbillingprofileinactive":   MoveValidationErrorCodeDestinationBillingProfileInactive,
		"destinationbillingprofilenotfound":   MoveValidationErrorCodeDestinationBillingProfileNotFound,
		"destinationbillingprofilepastdue":    MoveValidationErrorCodeDestinationBillingProfilePastDue,
		"destinationinvoicesectioninactive":   MoveValidationErrorCodeDestinationInvoiceSectionInactive,
		"destinationinvoicesectionnotfound":   MoveValidationErrorCodeDestinationInvoiceSectionNotFound,
		"insufficientpermissionondestination": MoveValidationErrorCodeInsufficientPermissionOnDestination,
		"insufficientpermissiononsource":      MoveValidationErrorCodeInsufficientPermissionOnSource,
		"invaliddestination":                  MoveValidationErrorCodeInvalidDestination,
		"invalidsource":                       MoveValidationErrorCodeInvalidSource,
		"marketplacenotenabledondestination":  MoveValidationErrorCodeMarketplaceNotEnabledOnDestination,
		"other":                               MoveValidationErrorCodeOther,
		"productinactive":                     MoveValidationErrorCodeProductInactive,
		"productnotfound":                     MoveValidationErrorCodeProductNotFound,
		"producttypenotsupported":             MoveValidationErrorCodeProductTypeNotSupported,
		"sourcebillingprofilepastdue":         MoveValidationErrorCodeSourceBillingProfilePastDue,
		"sourceinvoicesectioninactive":        MoveValidationErrorCodeSourceInvoiceSectionInactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MoveValidationErrorCode(input)
	return &out, nil
}

type ProductStatus string

const (
	ProductStatusActive    ProductStatus = "Active"
	ProductStatusAutoRenew ProductStatus = "AutoRenew"
	ProductStatusCanceled  ProductStatus = "Canceled"
	ProductStatusDeleted   ProductStatus = "Deleted"
	ProductStatusDisabled  ProductStatus = "Disabled"
	ProductStatusExpired   ProductStatus = "Expired"
	ProductStatusExpiring  ProductStatus = "Expiring"
	ProductStatusOther     ProductStatus = "Other"
	ProductStatusPastDue   ProductStatus = "PastDue"
	ProductStatusSuspended ProductStatus = "Suspended"
)

func PossibleValuesForProductStatus() []string {
	return []string{
		string(ProductStatusActive),
		string(ProductStatusAutoRenew),
		string(ProductStatusCanceled),
		string(ProductStatusDeleted),
		string(ProductStatusDisabled),
		string(ProductStatusExpired),
		string(ProductStatusExpiring),
		string(ProductStatusOther),
		string(ProductStatusPastDue),
		string(ProductStatusSuspended),
	}
}

func (s *ProductStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProductStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProductStatus(input string) (*ProductStatus, error) {
	vals := map[string]ProductStatus{
		"active":    ProductStatusActive,
		"autorenew": ProductStatusAutoRenew,
		"canceled":  ProductStatusCanceled,
		"deleted":   ProductStatusDeleted,
		"disabled":  ProductStatusDisabled,
		"expired":   ProductStatusExpired,
		"expiring":  ProductStatusExpiring,
		"other":     ProductStatusOther,
		"pastdue":   ProductStatusPastDue,
		"suspended": ProductStatusSuspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductStatus(input)
	return &out, nil
}
