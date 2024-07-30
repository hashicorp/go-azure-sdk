package intentcategorysettingdefinition

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

type ListIntentCategorySettingDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingDefinition
}

type ListIntentCategorySettingDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingDefinition
}

type ListIntentCategorySettingDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntentCategorySettingDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntentCategorySettingDefinitions ...
func (c IntentCategorySettingDefinitionClient) ListIntentCategorySettingDefinitions(ctx context.Context, id DeviceManagementIntentIdCategoryId) (result ListIntentCategorySettingDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntentCategorySettingDefinitionsCustomPager{},
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

// ListIntentCategorySettingDefinitionsComplete retrieves all the results into a single object
func (c IntentCategorySettingDefinitionClient) ListIntentCategorySettingDefinitionsComplete(ctx context.Context, id DeviceManagementIntentIdCategoryId) (ListIntentCategorySettingDefinitionsCompleteResult, error) {
	return c.ListIntentCategorySettingDefinitionsCompleteMatchingPredicate(ctx, id, DeviceManagementSettingDefinitionOperationPredicate{})
}

// ListIntentCategorySettingDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntentCategorySettingDefinitionClient) ListIntentCategorySettingDefinitionsCompleteMatchingPredicate(ctx context.Context, id DeviceManagementIntentIdCategoryId, predicate DeviceManagementSettingDefinitionOperationPredicate) (result ListIntentCategorySettingDefinitionsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingDefinition, 0)

	resp, err := c.ListIntentCategorySettingDefinitions(ctx, id)
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

	result = ListIntentCategorySettingDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
