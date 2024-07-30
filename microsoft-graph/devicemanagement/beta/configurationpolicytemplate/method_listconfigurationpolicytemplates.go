package configurationpolicytemplate

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

type ListConfigurationPolicyTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DeviceManagementConfigurationPolicyTemplate
}

type ListConfigurationPolicyTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DeviceManagementConfigurationPolicyTemplate
}

type ListConfigurationPolicyTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConfigurationPolicyTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConfigurationPolicyTemplates ...
func (c ConfigurationPolicyTemplateClient) ListConfigurationPolicyTemplates(ctx context.Context) (result ListConfigurationPolicyTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConfigurationPolicyTemplatesCustomPager{},
		Path:       "/deviceManagement/configurationPolicyTemplates",
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
		Values *[]beta.DeviceManagementConfigurationPolicyTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConfigurationPolicyTemplatesComplete retrieves all the results into a single object
func (c ConfigurationPolicyTemplateClient) ListConfigurationPolicyTemplatesComplete(ctx context.Context) (ListConfigurationPolicyTemplatesCompleteResult, error) {
	return c.ListConfigurationPolicyTemplatesCompleteMatchingPredicate(ctx, DeviceManagementConfigurationPolicyTemplateOperationPredicate{})
}

// ListConfigurationPolicyTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConfigurationPolicyTemplateClient) ListConfigurationPolicyTemplatesCompleteMatchingPredicate(ctx context.Context, predicate DeviceManagementConfigurationPolicyTemplateOperationPredicate) (result ListConfigurationPolicyTemplatesCompleteResult, err error) {
	items := make([]beta.DeviceManagementConfigurationPolicyTemplate, 0)

	resp, err := c.ListConfigurationPolicyTemplates(ctx)
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

	result = ListConfigurationPolicyTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
