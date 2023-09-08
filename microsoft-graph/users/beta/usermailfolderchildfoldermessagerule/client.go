package usermailfolderchildfoldermessagerule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderChildFolderMessageRuleClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderChildFolderMessageRuleClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderChildFolderMessageRuleClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderchildfoldermessagerule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderChildFolderMessageRuleClient: %+v", err)
	}

	return &UserMailFolderChildFolderMessageRuleClient{
		Client: client,
	}, nil
}
