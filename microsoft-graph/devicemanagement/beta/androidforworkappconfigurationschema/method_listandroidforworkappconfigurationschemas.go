package androidforworkappconfigurationschema

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

type ListAndroidForWorkAppConfigurationSchemasOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AndroidForWorkAppConfigurationSchema
}

type ListAndroidForWorkAppConfigurationSchemasCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AndroidForWorkAppConfigurationSchema
}

type ListAndroidForWorkAppConfigurationSchemasCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAndroidForWorkAppConfigurationSchemasCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAndroidForWorkAppConfigurationSchemas ...
func (c AndroidForWorkAppConfigurationSchemaClient) ListAndroidForWorkAppConfigurationSchemas(ctx context.Context) (result ListAndroidForWorkAppConfigurationSchemasOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAndroidForWorkAppConfigurationSchemasCustomPager{},
		Path:       "/deviceManagement/androidForWorkAppConfigurationSchemas",
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
		Values *[]beta.AndroidForWorkAppConfigurationSchema `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAndroidForWorkAppConfigurationSchemasComplete retrieves all the results into a single object
func (c AndroidForWorkAppConfigurationSchemaClient) ListAndroidForWorkAppConfigurationSchemasComplete(ctx context.Context) (ListAndroidForWorkAppConfigurationSchemasCompleteResult, error) {
	return c.ListAndroidForWorkAppConfigurationSchemasCompleteMatchingPredicate(ctx, AndroidForWorkAppConfigurationSchemaOperationPredicate{})
}

// ListAndroidForWorkAppConfigurationSchemasCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AndroidForWorkAppConfigurationSchemaClient) ListAndroidForWorkAppConfigurationSchemasCompleteMatchingPredicate(ctx context.Context, predicate AndroidForWorkAppConfigurationSchemaOperationPredicate) (result ListAndroidForWorkAppConfigurationSchemasCompleteResult, err error) {
	items := make([]beta.AndroidForWorkAppConfigurationSchema, 0)

	resp, err := c.ListAndroidForWorkAppConfigurationSchemas(ctx)
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

	result = ListAndroidForWorkAppConfigurationSchemasCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
