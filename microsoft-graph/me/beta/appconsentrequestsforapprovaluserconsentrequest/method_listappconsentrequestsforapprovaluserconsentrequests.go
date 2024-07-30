package appconsentrequestsforapprovaluserconsentrequest

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

type ListAppConsentRequestsForApprovalUserConsentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UserConsentRequest
}

type ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UserConsentRequest
}

type ListAppConsentRequestsForApprovalUserConsentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentRequestsForApprovalUserConsentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentRequestsForApprovalUserConsentRequests ...
func (c AppConsentRequestsForApprovalUserConsentRequestClient) ListAppConsentRequestsForApprovalUserConsentRequests(ctx context.Context, id MeAppConsentRequestsForApprovalId) (result ListAppConsentRequestsForApprovalUserConsentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppConsentRequestsForApprovalUserConsentRequestsCustomPager{},
		Path:       fmt.Sprintf("%s/userConsentRequests", id.ID()),
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
		Values *[]beta.UserConsentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentRequestsForApprovalUserConsentRequestsComplete retrieves all the results into a single object
func (c AppConsentRequestsForApprovalUserConsentRequestClient) ListAppConsentRequestsForApprovalUserConsentRequestsComplete(ctx context.Context, id MeAppConsentRequestsForApprovalId) (ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult, error) {
	return c.ListAppConsentRequestsForApprovalUserConsentRequestsCompleteMatchingPredicate(ctx, id, UserConsentRequestOperationPredicate{})
}

// ListAppConsentRequestsForApprovalUserConsentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentRequestsForApprovalUserConsentRequestClient) ListAppConsentRequestsForApprovalUserConsentRequestsCompleteMatchingPredicate(ctx context.Context, id MeAppConsentRequestsForApprovalId, predicate UserConsentRequestOperationPredicate) (result ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult, err error) {
	items := make([]beta.UserConsentRequest, 0)

	resp, err := c.ListAppConsentRequestsForApprovalUserConsentRequests(ctx, id)
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

	result = ListAppConsentRequestsForApprovalUserConsentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
