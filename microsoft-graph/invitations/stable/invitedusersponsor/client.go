package invitedusersponsor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedUserSponsorClient struct {
	Client *msgraph.Client
}

func NewInvitedUserSponsorClientWithBaseURI(sdkApi sdkEnv.Api) (*InvitedUserSponsorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "invitedusersponsor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitedUserSponsorClient: %+v", err)
	}

	return &InvitedUserSponsorClient{
		Client: client,
	}, nil
}
