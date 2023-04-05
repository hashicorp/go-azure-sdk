package dedicatedhsms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JsonWebKeyType string

const (
	JsonWebKeyTypeAllocating    JsonWebKeyType = "Allocating"
	JsonWebKeyTypeCheckingQuota JsonWebKeyType = "CheckingQuota"
	JsonWebKeyTypeConnecting    JsonWebKeyType = "Connecting"
	JsonWebKeyTypeDeleting      JsonWebKeyType = "Deleting"
	JsonWebKeyTypeFailed        JsonWebKeyType = "Failed"
	JsonWebKeyTypeProvisioning  JsonWebKeyType = "Provisioning"
	JsonWebKeyTypeSucceeded     JsonWebKeyType = "Succeeded"
)

func PossibleValuesForJsonWebKeyType() []string {
	return []string{
		string(JsonWebKeyTypeAllocating),
		string(JsonWebKeyTypeCheckingQuota),
		string(JsonWebKeyTypeConnecting),
		string(JsonWebKeyTypeDeleting),
		string(JsonWebKeyTypeFailed),
		string(JsonWebKeyTypeProvisioning),
		string(JsonWebKeyTypeSucceeded),
	}
}

type SkuName string

const (
	SkuNamePayShieldOneZeroKLMKOneCPSSixZero         SkuName = "payShield10K_LMK1_CPS60"
	SkuNamePayShieldOneZeroKLMKOneCPSTwoFiveZero     SkuName = "payShield10K_LMK1_CPS250"
	SkuNamePayShieldOneZeroKLMKOneCPSTwoFiveZeroZero SkuName = "payShield10K_LMK1_CPS2500"
	SkuNamePayShieldOneZeroKLMKTwoCPSSixZero         SkuName = "payShield10K_LMK2_CPS60"
	SkuNamePayShieldOneZeroKLMKTwoCPSTwoFiveZero     SkuName = "payShield10K_LMK2_CPS250"
	SkuNamePayShieldOneZeroKLMKTwoCPSTwoFiveZeroZero SkuName = "payShield10K_LMK2_CPS2500"
	SkuNameSafeNetLunaNetworkHSMASevenNineZero       SkuName = "SafeNet Luna Network HSM A790"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNamePayShieldOneZeroKLMKOneCPSSixZero),
		string(SkuNamePayShieldOneZeroKLMKOneCPSTwoFiveZero),
		string(SkuNamePayShieldOneZeroKLMKOneCPSTwoFiveZeroZero),
		string(SkuNamePayShieldOneZeroKLMKTwoCPSSixZero),
		string(SkuNamePayShieldOneZeroKLMKTwoCPSTwoFiveZero),
		string(SkuNamePayShieldOneZeroKLMKTwoCPSTwoFiveZeroZero),
		string(SkuNameSafeNetLunaNetworkHSMASevenNineZero),
	}
}
