// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resource

// NestedResourceOptions is a bag of optional settings that control a resource's behavior.
type NestedResourceOptions struct {
	// Ignore changes to any of the specified properties.
	IgnoreChanges []string `pulumi:"ignoreChanges"`
	// When provided with a resource ID, import indicates that this resource's provider should import its state from the cloud resource with the given ID. The inputs to the resource's constructor must align with the resource's current state. Once a resource has been imported, the import property must be removed from the resource's options.
	Import *string `pulumi:"import"`
	// When set to true, protect ensures this resource cannot be deleted.
	Protect *bool `pulumi:"protect"`
	// Changes to any of these property paths will force a replacement.  If this list includes `"*"`, changes to any properties will force a replacement.  Initialization errors from previous deployments will require replacement instead of update only if `"*"` is passed.
	ReplaceOnChanges []string `pulumi:"replaceOnChanges"`
}

func init() {
}
