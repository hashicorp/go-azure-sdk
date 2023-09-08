package groupeventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &GroupEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
