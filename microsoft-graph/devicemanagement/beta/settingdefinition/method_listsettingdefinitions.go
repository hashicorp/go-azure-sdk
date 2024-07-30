package settingdefinition

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

type ListSettingDefinitionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementSettingDefinition
}

type ListSettingDefinitionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementSettingDefinition
}

type ListSettingDefinitionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSettingDefinitionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSettingDefinitions ...
func (c SettingDefinitionClient) ListSettingDefinitions(ctx context.Context) (result ListSettingDefinitionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSettingDefinitionsCustomPager{},
		Path:       "/deviceManagement/settingDefinitions",
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

// ListSettingDefinitionsComplete retrieves all the results into a single object
func (c SettingDefinitionClient) ListSettingDefinitionsComplete(ctx context.Context) (ListSettingDefinitionsCompleteResult, error) {
	return c.ListSettingDefinitionsCompleteMatchingPredicate(ctx, DeviceManagementSettingDefinitionOperationPredicate{})
}

// ListSettingDefinitionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SettingDefinitionClient) ListSettingDefinitionsCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementSettingDefinitionOperationPredicate) (result ListSettingDefinitionsCompleteResult, err error) {
	items := make([]beta.DeviceManagementSettingDefinition, 0)

	resp, err := c.ListSettingDefinitions(ctx)
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

	result = ListSettingDefinitionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
