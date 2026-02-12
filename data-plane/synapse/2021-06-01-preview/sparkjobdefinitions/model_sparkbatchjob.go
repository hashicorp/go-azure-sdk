package sparkjobdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkBatchJob struct {
	AppId         *string                  `json:"appId,omitempty"`
	AppInfo       *map[string]string       `json:"appInfo,omitempty"`
	ArtifactId    *string                  `json:"artifactId,omitempty"`
	ErrorInfo     *[]SparkServiceError     `json:"errorInfo,omitempty"`
	Id            int64                    `json:"id"`
	JobType       *SparkJobType            `json:"jobType,omitempty"`
	LivyInfo      *SparkBatchJobState      `json:"livyInfo,omitempty"`
	Log           *[]string                `json:"log,omitempty"`
	Name          *string                  `json:"name,omitempty"`
	PluginInfo    *SparkServicePlugin      `json:"pluginInfo,omitempty"`
	Result        *SparkBatchJobResultType `json:"result,omitempty"`
	SchedulerInfo *SparkScheduler          `json:"schedulerInfo,omitempty"`
	SparkPoolName *string                  `json:"sparkPoolName,omitempty"`
	State         *string                  `json:"state,omitempty"`
	SubmitterId   *string                  `json:"submitterId,omitempty"`
	SubmitterName *string                  `json:"submitterName,omitempty"`
	Tags          *map[string]string       `json:"tags,omitempty"`
	WorkspaceName *string                  `json:"workspaceName,omitempty"`
}
