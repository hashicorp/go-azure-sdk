package eventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &EventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
