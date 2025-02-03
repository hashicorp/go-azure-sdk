package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceNameAvailabilityRequest struct {
	Name          string                 `json:"name"`
	ResourceGroup string                 `json:"resourceGroup"`
	Type          CheckNameResourceTypes `json:"type"`
}
