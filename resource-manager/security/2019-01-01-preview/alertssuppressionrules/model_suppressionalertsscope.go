package alertssuppressionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SuppressionAlertsScope struct {
	AllOf []ScopeElement `json:"allOf"`
}
