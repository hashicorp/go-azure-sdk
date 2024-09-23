package directoryroleeligibilityschedulerequestdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleRequestDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleRequestDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityschedulerequestdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleRequestDirectoryScopeClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleRequestDirectoryScopeClient{
		Client: client,
	}, nil
}
