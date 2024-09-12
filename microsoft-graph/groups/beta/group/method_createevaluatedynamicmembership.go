package group

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

type CreateEvaluateDynamicMembershipOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.EvaluateDynamicMembershipResult
}

// CreateEvaluateDynamicMembership - Invoke action evaluateDynamicMembership. Evaluate whether a user or device is or
// would be a member of a dynamic group. The membership rule is returned along with other details that were used in the
// evaluation. You can complete this operation in the following ways:
func (c GroupClient) CreateEvaluateDynamicMembership(ctx context.Context, id beta.GroupId, input CreateEvaluateDynamicMembershipRequest) (result CreateEvaluateDynamicMembershipOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/evaluateDynamicMembership", id.ID()),
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

	var model beta.EvaluateDynamicMembershipResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
