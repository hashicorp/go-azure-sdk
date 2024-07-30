package inviteduserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewInvitedUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*InvitedUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "inviteduserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitedUserServiceProvisioningErrorClient: %+v", err)
	}

	return &InvitedUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
