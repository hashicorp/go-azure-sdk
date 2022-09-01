package alertruletemplates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FusionAlertRuleTemplateProperties struct {
	Severity AlertSeverity   `json:"severity"`
	Tactics  *[]AttackTactic `json:"tactics,omitempty"`
}
