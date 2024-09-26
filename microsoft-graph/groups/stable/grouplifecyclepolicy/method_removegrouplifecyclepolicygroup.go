package grouplifecyclepolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveGroupLifecyclePolicyGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RemoveGroupLifecyclePolicyGroupResult
}

type RemoveGroupLifecyclePolicyGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveGroupLifecyclePolicyGroupOperationOptions() RemoveGroupLifecyclePolicyGroupOperationOptions {
	return RemoveGroupLifecyclePolicyGroupOperationOptions{}
}

func (o RemoveGroupLifecyclePolicyGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveGroupLifecyclePolicyGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveGroupLifecyclePolicyGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveGroupLifecyclePolicyGroup - Invoke action removeGroup. Removes a group from a lifecycle policy.
func (c GroupLifecyclePolicyClient) RemoveGroupLifecyclePolicyGroup(ctx context.Context, id stable.GroupIdGroupLifecyclePolicyId, input RemoveGroupLifecyclePolicyGroupRequest, options RemoveGroupLifecyclePolicyGroupOperationOptions) (result RemoveGroupLifecyclePolicyGroupOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/removeGroup", id.ID()),
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

	var model RemoveGroupLifecyclePolicyGroupResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
