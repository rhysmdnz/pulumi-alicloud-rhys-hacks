# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetControlPoliciesResult',
    'AwaitableGetControlPoliciesResult',
    'get_control_policies',
    'get_control_policies_output',
]

@pulumi.output_type
class GetControlPoliciesResult:
    """
    A collection of values returned by getControlPolicies.
    """
    def __init__(__self__, enable_details=None, id=None, ids=None, language=None, name_regex=None, names=None, output_file=None, policies=None, policy_type=None):
        if enable_details and not isinstance(enable_details, bool):
            raise TypeError("Expected argument 'enable_details' to be a bool")
        pulumi.set(__self__, "enable_details", enable_details)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if language and not isinstance(language, str):
            raise TypeError("Expected argument 'language' to be a str")
        pulumi.set(__self__, "language", language)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if policy_type and not isinstance(policy_type, str):
            raise TypeError("Expected argument 'policy_type' to be a str")
        pulumi.set(__self__, "policy_type", policy_type)

    @property
    @pulumi.getter(name="enableDetails")
    def enable_details(self) -> Optional[bool]:
        return pulumi.get(self, "enable_details")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ids(self) -> Sequence[str]:
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter
    def language(self) -> Optional[str]:
        return pulumi.get(self, "language")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def policies(self) -> Sequence['outputs.GetControlPoliciesPolicyResult']:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[str]:
        return pulumi.get(self, "policy_type")


class AwaitableGetControlPoliciesResult(GetControlPoliciesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetControlPoliciesResult(
            enable_details=self.enable_details,
            id=self.id,
            ids=self.ids,
            language=self.language,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            policies=self.policies,
            policy_type=self.policy_type)


def get_control_policies(enable_details: Optional[bool] = None,
                         ids: Optional[Sequence[str]] = None,
                         language: Optional[str] = None,
                         name_regex: Optional[str] = None,
                         output_file: Optional[str] = None,
                         policy_type: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetControlPoliciesResult:
    """
    This data source provides the Resource Manager Control Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.120.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.resourcemanager.get_control_policies(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstResourceManagerControlPolicyId", example.policies[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Control Policy IDs.
    :param str language: The language. Valid value `zh-CN`, `en`, and `ja`. Default value `zh-CN`.
    :param str name_regex: A regex string to filter results by Control Policy name.
    :param str policy_type: The type of policy.
    """
    __args__ = dict()
    __args__['enableDetails'] = enable_details
    __args__['ids'] = ids
    __args__['language'] = language
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['policyType'] = policy_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:resourcemanager/getControlPolicies:getControlPolicies', __args__, opts=opts, typ=GetControlPoliciesResult).value

    return AwaitableGetControlPoliciesResult(
        enable_details=__ret__.enable_details,
        id=__ret__.id,
        ids=__ret__.ids,
        language=__ret__.language,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        policies=__ret__.policies,
        policy_type=__ret__.policy_type)


@_utilities.lift_output_func(get_control_policies)
def get_control_policies_output(enable_details: Optional[pulumi.Input[Optional[bool]]] = None,
                                ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                language: Optional[pulumi.Input[Optional[str]]] = None,
                                name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                policy_type: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetControlPoliciesResult]:
    """
    This data source provides the Resource Manager Control Policies of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.120.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    example = alicloud.resourcemanager.get_control_policies(ids=["example_value"],
        name_regex="the_resource_name")
    pulumi.export("firstResourceManagerControlPolicyId", example.policies[0].id)
    ```


    :param bool enable_details: Default to `false`. Set it to `true` can output more details about resource attributes.
    :param Sequence[str] ids: A list of Control Policy IDs.
    :param str language: The language. Valid value `zh-CN`, `en`, and `ja`. Default value `zh-CN`.
    :param str name_regex: A regex string to filter results by Control Policy name.
    :param str policy_type: The type of policy.
    """
    ...