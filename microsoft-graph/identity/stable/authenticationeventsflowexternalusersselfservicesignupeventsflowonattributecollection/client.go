package authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient: %+v", err)
	}

	return &AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient{
		Client: client,
	}, nil
}
