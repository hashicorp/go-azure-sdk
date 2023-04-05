package configurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationSource string

const (
	ConfigurationSourceSystemNegativedefault ConfigurationSource = "system-default"
	ConfigurationSourceUserNegativeoverride  ConfigurationSource = "user-override"
)

func PossibleValuesForConfigurationSource() []string {
	return []string{
		string(ConfigurationSourceSystemNegativedefault),
		string(ConfigurationSourceUserNegativeoverride),
	}
}

type IsConfigPendingRestart string

const (
	IsConfigPendingRestartFalse IsConfigPendingRestart = "False"
	IsConfigPendingRestartTrue  IsConfigPendingRestart = "True"
)

func PossibleValuesForIsConfigPendingRestart() []string {
	return []string{
		string(IsConfigPendingRestartFalse),
		string(IsConfigPendingRestartTrue),
	}
}

type IsDynamicConfig string

const (
	IsDynamicConfigFalse IsDynamicConfig = "False"
	IsDynamicConfigTrue  IsDynamicConfig = "True"
)

func PossibleValuesForIsDynamicConfig() []string {
	return []string{
		string(IsDynamicConfigFalse),
		string(IsDynamicConfigTrue),
	}
}

type IsReadOnly string

const (
	IsReadOnlyFalse IsReadOnly = "False"
	IsReadOnlyTrue  IsReadOnly = "True"
)

func PossibleValuesForIsReadOnly() []string {
	return []string{
		string(IsReadOnlyFalse),
		string(IsReadOnlyTrue),
	}
}
