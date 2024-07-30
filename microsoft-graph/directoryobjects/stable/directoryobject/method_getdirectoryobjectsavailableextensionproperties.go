package directoryobject

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

type GetDirectoryObjectsAvailableExtensionPropertiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ExtensionProperty
}

type GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ExtensionProperty
}

type GetDirectoryObjectsAvailableExtensionPropertiesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDirectoryObjectsAvailableExtensionPropertiesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDirectoryObjectsAvailableExtensionProperties ...
func (c DirectoryObjectClient) GetDirectoryObjectsAvailableExtensionProperties(ctx context.Context, input GetDirectoryObjectsAvailableExtensionPropertiesRequest) (result GetDirectoryObjectsAvailableExtensionPropertiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetDirectoryObjectsAvailableExtensionPropertiesCustomPager{},
		Path:       "/directoryObjects/getAvailableExtensionProperties",
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

// GetDirectoryObjectsAvailableExtensionPropertiesComplete retrieves all the results into a single object
func (c DirectoryObjectClient) GetDirectoryObjectsAvailableExtensionPropertiesComplete(ctx context.Context, input GetDirectoryObjectsAvailableExtensionPropertiesRequest) (GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult, error) {
	return c.GetDirectoryObjectsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx, input, ExtensionPropertyOperationPredicate{})
}

// GetDirectoryObjectsAvailableExtensionPropertiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryObjectClient) GetDirectoryObjectsAvailableExtensionPropertiesCompleteMatchingPredicate(ctx context.Context, input GetDirectoryObjectsAvailableExtensionPropertiesRequest, predicate ExtensionPropertyOperationPredicate) (result GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult, err error) {
	items := make([]stable.ExtensionProperty, 0)

	resp, err := c.GetDirectoryObjectsAvailableExtensionProperties(ctx, input)
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

	result = GetDirectoryObjectsAvailableExtensionPropertiesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
