package trafficfilter

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticTrafficFilterResponse struct {
	Rulesets *[]ElasticTrafficFilter `json:"rulesets,omitempty"`
}
