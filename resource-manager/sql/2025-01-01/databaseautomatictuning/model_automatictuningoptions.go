package databaseautomatictuning

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticTuningOptions struct {
	ActualState  *AutomaticTuningOptionModeActual  `json:"actualState,omitempty"`
	DesiredState *AutomaticTuningOptionModeDesired `json:"desiredState,omitempty"`
	ReasonCode   *int64                            `json:"reasonCode,omitempty"`
	ReasonDesc   *AutomaticTuningDisabledReason    `json:"reasonDesc,omitempty"`
}
