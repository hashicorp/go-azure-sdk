package directoryroleassignmentschedulerequesttargetschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRequestTargetScheduleClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRequestTargetScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRequestTargetScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedulerequesttargetschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRequestTargetScheduleClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRequestTargetScheduleClient{
		Client: client,
	}, nil
}
