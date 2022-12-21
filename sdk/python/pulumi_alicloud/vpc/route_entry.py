# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RouteEntryArgs', 'RouteEntry']

@pulumi.input_type
class RouteEntryArgs:
    def __init__(__self__, *,
                 route_table_id: pulumi.Input[str],
                 destination_cidrblock: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nexthop_id: Optional[pulumi.Input[str]] = None,
                 nexthop_type: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RouteEntry resource.
        :param pulumi.Input[str] route_table_id: The ID of the route table.
        :param pulumi.Input[str] destination_cidrblock: The RouteEntry's target network segment.
        :param pulumi.Input[str] name: The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] nexthop_id: The route entry's next hop. ECS instance ID or VPC router interface ID.
        :param pulumi.Input[str] nexthop_type: The next hop type. Available values:
        :param pulumi.Input[str] router_id: This argument has been deprecated. Please use other arguments to launch a custom route entry.
        """
        pulumi.set(__self__, "route_table_id", route_table_id)
        if destination_cidrblock is not None:
            pulumi.set(__self__, "destination_cidrblock", destination_cidrblock)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nexthop_id is not None:
            pulumi.set(__self__, "nexthop_id", nexthop_id)
        if nexthop_type is not None:
            pulumi.set(__self__, "nexthop_type", nexthop_type)
        if router_id is not None:
            warnings.warn("""Attribute router_id has been deprecated and suggest removing it from your template.""", DeprecationWarning)
            pulumi.log.warn("""router_id is deprecated: Attribute router_id has been deprecated and suggest removing it from your template.""")
        if router_id is not None:
            pulumi.set(__self__, "router_id", router_id)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Input[str]:
        """
        The ID of the route table.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="destinationCidrblock")
    def destination_cidrblock(self) -> Optional[pulumi.Input[str]]:
        """
        The RouteEntry's target network segment.
        """
        return pulumi.get(self, "destination_cidrblock")

    @destination_cidrblock.setter
    def destination_cidrblock(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidrblock", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nexthopId")
    def nexthop_id(self) -> Optional[pulumi.Input[str]]:
        """
        The route entry's next hop. ECS instance ID or VPC router interface ID.
        """
        return pulumi.get(self, "nexthop_id")

    @nexthop_id.setter
    def nexthop_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nexthop_id", value)

    @property
    @pulumi.getter(name="nexthopType")
    def nexthop_type(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop type. Available values:
        """
        return pulumi.get(self, "nexthop_type")

    @nexthop_type.setter
    def nexthop_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nexthop_type", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[pulumi.Input[str]]:
        """
        This argument has been deprecated. Please use other arguments to launch a custom route entry.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router_id", value)


@pulumi.input_type
class _RouteEntryState:
    def __init__(__self__, *,
                 destination_cidrblock: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nexthop_id: Optional[pulumi.Input[str]] = None,
                 nexthop_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouteEntry resources.
        :param pulumi.Input[str] destination_cidrblock: The RouteEntry's target network segment.
        :param pulumi.Input[str] name: The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] nexthop_id: The route entry's next hop. ECS instance ID or VPC router interface ID.
        :param pulumi.Input[str] nexthop_type: The next hop type. Available values:
        :param pulumi.Input[str] route_table_id: The ID of the route table.
        :param pulumi.Input[str] router_id: This argument has been deprecated. Please use other arguments to launch a custom route entry.
        """
        if destination_cidrblock is not None:
            pulumi.set(__self__, "destination_cidrblock", destination_cidrblock)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nexthop_id is not None:
            pulumi.set(__self__, "nexthop_id", nexthop_id)
        if nexthop_type is not None:
            pulumi.set(__self__, "nexthop_type", nexthop_type)
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)
        if router_id is not None:
            warnings.warn("""Attribute router_id has been deprecated and suggest removing it from your template.""", DeprecationWarning)
            pulumi.log.warn("""router_id is deprecated: Attribute router_id has been deprecated and suggest removing it from your template.""")
        if router_id is not None:
            pulumi.set(__self__, "router_id", router_id)

    @property
    @pulumi.getter(name="destinationCidrblock")
    def destination_cidrblock(self) -> Optional[pulumi.Input[str]]:
        """
        The RouteEntry's target network segment.
        """
        return pulumi.get(self, "destination_cidrblock")

    @destination_cidrblock.setter
    def destination_cidrblock(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidrblock", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nexthopId")
    def nexthop_id(self) -> Optional[pulumi.Input[str]]:
        """
        The route entry's next hop. ECS instance ID or VPC router interface ID.
        """
        return pulumi.get(self, "nexthop_id")

    @nexthop_id.setter
    def nexthop_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nexthop_id", value)

    @property
    @pulumi.getter(name="nexthopType")
    def nexthop_type(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop type. Available values:
        """
        return pulumi.get(self, "nexthop_type")

    @nexthop_type.setter
    def nexthop_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nexthop_type", value)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the route table.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> Optional[pulumi.Input[str]]:
        """
        This argument has been deprecated. Please use other arguments to launch a custom route entry.
        """
        return pulumi.get(self, "router_id")

    @router_id.setter
    def router_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "router_id", value)


class RouteEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidrblock: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nexthop_id: Optional[pulumi.Input[str]] = None,
                 nexthop_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Router entry can be imported using the id, e.g (formatted as<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>).

        ```sh
         $ pulumi import alicloud:vpc/routeEntry:RouteEntry example vtb-123456:vrt-123456:0.0.0.0/0:NatGateway:ngw-123456
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidrblock: The RouteEntry's target network segment.
        :param pulumi.Input[str] name: The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] nexthop_id: The route entry's next hop. ECS instance ID or VPC router interface ID.
        :param pulumi.Input[str] nexthop_type: The next hop type. Available values:
        :param pulumi.Input[str] route_table_id: The ID of the route table.
        :param pulumi.Input[str] router_id: This argument has been deprecated. Please use other arguments to launch a custom route entry.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Router entry can be imported using the id, e.g (formatted as<route_table_id:router_id:destination_cidrblock:nexthop_type:nexthop_id>).

        ```sh
         $ pulumi import alicloud:vpc/routeEntry:RouteEntry example vtb-123456:vrt-123456:0.0.0.0/0:NatGateway:ngw-123456
        ```

        :param str resource_name: The name of the resource.
        :param RouteEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidrblock: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nexthop_id: Optional[pulumi.Input[str]] = None,
                 nexthop_type: Optional[pulumi.Input[str]] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 router_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RouteEntryArgs.__new__(RouteEntryArgs)

            __props__.__dict__["destination_cidrblock"] = destination_cidrblock
            __props__.__dict__["name"] = name
            __props__.__dict__["nexthop_id"] = nexthop_id
            __props__.__dict__["nexthop_type"] = nexthop_type
            if route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_id'")
            __props__.__dict__["route_table_id"] = route_table_id
            if router_id is not None and not opts.urn:
                warnings.warn("""Attribute router_id has been deprecated and suggest removing it from your template.""", DeprecationWarning)
                pulumi.log.warn("""router_id is deprecated: Attribute router_id has been deprecated and suggest removing it from your template.""")
            __props__.__dict__["router_id"] = router_id
        super(RouteEntry, __self__).__init__(
            'alicloud:vpc/routeEntry:RouteEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_cidrblock: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nexthop_id: Optional[pulumi.Input[str]] = None,
            nexthop_type: Optional[pulumi.Input[str]] = None,
            route_table_id: Optional[pulumi.Input[str]] = None,
            router_id: Optional[pulumi.Input[str]] = None) -> 'RouteEntry':
        """
        Get an existing RouteEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidrblock: The RouteEntry's target network segment.
        :param pulumi.Input[str] name: The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        :param pulumi.Input[str] nexthop_id: The route entry's next hop. ECS instance ID or VPC router interface ID.
        :param pulumi.Input[str] nexthop_type: The next hop type. Available values:
        :param pulumi.Input[str] route_table_id: The ID of the route table.
        :param pulumi.Input[str] router_id: This argument has been deprecated. Please use other arguments to launch a custom route entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouteEntryState.__new__(_RouteEntryState)

        __props__.__dict__["destination_cidrblock"] = destination_cidrblock
        __props__.__dict__["name"] = name
        __props__.__dict__["nexthop_id"] = nexthop_id
        __props__.__dict__["nexthop_type"] = nexthop_type
        __props__.__dict__["route_table_id"] = route_table_id
        __props__.__dict__["router_id"] = router_id
        return RouteEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationCidrblock")
    def destination_cidrblock(self) -> pulumi.Output[Optional[str]]:
        """
        The RouteEntry's target network segment.
        """
        return pulumi.get(self, "destination_cidrblock")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the route entry. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nexthopId")
    def nexthop_id(self) -> pulumi.Output[Optional[str]]:
        """
        The route entry's next hop. ECS instance ID or VPC router interface ID.
        """
        return pulumi.get(self, "nexthop_id")

    @property
    @pulumi.getter(name="nexthopType")
    def nexthop_type(self) -> pulumi.Output[Optional[str]]:
        """
        The next hop type. Available values:
        """
        return pulumi.get(self, "nexthop_type")

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the route table.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter(name="routerId")
    def router_id(self) -> pulumi.Output[str]:
        """
        This argument has been deprecated. Please use other arguments to launch a custom route entry.
        """
        return pulumi.get(self, "router_id")
