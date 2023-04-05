package application

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandLineSetting string

const (
	CommandLineSettingAllow      CommandLineSetting = "Allow"
	CommandLineSettingDoNotAllow CommandLineSetting = "DoNotAllow"
	CommandLineSettingRequire    CommandLineSetting = "Require"
)

func PossibleValuesForCommandLineSetting() []string {
	return []string{
		string(CommandLineSettingAllow),
		string(CommandLineSettingDoNotAllow),
		string(CommandLineSettingRequire),
	}
}

type RemoteApplicationType string

const (
	RemoteApplicationTypeInBuilt         RemoteApplicationType = "InBuilt"
	RemoteApplicationTypeMsixApplication RemoteApplicationType = "MsixApplication"
)

func PossibleValuesForRemoteApplicationType() []string {
	return []string{
		string(RemoteApplicationTypeInBuilt),
		string(RemoteApplicationTypeMsixApplication),
	}
}
