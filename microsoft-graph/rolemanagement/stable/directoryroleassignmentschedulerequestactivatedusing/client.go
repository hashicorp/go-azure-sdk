package directoryroleassignmentschedulerequestactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRequestActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRequestActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRequestActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedulerequestactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRequestActivatedUsingClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRequestActivatedUsingClient{
		Client: client,
	}, nil
}
