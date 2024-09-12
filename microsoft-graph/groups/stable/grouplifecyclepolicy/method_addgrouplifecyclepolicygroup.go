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

type AddGroupLifecyclePolicyGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *AddGroupLifecyclePolicyGroupResult
}

// AddGroupLifecyclePolicyGroup - Invoke action addGroup. Adds specific groups to a lifecycle policy. This action limits
// the group lifecycle policy to a set of groups only if the managedGroupTypes property of groupLifecyclePolicy is set
// to Selected.
func (c GroupLifecyclePolicyClient) AddGroupLifecyclePolicyGroup(ctx context.Context, id stable.GroupIdGroupLifecyclePolicyId, input AddGroupLifecyclePolicyGroupRequest) (result AddGroupLifecyclePolicyGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/addGroup", id.ID()),
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

	var model AddGroupLifecyclePolicyGroupResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
