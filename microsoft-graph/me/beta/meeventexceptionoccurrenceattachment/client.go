package meeventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeEventExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
