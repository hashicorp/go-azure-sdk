package homerealmdiscoverypolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateHomeRealmDiscoveryPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HomeRealmDiscoveryPolicy
}

type CreateHomeRealmDiscoveryPolicyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateHomeRealmDiscoveryPolicyOperationOptions() CreateHomeRealmDiscoveryPolicyOperationOptions {
	return CreateHomeRealmDiscoveryPolicyOperationOptions{}
}

func (o CreateHomeRealmDiscoveryPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateHomeRealmDiscoveryPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateHomeRealmDiscoveryPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateHomeRealmDiscoveryPolicy - Create homeRealmDiscoveryPolicy. Create a new homeRealmDiscoveryPolicy object.
func (c HomeRealmDiscoveryPolicyClient) CreateHomeRealmDiscoveryPolicy(ctx context.Context, input beta.HomeRealmDiscoveryPolicy, options CreateHomeRealmDiscoveryPolicyOperationOptions) (result CreateHomeRealmDiscoveryPolicyOperationResponse, err error) {
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
		Path:          "/policies/homeRealmDiscoveryPolicies",
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

	var model beta.HomeRealmDiscoveryPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
