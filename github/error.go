package github

import (
// "encoding/json"
)

type Error struct {
	Message          string `json:"message"`
	InvalidResources []InvalidResource
}

// GitHub API Reference:
// Error Name      Description
// ----------      -----------
// missing         This means a resource does not exist.
// missing_field   This means a required field on a resource has not been set.
// invalid         This means the formatting of a field is invalid. The documentation for that resource should be able to give you more specific information.
// already_exists  This means another resource has the same value as this field. This can happen in resources that must have some unique key (such as Label names).

type InvalidResource struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}
