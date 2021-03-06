// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/dcos/client-go/dcos/iam/client/groups"
	"github.com/dcos/client-go/dcos/iam/client/ldap"
	"github.com/dcos/client-go/dcos/iam/client/login"
	"github.com/dcos/client-go/dcos/iam/client/oidc"
	"github.com/dcos/client-go/dcos/iam/client/operations"
	"github.com/dcos/client-go/dcos/iam/client/permissions"
	"github.com/dcos/client-go/dcos/iam/client/saml"
	"github.com/dcos/client-go/dcos/iam/client/users"
)

// Default identity and access management HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/acs/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new identity and access management HTTP client.
func NewHTTPClient(formats strfmt.Registry) *IdentityAndAccessManagement {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new identity and access management HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *IdentityAndAccessManagement {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new identity and access management client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *IdentityAndAccessManagement {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(IdentityAndAccessManagement)
	cli.Transport = transport

	cli.Groups = groups.New(transport, formats)

	cli.Ldap = ldap.New(transport, formats)

	cli.Login = login.New(transport, formats)

	cli.Oidc = oidc.New(transport, formats)

	cli.Operations = operations.New(transport, formats)

	cli.Permissions = permissions.New(transport, formats)

	cli.Saml = saml.New(transport, formats)

	cli.Users = users.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// IdentityAndAccessManagement is a client for identity and access management
type IdentityAndAccessManagement struct {
	Groups *groups.Client

	Ldap *ldap.Client

	Login *login.Client

	Oidc *oidc.Client

	Operations *operations.Client

	Permissions *permissions.Client

	Saml *saml.Client

	Users *users.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *IdentityAndAccessManagement) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Groups.SetTransport(transport)

	c.Ldap.SetTransport(transport)

	c.Login.SetTransport(transport)

	c.Oidc.SetTransport(transport)

	c.Operations.SetTransport(transport)

	c.Permissions.SetTransport(transport)

	c.Saml.SetTransport(transport)

	c.Users.SetTransport(transport)

}
