package userpendingaccessreviewinstancedecisioninstancedefinition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *models.AccessReviewScheduleDefinition
}

// GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceDefinition ...
func (c UserPendingAccessReviewInstanceDecisionInstanceDefinitionClient) GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceDefinition(ctx context.Context, id UserPendingAccessReviewInstanceDecisionId) (result GetUserByIdPendingAccessReviewInstanceByIdDecisionByIdInstanceDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/instance/definition", id.ID()),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
