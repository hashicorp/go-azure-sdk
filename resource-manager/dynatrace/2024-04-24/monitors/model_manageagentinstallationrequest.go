package monitors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManageAgentInstallationRequest struct {
	Action                      Action            `json:"action"`
	ManageAgentInstallationList []ManageAgentList `json:"manageAgentInstallationList"`
}
