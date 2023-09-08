package memailfoldermessagerule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderMessageRuleClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderMessageRuleClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderMessageRuleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfoldermessagerule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderMessageRuleClient: %+v", err)
	}

	return &MeMailFolderMessageRuleClient{
		Client: client,
	}, nil
}
