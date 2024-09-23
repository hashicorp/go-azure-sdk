package directoryroleeligibilityschedulerequestappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleRequestAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleRequestAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleRequestAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityschedulerequestappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleRequestAppScopeClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleRequestAppScopeClient{
		Client: client,
	}, nil
}
