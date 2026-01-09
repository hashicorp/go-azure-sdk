package mailfolderchildfoldermessageattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderChildFolderMessageAttachmentClient struct {
	Client *msgraph.Client
}

func NewMailFolderChildFolderMessageAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderChildFolderMessageAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfolderchildfoldermessageattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderChildFolderMessageAttachmentClient: %+v", err)
	}

	return &MailFolderChildFolderMessageAttachmentClient{
		Client: client,
	}, nil
}
