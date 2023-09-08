package invitationinviteduser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationInvitedUserClient struct {
	Client *msgraph.Client
}

func NewInvitationInvitedUserClientWithBaseURI(api sdkEnv.Api) (*InvitationInvitedUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "invitationinviteduser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitationInvitedUserClient: %+v", err)
	}

	return &InvitationInvitedUserClient{
		Client: client,
	}, nil
}
