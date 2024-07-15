package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuntimeJavaJavaAgent struct {
	Enabled *bool                        `json:"enabled,omitempty"`
	Logging *RuntimeJavaJavaAgentLogging `json:"logging,omitempty"`
}
