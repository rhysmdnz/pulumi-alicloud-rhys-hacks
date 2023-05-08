# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServerGroupServerAttachmentArgs', 'ServerGroupServerAttachment']

@pulumi.input_type
class ServerGroupServerAttachmentArgs:
    def __init__(__self__, *,
                 port: pulumi.Input[int],
                 server_group_id: pulumi.Input[str],
                 server_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ServerGroupServerAttachment resource.
        :param pulumi.Input[int] port: The port that is used by the backend server. Valid values: `1` to `65535`.
        :param pulumi.Input[str] server_group_id: The ID of the server group.
        :param pulumi.Input[str] server_id: The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
        :param pulumi.Input[str] description: The description of the backend server.
        :param pulumi.Input[str] type: The type of backend server. Valid values: `ecs`, `eni`.
        :param pulumi.Input[int] weight: The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
        """
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "server_group_id", server_group_id)
        pulumi.set(__self__, "server_id", server_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port that is used by the backend server. Valid values: `1` to `65535`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverGroupId")
    def server_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the server group.
        """
        return pulumi.get(self, "server_group_id")

    @server_group_id.setter
    def server_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_group_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Input[str]:
        """
        The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the backend server.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of backend server. Valid values: `ecs`, `eni`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _ServerGroupServerAttachmentState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_group_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ServerGroupServerAttachment resources.
        :param pulumi.Input[str] description: The description of the backend server.
        :param pulumi.Input[int] port: The port that is used by the backend server. Valid values: `1` to `65535`.
        :param pulumi.Input[str] server_group_id: The ID of the server group.
        :param pulumi.Input[str] server_id: The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
        :param pulumi.Input[str] type: The type of backend server. Valid values: `ecs`, `eni`.
        :param pulumi.Input[int] weight: The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_group_id is not None:
            pulumi.set(__self__, "server_group_id", server_group_id)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the backend server.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port that is used by the backend server. Valid values: `1` to `65535`.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverGroupId")
    def server_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the server group.
        """
        return pulumi.get(self, "server_group_id")

    @server_group_id.setter
    def server_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_group_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of backend server. Valid values: `ecs`, `eni`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class ServerGroupServerAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_group_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        > **NOTE:** Available in v1.163.0+.

        For information about server group server attachment and how to use it, see [Configure a server group server attachment](https://www.alibabacloud.com/help/en/doc-detail/35218.html).

        > **NOTE:** Applying this resource may conflict with applying `slb.Listener`,
        and the `slb.Listener` block should use `depends_on = [alicloud_slb_server_group_server_attachment.xxx]` to avoid it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        slb_server_group_server_attachment = config.get("slbServerGroupServerAttachment")
        if slb_server_group_server_attachment is None:
            slb_server_group_server_attachment = "forSlbServerGroupServerAttachment"
        slb_server_group_server_attachment_count = config.get_float("slbServerGroupServerAttachmentCount")
        if slb_server_group_server_attachment_count is None:
            slb_server_group_server_attachment_count = 5
        server_attachment_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        server_attachment_instance_types = alicloud.ecs.get_instance_types(availability_zone=server_attachment_zones.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        server_attachment_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            most_recent=True,
            owners="system")
        server_attachment_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
        server_attachment_switches = alicloud.vpc.get_switches(vpc_id=server_attachment_networks.ids[0],
            zone_id=server_attachment_zones.zones[0].id)
        server_attachment_security_group = alicloud.ecs.SecurityGroup("serverAttachmentSecurityGroup", vpc_id=server_attachment_networks.ids[0])
        server_attachment_instance = []
        for range in [{"value": i} for i in range(0, slb_server_group_server_attachment_count)]:
            server_attachment_instance.append(alicloud.ecs.Instance(f"serverAttachmentInstance-{range['value']}",
                image_id=server_attachment_images.images[0].id,
                instance_type=server_attachment_instance_types.instance_types[0].id,
                instance_name=slb_server_group_server_attachment,
                security_groups=[__item.id for __item in [server_attachment_security_group]],
                internet_charge_type="PayByTraffic",
                internet_max_bandwidth_out=10,
                availability_zone=server_attachment_zones.zones[0].id,
                instance_charge_type="PostPaid",
                system_disk_category="cloud_efficiency",
                vswitch_id=server_attachment_switches.ids[0]))
        server_attachment_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("serverAttachmentApplicationLoadBalancer",
            load_balancer_name=slb_server_group_server_attachment,
            vswitch_id=server_attachment_switches.vswitches[0].id,
            load_balancer_spec="slb.s2.small",
            address_type="intranet")
        server_attachment_server_group = alicloud.slb.ServerGroup("serverAttachmentServerGroup", load_balancer_id=server_attachment_application_load_balancer.id)
        server_attachment_server_group_server_attachment = []
        for range in [{"value": i} for i in range(0, slb_server_group_server_attachment_count)]:
            server_attachment_server_group_server_attachment.append(alicloud.slb.ServerGroupServerAttachment(f"serverAttachmentServerGroupServerAttachment-{range['value']}",
                server_group_id=server_attachment_server_group.id,
                server_id=server_attachment_instance[range["index"]].id,
                port=8080,
                weight=0))
        server_attachment_listener = alicloud.slb.Listener("serverAttachmentListener",
            load_balancer_id=server_attachment_application_load_balancer.id,
            backend_port=80,
            frontend_port=80,
            protocol="tcp",
            bandwidth=10,
            scheduler="rr",
            server_group_id=server_attachment_server_group.id,
            opts=pulumi.ResourceOptions(depends_on=[server_attachment_server_group_server_attachment]))
        ```

        ## Import

        Load balancer backend server group server attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/serverGroupServerAttachment:ServerGroupServerAttachment example <server_group_id>:<server_id>:<port>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the backend server.
        :param pulumi.Input[int] port: The port that is used by the backend server. Valid values: `1` to `65535`.
        :param pulumi.Input[str] server_group_id: The ID of the server group.
        :param pulumi.Input[str] server_id: The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
        :param pulumi.Input[str] type: The type of backend server. Valid values: `ecs`, `eni`.
        :param pulumi.Input[int] weight: The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerGroupServerAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        > **NOTE:** Available in v1.163.0+.

        For information about server group server attachment and how to use it, see [Configure a server group server attachment](https://www.alibabacloud.com/help/en/doc-detail/35218.html).

        > **NOTE:** Applying this resource may conflict with applying `slb.Listener`,
        and the `slb.Listener` block should use `depends_on = [alicloud_slb_server_group_server_attachment.xxx]` to avoid it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        config = pulumi.Config()
        slb_server_group_server_attachment = config.get("slbServerGroupServerAttachment")
        if slb_server_group_server_attachment is None:
            slb_server_group_server_attachment = "forSlbServerGroupServerAttachment"
        slb_server_group_server_attachment_count = config.get_float("slbServerGroupServerAttachmentCount")
        if slb_server_group_server_attachment_count is None:
            slb_server_group_server_attachment_count = 5
        server_attachment_zones = alicloud.get_zones(available_disk_category="cloud_efficiency",
            available_resource_creation="VSwitch")
        server_attachment_instance_types = alicloud.ecs.get_instance_types(availability_zone=server_attachment_zones.zones[0].id,
            cpu_core_count=1,
            memory_size=2)
        server_attachment_images = alicloud.ecs.get_images(name_regex="^ubuntu_[0-9]+_[0-9]+_x64*",
            most_recent=True,
            owners="system")
        server_attachment_networks = alicloud.vpc.get_networks(name_regex="default-NODELETING")
        server_attachment_switches = alicloud.vpc.get_switches(vpc_id=server_attachment_networks.ids[0],
            zone_id=server_attachment_zones.zones[0].id)
        server_attachment_security_group = alicloud.ecs.SecurityGroup("serverAttachmentSecurityGroup", vpc_id=server_attachment_networks.ids[0])
        server_attachment_instance = []
        for range in [{"value": i} for i in range(0, slb_server_group_server_attachment_count)]:
            server_attachment_instance.append(alicloud.ecs.Instance(f"serverAttachmentInstance-{range['value']}",
                image_id=server_attachment_images.images[0].id,
                instance_type=server_attachment_instance_types.instance_types[0].id,
                instance_name=slb_server_group_server_attachment,
                security_groups=[__item.id for __item in [server_attachment_security_group]],
                internet_charge_type="PayByTraffic",
                internet_max_bandwidth_out=10,
                availability_zone=server_attachment_zones.zones[0].id,
                instance_charge_type="PostPaid",
                system_disk_category="cloud_efficiency",
                vswitch_id=server_attachment_switches.ids[0]))
        server_attachment_application_load_balancer = alicloud.slb.ApplicationLoadBalancer("serverAttachmentApplicationLoadBalancer",
            load_balancer_name=slb_server_group_server_attachment,
            vswitch_id=server_attachment_switches.vswitches[0].id,
            load_balancer_spec="slb.s2.small",
            address_type="intranet")
        server_attachment_server_group = alicloud.slb.ServerGroup("serverAttachmentServerGroup", load_balancer_id=server_attachment_application_load_balancer.id)
        server_attachment_server_group_server_attachment = []
        for range in [{"value": i} for i in range(0, slb_server_group_server_attachment_count)]:
            server_attachment_server_group_server_attachment.append(alicloud.slb.ServerGroupServerAttachment(f"serverAttachmentServerGroupServerAttachment-{range['value']}",
                server_group_id=server_attachment_server_group.id,
                server_id=server_attachment_instance[range["index"]].id,
                port=8080,
                weight=0))
        server_attachment_listener = alicloud.slb.Listener("serverAttachmentListener",
            load_balancer_id=server_attachment_application_load_balancer.id,
            backend_port=80,
            frontend_port=80,
            protocol="tcp",
            bandwidth=10,
            scheduler="rr",
            server_group_id=server_attachment_server_group.id,
            opts=pulumi.ResourceOptions(depends_on=[server_attachment_server_group_server_attachment]))
        ```

        ## Import

        Load balancer backend server group server attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:slb/serverGroupServerAttachment:ServerGroupServerAttachment example <server_group_id>:<server_id>:<port>
        ```

        :param str resource_name: The name of the resource.
        :param ServerGroupServerAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerGroupServerAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_group_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerGroupServerAttachmentArgs.__new__(ServerGroupServerAttachmentArgs)

            __props__.__dict__["description"] = description
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if server_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_group_id'")
            __props__.__dict__["server_group_id"] = server_group_id
            if server_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_id'")
            __props__.__dict__["server_id"] = server_id
            __props__.__dict__["type"] = type
            __props__.__dict__["weight"] = weight
        super(ServerGroupServerAttachment, __self__).__init__(
            'alicloud:slb/serverGroupServerAttachment:ServerGroupServerAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            server_group_id: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[int]] = None) -> 'ServerGroupServerAttachment':
        """
        Get an existing ServerGroupServerAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the backend server.
        :param pulumi.Input[int] port: The port that is used by the backend server. Valid values: `1` to `65535`.
        :param pulumi.Input[str] server_group_id: The ID of the server group.
        :param pulumi.Input[str] server_id: The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
        :param pulumi.Input[str] type: The type of backend server. Valid values: `ecs`, `eni`.
        :param pulumi.Input[int] weight: The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerGroupServerAttachmentState.__new__(_ServerGroupServerAttachmentState)

        __props__.__dict__["description"] = description
        __props__.__dict__["port"] = port
        __props__.__dict__["server_group_id"] = server_group_id
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["type"] = type
        __props__.__dict__["weight"] = weight
        return ServerGroupServerAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the backend server.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port that is used by the backend server. Valid values: `1` to `65535`.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="serverGroupId")
    def server_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the server group.
        """
        return pulumi.get(self, "server_group_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The ID of the backend server. You can specify the ID of an Elastic Compute Service (ECS) instance or an elastic network interface (ENI).
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of backend server. Valid values: `ecs`, `eni`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[int]:
        """
        The weight of the backend server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no requests are forwarded to the backend server.
        """
        return pulumi.get(self, "weight")

