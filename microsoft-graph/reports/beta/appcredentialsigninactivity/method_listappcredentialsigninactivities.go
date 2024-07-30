package appcredentialsigninactivity

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

type ListAppCredentialSignInActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AppCredentialSignInActivity
}

type ListAppCredentialSignInActivitiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AppCredentialSignInActivity
}

type ListAppCredentialSignInActivitiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAppCredentialSignInActivitiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAppCredentialSignInActivities ...
func (c AppCredentialSignInActivityClient) ListAppCredentialSignInActivities(ctx context.Context) (result ListAppCredentialSignInActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAppCredentialSignInActivitiesCustomPager{},
		Path:       "/reports/appCredentialSignInActivities",
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
		Values *[]beta.AppCredentialSignInActivity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAppCredentialSignInActivitiesComplete retrieves all the results into a single object
func (c AppCredentialSignInActivityClient) ListAppCredentialSignInActivitiesComplete(ctx context.Context) (ListAppCredentialSignInActivitiesCompleteResult, error) {
	return c.ListAppCredentialSignInActivitiesCompleteMatchingPredicate(ctx, AppCredentialSignInActivityOperationPredicate{})
}

// ListAppCredentialSignInActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AppCredentialSignInActivityClient) ListAppCredentialSignInActivitiesCompleteMatchingPredicate(ctx context.Context, predicate AppCredentialSignInActivityOperationPredicate) (result ListAppCredentialSignInActivitiesCompleteResult, err error) {
	items := make([]beta.AppCredentialSignInActivity, 0)

	resp, err := c.ListAppCredentialSignInActivities(ctx)
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

	result = ListAppCredentialSignInActivitiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
