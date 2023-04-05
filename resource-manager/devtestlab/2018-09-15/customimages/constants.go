package customimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomImageOsType string

const (
	CustomImageOsTypeLinux   CustomImageOsType = "Linux"
	CustomImageOsTypeNone    CustomImageOsType = "None"
	CustomImageOsTypeWindows CustomImageOsType = "Windows"
)

func PossibleValuesForCustomImageOsType() []string {
	return []string{
		string(CustomImageOsTypeLinux),
		string(CustomImageOsTypeNone),
		string(CustomImageOsTypeWindows),
	}
}

type LinuxOsState string

const (
	LinuxOsStateDeprovisionApplied   LinuxOsState = "DeprovisionApplied"
	LinuxOsStateDeprovisionRequested LinuxOsState = "DeprovisionRequested"
	LinuxOsStateNonDeprovisioned     LinuxOsState = "NonDeprovisioned"
)

func PossibleValuesForLinuxOsState() []string {
	return []string{
		string(LinuxOsStateDeprovisionApplied),
		string(LinuxOsStateDeprovisionRequested),
		string(LinuxOsStateNonDeprovisioned),
	}
}

type StorageType string

const (
	StorageTypePremium     StorageType = "Premium"
	StorageTypeStandard    StorageType = "Standard"
	StorageTypeStandardSSD StorageType = "StandardSSD"
)

func PossibleValuesForStorageType() []string {
	return []string{
		string(StorageTypePremium),
		string(StorageTypeStandard),
		string(StorageTypeStandardSSD),
	}
}

type WindowsOsState string

const (
	WindowsOsStateNonSysprepped    WindowsOsState = "NonSysprepped"
	WindowsOsStateSysprepApplied   WindowsOsState = "SysprepApplied"
	WindowsOsStateSysprepRequested WindowsOsState = "SysprepRequested"
)

func PossibleValuesForWindowsOsState() []string {
	return []string{
		string(WindowsOsStateNonSysprepped),
		string(WindowsOsStateSysprepApplied),
		string(WindowsOsStateSysprepRequested),
	}
}
