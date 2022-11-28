package virtualmachineimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineImageProperties struct {
	AutomaticOSUpgradeProperties *AutomaticOSUpgradeProperties `json:"automaticOSUpgradeProperties"`
	DataDiskImages               *[]DataDiskImage              `json:"dataDiskImages,omitempty"`
	Disallowed                   *DisallowedConfiguration      `json:"disallowed"`
	Features                     *[]VirtualMachineImageFeature `json:"features,omitempty"`
	HyperVGeneration             *HyperVGenerationTypes        `json:"hyperVGeneration,omitempty"`
	OsDiskImage                  *OSDiskImage                  `json:"osDiskImage"`
	Plan                         *PurchasePlan                 `json:"plan"`
}
