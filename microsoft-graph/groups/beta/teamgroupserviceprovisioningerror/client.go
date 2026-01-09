package teamgroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewTeamGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamgroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &TeamGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
