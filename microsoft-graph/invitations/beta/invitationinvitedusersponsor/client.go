package invitationinvitedusersponsor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationInvitedUserSponsorClient struct {
	Client *msgraph.Client
}

func NewInvitationInvitedUserSponsorClientWithBaseURI(api sdkEnv.Api) (*InvitationInvitedUserSponsorClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "invitationinvitedusersponsor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitationInvitedUserSponsorClient: %+v", err)
	}

	return &InvitationInvitedUserSponsorClient{
		Client: client,
	}, nil
}
