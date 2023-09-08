package usermailfoldermessagerule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderMessageRuleClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderMessageRuleClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderMessageRuleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfoldermessagerule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderMessageRuleClient: %+v", err)
	}

	return &UserMailFolderMessageRuleClient{
		Client: client,
	}, nil
}
