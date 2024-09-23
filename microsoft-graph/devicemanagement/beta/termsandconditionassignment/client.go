package termsandconditionassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionAssignmentClient struct {
	Client *msgraph.Client
}

func NewTermsAndConditionAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsAndConditionAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsandconditionassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsAndConditionAssignmentClient: %+v", err)
	}

	return &TermsAndConditionAssignmentClient{
		Client: client,
	}, nil
}
