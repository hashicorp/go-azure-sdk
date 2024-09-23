package directoryroleeligibilityscheduledirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduledirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleDirectoryScopeClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleDirectoryScopeClient{
		Client: client,
	}, nil
}
