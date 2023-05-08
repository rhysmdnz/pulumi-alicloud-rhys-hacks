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
    'GetForwardEntriesResult',
    'AwaitableGetForwardEntriesResult',
    'get_forward_entries',
    'get_forward_entries_output',
]

@pulumi.output_type
class GetForwardEntriesResult:
    """
    A collection of values returned by getForwardEntries.
    """
    def __init__(__self__, entries=None, external_ip=None, external_port=None, forward_entry_name=None, forward_table_id=None, id=None, ids=None, internal_ip=None, internal_port=None, ip_protocol=None, name_regex=None, names=None, output_file=None, status=None):
        if entries and not isinstance(entries, list):
            raise TypeError("Expected argument 'entries' to be a list")
        pulumi.set(__self__, "entries", entries)
        if external_ip and not isinstance(external_ip, str):
            raise TypeError("Expected argument 'external_ip' to be a str")
        pulumi.set(__self__, "external_ip", external_ip)
        if external_port and not isinstance(external_port, str):
            raise TypeError("Expected argument 'external_port' to be a str")
        pulumi.set(__self__, "external_port", external_port)
        if forward_entry_name and not isinstance(forward_entry_name, str):
            raise TypeError("Expected argument 'forward_entry_name' to be a str")
        pulumi.set(__self__, "forward_entry_name", forward_entry_name)
        if forward_table_id and not isinstance(forward_table_id, str):
            raise TypeError("Expected argument 'forward_table_id' to be a str")
        pulumi.set(__self__, "forward_table_id", forward_table_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ids and not isinstance(ids, list):
            raise TypeError("Expected argument 'ids' to be a list")
        pulumi.set(__self__, "ids", ids)
        if internal_ip and not isinstance(internal_ip, str):
            raise TypeError("Expected argument 'internal_ip' to be a str")
        pulumi.set(__self__, "internal_ip", internal_ip)
        if internal_port and not isinstance(internal_port, str):
            raise TypeError("Expected argument 'internal_port' to be a str")
        pulumi.set(__self__, "internal_port", internal_port)
        if ip_protocol and not isinstance(ip_protocol, str):
            raise TypeError("Expected argument 'ip_protocol' to be a str")
        pulumi.set(__self__, "ip_protocol", ip_protocol)
        if name_regex and not isinstance(name_regex, str):
            raise TypeError("Expected argument 'name_regex' to be a str")
        pulumi.set(__self__, "name_regex", name_regex)
        if names and not isinstance(names, list):
            raise TypeError("Expected argument 'names' to be a list")
        pulumi.set(__self__, "names", names)
        if output_file and not isinstance(output_file, str):
            raise TypeError("Expected argument 'output_file' to be a str")
        pulumi.set(__self__, "output_file", output_file)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def entries(self) -> Sequence['outputs.GetForwardEntriesEntryResult']:
        """
        A list of Forward Entries. Each element contains the following attributes:
        """
        return pulumi.get(self, "entries")

    @property
    @pulumi.getter(name="externalIp")
    def external_ip(self) -> Optional[str]:
        """
        The public IP address.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter(name="externalPort")
    def external_port(self) -> Optional[str]:
        """
        The public port.
        """
        return pulumi.get(self, "external_port")

    @property
    @pulumi.getter(name="forwardEntryName")
    def forward_entry_name(self) -> Optional[str]:
        """
        The name of forward entry.
        """
        return pulumi.get(self, "forward_entry_name")

    @property
    @pulumi.getter(name="forwardTableId")
    def forward_table_id(self) -> str:
        return pulumi.get(self, "forward_table_id")

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
        A list of Forward Entries IDs.
        """
        return pulumi.get(self, "ids")

    @property
    @pulumi.getter(name="internalIp")
    def internal_ip(self) -> Optional[str]:
        """
        The private IP address.
        """
        return pulumi.get(self, "internal_ip")

    @property
    @pulumi.getter(name="internalPort")
    def internal_port(self) -> Optional[str]:
        """
        The private port.
        """
        return pulumi.get(self, "internal_port")

    @property
    @pulumi.getter(name="ipProtocol")
    def ip_protocol(self) -> Optional[str]:
        """
        The protocol type.
        """
        return pulumi.get(self, "ip_protocol")

    @property
    @pulumi.getter(name="nameRegex")
    def name_regex(self) -> Optional[str]:
        return pulumi.get(self, "name_regex")

    @property
    @pulumi.getter
    def names(self) -> Sequence[str]:
        """
        A list of Forward Entries names.
        """
        return pulumi.get(self, "names")

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[str]:
        return pulumi.get(self, "output_file")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The status of forward entry.
        """
        return pulumi.get(self, "status")


class AwaitableGetForwardEntriesResult(GetForwardEntriesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetForwardEntriesResult(
            entries=self.entries,
            external_ip=self.external_ip,
            external_port=self.external_port,
            forward_entry_name=self.forward_entry_name,
            forward_table_id=self.forward_table_id,
            id=self.id,
            ids=self.ids,
            internal_ip=self.internal_ip,
            internal_port=self.internal_port,
            ip_protocol=self.ip_protocol,
            name_regex=self.name_regex,
            names=self.names,
            output_file=self.output_file,
            status=self.status)


def get_forward_entries(external_ip: Optional[str] = None,
                        external_port: Optional[str] = None,
                        forward_entry_name: Optional[str] = None,
                        forward_table_id: Optional[str] = None,
                        ids: Optional[Sequence[str]] = None,
                        internal_ip: Optional[str] = None,
                        internal_port: Optional[str] = None,
                        ip_protocol: Optional[str] = None,
                        name_regex: Optional[str] = None,
                        output_file: Optional[str] = None,
                        status: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetForwardEntriesResult:
    """
    This data source provides a list of Forward Entries owned by an Alibaba Cloud account.

    > **NOTE:** Available in 1.37.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "forward-entry-config-example-name"
    default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
    default_network = alicloud.vpc.Network("defaultNetwork",
        vpc_name=name,
        cidr_block="172.16.0.0/12")
    default_switch = alicloud.vpc.Switch("defaultSwitch",
        vpc_id=default_network.id,
        cidr_block="172.16.0.0/21",
        zone_id=default_zones.zones[0].id,
        vswitch_name=name)
    default_nat_gateway = alicloud.vpc.NatGateway("defaultNatGateway",
        vpc_id=default_network.id,
        internet_charge_type="PayByLcu",
        nat_gateway_name=name,
        nat_type="Enhanced",
        vswitch_id=default_switch.id)
    default_eip_address = alicloud.ecs.EipAddress("defaultEipAddress", address_name=name)
    default_eip_association = alicloud.ecs.EipAssociation("defaultEipAssociation",
        allocation_id=default_eip_address.id,
        instance_id=default_nat_gateway.id)
    default_forward_entry = alicloud.vpc.ForwardEntry("defaultForwardEntry",
        forward_table_id=default_nat_gateway.forward_table_ids,
        external_ip=default_eip_address.ip_address,
        external_port="80",
        ip_protocol="tcp",
        internal_ip="172.16.0.3",
        internal_port="8080")
    default_forward_entries = alicloud.vpc.get_forward_entries_output(forward_table_id=default_forward_entry.forward_table_id)
    ```


    :param str external_ip: The public IP address.
    :param str external_port: The public port.
    :param str forward_entry_name: The name of forward entry.
    :param str forward_table_id: The ID of the Forward table.
    :param Sequence[str] ids: A list of Forward Entries IDs.
    :param str internal_ip: The private IP address.
    :param str internal_port: The internal port.
    :param str ip_protocol: The ip protocol. Valid values: `any`,`tcp` and `udp`.
    :param str name_regex: A regex string to filter results by forward entry name.
    :param str status: The status of farward entry. Valid value `Available`, `Deleting` and `Pending`.
    """
    __args__ = dict()
    __args__['externalIp'] = external_ip
    __args__['externalPort'] = external_port
    __args__['forwardEntryName'] = forward_entry_name
    __args__['forwardTableId'] = forward_table_id
    __args__['ids'] = ids
    __args__['internalIp'] = internal_ip
    __args__['internalPort'] = internal_port
    __args__['ipProtocol'] = ip_protocol
    __args__['nameRegex'] = name_regex
    __args__['outputFile'] = output_file
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('alicloud:vpc/getForwardEntries:getForwardEntries', __args__, opts=opts, typ=GetForwardEntriesResult).value

    return AwaitableGetForwardEntriesResult(
        entries=__ret__.entries,
        external_ip=__ret__.external_ip,
        external_port=__ret__.external_port,
        forward_entry_name=__ret__.forward_entry_name,
        forward_table_id=__ret__.forward_table_id,
        id=__ret__.id,
        ids=__ret__.ids,
        internal_ip=__ret__.internal_ip,
        internal_port=__ret__.internal_port,
        ip_protocol=__ret__.ip_protocol,
        name_regex=__ret__.name_regex,
        names=__ret__.names,
        output_file=__ret__.output_file,
        status=__ret__.status)


@_utilities.lift_output_func(get_forward_entries)
def get_forward_entries_output(external_ip: Optional[pulumi.Input[Optional[str]]] = None,
                               external_port: Optional[pulumi.Input[Optional[str]]] = None,
                               forward_entry_name: Optional[pulumi.Input[Optional[str]]] = None,
                               forward_table_id: Optional[pulumi.Input[str]] = None,
                               ids: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                               internal_ip: Optional[pulumi.Input[Optional[str]]] = None,
                               internal_port: Optional[pulumi.Input[Optional[str]]] = None,
                               ip_protocol: Optional[pulumi.Input[Optional[str]]] = None,
                               name_regex: Optional[pulumi.Input[Optional[str]]] = None,
                               output_file: Optional[pulumi.Input[Optional[str]]] = None,
                               status: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetForwardEntriesResult]:
    """
    This data source provides a list of Forward Entries owned by an Alibaba Cloud account.

    > **NOTE:** Available in 1.37.0+.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_alicloud as alicloud

    config = pulumi.Config()
    name = config.get("name")
    if name is None:
        name = "forward-entry-config-example-name"
    default_zones = alicloud.get_zones(available_resource_creation="VSwitch")
    default_network = alicloud.vpc.Network("defaultNetwork",
        vpc_name=name,
        cidr_block="172.16.0.0/12")
    default_switch = alicloud.vpc.Switch("defaultSwitch",
        vpc_id=default_network.id,
        cidr_block="172.16.0.0/21",
        zone_id=default_zones.zones[0].id,
        vswitch_name=name)
    default_nat_gateway = alicloud.vpc.NatGateway("defaultNatGateway",
        vpc_id=default_network.id,
        internet_charge_type="PayByLcu",
        nat_gateway_name=name,
        nat_type="Enhanced",
        vswitch_id=default_switch.id)
    default_eip_address = alicloud.ecs.EipAddress("defaultEipAddress", address_name=name)
    default_eip_association = alicloud.ecs.EipAssociation("defaultEipAssociation",
        allocation_id=default_eip_address.id,
        instance_id=default_nat_gateway.id)
    default_forward_entry = alicloud.vpc.ForwardEntry("defaultForwardEntry",
        forward_table_id=default_nat_gateway.forward_table_ids,
        external_ip=default_eip_address.ip_address,
        external_port="80",
        ip_protocol="tcp",
        internal_ip="172.16.0.3",
        internal_port="8080")
    default_forward_entries = alicloud.vpc.get_forward_entries_output(forward_table_id=default_forward_entry.forward_table_id)
    ```


    :param str external_ip: The public IP address.
    :param str external_port: The public port.
    :param str forward_entry_name: The name of forward entry.
    :param str forward_table_id: The ID of the Forward table.
    :param Sequence[str] ids: A list of Forward Entries IDs.
    :param str internal_ip: The private IP address.
    :param str internal_port: The internal port.
    :param str ip_protocol: The ip protocol. Valid values: `any`,`tcp` and `udp`.
    :param str name_regex: A regex string to filter results by forward entry name.
    :param str status: The status of farward entry. Valid value `Available`, `Deleting` and `Pending`.
    """
    ...
