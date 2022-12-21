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
    'GetAccessRulesResult',
    'AwaitableGetAccessRulesResult',
    'get_access_rules',
    'get_access_rules_output',
]

@pulumi.output_type
class GetAccessRulesResult:
    """
    A collection of values returned by getAccessRules.
    """
    def __init__(__self__, access_group_id=None, id=None, ids=None, output_file=None, rules=None):
        if access_group_id and not isinstance(access_group_id, str):
            raise TypeError("Expected argument 'access_group_id' to be a str")
        pulumi.set(__self__, "access_group_id", access_group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)

    @property
    @pulumi.getter(name="accessGroupId")
    def access_group_id(self) -> str:
        return pulumi.get(self, "access_group_id")

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
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetAccessRulesRuleResult']:
        return pulumi.get(self, "rules")


class AwaitableGetAccessRulesResult(GetAccessRulesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessRulesResult(
            access_group_id=self.access_group_id,
            id=self.id,
            ids=self.ids,
            output_file=self.output_file,
            rules=self.rules)


def get_access_rules(access_group_id: Optional[str] = None,
                     ids: Optional[Sequence[str]] = None,
                     output_file: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessRulesResult:
    """
    This data source provides the Dfs Access Rules of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.140.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.dfs.get_access_rules(access_group_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("dfsAccessRuleId1", ids.rules[0].id)
    ```


    :param str access_group_id: The resource ID of the Access Group.
    :param Sequence[str] ids: A list of Access Rule IDs.
    """
    __args__ = dict()
    __args__['accessGroupId'] = access_group_id
    __args__['ids'] = ids
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:dfs/getAccessRules:getAccessRules', __args__, opts=opts, typ=GetAccessRulesResult).value

    return AwaitableGetAccessRulesResult(
        access_group_id=__ret__.access_group_id,
        id=__ret__.id,
        ids=__ret__.ids,
        output_file=__ret__.output_file,
        rules=__ret__.rules)


@_utilities.lift_output_func(get_access_rules)
def get_access_rules_output(access_group_id: Optional[pulumi.Input[str]] = None,
                            ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                            output_file: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessRulesResult]:
    """
    This data source provides the Dfs Access Rules of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.140.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.dfs.get_access_rules(access_group_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("dfsAccessRuleId1", ids.rules[0].id)
    ```


    :param str access_group_id: The resource ID of the Access Group.
    :param Sequence[str] ids: A list of Access Rule IDs.
    """
    ...