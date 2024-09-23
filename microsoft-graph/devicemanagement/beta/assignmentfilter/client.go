package assignmentfilter

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterClient struct {
	Client *msgraph.Client
}

func NewAssignmentFilterClientWithBaseURI(sdkApi sdkEnv.Api) (*AssignmentFilterClient, error) {
	client, err := msgraph.NewClient(sdkApi, "assignmentfilter", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssignmentFilterClient: %+v", err)
	}

	return &AssignmentFilterClient{
		Client: client,
	}, nil
}
