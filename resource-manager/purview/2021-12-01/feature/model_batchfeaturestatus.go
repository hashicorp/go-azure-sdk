package feature

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchFeatureStatus struct {
	Features *map[string]bool `json:"features,omitempty"`
}
