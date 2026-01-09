package mailfoldermessageattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailFolderMessageAttachmentClient struct {
	Client *msgraph.Client
}

func NewMailFolderMessageAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*MailFolderMessageAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mailfoldermessageattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MailFolderMessageAttachmentClient: %+v", err)
	}

	return &MailFolderMessageAttachmentClient{
		Client: client,
	}, nil
}
