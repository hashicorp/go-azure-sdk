package quota

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageAndQuotaDetails struct {
	ModelCollection *string               `json:"modelCollection,omitempty"`
	Quota           *int64                `json:"quota,omitempty"`
	UsageDetails    *[]PTUDeploymentUsage `json:"usageDetails,omitempty"`
}
