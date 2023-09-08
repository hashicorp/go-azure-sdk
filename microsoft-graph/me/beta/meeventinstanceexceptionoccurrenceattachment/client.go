package meeventinstanceexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventInstanceExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeEventInstanceExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeEventInstanceExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventinstanceexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventInstanceExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &MeEventInstanceExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
