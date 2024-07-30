package mobileappintentandstate

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

type ListMobileAppIntentAndStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.MobileAppIntentAndState
}

type ListMobileAppIntentAndStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.MobileAppIntentAndState
}

type ListMobileAppIntentAndStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListMobileAppIntentAndStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListMobileAppIntentAndStates ...
func (c MobileAppIntentAndStateClient) ListMobileAppIntentAndStates(ctx context.Context, id UserId) (result ListMobileAppIntentAndStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListMobileAppIntentAndStatesCustomPager{},
		Path:       fmt.Sprintf("%s/mobileAppIntentAndStates", id.ID()),
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
		Values *[]beta.MobileAppIntentAndState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMobileAppIntentAndStatesComplete retrieves all the results into a single object
func (c MobileAppIntentAndStateClient) ListMobileAppIntentAndStatesComplete(ctx context.Context, id UserId) (ListMobileAppIntentAndStatesCompleteResult, error) {
	return c.ListMobileAppIntentAndStatesCompleteMatchingPredicate(ctx, id, MobileAppIntentAndStateOperationPredicate{})
}

// ListMobileAppIntentAndStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MobileAppIntentAndStateClient) ListMobileAppIntentAndStatesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate MobileAppIntentAndStateOperationPredicate) (result ListMobileAppIntentAndStatesCompleteResult, err error) {
	items := make([]beta.MobileAppIntentAndState, 0)

	resp, err := c.ListMobileAppIntentAndStates(ctx, id)
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

	result = ListMobileAppIntentAndStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
