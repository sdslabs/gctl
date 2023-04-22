/*
 * Gasper Master API
 *
 * Handles authentication, creation/management of applications, databases, users and also provides a superuser API. <br><br> Only a superuser can avail the superuser API. A superuser can **grant/revoke** superuser privileges to other users. A default  superuser is created every time a Gasper instance is launched whose credentials are defined in the `admin` section of `config.toml`, the main configuration file. A sample configuration file is available [here](https://github.com/sdslabs/gasper/blob/develop/config.sample.toml#L37).<br><br> **Note:-** Normally the applications and databases can only be managed by their owners but the superuser can bypass that check.<br><br> **PS:-** If you want to programmatically generate a client for this API, you can find the corresponding OpenAPI specifications [here](https://github.com/sdslabs/gasper/tree/develop/docs/content/api/specs). We recommend using [OpenAPI-Generator](https://openapi-generator.tech/) for generating clients.
 *
 * API version: 1.0
 * Contact: contact@sdslabs.co.in
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	Success bool `json:"success,omitempty"`
	Data []Instances `json:"data,omitempty"`
}

// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	Success bool `json:"success,omitempty"`
}

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	Success bool `json:"success,omitempty"`
	Data []CreatedApplication `json:"data,omitempty"`
}

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	Success bool      `json:"success,omitempty"`
	Expire  time.Time `json:"expire,omitempty"`
}

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Success bool `json:"success,omitempty"`
	Data []string `json:"data,omitempty"`
}

// InlineResponse2006 struct for InlineResponse2006
type InlineResponse2006 struct {
	Success bool `json:"success,omitempty"`
	Data []Metrics `json:"data,omitempty"`
}

// InlineResponse2007 struct for InlineResponse2007
type InlineResponse2007 struct {
	Success bool `json:"success,omitempty"`
	Data []CreatedDatabase `json:"data,omitempty"`
}

// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	// Email id of Gasper Github user
	GitURL string `json:"giturl" bson:"giturl"`
}

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	// PAT for pushing code to repository
	PAT string `json:"pat" bson:"pat"`
	// Username of Gasper Github user
	Username string `json:"username" bson:"username"`
	// Email id of Gasper Github user
	Email string `json:"email" bson:"email"`
}
