// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type RequestOptions struct {
	ContentType         string
	ExpectedStatusCodes []int
	HttpMethod          string
	OptionsObject       Options
	Pager               odata.Pager
	Path                string
}

func (ro RequestOptions) Validate() error {
	if ro.ContentType == "" {
		return fmt.Errorf("missing `ContentType`")
	}
	if len(ro.ExpectedStatusCodes) == 0 {
		return fmt.Errorf("missing `ExpectedStatusCodes`")
	}
	if ro.HttpMethod == "" {
		return fmt.Errorf("missing `HttpMethod`")
	}
	if ro.Path == "" {
		return fmt.Errorf("missing `Path`")
	}
	return nil
}
