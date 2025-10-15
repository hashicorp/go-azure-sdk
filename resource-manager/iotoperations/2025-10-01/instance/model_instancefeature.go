package instance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstanceFeature struct {
	Mode     *InstanceFeatureMode        `json:"mode,omitempty"`
	Settings *map[string]OperationalMode `json:"settings,omitempty"`
}
