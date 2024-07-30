package intentuserstate

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

type ListIntentUserStatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementIntentUserState
}

type ListIntentUserStatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementIntentUserState
}

type ListIntentUserStatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentUserStatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntentUserStates ...
func (c IntentUserStateClient) ListIntentUserStates(ctx context.Context, id DeviceManagementIntentId) (result ListIntentUserStatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntentUserStatesCustomPager{},
		Path:       fmt.Sprintf("%s/userStates", id.ID()),
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
		Values *[]beta.DeviceManagementIntentUserState `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntentUserStatesComplete retrieves all the results into a single object
func (c IntentUserStateClient) ListIntentUserStatesComplete(ctx context.Context, id DeviceManagementIntentId) (ListIntentUserStatesCompleteResult, error) {
	return c.ListIntentUserStatesCompleteMatchingPredicate(ctx, id, DeviceManagementIntentUserStateOperationPredicate{})
}

// ListIntentUserStatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentUserStateClient) ListIntentUserStatesCompleteMatchingPredicate(ctx context.Context, id DeviceManagementIntentId, predicate DeviceManagementIntentUserStateOperationPredicate) (result ListIntentUserStatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementIntentUserState, 0)

	resp, err := c.ListIntentUserStates(ctx, id)
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

	result = ListIntentUserStatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
