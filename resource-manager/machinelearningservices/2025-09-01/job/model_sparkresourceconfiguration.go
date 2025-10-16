package job

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkResourceConfiguration struct {
	InstanceType   *string `json:"instanceType,omitempty"`
	RuntimeVersion *string `json:"runtimeVersion,omitempty"`
}
