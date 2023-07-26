package cognitiveservicescommitmentplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentPlansListPlansByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CommitmentPlan
}

type CommitmentPlansListPlansByResourceGroupCompleteResult struct {
	Items []CommitmentPlan
}

// CommitmentPlansListPlansByResourceGroup ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result CommitmentPlansListPlansByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.CognitiveServices/commitmentPlans", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]CommitmentPlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// CommitmentPlansListPlansByResourceGroupComplete retrieves all the results into a single object
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (CommitmentPlansListPlansByResourceGroupCompleteResult, error) {
	return c.CommitmentPlansListPlansByResourceGroupCompleteMatchingPredicate(ctx, id, CommitmentPlanOperationPredicate{})
}

// CommitmentPlansListPlansByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansListPlansByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate CommitmentPlanOperationPredicate) (result CommitmentPlansListPlansByResourceGroupCompleteResult, err error) {
	items := make([]CommitmentPlan, 0)

	resp, err := c.CommitmentPlansListPlansByResourceGroup(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = CommitmentPlansListPlansByResourceGroupCompleteResult{
		Items: items,
	}
	return
}
