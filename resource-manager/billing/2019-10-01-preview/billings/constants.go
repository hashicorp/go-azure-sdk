package billings

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductTransferValidationErrorCode string

const (
	ProductTransferValidationErrorCodeCrossBillingAccountNotAllowed            ProductTransferValidationErrorCode = "CrossBillingAccountNotAllowed"
	ProductTransferValidationErrorCodeDestinationBillingProfilePastDue         ProductTransferValidationErrorCode = "DestinationBillingProfilePastDue"
	ProductTransferValidationErrorCodeInsufficientPermissionOnDestination      ProductTransferValidationErrorCode = "InsufficientPermissionOnDestination"
	ProductTransferValidationErrorCodeInsufficientPermissionOnSource           ProductTransferValidationErrorCode = "InsufficientPermissionOnSource"
	ProductTransferValidationErrorCodeInvalidSource                            ProductTransferValidationErrorCode = "InvalidSource"
	ProductTransferValidationErrorCodeNotAvailableForDestinationMarket         ProductTransferValidationErrorCode = "NotAvailableForDestinationMarket"
	ProductTransferValidationErrorCodeOneTimePurchaseProductTransferNotAllowed ProductTransferValidationErrorCode = "OneTimePurchaseProductTransferNotAllowed"
	ProductTransferValidationErrorCodeProductNotActive                         ProductTransferValidationErrorCode = "ProductNotActive"
	ProductTransferValidationErrorCodeProductTypeNotSupported                  ProductTransferValidationErrorCode = "ProductTypeNotSupported"
)

func PossibleValuesForProductTransferValidationErrorCode() []string {
	return []string{
		string(ProductTransferValidationErrorCodeCrossBillingAccountNotAllowed),
		string(ProductTransferValidationErrorCodeDestinationBillingProfilePastDue),
		string(ProductTransferValidationErrorCodeInsufficientPermissionOnDestination),
		string(ProductTransferValidationErrorCodeInsufficientPermissionOnSource),
		string(ProductTransferValidationErrorCodeInvalidSource),
		string(ProductTransferValidationErrorCodeNotAvailableForDestinationMarket),
		string(ProductTransferValidationErrorCodeOneTimePurchaseProductTransferNotAllowed),
		string(ProductTransferValidationErrorCodeProductNotActive),
		string(ProductTransferValidationErrorCodeProductTypeNotSupported),
	}
}

func (s *ProductTransferValidationErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProductTransferValidationErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProductTransferValidationErrorCode(input string) (*ProductTransferValidationErrorCode, error) {
	vals := map[string]ProductTransferValidationErrorCode{
		"crossbillingaccountnotallowed":            ProductTransferValidationErrorCodeCrossBillingAccountNotAllowed,
		"destinationbillingprofilepastdue":         ProductTransferValidationErrorCodeDestinationBillingProfilePastDue,
		"insufficientpermissionondestination":      ProductTransferValidationErrorCodeInsufficientPermissionOnDestination,
		"insufficientpermissiononsource":           ProductTransferValidationErrorCodeInsufficientPermissionOnSource,
		"invalidsource":                            ProductTransferValidationErrorCodeInvalidSource,
		"notavailablefordestinationmarket":         ProductTransferValidationErrorCodeNotAvailableForDestinationMarket,
		"onetimepurchaseproducttransfernotallowed": ProductTransferValidationErrorCodeOneTimePurchaseProductTransferNotAllowed,
		"productnotactive":                         ProductTransferValidationErrorCodeProductNotActive,
		"producttypenotsupported":                  ProductTransferValidationErrorCodeProductTypeNotSupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProductTransferValidationErrorCode(input)
	return &out, nil
}

type SubscriptionTransferValidationErrorCode string

const (
	SubscriptionTransferValidationErrorCodeBillingAccountInactive              SubscriptionTransferValidationErrorCode = "BillingAccountInactive"
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
	SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination  SubscriptionTransferValidationErrorCode = "MarketplaceNotEnabledOnDestination"
	SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket    SubscriptionTransferValidationErrorCode = "NotAvailableForDestinationMarket"
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
		string(SubscriptionTransferValidationErrorCodeBillingAccountInactive),
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
		string(SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination),
		string(SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket),
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
		"billingaccountinactive":              SubscriptionTransferValidationErrorCodeBillingAccountInactive,
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
		"marketplacenotenabledondestination":  SubscriptionTransferValidationErrorCodeMarketplaceNotEnabledOnDestination,
		"notavailablefordestinationmarket":    SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket,
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
