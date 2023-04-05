package machines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentModeTypes string

const (
	AssessmentModeTypesAutomaticByPlatform AssessmentModeTypes = "AutomaticByPlatform"
	AssessmentModeTypesImageDefault        AssessmentModeTypes = "ImageDefault"
)

func PossibleValuesForAssessmentModeTypes() []string {
	return []string{
		string(AssessmentModeTypesAutomaticByPlatform),
		string(AssessmentModeTypesImageDefault),
	}
}

type InstanceViewTypes string

const (
	InstanceViewTypesInstanceView InstanceViewTypes = "instanceView"
)

func PossibleValuesForInstanceViewTypes() []string {
	return []string{
		string(InstanceViewTypesInstanceView),
	}
}

type PatchModeTypes string

const (
	PatchModeTypesAutomaticByOS       PatchModeTypes = "AutomaticByOS"
	PatchModeTypesAutomaticByPlatform PatchModeTypes = "AutomaticByPlatform"
	PatchModeTypesImageDefault        PatchModeTypes = "ImageDefault"
	PatchModeTypesManual              PatchModeTypes = "Manual"
)

func PossibleValuesForPatchModeTypes() []string {
	return []string{
		string(PatchModeTypesAutomaticByOS),
		string(PatchModeTypesAutomaticByPlatform),
		string(PatchModeTypesImageDefault),
		string(PatchModeTypesManual),
	}
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}

type StatusTypes string

const (
	StatusTypesConnected    StatusTypes = "Connected"
	StatusTypesDisconnected StatusTypes = "Disconnected"
	StatusTypesError        StatusTypes = "Error"
)

func PossibleValuesForStatusTypes() []string {
	return []string{
		string(StatusTypesConnected),
		string(StatusTypesDisconnected),
		string(StatusTypesError),
	}
}
