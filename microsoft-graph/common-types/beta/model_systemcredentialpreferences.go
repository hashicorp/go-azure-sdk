package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SystemCredentialPreferences struct {
	ExcludeTargets *[]ExcludeTarget                  `json:"excludeTargets,omitempty"`
	IncludeTargets *[]IncludeTarget                  `json:"includeTargets,omitempty"`
	ODataType      *string                           `json:"@odata.type,omitempty"`
	State          *SystemCredentialPreferencesState `json:"state,omitempty"`
}
