// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi-awsx/sdk/go/awsx/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Role and policy attachments with default setup unless explicitly skipped or an existing role ARN provided.
type DefaultRoleWithPolicy struct {
	// Args to use when creating the role and policies. Can't be specified if `roleArn` is used.
	Args *RoleWithPolicy `pulumi:"args"`
	// Resource options to use for the role. Can't be specified if `roleArn` is used.
	Opts *resource.NestedResourceOptions `pulumi:"opts"`
	// ARN of existing role to use instead of creating a new role. Cannot be used in combination with `args` or `opts`.
	RoleArn *string `pulumi:"roleArn"`
	// Skips creation of the role if set to `true`.
	Skip *bool `pulumi:"skip"`
}

// DefaultRoleWithPolicyInput is an input type that accepts DefaultRoleWithPolicyArgs and DefaultRoleWithPolicyOutput values.
// You can construct a concrete instance of `DefaultRoleWithPolicyInput` via:
//
//          DefaultRoleWithPolicyArgs{...}
type DefaultRoleWithPolicyInput interface {
	pulumi.Input

	ToDefaultRoleWithPolicyOutput() DefaultRoleWithPolicyOutput
	ToDefaultRoleWithPolicyOutputWithContext(context.Context) DefaultRoleWithPolicyOutput
}

// Role and policy attachments with default setup unless explicitly skipped or an existing role ARN provided.
type DefaultRoleWithPolicyArgs struct {
	// Args to use when creating the role and policies. Can't be specified if `roleArn` is used.
	Args *RoleWithPolicyArgs `pulumi:"args"`
	// Resource options to use for the role. Can't be specified if `roleArn` is used.
	Opts *resource.NestedResourceOptionsArgs `pulumi:"opts"`
	// ARN of existing role to use instead of creating a new role. Cannot be used in combination with `args` or `opts`.
	RoleArn pulumi.StringPtrInput `pulumi:"roleArn"`
	// Skips creation of the role if set to `true`.
	Skip *bool `pulumi:"skip"`
}

func (DefaultRoleWithPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultRoleWithPolicy)(nil)).Elem()
}

func (i DefaultRoleWithPolicyArgs) ToDefaultRoleWithPolicyOutput() DefaultRoleWithPolicyOutput {
	return i.ToDefaultRoleWithPolicyOutputWithContext(context.Background())
}

func (i DefaultRoleWithPolicyArgs) ToDefaultRoleWithPolicyOutputWithContext(ctx context.Context) DefaultRoleWithPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRoleWithPolicyOutput)
}

func (i DefaultRoleWithPolicyArgs) ToDefaultRoleWithPolicyPtrOutput() DefaultRoleWithPolicyPtrOutput {
	return i.ToDefaultRoleWithPolicyPtrOutputWithContext(context.Background())
}

func (i DefaultRoleWithPolicyArgs) ToDefaultRoleWithPolicyPtrOutputWithContext(ctx context.Context) DefaultRoleWithPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRoleWithPolicyOutput).ToDefaultRoleWithPolicyPtrOutputWithContext(ctx)
}

// DefaultRoleWithPolicyPtrInput is an input type that accepts DefaultRoleWithPolicyArgs, DefaultRoleWithPolicyPtr and DefaultRoleWithPolicyPtrOutput values.
// You can construct a concrete instance of `DefaultRoleWithPolicyPtrInput` via:
//
//          DefaultRoleWithPolicyArgs{...}
//
//  or:
//
//          nil
type DefaultRoleWithPolicyPtrInput interface {
	pulumi.Input

	ToDefaultRoleWithPolicyPtrOutput() DefaultRoleWithPolicyPtrOutput
	ToDefaultRoleWithPolicyPtrOutputWithContext(context.Context) DefaultRoleWithPolicyPtrOutput
}

type defaultRoleWithPolicyPtrType DefaultRoleWithPolicyArgs

func DefaultRoleWithPolicyPtr(v *DefaultRoleWithPolicyArgs) DefaultRoleWithPolicyPtrInput {
	return (*defaultRoleWithPolicyPtrType)(v)
}

func (*defaultRoleWithPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRoleWithPolicy)(nil)).Elem()
}

func (i *defaultRoleWithPolicyPtrType) ToDefaultRoleWithPolicyPtrOutput() DefaultRoleWithPolicyPtrOutput {
	return i.ToDefaultRoleWithPolicyPtrOutputWithContext(context.Background())
}

func (i *defaultRoleWithPolicyPtrType) ToDefaultRoleWithPolicyPtrOutputWithContext(ctx context.Context) DefaultRoleWithPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultRoleWithPolicyPtrOutput)
}

// Role and policy attachments with default setup unless explicitly skipped or an existing role ARN provided.
type DefaultRoleWithPolicyOutput struct{ *pulumi.OutputState }

func (DefaultRoleWithPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultRoleWithPolicy)(nil)).Elem()
}

func (o DefaultRoleWithPolicyOutput) ToDefaultRoleWithPolicyOutput() DefaultRoleWithPolicyOutput {
	return o
}

func (o DefaultRoleWithPolicyOutput) ToDefaultRoleWithPolicyOutputWithContext(ctx context.Context) DefaultRoleWithPolicyOutput {
	return o
}

func (o DefaultRoleWithPolicyOutput) ToDefaultRoleWithPolicyPtrOutput() DefaultRoleWithPolicyPtrOutput {
	return o.ToDefaultRoleWithPolicyPtrOutputWithContext(context.Background())
}

func (o DefaultRoleWithPolicyOutput) ToDefaultRoleWithPolicyPtrOutputWithContext(ctx context.Context) DefaultRoleWithPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultRoleWithPolicy) *DefaultRoleWithPolicy {
		return &v
	}).(DefaultRoleWithPolicyPtrOutput)
}

// Args to use when creating the role and policies. Can't be specified if `roleArn` is used.
func (o DefaultRoleWithPolicyOutput) Args() RoleWithPolicyPtrOutput {
	return o.ApplyT(func(v DefaultRoleWithPolicy) *RoleWithPolicy { return v.Args }).(RoleWithPolicyPtrOutput)
}

// Resource options to use for the role. Can't be specified if `roleArn` is used.
func (o DefaultRoleWithPolicyOutput) Opts() resource.NestedResourceOptionsPtrOutput {
	return o.ApplyT(func(v DefaultRoleWithPolicy) *resource.NestedResourceOptions { return v.Opts }).(resource.NestedResourceOptionsPtrOutput)
}

// ARN of existing role to use instead of creating a new role. Cannot be used in combination with `args` or `opts`.
func (o DefaultRoleWithPolicyOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DefaultRoleWithPolicy) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// Skips creation of the role if set to `true`.
func (o DefaultRoleWithPolicyOutput) Skip() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DefaultRoleWithPolicy) *bool { return v.Skip }).(pulumi.BoolPtrOutput)
}

type DefaultRoleWithPolicyPtrOutput struct{ *pulumi.OutputState }

func (DefaultRoleWithPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultRoleWithPolicy)(nil)).Elem()
}

func (o DefaultRoleWithPolicyPtrOutput) ToDefaultRoleWithPolicyPtrOutput() DefaultRoleWithPolicyPtrOutput {
	return o
}

func (o DefaultRoleWithPolicyPtrOutput) ToDefaultRoleWithPolicyPtrOutputWithContext(ctx context.Context) DefaultRoleWithPolicyPtrOutput {
	return o
}

func (o DefaultRoleWithPolicyPtrOutput) Elem() DefaultRoleWithPolicyOutput {
	return o.ApplyT(func(v *DefaultRoleWithPolicy) DefaultRoleWithPolicy {
		if v != nil {
			return *v
		}
		var ret DefaultRoleWithPolicy
		return ret
	}).(DefaultRoleWithPolicyOutput)
}

// Args to use when creating the role and policies. Can't be specified if `roleArn` is used.
func (o DefaultRoleWithPolicyPtrOutput) Args() RoleWithPolicyPtrOutput {
	return o.ApplyT(func(v *DefaultRoleWithPolicy) *RoleWithPolicy {
		if v == nil {
			return nil
		}
		return v.Args
	}).(RoleWithPolicyPtrOutput)
}

// Resource options to use for the role. Can't be specified if `roleArn` is used.
func (o DefaultRoleWithPolicyPtrOutput) Opts() resource.NestedResourceOptionsPtrOutput {
	return o.ApplyT(func(v *DefaultRoleWithPolicy) *resource.NestedResourceOptions {
		if v == nil {
			return nil
		}
		return v.Opts
	}).(resource.NestedResourceOptionsPtrOutput)
}

// ARN of existing role to use instead of creating a new role. Cannot be used in combination with `args` or `opts`.
func (o DefaultRoleWithPolicyPtrOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DefaultRoleWithPolicy) *string {
		if v == nil {
			return nil
		}
		return v.RoleArn
	}).(pulumi.StringPtrOutput)
}

// Skips creation of the role if set to `true`.
func (o DefaultRoleWithPolicyPtrOutput) Skip() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DefaultRoleWithPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Skip
	}).(pulumi.BoolPtrOutput)
}

// The set of arguments for constructing a Role resource and Policy attachments.
type RoleWithPolicy struct {
	// Description of the role.
	Description *string `pulumi:"description"`
	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies *bool `pulumi:"forceDetachPolicies"`
	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, this provider will not manage any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies added out of band on `apply`.
	InlinePolicies []iam.RoleInlinePolicy `pulumi:"inlinePolicies"`
	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, this provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
	ManagedPolicyArns []string `pulumi:"managedPolicyArns"`
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// Name of the role policy.
	Name *string `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path *string `pulumi:"path"`
	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary *string `pulumi:"permissionsBoundary"`
	// ARNs of the policies to attach to the created role.
	PolicyArns []string `pulumi:"policyArns"`
	// Key-value mapping of tags for the IAM role. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// RoleWithPolicyInput is an input type that accepts RoleWithPolicyArgs and RoleWithPolicyOutput values.
// You can construct a concrete instance of `RoleWithPolicyInput` via:
//
//          RoleWithPolicyArgs{...}
type RoleWithPolicyInput interface {
	pulumi.Input

	ToRoleWithPolicyOutput() RoleWithPolicyOutput
	ToRoleWithPolicyOutputWithContext(context.Context) RoleWithPolicyOutput
}

// The set of arguments for constructing a Role resource and Policy attachments.
type RoleWithPolicyArgs struct {
	// Description of the role.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
	ForceDetachPolicies pulumi.BoolPtrInput `pulumi:"forceDetachPolicies"`
	// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, this provider will not manage any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies added out of band on `apply`.
	InlinePolicies iam.RoleInlinePolicyArrayInput `pulumi:"inlinePolicies"`
	// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, this provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
	ManagedPolicyArns pulumi.StringArrayInput `pulumi:"managedPolicyArns"`
	// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
	MaxSessionDuration pulumi.IntPtrInput `pulumi:"maxSessionDuration"`
	// Name of the role policy.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
	// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path pulumi.StringPtrInput `pulumi:"path"`
	// ARN of the policy that is used to set the permissions boundary for the role.
	PermissionsBoundary pulumi.StringPtrInput `pulumi:"permissionsBoundary"`
	// ARNs of the policies to attach to the created role.
	PolicyArns pulumi.StringArrayInput `pulumi:"policyArns"`
	// Key-value mapping of tags for the IAM role. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (RoleWithPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleWithPolicy)(nil)).Elem()
}

func (i RoleWithPolicyArgs) ToRoleWithPolicyOutput() RoleWithPolicyOutput {
	return i.ToRoleWithPolicyOutputWithContext(context.Background())
}

func (i RoleWithPolicyArgs) ToRoleWithPolicyOutputWithContext(ctx context.Context) RoleWithPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleWithPolicyOutput)
}

func (i RoleWithPolicyArgs) ToRoleWithPolicyPtrOutput() RoleWithPolicyPtrOutput {
	return i.ToRoleWithPolicyPtrOutputWithContext(context.Background())
}

func (i RoleWithPolicyArgs) ToRoleWithPolicyPtrOutputWithContext(ctx context.Context) RoleWithPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleWithPolicyOutput).ToRoleWithPolicyPtrOutputWithContext(ctx)
}

// RoleWithPolicyPtrInput is an input type that accepts RoleWithPolicyArgs, RoleWithPolicyPtr and RoleWithPolicyPtrOutput values.
// You can construct a concrete instance of `RoleWithPolicyPtrInput` via:
//
//          RoleWithPolicyArgs{...}
//
//  or:
//
//          nil
type RoleWithPolicyPtrInput interface {
	pulumi.Input

	ToRoleWithPolicyPtrOutput() RoleWithPolicyPtrOutput
	ToRoleWithPolicyPtrOutputWithContext(context.Context) RoleWithPolicyPtrOutput
}

type roleWithPolicyPtrType RoleWithPolicyArgs

func RoleWithPolicyPtr(v *RoleWithPolicyArgs) RoleWithPolicyPtrInput {
	return (*roleWithPolicyPtrType)(v)
}

func (*roleWithPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleWithPolicy)(nil)).Elem()
}

func (i *roleWithPolicyPtrType) ToRoleWithPolicyPtrOutput() RoleWithPolicyPtrOutput {
	return i.ToRoleWithPolicyPtrOutputWithContext(context.Background())
}

func (i *roleWithPolicyPtrType) ToRoleWithPolicyPtrOutputWithContext(ctx context.Context) RoleWithPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleWithPolicyPtrOutput)
}

// The set of arguments for constructing a Role resource and Policy attachments.
type RoleWithPolicyOutput struct{ *pulumi.OutputState }

func (RoleWithPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleWithPolicy)(nil)).Elem()
}

func (o RoleWithPolicyOutput) ToRoleWithPolicyOutput() RoleWithPolicyOutput {
	return o
}

func (o RoleWithPolicyOutput) ToRoleWithPolicyOutputWithContext(ctx context.Context) RoleWithPolicyOutput {
	return o
}

func (o RoleWithPolicyOutput) ToRoleWithPolicyPtrOutput() RoleWithPolicyPtrOutput {
	return o.ToRoleWithPolicyPtrOutputWithContext(context.Background())
}

func (o RoleWithPolicyOutput) ToRoleWithPolicyPtrOutputWithContext(ctx context.Context) RoleWithPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoleWithPolicy) *RoleWithPolicy {
		return &v
	}).(RoleWithPolicyPtrOutput)
}

// Description of the role.
func (o RoleWithPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleWithPolicy) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
func (o RoleWithPolicyOutput) ForceDetachPolicies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RoleWithPolicy) *bool { return v.ForceDetachPolicies }).(pulumi.BoolPtrOutput)
}

// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, this provider will not manage any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies added out of band on `apply`.
func (o RoleWithPolicyOutput) InlinePolicies() iam.RoleInlinePolicyArrayOutput {
	return o.ApplyT(func(v RoleWithPolicy) []iam.RoleInlinePolicy { return v.InlinePolicies }).(iam.RoleInlinePolicyArrayOutput)
}

// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, this provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
func (o RoleWithPolicyOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoleWithPolicy) []string { return v.ManagedPolicyArns }).(pulumi.StringArrayOutput)
}

// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
func (o RoleWithPolicyOutput) MaxSessionDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoleWithPolicy) *int { return v.MaxSessionDuration }).(pulumi.IntPtrOutput)
}

// Name of the role policy.
func (o RoleWithPolicyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleWithPolicy) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
func (o RoleWithPolicyOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleWithPolicy) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
func (o RoleWithPolicyOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleWithPolicy) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// ARN of the policy that is used to set the permissions boundary for the role.
func (o RoleWithPolicyOutput) PermissionsBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleWithPolicy) *string { return v.PermissionsBoundary }).(pulumi.StringPtrOutput)
}

// ARNs of the policies to attach to the created role.
func (o RoleWithPolicyOutput) PolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoleWithPolicy) []string { return v.PolicyArns }).(pulumi.StringArrayOutput)
}

// Key-value mapping of tags for the IAM role. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o RoleWithPolicyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RoleWithPolicy) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type RoleWithPolicyPtrOutput struct{ *pulumi.OutputState }

func (RoleWithPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleWithPolicy)(nil)).Elem()
}

func (o RoleWithPolicyPtrOutput) ToRoleWithPolicyPtrOutput() RoleWithPolicyPtrOutput {
	return o
}

func (o RoleWithPolicyPtrOutput) ToRoleWithPolicyPtrOutputWithContext(ctx context.Context) RoleWithPolicyPtrOutput {
	return o
}

func (o RoleWithPolicyPtrOutput) Elem() RoleWithPolicyOutput {
	return o.ApplyT(func(v *RoleWithPolicy) RoleWithPolicy {
		if v != nil {
			return *v
		}
		var ret RoleWithPolicy
		return ret
	}).(RoleWithPolicyOutput)
}

// Description of the role.
func (o RoleWithPolicyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleWithPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
func (o RoleWithPolicyPtrOutput) ForceDetachPolicies() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RoleWithPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.ForceDetachPolicies
	}).(pulumi.BoolPtrOutput)
}

// Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, this provider will not manage any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies added out of band on `apply`.
func (o RoleWithPolicyPtrOutput) InlinePolicies() iam.RoleInlinePolicyArrayOutput {
	return o.ApplyT(func(v *RoleWithPolicy) []iam.RoleInlinePolicy {
		if v == nil {
			return nil
		}
		return v.InlinePolicies
	}).(iam.RoleInlinePolicyArrayOutput)
}

// Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, this provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
func (o RoleWithPolicyPtrOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RoleWithPolicy) []string {
		if v == nil {
			return nil
		}
		return v.ManagedPolicyArns
	}).(pulumi.StringArrayOutput)
}

// Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
func (o RoleWithPolicyPtrOutput) MaxSessionDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RoleWithPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxSessionDuration
	}).(pulumi.IntPtrOutput)
}

// Name of the role policy.
func (o RoleWithPolicyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleWithPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

// Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
func (o RoleWithPolicyPtrOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleWithPolicy) *string {
		if v == nil {
			return nil
		}
		return v.NamePrefix
	}).(pulumi.StringPtrOutput)
}

// Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
func (o RoleWithPolicyPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleWithPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

// ARN of the policy that is used to set the permissions boundary for the role.
func (o RoleWithPolicyPtrOutput) PermissionsBoundary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RoleWithPolicy) *string {
		if v == nil {
			return nil
		}
		return v.PermissionsBoundary
	}).(pulumi.StringPtrOutput)
}

// ARNs of the policies to attach to the created role.
func (o RoleWithPolicyPtrOutput) PolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RoleWithPolicy) []string {
		if v == nil {
			return nil
		}
		return v.PolicyArns
	}).(pulumi.StringArrayOutput)
}

// Key-value mapping of tags for the IAM role. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
func (o RoleWithPolicyPtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *RoleWithPolicy) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultRoleWithPolicyInput)(nil)).Elem(), DefaultRoleWithPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultRoleWithPolicyPtrInput)(nil)).Elem(), DefaultRoleWithPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleWithPolicyInput)(nil)).Elem(), RoleWithPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleWithPolicyPtrInput)(nil)).Elem(), RoleWithPolicyArgs{})
	pulumi.RegisterOutputType(DefaultRoleWithPolicyOutput{})
	pulumi.RegisterOutputType(DefaultRoleWithPolicyPtrOutput{})
	pulumi.RegisterOutputType(RoleWithPolicyOutput{})
	pulumi.RegisterOutputType(RoleWithPolicyPtrOutput{})
}