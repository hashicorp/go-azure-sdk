package usereventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &UserEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
