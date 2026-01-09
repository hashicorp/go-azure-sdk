package authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollectiononattributecollectionexternalusersselfservicesignupattribute

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions() AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions {
	return AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions{}
}

func (o AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRef - Add attribute
// (to user flow). Add an attribute to an external identities self-service sign up user flow that's represented by an
// externalUsersSelfServiceSignupEventsFlow object. You can add both custom and built-in attributes to a user flow. The
// attribute is added to both the attributeCollection> attributes and attributeCollection> attributeCollectionPage >
// views collections on the user flow. In the views collection, the attribute is assigned the default settings. You can
// PATCH the user flow to customize the settings of the attribute on the views object, for example, marking it as
// required or updating the allowed input types.
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionOnAttributeCollectionExternalUsersSelfServiceSignUpAttributeClient) AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRef(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, input stable.ReferenceCreate, options AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationOptions) (result AddAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalAttributeRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/$ref", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
