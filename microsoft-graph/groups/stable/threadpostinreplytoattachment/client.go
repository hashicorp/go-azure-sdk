package threadpostinreplytoattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadPostInReplyToAttachmentClient struct {
	Client *msgraph.Client
}

func NewThreadPostInReplyToAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadPostInReplyToAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "threadpostinreplytoattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadPostInReplyToAttachmentClient: %+v", err)
	}

	return &ThreadPostInReplyToAttachmentClient{
		Client: client,
	}, nil
}
