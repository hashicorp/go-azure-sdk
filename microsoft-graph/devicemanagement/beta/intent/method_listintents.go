package intent

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

type ListIntentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementIntent
}

type ListIntentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementIntent
}

type ListIntentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntents ...
func (c IntentClient) ListIntents(ctx context.Context) (result ListIntentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntentsCustomPager{},
		Path:       "/deviceManagement/intents",
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
		Values *[]beta.DeviceManagementIntent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntentsComplete retrieves all the results into a single object
func (c IntentClient) ListIntentsComplete(ctx context.Context) (ListIntentsCompleteResult, error) {
	return c.ListIntentsCompleteMatchingPredicate(ctx, DeviceManagementIntentOperationPredicate{})
}

// ListIntentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentClient) ListIntentsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementIntentOperationPredicate) (result ListIntentsCompleteResult, err error) {
	items := make([]beta.DeviceManagementIntent, 0)

	resp, err := c.ListIntents(ctx)
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

	result = ListIntentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
