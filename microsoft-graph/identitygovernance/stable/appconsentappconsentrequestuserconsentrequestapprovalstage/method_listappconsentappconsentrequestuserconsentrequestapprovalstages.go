package appconsentappconsentrequestuserconsentrequestapprovalstage

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

type ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ApprovalStage
}

type ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ApprovalStage
}

type ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentAppConsentRequestUserConsentRequestApprovalStages ...
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStageClient) ListAppConsentAppConsentRequestUserConsentRequestApprovalStages(ctx context.Context, id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId) (result ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCustomPager{},
		Path:       fmt.Sprintf("%s/approval/stages", id.ID()),
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
		Values *[]stable.ApprovalStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesComplete retrieves all the results into a single object
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStageClient) ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesComplete(ctx context.Context, id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId) (ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCompleteResult, error) {
	return c.ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCompleteMatchingPredicate(ctx, id, ApprovalStageOperationPredicate{})
}

// ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentAppConsentRequestUserConsentRequestApprovalStageClient) ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceAppConsentAppConsentRequestIdUserConsentRequestId, predicate ApprovalStageOperationPredicate) (result ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCompleteResult, err error) {
	items := make([]stable.ApprovalStage, 0)

	resp, err := c.ListAppConsentAppConsentRequestUserConsentRequestApprovalStages(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListAppConsentAppConsentRequestUserConsentRequestApprovalStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
