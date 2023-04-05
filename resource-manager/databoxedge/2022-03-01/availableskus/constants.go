package availableskus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShipmentType string

const (
	ShipmentTypeNotApplicable     ShipmentType = "NotApplicable"
	ShipmentTypeSelfPickup        ShipmentType = "SelfPickup"
	ShipmentTypeShippedToCustomer ShipmentType = "ShippedToCustomer"
)

func PossibleValuesForShipmentType() []string {
	return []string{
		string(ShipmentTypeNotApplicable),
		string(ShipmentTypeSelfPickup),
		string(ShipmentTypeShippedToCustomer),
	}
}

type SkuAvailability string

const (
	SkuAvailabilityAvailable   SkuAvailability = "Available"
	SkuAvailabilityUnavailable SkuAvailability = "Unavailable"
)

func PossibleValuesForSkuAvailability() []string {
	return []string{
		string(SkuAvailabilityAvailable),
		string(SkuAvailabilityUnavailable),
	}
}

type SkuName string

const (
	SkuNameEPTwoOneTwoEightGPUOneMxOneW   SkuName = "EP2_128_GPU1_Mx1_W"
	SkuNameEPTwoOneTwoEightOneTFourMxOneW SkuName = "EP2_128_1T4_Mx1_W"
	SkuNameEPTwoSixFourMxOneW             SkuName = "EP2_64_Mx1_W"
	SkuNameEPTwoSixFourOneVPUW            SkuName = "EP2_64_1VPU_W"
	SkuNameEPTwoTwoFiveSixGPUTwoMxOne     SkuName = "EP2_256_GPU2_Mx1"
	SkuNameEPTwoTwoFiveSixTwoTFourW       SkuName = "EP2_256_2T4_W"
	SkuNameEdge                           SkuName = "Edge"
	SkuNameEdgeMRMini                     SkuName = "EdgeMR_Mini"
	SkuNameEdgeMRTCP                      SkuName = "EdgeMR_TCP"
	SkuNameEdgePBase                      SkuName = "EdgeP_Base"
	SkuNameEdgePHigh                      SkuName = "EdgeP_High"
	SkuNameEdgePRBase                     SkuName = "EdgePR_Base"
	SkuNameEdgePRBaseUPS                  SkuName = "EdgePR_Base_UPS"
	SkuNameGPU                            SkuName = "GPU"
	SkuNameGateway                        SkuName = "Gateway"
	SkuNameManagement                     SkuName = "Management"
	SkuNameRCALarge                       SkuName = "RCA_Large"
	SkuNameRCASmall                       SkuName = "RCA_Small"
	SkuNameRDC                            SkuName = "RDC"
	SkuNameTCALarge                       SkuName = "TCA_Large"
	SkuNameTCASmall                       SkuName = "TCA_Small"
	SkuNameTDC                            SkuName = "TDC"
	SkuNameTEAFourNodeHeater              SkuName = "TEA_4Node_Heater"
	SkuNameTEAFourNodeUPSHeater           SkuName = "TEA_4Node_UPS_Heater"
	SkuNameTEAOneNode                     SkuName = "TEA_1Node"
	SkuNameTEAOneNodeHeater               SkuName = "TEA_1Node_Heater"
	SkuNameTEAOneNodeUPS                  SkuName = "TEA_1Node_UPS"
	SkuNameTEAOneNodeUPSHeater            SkuName = "TEA_1Node_UPS_Heater"
	SkuNameTMA                            SkuName = "TMA"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameEPTwoOneTwoEightGPUOneMxOneW),
		string(SkuNameEPTwoOneTwoEightOneTFourMxOneW),
		string(SkuNameEPTwoSixFourMxOneW),
		string(SkuNameEPTwoSixFourOneVPUW),
		string(SkuNameEPTwoTwoFiveSixGPUTwoMxOne),
		string(SkuNameEPTwoTwoFiveSixTwoTFourW),
		string(SkuNameEdge),
		string(SkuNameEdgeMRMini),
		string(SkuNameEdgeMRTCP),
		string(SkuNameEdgePBase),
		string(SkuNameEdgePHigh),
		string(SkuNameEdgePRBase),
		string(SkuNameEdgePRBaseUPS),
		string(SkuNameGPU),
		string(SkuNameGateway),
		string(SkuNameManagement),
		string(SkuNameRCALarge),
		string(SkuNameRCASmall),
		string(SkuNameRDC),
		string(SkuNameTCALarge),
		string(SkuNameTCASmall),
		string(SkuNameTDC),
		string(SkuNameTEAFourNodeHeater),
		string(SkuNameTEAFourNodeUPSHeater),
		string(SkuNameTEAOneNode),
		string(SkuNameTEAOneNodeHeater),
		string(SkuNameTEAOneNodeUPS),
		string(SkuNameTEAOneNodeUPSHeater),
		string(SkuNameTMA),
	}
}

type SkuSignupOption string

const (
	SkuSignupOptionAvailable SkuSignupOption = "Available"
	SkuSignupOptionNone      SkuSignupOption = "None"
)

func PossibleValuesForSkuSignupOption() []string {
	return []string{
		string(SkuSignupOptionAvailable),
		string(SkuSignupOptionNone),
	}
}

type SkuTier string

const (
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierStandard),
	}
}

type SkuVersion string

const (
	SkuVersionPreview SkuVersion = "Preview"
	SkuVersionStable  SkuVersion = "Stable"
)

func PossibleValuesForSkuVersion() []string {
	return []string{
		string(SkuVersionPreview),
		string(SkuVersionStable),
	}
}
