package directoryroleeligibilityscheduleappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleAppScopeClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleAppScopeClient{
		Client: client,
	}, nil
}
