// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GitLabIdentityProviderApplyConfiguration represents an declarative configuration of the GitLabIdentityProvider type for use
// with apply.
type GitLabIdentityProviderApplyConfiguration struct {
	ClientID     *string                                   `json:"clientID,omitempty"`
	ClientSecret *SecretNameReferenceApplyConfiguration    `json:"clientSecret,omitempty"`
	URL          *string                                   `json:"url,omitempty"`
	CA           *ConfigMapNameReferenceApplyConfiguration `json:"ca,omitempty"`
}

// GitLabIdentityProviderApplyConfiguration constructs an declarative configuration of the GitLabIdentityProvider type for use with
// apply.
func GitLabIdentityProvider() *GitLabIdentityProviderApplyConfiguration {
	return &GitLabIdentityProviderApplyConfiguration{}
}

// WithClientID sets the ClientID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientID field is set to the value of the last call.
func (b *GitLabIdentityProviderApplyConfiguration) WithClientID(value string) *GitLabIdentityProviderApplyConfiguration {
	b.ClientID = &value
	return b
}

// WithClientSecret sets the ClientSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientSecret field is set to the value of the last call.
func (b *GitLabIdentityProviderApplyConfiguration) WithClientSecret(value *SecretNameReferenceApplyConfiguration) *GitLabIdentityProviderApplyConfiguration {
	b.ClientSecret = value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *GitLabIdentityProviderApplyConfiguration) WithURL(value string) *GitLabIdentityProviderApplyConfiguration {
	b.URL = &value
	return b
}

// WithCA sets the CA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CA field is set to the value of the last call.
func (b *GitLabIdentityProviderApplyConfiguration) WithCA(value *ConfigMapNameReferenceApplyConfiguration) *GitLabIdentityProviderApplyConfiguration {
	b.CA = value
	return b
}