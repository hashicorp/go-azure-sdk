package groupthreadpostinreplytoextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupThreadPostInReplyToExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupThreadPostInReplyToExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupThreadPostInReplyToExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupthreadpostinreplytoextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupThreadPostInReplyToExtensionClient: %+v", err)
	}

	return &GroupThreadPostInReplyToExtensionClient{
		Client: client,
	}, nil
}
