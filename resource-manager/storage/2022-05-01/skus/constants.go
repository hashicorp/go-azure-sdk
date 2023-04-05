package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Kind string

const (
	KindBlobStorage      Kind = "BlobStorage"
	KindBlockBlobStorage Kind = "BlockBlobStorage"
	KindFileStorage      Kind = "FileStorage"
	KindStorage          Kind = "Storage"
	KindStorageVTwo      Kind = "StorageV2"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindBlobStorage),
		string(KindBlockBlobStorage),
		string(KindFileStorage),
		string(KindStorage),
		string(KindStorageVTwo),
	}
}

type ReasonCode string

const (
	ReasonCodeNotAvailableForSubscription ReasonCode = "NotAvailableForSubscription"
	ReasonCodeQuotaId                     ReasonCode = "QuotaId"
)

func PossibleValuesForReasonCode() []string {
	return []string{
		string(ReasonCodeNotAvailableForSubscription),
		string(ReasonCodeQuotaId),
	}
}

type SkuName string

const (
	SkuNamePremiumLRS     SkuName = "Premium_LRS"
	SkuNamePremiumZRS     SkuName = "Premium_ZRS"
	SkuNameStandardGRS    SkuName = "Standard_GRS"
	SkuNameStandardGZRS   SkuName = "Standard_GZRS"
	SkuNameStandardLRS    SkuName = "Standard_LRS"
	SkuNameStandardRAGRS  SkuName = "Standard_RAGRS"
	SkuNameStandardRAGZRS SkuName = "Standard_RAGZRS"
	SkuNameStandardZRS    SkuName = "Standard_ZRS"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePremiumLRS),
		string(SkuNamePremiumZRS),
		string(SkuNameStandardGRS),
		string(SkuNameStandardGZRS),
		string(SkuNameStandardLRS),
		string(SkuNameStandardRAGRS),
		string(SkuNameStandardRAGZRS),
		string(SkuNameStandardZRS),
	}
}

type SkuTier string

const (
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}
