package mailfolderchildfoldermessagerule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderChildFolderMessageRuleClient struct {
	Client *msgraph.Client
}

func NewMailFolderChildFolderMessageRuleClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderChildFolderMessageRuleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderchildfoldermessagerule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderChildFolderMessageRuleClient: %+v", err)
	}

	return &MailFolderChildFolderMessageRuleClient{
		Client: client,
	}, nil
}
