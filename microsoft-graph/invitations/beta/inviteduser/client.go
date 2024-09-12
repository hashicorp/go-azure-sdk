package inviteduser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedUserClient struct {
	Client *msgraph.Client
}

func NewInvitedUserClientWithBaseURI(sdkApi sdkEnv.Api) (*InvitedUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "inviteduser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitedUserClient: %+v", err)
	}

	return &InvitedUserClient{
		Client: client,
	}, nil
}
