package kqlscripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KqlScriptContent struct {
	CurrentConnection *KqlScriptContentCurrentConnection `json:"currentConnection,omitempty"`
	Metadata          *KqlScriptContentMetadata          `json:"metadata,omitempty"`
	Query             *string                            `json:"query,omitempty"`
}
