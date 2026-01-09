package directoryroleassignmentscheduleinstanceactivatedusing

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleInstanceActivatedUsingClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleInstanceActivatedUsingClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleInstanceActivatedUsingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleinstanceactivatedusing", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleInstanceActivatedUsingClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleInstanceActivatedUsingClient{
		Client: client,
	}, nil
}
