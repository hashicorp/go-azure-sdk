package grouplifecyclepolicy

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

type RenewGroupLifecyclePoliciesGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RenewGroupLifecyclePoliciesGroupResult
}

type RenewGroupLifecyclePoliciesGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRenewGroupLifecyclePoliciesGroupOperationOptions() RenewGroupLifecyclePoliciesGroupOperationOptions {
	return RenewGroupLifecyclePoliciesGroupOperationOptions{}
}

func (o RenewGroupLifecyclePoliciesGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RenewGroupLifecyclePoliciesGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RenewGroupLifecyclePoliciesGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RenewGroupLifecyclePoliciesGroup - Invoke action renewGroup. Renew a group's expiration. When a group is renewed, the
// group expiration is extended by the number of days defined in the policy.
func (c GroupLifecyclePolicyClient) RenewGroupLifecyclePoliciesGroup(ctx context.Context, id beta.GroupId, input RenewGroupLifecyclePoliciesGroupRequest, options RenewGroupLifecyclePoliciesGroupOperationOptions) (result RenewGroupLifecyclePoliciesGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/groupLifecyclePolicies/renewGroup", id.ID()),
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

	var model RenewGroupLifecyclePoliciesGroupResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
