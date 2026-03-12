package csmpublishingcredentialspoliciesentities

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

type WebAppsListBasicPublishingCredentialsPoliciesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CsmPublishingCredentialsPoliciesEntity
}

type WebAppsListBasicPublishingCredentialsPoliciesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CsmPublishingCredentialsPoliciesEntity
}

type WebAppsListBasicPublishingCredentialsPoliciesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListBasicPublishingCredentialsPoliciesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListBasicPublishingCredentialsPolicies ...
func (c CsmPublishingCredentialsPoliciesEntitiesClient) WebAppsListBasicPublishingCredentialsPolicies(ctx context.Context, id commonids.AppServiceId) (result WebAppsListBasicPublishingCredentialsPoliciesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListBasicPublishingCredentialsPoliciesCustomPager{},
		Path:       fmt.Sprintf("%s/basicPublishingCredentialsPolicies", id.ID()),
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
		Values *[]CsmPublishingCredentialsPoliciesEntity `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListBasicPublishingCredentialsPoliciesComplete retrieves all the results into a single object
func (c CsmPublishingCredentialsPoliciesEntitiesClient) WebAppsListBasicPublishingCredentialsPoliciesComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListBasicPublishingCredentialsPoliciesCompleteResult, error) {
	return c.WebAppsListBasicPublishingCredentialsPoliciesCompleteMatchingPredicate(ctx, id, CsmPublishingCredentialsPoliciesEntityOperationPredicate{})
}

// WebAppsListBasicPublishingCredentialsPoliciesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CsmPublishingCredentialsPoliciesEntitiesClient) WebAppsListBasicPublishingCredentialsPoliciesCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate CsmPublishingCredentialsPoliciesEntityOperationPredicate) (result WebAppsListBasicPublishingCredentialsPoliciesCompleteResult, err error) {
	items := make([]CsmPublishingCredentialsPoliciesEntity, 0)

	resp, err := c.WebAppsListBasicPublishingCredentialsPolicies(ctx, id)
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

	result = WebAppsListBasicPublishingCredentialsPoliciesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
