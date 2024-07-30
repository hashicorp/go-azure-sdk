package appconsentrequestsforapproval

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

type ListAppConsentRequestsForApprovalsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppConsentRequest
}

type ListAppConsentRequestsForApprovalsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppConsentRequest
}

type ListAppConsentRequestsForApprovalsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestsForApprovalsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestsForApprovals ...
func (c AppConsentRequestsForApprovalClient) ListAppConsentRequestsForApprovals(ctx context.Context, id UserId) (result ListAppConsentRequestsForApprovalsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppConsentRequestsForApprovalsCustomPager{},
		Path:       fmt.Sprintf("%s/appConsentRequestsForApproval", id.ID()),
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
		Values *[]beta.AppConsentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentRequestsForApprovalsComplete retrieves all the results into a single object
func (c AppConsentRequestsForApprovalClient) ListAppConsentRequestsForApprovalsComplete(ctx context.Context, id UserId) (ListAppConsentRequestsForApprovalsCompleteResult, error) {
	return c.ListAppConsentRequestsForApprovalsCompleteMatchingPredicate(ctx, id, AppConsentRequestOperationPredicate{})
}

// ListAppConsentRequestsForApprovalsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentRequestsForApprovalClient) ListAppConsentRequestsForApprovalsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate AppConsentRequestOperationPredicate) (result ListAppConsentRequestsForApprovalsCompleteResult, err error) {
	items := make([]beta.AppConsentRequest, 0)

	resp, err := c.ListAppConsentRequestsForApprovals(ctx, id)
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

	result = ListAppConsentRequestsForApprovalsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
