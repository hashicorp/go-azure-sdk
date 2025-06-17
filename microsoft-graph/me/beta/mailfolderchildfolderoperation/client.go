package mailfolderchildfolderoperation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderChildFolderOperationClient struct {
	Client *msgraph.Client
}

func NewMailFolderChildFolderOperationClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderChildFolderOperationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderchildfolderoperation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderChildFolderOperationClient: %+v", err)
	}

	return &MailFolderChildFolderOperationClient{
		Client: client,
	}, nil
}
