package joinedteamgroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinedTeamGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewJoinedTeamGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*JoinedTeamGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "joinedteamgroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JoinedTeamGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &JoinedTeamGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
