package groupconversationthread

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupConversationThreadClient struct {
	Client *msgraph.Client
}

func NewGroupConversationThreadClientWithBaseURI(api sdkEnv.Api) (*GroupConversationThreadClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupconversationthread", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupConversationThreadClient: %+v", err)
	}

	return &GroupConversationThreadClient{
		Client: client,
	}, nil
}
