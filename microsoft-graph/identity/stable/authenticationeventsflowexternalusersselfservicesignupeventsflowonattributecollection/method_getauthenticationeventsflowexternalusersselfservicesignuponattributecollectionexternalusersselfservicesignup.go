package authenticationeventsflowexternalusersselfservicesignupeventsflowonattributecollection

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

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnAttributeCollectionExternalUsersSelfServiceSignUp
}

type GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions() GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions {
	return GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions{}
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUp - Get
// the item of type microsoft.graph.onAttributeCollectionHandler as
// microsoft.graph.onAttributeCollectionExternalUsersSelfServiceSignUp
func (c AuthenticationEventsFlowExternalUsersSelfServiceSignUpEventsFlowOnAttributeCollectionClient) GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUp(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, options GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationOptions) (result GetAuthenticationEventsFlowExternalUsersSelfServiceSignUpOnAttributeCollectionExternalUsersSelfServiceSignUpOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/onAttributeCollectionExternalUsersSelfServiceSignUp", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model stable.OnAttributeCollectionExternalUsersSelfServiceSignUp
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
