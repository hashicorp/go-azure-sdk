package teamownerserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamOwnerServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewTeamOwnerServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamOwnerServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamownerserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamOwnerServiceProvisioningErrorClient: %+v", err)
	}

	return &TeamOwnerServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
