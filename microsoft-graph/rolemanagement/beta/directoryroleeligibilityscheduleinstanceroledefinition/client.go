package directoryroleeligibilityscheduleinstanceroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleInstanceRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleinstanceroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleInstanceRoleDefinitionClient{
		Client: client,
	}, nil
}
