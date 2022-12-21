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
    'GetCommonBandwidthPackagesResult',
    'AwaitableGetCommonBandwidthPackagesResult',
    'get_common_bandwidth_packages',
    'get_common_bandwidth_packages_output',
]

@pulumi.output_type
class GetCommonBandwidthPackagesResult:
    """
    A collection of values returned by getCommonBandwidthPackages.
    """
    def __init__(__self__, bandwidth_package_name=None, dry_run=None, id=None, ids=None, include_reservation_data=None, name_regex=None, names=None, output_file=None, packages=None, resource_group_id=None, status=None):
        if bandwidth_package_name and not isinstance(bandwidth_package_name, str):
            raise TypeError("Expected argument 'bandwidth_package_name' to be a str")
        pulumi.set(__self__, "bandwidth_package_name", bandwidth_package_name)
        if dry_run and not isinstance(dry_run, bool):
            raise TypeError("Expected argument 'dry_run' to be a bool")
        pulumi.set(__self__, "dry_run", dry_run)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if include_reservation_data and not isinstance(include_reservation_data, bool):
            raise TypeError("Expected argument 'include_reservation_data' to be a bool")
        pulumi.set(__self__, "include_reservation_data", include_reservation_data)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if packages and not isinstance(packages, list):
            raise TypeError("Expected argument 'packages' to be a list")
        pulumi.set(__self__, "packages", packages)
        if resource_group_id and not isinstance(resource_group_id, str):
            raise TypeError("Expected argument 'resource_group_id' to be a str")
        pulumi.set(__self__, "resource_group_id", resource_group_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="bandwidthPackageName")
    def bandwidth_package_name(self) -> Optional[str]:
        """
        The name of bandwidth package.
        """
        return pulumi.get(self, "bandwidth_package_name")

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[bool]:
        return pulumi.get(self, "dry_run")

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
        """
        (Optional) A list of Common Bandwidth Packages IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="includeReservationData")
    def include_reservation_data(self) -> Optional[bool]:
        return pulumi.get(self, "include_reservation_data")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Common Bandwidth Packages names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def packages(self) -> Sequence['outputs.GetCommonBandwidthPackagesPackageResult']:
        """
        A list of Common Bandwidth Packages. Each element contains the following attributes:
        """
        return pulumi.get(self, "packages")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The Id of resource group which the common bandwidth package belongs.
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Status of the Common Bandwidth Package.
        """
        return pulumi.get(self, "status")


class AwaitableGetCommonBandwidthPackagesResult(GetCommonBandwidthPackagesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCommonBandwidthPackagesResult(
            bandwidth_package_name=self.bandwidth_package_name,
            dry_run=self.dry_run,
            id=self.id,
            ids=self.ids,
            include_reservation_data=self.include_reservation_data,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            packages=self.packages,
            resource_group_id=self.resource_group_id,
            status=self.status)


def get_common_bandwidth_packages(bandwidth_package_name: Optional[str] = None,
                                  dry_run: Optional[bool] = None,
                                  ids: Optional[Sequence[str]] = None,
                                  include_reservation_data: Optional[bool] = None,
                                  name_regex: Optional[str] = None,
                                  output_file: Optional[str] = None,
                                  resource_group_id: Optional[str] = None,
                                  status: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCommonBandwidthPackagesResult:
    """
    This data source provides a list of Common Bandwidth Packages owned by an Alibaba Cloud account.

    > **NOTE:** Available in 1.36.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    foo_common_bandwith_package = alicloud.vpc.CommonBandwithPackage("fooCommonBandwithPackage",
        bandwidth="2",
        description="tf-testAcc-CommonBandwidthPackage")
    foo_common_bandwidth_packages = alicloud.vpc.get_common_bandwidth_packages_output(ids=[foo_common_bandwith_package.id],
        name_regex="^tf-testAcc.*")
    ```
    ## Public ip addresses Block

      The public ip addresses mapping supports the following:

      * `ip_address`   - The address of the EIP.
      * `allocation_id` - The ID of the EIP instance.
      * `bandwidth_package_ip_relation_status` - The IP relation status of bandwidth package.


    :param str bandwidth_package_name: The name of bandwidth package.
    :param bool dry_run: Specifies whether to precheck only the request.
    :param Sequence[str] ids: A list of Common Bandwidth Packages IDs.
    :param bool include_reservation_data: Specifies whether to return data of orders that have not taken effect.
    :param str name_regex: A regex string to filter results by name.
    :param str resource_group_id: The Id of resource group which the common bandwidth package belongs.
    :param str status: The status of bandwidth package. Valid values: `Available` and `Pending`.
    """
    __args__ = dict()
    __args__['bandwidthPackageName'] = bandwidth_package_name
    __args__['dryRun'] = dry_run
    __args__['ids'] = ids
    __args__['includeReservationData'] = include_reservation_data
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['resourceGroupId'] = resource_group_id
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getCommonBandwidthPackages:getCommonBandwidthPackages', __args__, opts=opts, typ=GetCommonBandwidthPackagesResult).value

    return AwaitableGetCommonBandwidthPackagesResult(
        bandwidth_package_name=__ret__.bandwidth_package_name,
        dry_run=__ret__.dry_run,
        id=__ret__.id,
        ids=__ret__.ids,
        include_reservation_data=__ret__.include_reservation_data,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        packages=__ret__.packages,
        resource_group_id=__ret__.resource_group_id,
        status=__ret__.status)


@_utilities.lift_output_func(get_common_bandwidth_packages)
def get_common_bandwidth_packages_output(bandwidth_package_name: Optional[pulumi.Input[Optional[str]]] = None,
                                         dry_run: Optional[pulumi.Input[Optional[bool]]] = None,
                                         ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                         include_reservation_data: Optional[pulumi.Input[Optional[bool]]] = None,
                                         name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                         output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                         resource_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                         status: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCommonBandwidthPackagesResult]:
    """
    This data source provides a list of Common Bandwidth Packages owned by an Alibaba Cloud account.

    > **NOTE:** Available in 1.36.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    foo_common_bandwith_package = alicloud.vpc.CommonBandwithPackage("fooCommonBandwithPackage",
        bandwidth="2",
        description="tf-testAcc-CommonBandwidthPackage")
    foo_common_bandwidth_packages = alicloud.vpc.get_common_bandwidth_packages_output(ids=[foo_common_bandwith_package.id],
        name_regex="^tf-testAcc.*")
    ```
    ## Public ip addresses Block

      The public ip addresses mapping supports the following:

      * `ip_address`   - The address of the EIP.
      * `allocation_id` - The ID of the EIP instance.
      * `bandwidth_package_ip_relation_status` - The IP relation status of bandwidth package.


    :param str bandwidth_package_name: The name of bandwidth package.
    :param bool dry_run: Specifies whether to precheck only the request.
    :param Sequence[str] ids: A list of Common Bandwidth Packages IDs.
    :param bool include_reservation_data: Specifies whether to return data of orders that have not taken effect.
    :param str name_regex: A regex string to filter results by name.
    :param str resource_group_id: The Id of resource group which the common bandwidth package belongs.
    :param str status: The status of bandwidth package. Valid values: `Available` and `Pending`.
    """
    ...