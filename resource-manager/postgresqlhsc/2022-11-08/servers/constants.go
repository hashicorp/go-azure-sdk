package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerRole string

const (
	ServerRoleCoordinator ServerRole = "Coordinator"
	ServerRoleWorker      ServerRole = "Worker"
)

func PossibleValuesForServerRole() []string {
	return []string{
		string(ServerRoleCoordinator),
		string(ServerRoleWorker),
	}
}
