package managedhsmkeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletionRecoveryLevel string

const (
	DeletionRecoveryLevelPurgeable                                DeletionRecoveryLevel = "Purgeable"
	DeletionRecoveryLevelRecoverable                              DeletionRecoveryLevel = "Recoverable"
	DeletionRecoveryLevelRecoverablePositiveProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"
	DeletionRecoveryLevelRecoverablePositivePurgeable             DeletionRecoveryLevel = "Recoverable+Purgeable"
)

func PossibleValuesForDeletionRecoveryLevel() []string {
	return []string{
		string(DeletionRecoveryLevelPurgeable),
		string(DeletionRecoveryLevelRecoverable),
		string(DeletionRecoveryLevelRecoverablePositiveProtectedSubscription),
		string(DeletionRecoveryLevelRecoverablePositivePurgeable),
	}
}

type JsonWebKeyCurveName string

const (
	JsonWebKeyCurveNamePNegativeFiveTwoOne     JsonWebKeyCurveName = "P-521"
	JsonWebKeyCurveNamePNegativeThreeEightFour JsonWebKeyCurveName = "P-384"
	JsonWebKeyCurveNamePNegativeTwoFiveSix     JsonWebKeyCurveName = "P-256"
	JsonWebKeyCurveNamePNegativeTwoFiveSixK    JsonWebKeyCurveName = "P-256K"
)

func PossibleValuesForJsonWebKeyCurveName() []string {
	return []string{
		string(JsonWebKeyCurveNamePNegativeFiveTwoOne),
		string(JsonWebKeyCurveNamePNegativeThreeEightFour),
		string(JsonWebKeyCurveNamePNegativeTwoFiveSix),
		string(JsonWebKeyCurveNamePNegativeTwoFiveSixK),
	}
}

type JsonWebKeyOperation string

const (
	JsonWebKeyOperationDecrypt   JsonWebKeyOperation = "decrypt"
	JsonWebKeyOperationEncrypt   JsonWebKeyOperation = "encrypt"
	JsonWebKeyOperationImport    JsonWebKeyOperation = "import"
	JsonWebKeyOperationRelease   JsonWebKeyOperation = "release"
	JsonWebKeyOperationSign      JsonWebKeyOperation = "sign"
	JsonWebKeyOperationUnwrapKey JsonWebKeyOperation = "unwrapKey"
	JsonWebKeyOperationVerify    JsonWebKeyOperation = "verify"
	JsonWebKeyOperationWrapKey   JsonWebKeyOperation = "wrapKey"
)

func PossibleValuesForJsonWebKeyOperation() []string {
	return []string{
		string(JsonWebKeyOperationDecrypt),
		string(JsonWebKeyOperationEncrypt),
		string(JsonWebKeyOperationImport),
		string(JsonWebKeyOperationRelease),
		string(JsonWebKeyOperationSign),
		string(JsonWebKeyOperationUnwrapKey),
		string(JsonWebKeyOperationVerify),
		string(JsonWebKeyOperationWrapKey),
	}
}

type JsonWebKeyType string

const (
	JsonWebKeyTypeEC             JsonWebKeyType = "EC"
	JsonWebKeyTypeECNegativeHSM  JsonWebKeyType = "EC-HSM"
	JsonWebKeyTypeRSA            JsonWebKeyType = "RSA"
	JsonWebKeyTypeRSANegativeHSM JsonWebKeyType = "RSA-HSM"
)

func PossibleValuesForJsonWebKeyType() []string {
	return []string{
		string(JsonWebKeyTypeEC),
		string(JsonWebKeyTypeECNegativeHSM),
		string(JsonWebKeyTypeRSA),
		string(JsonWebKeyTypeRSANegativeHSM),
	}
}

type KeyRotationPolicyActionType string

const (
	KeyRotationPolicyActionTypeNotify KeyRotationPolicyActionType = "notify"
	KeyRotationPolicyActionTypeRotate KeyRotationPolicyActionType = "rotate"
)

func PossibleValuesForKeyRotationPolicyActionType() []string {
	return []string{
		string(KeyRotationPolicyActionTypeNotify),
		string(KeyRotationPolicyActionTypeRotate),
	}
}
