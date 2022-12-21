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
    'GetGatewayFileSharesResult',
    'AwaitableGetGatewayFileSharesResult',
    'get_gateway_file_shares',
    'get_gateway_file_shares_output',
]

@pulumi.output_type
class GetGatewayFileSharesResult:
    """
    A collection of values returned by getGatewayFileShares.
    """
    def __init__(__self__, gateway_id=None, id=None, ids=None, name_regex=None, names=None, output_file=None, shares=None):
        if gateway_id and not isinstance(gateway_id, str):
            raise TypeError("Expected argument 'gateway_id' to be a str")
        pulumi.set(__self__, "gateway_id", gateway_id)
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
        if shares and not isinstance(shares, list):
            raise TypeError("Expected argument 'shares' to be a list")
        pulumi.set(__self__, "shares", shares)

    @property
    @pulumi.getter(name="gatewayId")
    def gateway_id(self) -> str:
        return pulumi.get(self, "gateway_id")

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

    @property
    @pulumi.getter
    def shares(self) -> Sequence['outputs.GetGatewayFileSharesShareResult']:
        return pulumi.get(self, "shares")


class AwaitableGetGatewayFileSharesResult(GetGatewayFileSharesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGatewayFileSharesResult(
            gateway_id=self.gateway_id,
            id=self.id,
            ids=self.ids,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            shares=self.shares)


def get_gateway_file_shares(gateway_id: Optional[str] = None,
                            ids: Optional[Sequence[str]] = None,
                            name_regex: Optional[str] = None,
                            output_file: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGatewayFileSharesResult:
    """
    This data source provides the Cloud Storage Gateway Gateway File Shares of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.144.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cloudstoragegateway.get_gateway_file_shares(gateway_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cloudStorageGatewayGatewayFileShareId1", ids.shares[0].id)
    name_regex = alicloud.cloudstoragegateway.get_gateway_file_shares(gateway_id="example_value",
        name_regex="^my-GatewayFileShare")
    pulumi.export("cloudStorageGatewayGatewayFileShareId2", name_regex.shares[0].id)
    ```


    :param str gateway_id: The ID of the gateway.
    :param Sequence[str] ids: A list of Gateway File Share IDs.
    :param str name_regex: A regex string to filter results by Gateway File Share name.
    """
    __args__ = dict()
    __args__['gatewayId'] = gateway_id
    __args__['ids'] = ids
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:cloudstoragegateway/getGatewayFileShares:getGatewayFileShares', __args__, opts=opts, typ=GetGatewayFileSharesResult).value

    return AwaitableGetGatewayFileSharesResult(
        gateway_id=__ret__.gateway_id,
        id=__ret__.id,
        ids=__ret__.ids,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        shares=__ret__.shares)


@_utilities.lift_output_func(get_gateway_file_shares)
def get_gateway_file_shares_output(gateway_id: Optional[pulumi.Input[str]] = None,
                                   ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                   name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                                   output_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGatewayFileSharesResult]:
    """
    This data source provides the Cloud Storage Gateway Gateway File Shares of the current Alibaba Cloud user.

    > **NOTE:** Available in v1.144.0+.

    ## Example Usage

    Basic Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    ids = alicloud.cloudstoragegateway.get_gateway_file_shares(gateway_id="example_value",
        ids=[
            "example_value-1",
            "example_value-2",
        ])
    pulumi.export("cloudStorageGatewayGatewayFileShareId1", ids.shares[0].id)
    name_regex = alicloud.cloudstoragegateway.get_gateway_file_shares(gateway_id="example_value",
        name_regex="^my-GatewayFileShare")
    pulumi.export("cloudStorageGatewayGatewayFileShareId2", name_regex.shares[0].id)
    ```


    :param str gateway_id: The ID of the gateway.
    :param Sequence[str] ids: A list of Gateway File Share IDs.
    :param str name_regex: A regex string to filter results by Gateway File Share name.
    """
    ...