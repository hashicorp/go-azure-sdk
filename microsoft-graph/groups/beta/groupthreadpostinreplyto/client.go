package groupthreadpostinreplyto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostInReplyToClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostInReplyToClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostInReplyToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpostinreplyto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostInReplyToClient: %+v", err)
	}

	return &GroupThreadPostInReplyToClient{
		Client: client,
	}, nil
}
