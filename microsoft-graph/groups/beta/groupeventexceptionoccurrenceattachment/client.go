package groupeventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupEventExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &GroupEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
