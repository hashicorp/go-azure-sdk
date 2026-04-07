package raitoollabels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiToolLabelPropertiesProjectScopesItem struct {
	LabelValues map[string]string `json:"labelValues"`
	Project     string            `json:"project"`
}
