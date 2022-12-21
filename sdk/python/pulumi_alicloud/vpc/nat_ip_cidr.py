# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NatIpCidrArgs', 'NatIpCidr']

@pulumi.input_type
class NatIpCidrArgs:
    def __init__(__self__, *,
                 nat_gateway_id: pulumi.Input[str],
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 nat_ip_cidr: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_description: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NatIpCidr resource.
        :param pulumi.Input[str] nat_gateway_id: The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        :param pulumi.Input[bool] dry_run: Specifies whether to precheck this request only. Valid values: `true` and `false`.
        :param pulumi.Input[str] nat_ip_cidr: - The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        :param pulumi.Input[str] nat_ip_cidr_description: The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] nat_ip_cidr_name: The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        """
        pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if nat_ip_cidr is not None:
            pulumi.set(__self__, "nat_ip_cidr", nat_ip_cidr)
        if nat_ip_cidr_description is not None:
            pulumi.set(__self__, "nat_ip_cidr_description", nat_ip_cidr_description)
        if nat_ip_cidr_name is not None:
            pulumi.set(__self__, "nat_ip_cidr_name", nat_ip_cidr_name)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Input[str]:
        """
        The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to precheck this request only. Valid values: `true` and `false`.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="natIpCidr")
    def nat_ip_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        - The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        """
        return pulumi.get(self, "nat_ip_cidr")

    @nat_ip_cidr.setter
    def nat_ip_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_ip_cidr", value)

    @property
    @pulumi.getter(name="natIpCidrDescription")
    def nat_ip_cidr_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "nat_ip_cidr_description")

    @nat_ip_cidr_description.setter
    def nat_ip_cidr_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_ip_cidr_description", value)

    @property
    @pulumi.getter(name="natIpCidrName")
    def nat_ip_cidr_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "nat_ip_cidr_name")

    @nat_ip_cidr_name.setter
    def nat_ip_cidr_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_ip_cidr_name", value)


@pulumi.input_type
class _NatIpCidrState:
    def __init__(__self__, *,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_description: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NatIpCidr resources.
        :param pulumi.Input[bool] dry_run: Specifies whether to precheck this request only. Valid values: `true` and `false`.
        :param pulumi.Input[str] nat_gateway_id: The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        :param pulumi.Input[str] nat_ip_cidr: - The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        :param pulumi.Input[str] nat_ip_cidr_description: The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] nat_ip_cidr_name: The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] status: The status of the CIDR block of the NAT gateway. Valid values: `Available`.
        """
        if dry_run is not None:
            pulumi.set(__self__, "dry_run", dry_run)
        if nat_gateway_id is not None:
            pulumi.set(__self__, "nat_gateway_id", nat_gateway_id)
        if nat_ip_cidr is not None:
            pulumi.set(__self__, "nat_ip_cidr", nat_ip_cidr)
        if nat_ip_cidr_description is not None:
            pulumi.set(__self__, "nat_ip_cidr_description", nat_ip_cidr_description)
        if nat_ip_cidr_name is not None:
            pulumi.set(__self__, "nat_ip_cidr_name", nat_ip_cidr_name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether to precheck this request only. Valid values: `true` and `false`.
        """
        return pulumi.get(self, "dry_run")

    @dry_run.setter
    def dry_run(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dry_run", value)

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        """
        return pulumi.get(self, "nat_gateway_id")

    @nat_gateway_id.setter
    def nat_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_gateway_id", value)

    @property
    @pulumi.getter(name="natIpCidr")
    def nat_ip_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        - The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        """
        return pulumi.get(self, "nat_ip_cidr")

    @nat_ip_cidr.setter
    def nat_ip_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_ip_cidr", value)

    @property
    @pulumi.getter(name="natIpCidrDescription")
    def nat_ip_cidr_description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "nat_ip_cidr_description")

    @nat_ip_cidr_description.setter
    def nat_ip_cidr_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_ip_cidr_description", value)

    @property
    @pulumi.getter(name="natIpCidrName")
    def nat_ip_cidr_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "nat_ip_cidr_name")

    @nat_ip_cidr_name.setter
    def nat_ip_cidr_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_ip_cidr_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the CIDR block of the NAT gateway. Valid values: `Available`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class NatIpCidr(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_description: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a VPC Nat Ip Cidr resource.

        For information about VPC Nat Ip Cidr and how to use it, see [What is Nat Ip Cidr](https://www.alibabacloud.com/help/doc-detail/281972.htm).

        > **NOTE:** Available in v1.136.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name="example_value",
            cidr_block="172.16.0.0/12")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vpc_id=alicloud_vpc["default"]["id"],
            cidr_block="172.16.0.0/21",
            zone_id=example_zones.zones[0].id,
            vswitch_name=var["name"])
        example_nat_gateway = alicloud.vpc.NatGateway("exampleNatGateway",
            vpc_id=alicloud_vpc["default"]["id"],
            internet_charge_type="PayByLcu",
            nat_gateway_name="example_value",
            description="example_value",
            nat_type="Enhanced",
            vswitch_id=example_switch.id,
            network_type="intranet")
        example_nat_ip_cidr = alicloud.vpc.NatIpCidr("exampleNatIpCidr",
            nat_gateway_id=example_nat_gateway.id,
            nat_ip_cidr_name="example_value",
            nat_ip_cidr="example_value")
        ```

        ## Import

        VPC Nat Ip Cidr can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/natIpCidr:NatIpCidr example <nat_gateway_id>:<nat_ip_cidr>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: Specifies whether to precheck this request only. Valid values: `true` and `false`.
        :param pulumi.Input[str] nat_gateway_id: The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        :param pulumi.Input[str] nat_ip_cidr: - The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        :param pulumi.Input[str] nat_ip_cidr_description: The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] nat_ip_cidr_name: The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NatIpCidrArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC Nat Ip Cidr resource.

        For information about VPC Nat Ip Cidr and how to use it, see [What is Nat Ip Cidr](https://www.alibabacloud.com/help/doc-detail/281972.htm).

        > **NOTE:** Available in v1.136.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example_zones = alicloud.get_zones(available_resource_creation="VSwitch")
        example_network = alicloud.vpc.Network("exampleNetwork",
            vpc_name="example_value",
            cidr_block="172.16.0.0/12")
        example_switch = alicloud.vpc.Switch("exampleSwitch",
            vpc_id=alicloud_vpc["default"]["id"],
            cidr_block="172.16.0.0/21",
            zone_id=example_zones.zones[0].id,
            vswitch_name=var["name"])
        example_nat_gateway = alicloud.vpc.NatGateway("exampleNatGateway",
            vpc_id=alicloud_vpc["default"]["id"],
            internet_charge_type="PayByLcu",
            nat_gateway_name="example_value",
            description="example_value",
            nat_type="Enhanced",
            vswitch_id=example_switch.id,
            network_type="intranet")
        example_nat_ip_cidr = alicloud.vpc.NatIpCidr("exampleNatIpCidr",
            nat_gateway_id=example_nat_gateway.id,
            nat_ip_cidr_name="example_value",
            nat_ip_cidr="example_value")
        ```

        ## Import

        VPC Nat Ip Cidr can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:vpc/natIpCidr:NatIpCidr example <nat_gateway_id>:<nat_ip_cidr>
        ```

        :param str resource_name: The name of the resource.
        :param NatIpCidrArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NatIpCidrArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dry_run: Optional[pulumi.Input[bool]] = None,
                 nat_gateway_id: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_description: Optional[pulumi.Input[str]] = None,
                 nat_ip_cidr_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NatIpCidrArgs.__new__(NatIpCidrArgs)

            __props__.__dict__["dry_run"] = dry_run
            if nat_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'nat_gateway_id'")
            __props__.__dict__["nat_gateway_id"] = nat_gateway_id
            __props__.__dict__["nat_ip_cidr"] = nat_ip_cidr
            __props__.__dict__["nat_ip_cidr_description"] = nat_ip_cidr_description
            __props__.__dict__["nat_ip_cidr_name"] = nat_ip_cidr_name
            __props__.__dict__["status"] = None
        super(NatIpCidr, __self__).__init__(
            'alicloud:vpc/natIpCidr:NatIpCidr',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dry_run: Optional[pulumi.Input[bool]] = None,
            nat_gateway_id: Optional[pulumi.Input[str]] = None,
            nat_ip_cidr: Optional[pulumi.Input[str]] = None,
            nat_ip_cidr_description: Optional[pulumi.Input[str]] = None,
            nat_ip_cidr_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'NatIpCidr':
        """
        Get an existing NatIpCidr resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] dry_run: Specifies whether to precheck this request only. Valid values: `true` and `false`.
        :param pulumi.Input[str] nat_gateway_id: The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        :param pulumi.Input[str] nat_ip_cidr: - The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        :param pulumi.Input[str] nat_ip_cidr_description: The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] nat_ip_cidr_name: The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        :param pulumi.Input[str] status: The status of the CIDR block of the NAT gateway. Valid values: `Available`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NatIpCidrState.__new__(_NatIpCidrState)

        __props__.__dict__["dry_run"] = dry_run
        __props__.__dict__["nat_gateway_id"] = nat_gateway_id
        __props__.__dict__["nat_ip_cidr"] = nat_ip_cidr
        __props__.__dict__["nat_ip_cidr_description"] = nat_ip_cidr_description
        __props__.__dict__["nat_ip_cidr_name"] = nat_ip_cidr_name
        __props__.__dict__["status"] = status
        return NatIpCidr(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dryRun")
    def dry_run(self) -> pulumi.Output[bool]:
        """
        Specifies whether to precheck this request only. Valid values: `true` and `false`.
        """
        return pulumi.get(self, "dry_run")

    @property
    @pulumi.getter(name="natGatewayId")
    def nat_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the Virtual Private Cloud (VPC) NAT gateway where you want to create the NAT CIDR block.
        """
        return pulumi.get(self, "nat_gateway_id")

    @property
    @pulumi.getter(name="natIpCidr")
    def nat_ip_cidr(self) -> pulumi.Output[Optional[str]]:
        """
        - The NAT CIDR block to be created. The CIDR block must meet the following conditions: It must be `10.0.0.0/8`, `172.16.0.0/12`, `192.168.0.0/16`, or one of their subnets. The subnet mask must be `16` to `32` bits in lengths. To use a public CIDR block as the NAT CIDR block, the VPC to which the VPC NAT gateway belongs must be authorized to use public CIDR blocks. For more information, see [Create a VPC NAT gateway](https://www.alibabacloud.com/help/doc-detail/268230.htm).
        """
        return pulumi.get(self, "nat_ip_cidr")

    @property
    @pulumi.getter(name="natIpCidrDescription")
    def nat_ip_cidr_description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the NAT CIDR block. The description must be `2` to `256` characters in length. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "nat_ip_cidr_description")

    @property
    @pulumi.getter(name="natIpCidrName")
    def nat_ip_cidr_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the NAT CIDR block. The name must be `2` to `128` characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It must start with a letter. It must start with a letter but cannot start with `http://` or `https://`.
        """
        return pulumi.get(self, "nat_ip_cidr_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the CIDR block of the NAT gateway. Valid values: `Available`.
        """
        return pulumi.get(self, "status")
