package usereventexceptionoccurrenceattachment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExceptionOccurrenceAttachmentClient struct {
	Client *msgraph.Client
}

func NewUserEventExceptionOccurrenceAttachmentClientWithBaseURI(api sdkEnv.Api) (*UserEventExceptionOccurrenceAttachmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventexceptionoccurrenceattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExceptionOccurrenceAttachmentClient: %+v", err)
	}

	return &UserEventExceptionOccurrenceAttachmentClient{
		Client: client,
	}, nil
}
