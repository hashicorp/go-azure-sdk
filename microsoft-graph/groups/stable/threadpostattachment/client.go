package threadpostattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreadPostAttachmentClient struct {
	Client *msgraph.Client
}

func NewThreadPostAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ThreadPostAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "threadpostattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ThreadPostAttachmentClient: %+v", err)
	}

	return &ThreadPostAttachmentClient{
		Client: client,
	}, nil
}
