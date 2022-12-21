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
    'GetApplicationGroupsResult',
    'AwaitableGetApplicationGroupsResult',
    'get_application_groups',
    'get_application_groups_output',
]

@pulumi.output_type
class GetApplicationGroupsResult:
    """
    A collection of values returned by getApplicationGroups.
    """
    def __init__(__self__, application_name=None, deploy_region_id=None, groups=None, id=None, ids=None, name_regex=None, names=None, output_file=None):
        if application_name and not isinstance(application_name, str):
            raise TypeError("Expected argument 'application_name' to be a str")
        pulumi.set(__self__, "application_name", application_name)
        if deploy_region_id and not isinstance(deploy_region_id, str):
            raise TypeError("Expected argument 'deploy_region_id' to be a str")
        pulumi.set(__self__, "deploy_region_id", deploy_region_id)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)

    @property
    @pulumi.getter(name="applicationName")
    def application_name(self) -> str:
        return pulumi.get(self, "application_name")

    @property
    @pulumi.getter(name="deployRegionId")
    def deploy_region_id(self) -> Optional[str]:
        return pulumi.get(self, "deploy_region_id")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetApplicationGroupsGroupResult']:
        return pulumi.get(self, "groups")

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


class AwaitableGetApplicationGroupsResult(GetApplicationGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationGroupsResult(
            application_name=self.application_name,
            deploy_region_id=self.deploy_region_id,
            groups=self.groups,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file)


def get_application_groups(application_name: Optional[str] = None,
                           deploy_region_id: Optional[str] = None,
                           ids: Optional[Sequence[str]] = None,
                           name_regex: Optional[str] = None,
                           output_file: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationGroupsResult:
    """
    This data source provides the Oos Application Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.146.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.oos.get_application_groups(application_name="example_value",
        ids=[
            "my-ApplicationGroup-1",
            "my-ApplicationGroup-2",
        ])
    pulumi.export("oosApplicationGroupId1", ids.groups[0].id)
    name_regex = alicloud.oos.get_application_groups(application_name="example_value",
        name_regex="^my-ApplicationGroup")
    pulumi.export("oosApplicationGroupId2", name_regex.groups[0].id)
    ```


    :param str application_name: The name of the Application.
    :param str deploy_region_id: The region ID of the deployment.
    :param Sequence[str] ids: A list of Application Group IDs. Its element value is same as Application Group Name.
    :param str name_regex: A regex string to filter results by Application Group name.
    """
    __args__ = dict()
    __args__['applicationName'] = application_name
    __args__['deployRegionId'] = deploy_region_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:oos/getApplicationGroups:getApplicationGroups', __args__, opts=opts, typ=GetApplicationGroupsResult).value

    return AwaitableGetApplicationGroupsResult(
        application_name=__ret__.application_name,
        deploy_region_id=__ret__.deploy_region_id,
        groups=__ret__.groups,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file)


@_utilities.lift_output_func(get_application_groups)
def get_application_groups_output(application_name: Optional[pulumi.Input[str]] = None,
                                  deploy_region_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                  name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                  output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetApplicationGroupsResult]:
    """
    This data source provides the Oos Application Groups of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.146.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.oos.get_application_groups(application_name="example_value",
        ids=[
            "my-ApplicationGroup-1",
            "my-ApplicationGroup-2",
        ])
    pulumi.export("oosApplicationGroupId1", ids.groups[0].id)
    name_regex = alicloud.oos.get_application_groups(application_name="example_value",
        name_regex="^my-ApplicationGroup")
    pulumi.export("oosApplicationGroupId2", name_regex.groups[0].id)
    ```


    :param str application_name: The name of the Application.
    :param str deploy_region_id: The region ID of the deployment.
    :param Sequence[str] ids: A list of Application Group IDs. Its element value is same as Application Group Name.
    :param str name_regex: A regex string to filter results by Application Group name.
    """
    ...