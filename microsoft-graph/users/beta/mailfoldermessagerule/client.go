package mailfoldermessagerule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderMessageRuleClient struct {
	Client *msgraph.Client
}

func NewMailFolderMessageRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderMessageRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfoldermessagerule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderMessageRuleClient: %+v", err)
	}

	return &MailFolderMessageRuleClient{
		Client: client,
	}, nil
}
