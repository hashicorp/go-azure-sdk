package groupeventexceptionoccurrenceinstanceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventExceptionOccurrenceInstanceAttachmentClient struct {
	Client *msgraph.Client
}

func NewGroupEventExceptionOccurrenceInstanceAttachmentClientWithBaseURI(api sdkEnv.Api) (*GroupEventExceptionOccurrenceInstanceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventexceptionoccurrenceinstanceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventExceptionOccurrenceInstanceAttachmentClient: %+v", err)
	}

	return &GroupEventExceptionOccurrenceInstanceAttachmentClient{
		Client: client,
	}, nil
}
