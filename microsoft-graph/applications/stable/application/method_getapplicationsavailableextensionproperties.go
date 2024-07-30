package application

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

type GetApplicationsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExtensionProperty
}

type GetApplicationsAvailableExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExtensionProperty
}

type GetApplicationsAvailableExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetApplicationsAvailableExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetApplicationsAvailableExtensionProperties ...
func (c ApplicationClient) GetApplicationsAvailableExtensionProperties(ctx context.Context, input GetApplicationsAvailableExtensionPropertiesRequest) (result GetApplicationsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetApplicationsAvailableExtensionPropertiesCustomPager{},
		Path:       "/applications/getAvailableExtensionProperties",
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
		Values *[]stable.ExtensionProperty `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetApplicationsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c ApplicationClient) GetApplicationsAvailableExtensionPropertiesComplete(ctx context.Context, input GetApplicationsAvailableExtensionPropertiesRequest) (GetApplicationsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetApplicationsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, input, ExtensionPropertyOperationPredicate{})
}

// GetApplicationsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ApplicationClient) GetApplicationsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, input GetApplicationsAvailableExtensionPropertiesRequest, predicate ExtensionPropertyOperationPredicate) (result GetApplicationsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]stable.ExtensionProperty, 0)

	resp, err := c.GetApplicationsAvailableExtensionProperties(ctx, input)
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

	result = GetApplicationsAvailableExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
