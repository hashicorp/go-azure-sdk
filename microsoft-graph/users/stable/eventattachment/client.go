package eventattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventAttachmentClient struct {
	Client *msgraph.Client
}

func NewEventAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EventAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventAttachmentClient: %+v", err)
	}

	return &EventAttachmentClient{
		Client: client,
	}, nil
}
