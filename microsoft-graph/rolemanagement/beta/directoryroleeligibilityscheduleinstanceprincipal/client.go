package directoryroleeligibilityscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilityScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilityScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilityScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilityScheduleInstancePrincipalClient: %+v", err)
	}

	return &DirectoryRoleEligibilityScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
