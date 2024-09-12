package directoryroleeligibilityscheduleroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleRoleDefinitionClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleRoleDefinitionClient{
		Client: client,
	}, nil
}
