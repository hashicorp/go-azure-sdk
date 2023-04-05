package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureScaleType string

const (
	AzureScaleTypeAutomatic AzureScaleType = "automatic"
	AzureScaleTypeManual    AzureScaleType = "manual"
	AzureScaleTypeNone      AzureScaleType = "none"
)

func PossibleValuesForAzureScaleType() []string {
	return []string{
		string(AzureScaleTypeAutomatic),
		string(AzureScaleTypeManual),
		string(AzureScaleTypeNone),
	}
}

type AzureSkuName string

const (
	AzureSkuNameDevNoSLAStandardDOneOneVTwo              AzureSkuName = "Dev(No SLA)_Standard_D11_v2"
	AzureSkuNameDevNoSLAStandardETwoaVFour               AzureSkuName = "Dev(No SLA)_Standard_E2a_v4"
	AzureSkuNameStandardDOneFourVTwo                     AzureSkuName = "Standard_D14_v2"
	AzureSkuNameStandardDOneOneVTwo                      AzureSkuName = "Standard_D11_v2"
	AzureSkuNameStandardDOneSixdVFive                    AzureSkuName = "Standard_D16d_v5"
	AzureSkuNameStandardDOneThreeVTwo                    AzureSkuName = "Standard_D13_v2"
	AzureSkuNameStandardDOneTwoVTwo                      AzureSkuName = "Standard_D12_v2"
	AzureSkuNameStandardDSOneFourVTwoPositiveFourTBPS    AzureSkuName = "Standard_DS14_v2+4TB_PS"
	AzureSkuNameStandardDSOneFourVTwoPositiveThreeTBPS   AzureSkuName = "Standard_DS14_v2+3TB_PS"
	AzureSkuNameStandardDSOneThreeVTwoPositiveOneTBPS    AzureSkuName = "Standard_DS13_v2+1TB_PS"
	AzureSkuNameStandardDSOneThreeVTwoPositiveTwoTBPS    AzureSkuName = "Standard_DS13_v2+2TB_PS"
	AzureSkuNameStandardDThreeTwodVFive                  AzureSkuName = "Standard_D32d_v5"
	AzureSkuNameStandardDThreeTwodVFour                  AzureSkuName = "Standard_D32d_v4"
	AzureSkuNameStandardECEightadsVFive                  AzureSkuName = "Standard_EC8ads_v5"
	AzureSkuNameStandardECEightasVFivePositiveOneTBPS    AzureSkuName = "Standard_EC8as_v5+1TB_PS"
	AzureSkuNameStandardECEightasVFivePositiveTwoTBPS    AzureSkuName = "Standard_EC8as_v5+2TB_PS"
	AzureSkuNameStandardECOneSixadsVFive                 AzureSkuName = "Standard_EC16ads_v5"
	AzureSkuNameStandardECOneSixasVFivePositiveFourTBPS  AzureSkuName = "Standard_EC16as_v5+4TB_PS"
	AzureSkuNameStandardECOneSixasVFivePositiveThreeTBPS AzureSkuName = "Standard_EC16as_v5+3TB_PS"
	AzureSkuNameStandardEEightZeroidsVFour               AzureSkuName = "Standard_E80ids_v4"
	AzureSkuNameStandardEEightaVFour                     AzureSkuName = "Standard_E8a_v4"
	AzureSkuNameStandardEEightadsVFive                   AzureSkuName = "Standard_E8ads_v5"
	AzureSkuNameStandardEEightasVFivePositiveOneTBPS     AzureSkuName = "Standard_E8as_v5+1TB_PS"
	AzureSkuNameStandardEEightasVFivePositiveTwoTBPS     AzureSkuName = "Standard_E8as_v5+2TB_PS"
	AzureSkuNameStandardEEightasVFourPositiveOneTBPS     AzureSkuName = "Standard_E8as_v4+1TB_PS"
	AzureSkuNameStandardEEightasVFourPositiveTwoTBPS     AzureSkuName = "Standard_E8as_v4+2TB_PS"
	AzureSkuNameStandardEEightdVFive                     AzureSkuName = "Standard_E8d_v5"
	AzureSkuNameStandardEEightdVFour                     AzureSkuName = "Standard_E8d_v4"
	AzureSkuNameStandardEEightsVFivePositiveOneTBPS      AzureSkuName = "Standard_E8s_v5+1TB_PS"
	AzureSkuNameStandardEEightsVFivePositiveTwoTBPS      AzureSkuName = "Standard_E8s_v5+2TB_PS"
	AzureSkuNameStandardEEightsVFourPositiveOneTBPS      AzureSkuName = "Standard_E8s_v4+1TB_PS"
	AzureSkuNameStandardEEightsVFourPositiveTwoTBPS      AzureSkuName = "Standard_E8s_v4+2TB_PS"
	AzureSkuNameStandardEFouraVFour                      AzureSkuName = "Standard_E4a_v4"
	AzureSkuNameStandardEFouradsVFive                    AzureSkuName = "Standard_E4ads_v5"
	AzureSkuNameStandardEFourdVFive                      AzureSkuName = "Standard_E4d_v5"
	AzureSkuNameStandardEFourdVFour                      AzureSkuName = "Standard_E4d_v4"
	AzureSkuNameStandardEOneSixaVFour                    AzureSkuName = "Standard_E16a_v4"
	AzureSkuNameStandardEOneSixadsVFive                  AzureSkuName = "Standard_E16ads_v5"
	AzureSkuNameStandardEOneSixasVFivePositiveFourTBPS   AzureSkuName = "Standard_E16as_v5+4TB_PS"
	AzureSkuNameStandardEOneSixasVFivePositiveThreeTBPS  AzureSkuName = "Standard_E16as_v5+3TB_PS"
	AzureSkuNameStandardEOneSixasVFourPositiveFourTBPS   AzureSkuName = "Standard_E16as_v4+4TB_PS"
	AzureSkuNameStandardEOneSixasVFourPositiveThreeTBPS  AzureSkuName = "Standard_E16as_v4+3TB_PS"
	AzureSkuNameStandardEOneSixdVFive                    AzureSkuName = "Standard_E16d_v5"
	AzureSkuNameStandardEOneSixdVFour                    AzureSkuName = "Standard_E16d_v4"
	AzureSkuNameStandardEOneSixsVFivePositiveFourTBPS    AzureSkuName = "Standard_E16s_v5+4TB_PS"
	AzureSkuNameStandardEOneSixsVFivePositiveThreeTBPS   AzureSkuName = "Standard_E16s_v5+3TB_PS"
	AzureSkuNameStandardEOneSixsVFourPositiveFourTBPS    AzureSkuName = "Standard_E16s_v4+4TB_PS"
	AzureSkuNameStandardEOneSixsVFourPositiveThreeTBPS   AzureSkuName = "Standard_E16s_v4+3TB_PS"
	AzureSkuNameStandardESixFouriVThree                  AzureSkuName = "Standard_E64i_v3"
	AzureSkuNameStandardETwoaVFour                       AzureSkuName = "Standard_E2a_v4"
	AzureSkuNameStandardETwoadsVFive                     AzureSkuName = "Standard_E2ads_v5"
	AzureSkuNameStandardETwodVFive                       AzureSkuName = "Standard_E2d_v5"
	AzureSkuNameStandardETwodVFour                       AzureSkuName = "Standard_E2d_v4"
	AzureSkuNameStandardLEightasVThree                   AzureSkuName = "Standard_L8as_v3"
	AzureSkuNameStandardLEights                          AzureSkuName = "Standard_L8s"
	AzureSkuNameStandardLEightsVThree                    AzureSkuName = "Standard_L8s_v3"
	AzureSkuNameStandardLEightsVTwo                      AzureSkuName = "Standard_L8s_v2"
	AzureSkuNameStandardLFours                           AzureSkuName = "Standard_L4s"
	AzureSkuNameStandardLOneSixasVThree                  AzureSkuName = "Standard_L16as_v3"
	AzureSkuNameStandardLOneSixs                         AzureSkuName = "Standard_L16s"
	AzureSkuNameStandardLOneSixsVThree                   AzureSkuName = "Standard_L16s_v3"
	AzureSkuNameStandardLOneSixsVTwo                     AzureSkuName = "Standard_L16s_v2"
	AzureSkuNameStandardLThreeTwoasVThree                AzureSkuName = "Standard_L32as_v3"
	AzureSkuNameStandardLThreeTwosVThree                 AzureSkuName = "Standard_L32s_v3"
)

func PossibleValuesForAzureSkuName() []string {
	return []string{
		string(AzureSkuNameDevNoSLAStandardDOneOneVTwo),
		string(AzureSkuNameDevNoSLAStandardETwoaVFour),
		string(AzureSkuNameStandardDOneFourVTwo),
		string(AzureSkuNameStandardDOneOneVTwo),
		string(AzureSkuNameStandardDOneSixdVFive),
		string(AzureSkuNameStandardDOneThreeVTwo),
		string(AzureSkuNameStandardDOneTwoVTwo),
		string(AzureSkuNameStandardDSOneFourVTwoPositiveFourTBPS),
		string(AzureSkuNameStandardDSOneFourVTwoPositiveThreeTBPS),
		string(AzureSkuNameStandardDSOneThreeVTwoPositiveOneTBPS),
		string(AzureSkuNameStandardDSOneThreeVTwoPositiveTwoTBPS),
		string(AzureSkuNameStandardDThreeTwodVFive),
		string(AzureSkuNameStandardDThreeTwodVFour),
		string(AzureSkuNameStandardECEightadsVFive),
		string(AzureSkuNameStandardECEightasVFivePositiveOneTBPS),
		string(AzureSkuNameStandardECEightasVFivePositiveTwoTBPS),
		string(AzureSkuNameStandardECOneSixadsVFive),
		string(AzureSkuNameStandardECOneSixasVFivePositiveFourTBPS),
		string(AzureSkuNameStandardECOneSixasVFivePositiveThreeTBPS),
		string(AzureSkuNameStandardEEightZeroidsVFour),
		string(AzureSkuNameStandardEEightaVFour),
		string(AzureSkuNameStandardEEightadsVFive),
		string(AzureSkuNameStandardEEightasVFivePositiveOneTBPS),
		string(AzureSkuNameStandardEEightasVFivePositiveTwoTBPS),
		string(AzureSkuNameStandardEEightasVFourPositiveOneTBPS),
		string(AzureSkuNameStandardEEightasVFourPositiveTwoTBPS),
		string(AzureSkuNameStandardEEightdVFive),
		string(AzureSkuNameStandardEEightdVFour),
		string(AzureSkuNameStandardEEightsVFivePositiveOneTBPS),
		string(AzureSkuNameStandardEEightsVFivePositiveTwoTBPS),
		string(AzureSkuNameStandardEEightsVFourPositiveOneTBPS),
		string(AzureSkuNameStandardEEightsVFourPositiveTwoTBPS),
		string(AzureSkuNameStandardEFouraVFour),
		string(AzureSkuNameStandardEFouradsVFive),
		string(AzureSkuNameStandardEFourdVFive),
		string(AzureSkuNameStandardEFourdVFour),
		string(AzureSkuNameStandardEOneSixaVFour),
		string(AzureSkuNameStandardEOneSixadsVFive),
		string(AzureSkuNameStandardEOneSixasVFivePositiveFourTBPS),
		string(AzureSkuNameStandardEOneSixasVFivePositiveThreeTBPS),
		string(AzureSkuNameStandardEOneSixasVFourPositiveFourTBPS),
		string(AzureSkuNameStandardEOneSixasVFourPositiveThreeTBPS),
		string(AzureSkuNameStandardEOneSixdVFive),
		string(AzureSkuNameStandardEOneSixdVFour),
		string(AzureSkuNameStandardEOneSixsVFivePositiveFourTBPS),
		string(AzureSkuNameStandardEOneSixsVFivePositiveThreeTBPS),
		string(AzureSkuNameStandardEOneSixsVFourPositiveFourTBPS),
		string(AzureSkuNameStandardEOneSixsVFourPositiveThreeTBPS),
		string(AzureSkuNameStandardESixFouriVThree),
		string(AzureSkuNameStandardETwoaVFour),
		string(AzureSkuNameStandardETwoadsVFive),
		string(AzureSkuNameStandardETwodVFive),
		string(AzureSkuNameStandardETwodVFour),
		string(AzureSkuNameStandardLEightasVThree),
		string(AzureSkuNameStandardLEights),
		string(AzureSkuNameStandardLEightsVThree),
		string(AzureSkuNameStandardLEightsVTwo),
		string(AzureSkuNameStandardLFours),
		string(AzureSkuNameStandardLOneSixasVThree),
		string(AzureSkuNameStandardLOneSixs),
		string(AzureSkuNameStandardLOneSixsVThree),
		string(AzureSkuNameStandardLOneSixsVTwo),
		string(AzureSkuNameStandardLThreeTwoasVThree),
		string(AzureSkuNameStandardLThreeTwosVThree),
	}
}

type AzureSkuTier string

const (
	AzureSkuTierBasic    AzureSkuTier = "Basic"
	AzureSkuTierStandard AzureSkuTier = "Standard"
)

func PossibleValuesForAzureSkuTier() []string {
	return []string{
		string(AzureSkuTierBasic),
		string(AzureSkuTierStandard),
	}
}

type ClusterNetworkAccessFlag string

const (
	ClusterNetworkAccessFlagDisabled ClusterNetworkAccessFlag = "Disabled"
	ClusterNetworkAccessFlagEnabled  ClusterNetworkAccessFlag = "Enabled"
)

func PossibleValuesForClusterNetworkAccessFlag() []string {
	return []string{
		string(ClusterNetworkAccessFlagDisabled),
		string(ClusterNetworkAccessFlagEnabled),
	}
}

type ClusterType string

const (
	ClusterTypeMicrosoftPointKustoClusters ClusterType = "Microsoft.Kusto/clusters"
)

func PossibleValuesForClusterType() []string {
	return []string{
		string(ClusterTypeMicrosoftPointKustoClusters),
	}
}

type DatabaseShareOrigin string

const (
	DatabaseShareOriginDataShare DatabaseShareOrigin = "DataShare"
	DatabaseShareOriginDirect    DatabaseShareOrigin = "Direct"
	DatabaseShareOriginOther     DatabaseShareOrigin = "Other"
)

func PossibleValuesForDatabaseShareOrigin() []string {
	return []string{
		string(DatabaseShareOriginDataShare),
		string(DatabaseShareOriginDirect),
		string(DatabaseShareOriginOther),
	}
}

type EngineType string

const (
	EngineTypeVThree EngineType = "V3"
	EngineTypeVTwo   EngineType = "V2"
)

func PossibleValuesForEngineType() []string {
	return []string{
		string(EngineTypeVThree),
		string(EngineTypeVTwo),
	}
}

type LanguageExtensionImageName string

const (
	LanguageExtensionImageNamePythonThreeOneZeroEight LanguageExtensionImageName = "Python3_10_8"
	LanguageExtensionImageNamePythonThreeSixFive      LanguageExtensionImageName = "Python3_6_5"
	LanguageExtensionImageNameR                       LanguageExtensionImageName = "R"
)

func PossibleValuesForLanguageExtensionImageName() []string {
	return []string{
		string(LanguageExtensionImageNamePythonThreeOneZeroEight),
		string(LanguageExtensionImageNamePythonThreeSixFive),
		string(LanguageExtensionImageNameR),
	}
}

type LanguageExtensionName string

const (
	LanguageExtensionNamePYTHON LanguageExtensionName = "PYTHON"
	LanguageExtensionNameR      LanguageExtensionName = "R"
)

func PossibleValuesForLanguageExtensionName() []string {
	return []string{
		string(LanguageExtensionNamePYTHON),
		string(LanguageExtensionNameR),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
	}
}

type PublicIPType string

const (
	PublicIPTypeDualStack PublicIPType = "DualStack"
	PublicIPTypeIPvFour   PublicIPType = "IPv4"
)

func PossibleValuesForPublicIPType() []string {
	return []string{
		string(PublicIPTypeDualStack),
		string(PublicIPTypeIPvFour),
	}
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
		string(ReasonInvalid),
	}
}

type State string

const (
	StateCreating    State = "Creating"
	StateDeleted     State = "Deleted"
	StateDeleting    State = "Deleting"
	StateRunning     State = "Running"
	StateStarting    State = "Starting"
	StateStopped     State = "Stopped"
	StateStopping    State = "Stopping"
	StateUnavailable State = "Unavailable"
	StateUpdating    State = "Updating"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateCreating),
		string(StateDeleted),
		string(StateDeleting),
		string(StateRunning),
		string(StateStarting),
		string(StateStopped),
		string(StateStopping),
		string(StateUnavailable),
		string(StateUpdating),
	}
}
