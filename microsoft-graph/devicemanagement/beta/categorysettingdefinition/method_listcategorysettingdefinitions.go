package categorysettingdefinition

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

type ListCategorySettingDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingDefinition
}

type ListCategorySettingDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingDefinition
}

type ListCategorySettingDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCategorySettingDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCategorySettingDefinitions ...
func (c CategorySettingDefinitionClient) ListCategorySettingDefinitions(ctx context.Context, id DeviceManagementCategoryId) (result ListCategorySettingDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCategorySettingDefinitionsCustomPager{},
		Path:       fmt.Sprintf("%s/settingDefinitions", id.ID()),
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
		Values *[]beta.DeviceManagementSettingDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCategorySettingDefinitionsComplete retrieves all the results into a single object
func (c CategorySettingDefinitionClient) ListCategorySettingDefinitionsComplete(ctx context.Context, id DeviceManagementCategoryId) (ListCategorySettingDefinitionsCompleteResult, error) {
	return c.ListCategorySettingDefinitionsCompleteMatchingPredicate(ctx, id, DeviceManagementSettingDefinitionOperationPredicate{})
}

// ListCategorySettingDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CategorySettingDefinitionClient) ListCategorySettingDefinitionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementCategoryId, predicate DeviceManagementSettingDefinitionOperationPredicate) (result ListCategorySettingDefinitionsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingDefinition, 0)

	resp, err := c.ListCategorySettingDefinitions(ctx, id)
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

	result = ListCategorySettingDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
