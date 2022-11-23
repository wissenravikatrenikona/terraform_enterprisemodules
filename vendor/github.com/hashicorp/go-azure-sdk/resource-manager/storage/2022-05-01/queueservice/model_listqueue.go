package queueservice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListQueue struct {
	Id         *string              `json:"id,omitempty"`
	Name       *string              `json:"name,omitempty"`
	Properties *ListQueueProperties `json:"properties"`
	Type       *string              `json:"type,omitempty"`
}
