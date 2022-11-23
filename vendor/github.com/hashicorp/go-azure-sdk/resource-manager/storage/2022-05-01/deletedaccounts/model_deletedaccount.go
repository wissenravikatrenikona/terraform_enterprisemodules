package deletedaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedAccount struct {
	Id         *string                   `json:"id,omitempty"`
	Name       *string                   `json:"name,omitempty"`
	Properties *DeletedAccountProperties `json:"properties"`
	Type       *string                   `json:"type,omitempty"`
}
