package usereventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewUserEventExceptionOccurrenceInstanceClientWithBaseURI(api sdkEnv.Api) (*UserEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &UserEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
