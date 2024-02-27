package grafanaplugin

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GrafanaAvailablePluginOperationPredicate struct {
	Name     *string
	PluginId *string
}

func (p GrafanaAvailablePluginOperationPredicate) Matches(input GrafanaAvailablePlugin) bool {

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.PluginId != nil && (input.PluginId == nil || *p.PluginId != *input.PluginId) {
		return false
	}

	return true
}
