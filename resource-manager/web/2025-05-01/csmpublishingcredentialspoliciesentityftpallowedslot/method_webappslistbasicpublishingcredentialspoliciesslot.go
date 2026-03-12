package csmpublishingcredentialspoliciesentityftpallowedslot

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListBasicPublishingCredentialsPoliciesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CsmPublishingCredentialsPoliciesEntity
}

type WebAppsListBasicPublishingCredentialsPoliciesSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CsmPublishingCredentialsPoliciesEntity
}

type WebAppsListBasicPublishingCredentialsPoliciesSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListBasicPublishingCredentialsPoliciesSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListBasicPublishingCredentialsPoliciesSlot ...
func (c CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient) WebAppsListBasicPublishingCredentialsPoliciesSlot(ctx context.Context, id SlotId) (result WebAppsListBasicPublishingCredentialsPoliciesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListBasicPublishingCredentialsPoliciesSlotCustomPager{},
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

// WebAppsListBasicPublishingCredentialsPoliciesSlotComplete retrieves all the results into a single object
func (c CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient) WebAppsListBasicPublishingCredentialsPoliciesSlotComplete(ctx context.Context, id SlotId) (WebAppsListBasicPublishingCredentialsPoliciesSlotCompleteResult, error) {
	return c.WebAppsListBasicPublishingCredentialsPoliciesSlotCompleteMatchingPredicate(ctx, id, CsmPublishingCredentialsPoliciesEntityOperationPredicate{})
}

// WebAppsListBasicPublishingCredentialsPoliciesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CsmPublishingCredentialsPoliciesEntityFtpAllowedSlotClient) WebAppsListBasicPublishingCredentialsPoliciesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate CsmPublishingCredentialsPoliciesEntityOperationPredicate) (result WebAppsListBasicPublishingCredentialsPoliciesSlotCompleteResult, err error) {
	items := make([]CsmPublishingCredentialsPoliciesEntity, 0)

	resp, err := c.WebAppsListBasicPublishingCredentialsPoliciesSlot(ctx, id)
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

	result = WebAppsListBasicPublishingCredentialsPoliciesSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
