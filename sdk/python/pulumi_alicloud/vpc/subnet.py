# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SubnetArgs', 'Subnet']

@pulumi.input_type
class SubnetArgs:
    def __init__(__self__, *,
                 cidr_block: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 ipv6_cidr_block_mask: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Subnet resource.
        """
        pulumi.set(__self__, "cidr_block", cidr_block)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if availability_zone is not None:
            warnings.warn("""Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""", DeprecationWarning)
            pulumi.log.warn("""availability_zone is deprecated: Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""")
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_ipv6 is not None:
            pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if ipv6_cidr_block_mask is not None:
            pulumi.set(__self__, "ipv6_cidr_block_mask", ipv6_cidr_block_mask)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vswitch_name is not None:
            pulumi.set(__self__, "vswitch_name", vswitch_name)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_ipv6")

    @enable_ipv6.setter
    def enable_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ipv6", value)

    @property
    @pulumi.getter(name="ipv6CidrBlockMask")
    def ipv6_cidr_block_mask(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ipv6_cidr_block_mask")

    @ipv6_cidr_block_mask.setter
    def ipv6_cidr_block_mask(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipv6_cidr_block_mask", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vswitchName")
    def vswitch_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vswitch_name")

    @vswitch_name.setter
    def vswitch_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


@pulumi.input_type
class _SubnetState:
    def __init__(__self__, *,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr_block_mask: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Subnet resources.
        """
        if availability_zone is not None:
            warnings.warn("""Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""", DeprecationWarning)
            pulumi.log.warn("""availability_zone is deprecated: Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""")
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if cidr_block is not None:
            pulumi.set(__self__, "cidr_block", cidr_block)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_ipv6 is not None:
            pulumi.set(__self__, "enable_ipv6", enable_ipv6)
        if ipv6_cidr_block is not None:
            pulumi.set(__self__, "ipv6_cidr_block", ipv6_cidr_block)
        if ipv6_cidr_block_mask is not None:
            pulumi.set(__self__, "ipv6_cidr_block_mask", ipv6_cidr_block_mask)
        if name is not None:
            warnings.warn("""Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""", DeprecationWarning)
            pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""")
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vswitch_name is not None:
            pulumi.set(__self__, "vswitch_name", vswitch_name)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "cidr_block")

    @cidr_block.setter
    def cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr_block", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enable_ipv6")

    @enable_ipv6.setter
    def enable_ipv6(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_ipv6", value)

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ipv6_cidr_block")

    @ipv6_cidr_block.setter
    def ipv6_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipv6_cidr_block", value)

    @property
    @pulumi.getter(name="ipv6CidrBlockMask")
    def ipv6_cidr_block_mask(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "ipv6_cidr_block_mask")

    @ipv6_cidr_block_mask.setter
    def ipv6_cidr_block_mask(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ipv6_cidr_block_mask", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vswitchName")
    def vswitch_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vswitch_name")

    @vswitch_name.setter
    def vswitch_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_name", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)


warnings.warn("""This resource has been deprecated and replaced by the Switch resource.""", DeprecationWarning)


class Subnet(pulumi.CustomResource):
    warnings.warn("""This resource has been deprecated and replaced by the Switch resource.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 ipv6_cidr_block_mask: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Subnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubnetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Subnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SubnetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_ipv6: Optional[pulumi.Input[bool]] = None,
                 ipv6_cidr_block_mask: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vswitch_name: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""Subnet is deprecated: This resource has been deprecated and replaced by the Switch resource.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetArgs.__new__(SubnetArgs)

            if availability_zone is not None and not opts.urn:
                warnings.warn("""Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""", DeprecationWarning)
                pulumi.log.warn("""availability_zone is deprecated: Field 'availability_zone' has been deprecated from provider version 1.119.0. New field 'zone_id' instead.""")
            __props__.__dict__["availability_zone"] = availability_zone
            if cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'cidr_block'")
            __props__.__dict__["cidr_block"] = cidr_block
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_ipv6"] = enable_ipv6
            __props__.__dict__["ipv6_cidr_block_mask"] = ipv6_cidr_block_mask
            if name is not None and not opts.urn:
                warnings.warn("""Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""", DeprecationWarning)
                pulumi.log.warn("""name is deprecated: Field 'name' has been deprecated from provider version 1.119.0. New field 'vswitch_name' instead.""")
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["vswitch_name"] = vswitch_name
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["ipv6_cidr_block"] = None
            __props__.__dict__["status"] = None
        super(Subnet, __self__).__init__(
            'alicloud:vpc/subnet:Subnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable_ipv6: Optional[pulumi.Input[bool]] = None,
            ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
            ipv6_cidr_block_mask: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vswitch_name: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'Subnet':
        """
        Get an existing Subnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubnetState.__new__(_SubnetState)

        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["cidr_block"] = cidr_block
        __props__.__dict__["description"] = description
        __props__.__dict__["enable_ipv6"] = enable_ipv6
        __props__.__dict__["ipv6_cidr_block"] = ipv6_cidr_block
        __props__.__dict__["ipv6_cidr_block_mask"] = ipv6_cidr_block_mask
        __props__.__dict__["name"] = name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vswitch_name"] = vswitch_name
        __props__.__dict__["zone_id"] = zone_id
        return Subnet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="cidrBlock")
    def cidr_block(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cidr_block")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableIpv6")
    def enable_ipv6(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enable_ipv6")

    @property
    @pulumi.getter(name="ipv6CidrBlock")
    def ipv6_cidr_block(self) -> pulumi.Output[str]:
        return pulumi.get(self, "ipv6_cidr_block")

    @property
    @pulumi.getter(name="ipv6CidrBlockMask")
    def ipv6_cidr_block_mask(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "ipv6_cidr_block_mask")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vswitchName")
    def vswitch_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vswitch_name")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "zone_id")

