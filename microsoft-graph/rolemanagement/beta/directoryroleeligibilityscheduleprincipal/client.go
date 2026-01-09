package directoryroleeligibilityscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleEligibilitySchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleEligibilitySchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleEligibilitySchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleeligibilityscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleEligibilitySchedulePrincipalClient: %+v", err)
	}

	return &DirectoryRoleEligibilitySchedulePrincipalClient{
		Client: client,
	}, nil
}
