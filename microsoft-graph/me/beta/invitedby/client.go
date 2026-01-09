package invitedby

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedByClient struct {
	Client *msgraph.Client
}

func NewInvitedByClientWithBaseURI(sdkApi sdkEnv.Api) (*InvitedByClient, error) {
	client, err := msgraph.NewClient(sdkApi, "invitedby", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitedByClient: %+v", err)
	}

	return &InvitedByClient{
		Client: client,
	}, nil
}
