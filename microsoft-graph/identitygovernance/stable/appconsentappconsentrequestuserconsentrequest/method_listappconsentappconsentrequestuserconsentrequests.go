package appconsentappconsentrequestuserconsentrequest

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

type ListAppConsentAppConsentRequestUserConsentRequestsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserConsentRequest
}

type ListAppConsentAppConsentRequestUserConsentRequestsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserConsentRequest
}

type ListAppConsentAppConsentRequestUserConsentRequestsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppConsentAppConsentRequestUserConsentRequestsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppConsentAppConsentRequestUserConsentRequests ...
func (c AppConsentAppConsentRequestUserConsentRequestClient) ListAppConsentAppConsentRequestUserConsentRequests(ctx context.Context, id IdentityGovernanceAppConsentAppConsentRequestId) (result ListAppConsentAppConsentRequestUserConsentRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppConsentAppConsentRequestUserConsentRequestsCustomPager{},
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
		Values *[]stable.UserConsentRequest `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppConsentAppConsentRequestUserConsentRequestsComplete retrieves all the results into a single object
func (c AppConsentAppConsentRequestUserConsentRequestClient) ListAppConsentAppConsentRequestUserConsentRequestsComplete(ctx context.Context, id IdentityGovernanceAppConsentAppConsentRequestId) (ListAppConsentAppConsentRequestUserConsentRequestsCompleteResult, error) {
	return c.ListAppConsentAppConsentRequestUserConsentRequestsCompleteMatchingPredicate(ctx, id, UserConsentRequestOperationPredicate{})
}

// ListAppConsentAppConsentRequestUserConsentRequestsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppConsentAppConsentRequestUserConsentRequestClient) ListAppConsentAppConsentRequestUserConsentRequestsCompleteMatchingPredicate(ctx context.Context, id IdentityGovernanceAppConsentAppConsentRequestId, predicate UserConsentRequestOperationPredicate) (result ListAppConsentAppConsentRequestUserConsentRequestsCompleteResult, err error) {
	items := make([]stable.UserConsentRequest, 0)

	resp, err := c.ListAppConsentAppConsentRequestUserConsentRequests(ctx, id)
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

	result = ListAppConsentAppConsentRequestUserConsentRequestsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
