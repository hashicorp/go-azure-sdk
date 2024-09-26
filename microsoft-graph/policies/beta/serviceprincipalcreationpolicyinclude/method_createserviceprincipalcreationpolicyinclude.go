package serviceprincipalcreationpolicyinclude

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateServicePrincipalCreationPolicyIncludeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ServicePrincipalCreationConditionSet
}

type CreateServicePrincipalCreationPolicyIncludeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateServicePrincipalCreationPolicyIncludeOperationOptions() CreateServicePrincipalCreationPolicyIncludeOperationOptions {
	return CreateServicePrincipalCreationPolicyIncludeOperationOptions{}
}

func (o CreateServicePrincipalCreationPolicyIncludeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateServicePrincipalCreationPolicyIncludeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateServicePrincipalCreationPolicyIncludeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateServicePrincipalCreationPolicyInclude - Create new navigation property to includes for policies
func (c ServicePrincipalCreationPolicyIncludeClient) CreateServicePrincipalCreationPolicyInclude(ctx context.Context, id beta.PolicyServicePrincipalCreationPolicyId, input beta.ServicePrincipalCreationConditionSet, options CreateServicePrincipalCreationPolicyIncludeOperationOptions) (result CreateServicePrincipalCreationPolicyIncludeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/includes", id.ID()),
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

	var model beta.ServicePrincipalCreationConditionSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
