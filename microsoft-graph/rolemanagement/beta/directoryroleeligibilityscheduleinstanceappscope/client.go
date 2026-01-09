package directoryroleeligibilityscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleInstanceAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleInstanceAppScopeClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
