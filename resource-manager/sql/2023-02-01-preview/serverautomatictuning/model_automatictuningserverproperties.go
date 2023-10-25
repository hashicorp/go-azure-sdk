package serverautomatictuning

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticTuningServerProperties struct {
	ActualState  *AutomaticTuningServerMode               `json:"actualState,omitempty"`
	DesiredState *AutomaticTuningServerMode               `json:"desiredState,omitempty"`
	Options      *map[string]AutomaticTuningServerOptions `json:"options,omitempty"`
}
