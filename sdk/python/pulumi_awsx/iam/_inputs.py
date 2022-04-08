# coding=utf-8
# *** WARNING: this file was generated by pulumi-gen-awsx. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from .. import resource as _resource
import pulumi_aws

__all__ = [
    'DefaultRoleWithPolicyArgs',
    'RoleWithPolicyArgs',
]

@pulumi.input_type
class DefaultRoleWithPolicyArgs:
    def __init__(__self__, *,
                 args: Optional['RoleWithPolicyArgs'] = None,
                 opts: Optional['_resource.NestedResourceOptionsArgs'] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 skip: Optional[bool] = None):
        """
        Role and policy attachments with default setup unless explicitly skipped or an existing role ARN provided.
        :param 'RoleWithPolicyArgs' args: Args to use when creating the role and policies. Can't be specified if `roleArn` is used.
        :param '_resource.NestedResourceOptionsArgs' opts: Resource options to use for the role. Can't be specified if `roleArn` is used.
        :param pulumi.Input[str] role_arn: ARN of existing role to use instead of creating a new role. Cannot be used in combination with `args` or `opts`.
        :param bool skip: Skips creation of the role if set to `true`.
        """
        if args is not None:
            pulumi.set(__self__, "args", args)
        if opts is not None:
            pulumi.set(__self__, "opts", opts)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if skip is not None:
            pulumi.set(__self__, "skip", skip)

    @property
    @pulumi.getter
    def args(self) -> Optional['RoleWithPolicyArgs']:
        """
        Args to use when creating the role and policies. Can't be specified if `roleArn` is used.
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional['RoleWithPolicyArgs']):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter
    def opts(self) -> Optional['_resource.NestedResourceOptionsArgs']:
        """
        Resource options to use for the role. Can't be specified if `roleArn` is used.
        """
        return pulumi.get(self, "opts")

    @opts.setter
    def opts(self, value: Optional['_resource.NestedResourceOptionsArgs']):
        pulumi.set(self, "opts", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of existing role to use instead of creating a new role. Cannot be used in combination with `args` or `opts`.
        """
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter
    def skip(self) -> Optional[bool]:
        """
        Skips creation of the role if set to `true`.
        """
        return pulumi.get(self, "skip")

    @skip.setter
    def skip(self, value: Optional[bool]):
        pulumi.set(self, "skip", value)


@pulumi.input_type
class RoleWithPolicyArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 force_detach_policies: Optional[pulumi.Input[bool]] = None,
                 inline_policies: Optional[pulumi.Input[Sequence[pulumi.Input['pulumi_aws.iam.RoleInlinePolicyArgs']]]] = None,
                 managed_policy_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 permissions_boundary: Optional[pulumi.Input[str]] = None,
                 policy_arns: Optional[Sequence[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Role resource and Policy attachments.
        :param pulumi.Input[str] description: Description of the role.
        :param pulumi.Input[bool] force_detach_policies: Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input['pulumi_aws.iam.RoleInlinePolicyArgs']]] inline_policies: Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, this provider will not manage any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies added out of band on `apply`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] managed_policy_arns: Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, this provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
        :param pulumi.Input[int] max_session_duration: Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
        :param pulumi.Input[str] name: Name of the role policy.
        :param pulumi.Input[str] name_prefix: Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[str] path: Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
        :param pulumi.Input[str] permissions_boundary: ARN of the policy that is used to set the permissions boundary for the role.
        :param Sequence[str] policy_arns: ARNs of the policies to attach to the created role.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of tags for the IAM role. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if force_detach_policies is not None:
            pulumi.set(__self__, "force_detach_policies", force_detach_policies)
        if inline_policies is not None:
            pulumi.set(__self__, "inline_policies", inline_policies)
        if managed_policy_arns is not None:
            pulumi.set(__self__, "managed_policy_arns", managed_policy_arns)
        if max_session_duration is not None:
            pulumi.set(__self__, "max_session_duration", max_session_duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if permissions_boundary is not None:
            pulumi.set(__self__, "permissions_boundary", permissions_boundary)
        if policy_arns is not None:
            pulumi.set(__self__, "policy_arns", policy_arns)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="forceDetachPolicies")
    def force_detach_policies(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
        """
        return pulumi.get(self, "force_detach_policies")

    @force_detach_policies.setter
    def force_detach_policies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_detach_policies", value)

    @property
    @pulumi.getter(name="inlinePolicies")
    def inline_policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['pulumi_aws.iam.RoleInlinePolicyArgs']]]]:
        """
        Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. See below. If no blocks are configured, this provider will not manage any inline policies in this resource. Configuring one empty block (i.e., `inline_policy {}`) will cause the provider to remove _all_ inline policies added out of band on `apply`.
        """
        return pulumi.get(self, "inline_policies")

    @inline_policies.setter
    def inline_policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['pulumi_aws.iam.RoleInlinePolicyArgs']]]]):
        pulumi.set(self, "inline_policies", value)

    @property
    @pulumi.getter(name="managedPolicyArns")
    def managed_policy_arns(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, this provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managed_policy_arns = []`) will cause the provider to remove _all_ managed policy attachments.
        """
        return pulumi.get(self, "managed_policy_arns")

    @managed_policy_arns.setter
    def managed_policy_arns(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "managed_policy_arns", value)

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
        """
        return pulumi.get(self, "max_session_duration")

    @max_session_duration.setter
    def max_session_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_session_duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the role policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="permissionsBoundary")
    def permissions_boundary(self) -> Optional[pulumi.Input[str]]:
        """
        ARN of the policy that is used to set the permissions boundary for the role.
        """
        return pulumi.get(self, "permissions_boundary")

    @permissions_boundary.setter
    def permissions_boundary(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permissions_boundary", value)

    @property
    @pulumi.getter(name="policyArns")
    def policy_arns(self) -> Optional[Sequence[str]]:
        """
        ARNs of the policies to attach to the created role.
        """
        return pulumi.get(self, "policy_arns")

    @policy_arns.setter
    def policy_arns(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "policy_arns", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value mapping of tags for the IAM role. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


