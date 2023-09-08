package invitation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationClient struct {
	Client *msgraph.Client
}

func NewInvitationClientWithBaseURI(api sdkEnv.Api) (*InvitationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "invitation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InvitationClient: %+v", err)
	}

	return &InvitationClient{
		Client: client,
	}, nil
}
