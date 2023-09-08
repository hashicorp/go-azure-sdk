package meeventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewMeEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*MeEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &MeEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
