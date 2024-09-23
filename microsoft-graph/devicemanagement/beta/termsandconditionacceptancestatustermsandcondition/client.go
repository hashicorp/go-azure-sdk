package termsandconditionacceptancestatustermsandcondition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermsAndConditionAcceptanceStatusTermsAndConditionClient struct {
	Client *msgraph.Client
}

func NewTermsAndConditionAcceptanceStatusTermsAndConditionClientWithBaseURI(sdkApi sdkEnv.Api) (*TermsAndConditionAcceptanceStatusTermsAndConditionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "termsandconditionacceptancestatustermsandcondition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TermsAndConditionAcceptanceStatusTermsAndConditionClient: %+v", err)
	}

	return &TermsAndConditionAcceptanceStatusTermsAndConditionClient{
		Client: client,
	}, nil
}
