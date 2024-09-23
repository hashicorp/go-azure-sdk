package eventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewEventExceptionOccurrenceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &EventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
