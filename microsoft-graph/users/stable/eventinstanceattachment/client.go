package eventinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewEventInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceAttachmentClient: %+v", err)
	}

	return &EventInstanceAttachmentClient{
		Client: client,
	}, nil
}
