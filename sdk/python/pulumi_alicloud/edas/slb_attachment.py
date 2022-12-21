# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SlbAttachmentArgs', 'SlbAttachment']

@pulumi.input_type
class SlbAttachmentArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 slb_id: pulumi.Input[str],
                 slb_ip: pulumi.Input[str],
                 type: pulumi.Input[str],
                 listener_port: Optional[pulumi.Input[int]] = None,
                 vserver_group_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SlbAttachment resource.
        :param pulumi.Input[str] app_id: The ID of the application to which you want to bind an SLB instance.
        :param pulumi.Input[str] slb_id: The ID of the SLB instance that is going to be bound.
        :param pulumi.Input[str] slb_ip: The IP address that is allocated to the bound SLB instance.
        :param pulumi.Input[str] type: The type of the bound SLB instance.
        :param pulumi.Input[int] listener_port: The listening port for the bound SLB instance.
        :param pulumi.Input[str] vserver_group_id: The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "slb_id", slb_id)
        pulumi.set(__self__, "slb_ip", slb_ip)
        pulumi.set(__self__, "type", type)
        if listener_port is not None:
            pulumi.set(__self__, "listener_port", listener_port)
        if vserver_group_id is not None:
            pulumi.set(__self__, "vserver_group_id", vserver_group_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        """
        The ID of the application to which you want to bind an SLB instance.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="slbId")
    def slb_id(self) -> pulumi.Input[str]:
        """
        The ID of the SLB instance that is going to be bound.
        """
        return pulumi.get(self, "slb_id")

    @slb_id.setter
    def slb_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "slb_id", value)

    @property
    @pulumi.getter(name="slbIp")
    def slb_ip(self) -> pulumi.Input[str]:
        """
        The IP address that is allocated to the bound SLB instance.
        """
        return pulumi.get(self, "slb_ip")

    @slb_ip.setter
    def slb_ip(self, value: pulumi.Input[str]):
        pulumi.set(self, "slb_ip", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the bound SLB instance.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="listenerPort")
    def listener_port(self) -> Optional[pulumi.Input[int]]:
        """
        The listening port for the bound SLB instance.
        """
        return pulumi.get(self, "listener_port")

    @listener_port.setter
    def listener_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listener_port", value)

    @property
    @pulumi.getter(name="vserverGroupId")
    def vserver_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        """
        return pulumi.get(self, "vserver_group_id")

    @vserver_group_id.setter
    def vserver_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vserver_group_id", value)


@pulumi.input_type
class _SlbAttachmentState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[str]] = None,
                 listener_port: Optional[pulumi.Input[int]] = None,
                 slb_id: Optional[pulumi.Input[str]] = None,
                 slb_ip: Optional[pulumi.Input[str]] = None,
                 slb_status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vserver_group_id: Optional[pulumi.Input[str]] = None,
                 vswitch_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SlbAttachment resources.
        :param pulumi.Input[str] app_id: The ID of the application to which you want to bind an SLB instance.
        :param pulumi.Input[int] listener_port: The listening port for the bound SLB instance.
        :param pulumi.Input[str] slb_id: The ID of the SLB instance that is going to be bound.
        :param pulumi.Input[str] slb_ip: The IP address that is allocated to the bound SLB instance.
        :param pulumi.Input[str] slb_status: Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and forward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.
        :param pulumi.Input[str] type: The type of the bound SLB instance.
        :param pulumi.Input[str] vserver_group_id: The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        :param pulumi.Input[str] vswitch_id: VPC related vswitch ID.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if listener_port is not None:
            pulumi.set(__self__, "listener_port", listener_port)
        if slb_id is not None:
            pulumi.set(__self__, "slb_id", slb_id)
        if slb_ip is not None:
            pulumi.set(__self__, "slb_ip", slb_ip)
        if slb_status is not None:
            pulumi.set(__self__, "slb_status", slb_status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vserver_group_id is not None:
            pulumi.set(__self__, "vserver_group_id", vserver_group_id)
        if vswitch_id is not None:
            pulumi.set(__self__, "vswitch_id", vswitch_id)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the application to which you want to bind an SLB instance.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter(name="listenerPort")
    def listener_port(self) -> Optional[pulumi.Input[int]]:
        """
        The listening port for the bound SLB instance.
        """
        return pulumi.get(self, "listener_port")

    @listener_port.setter
    def listener_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "listener_port", value)

    @property
    @pulumi.getter(name="slbId")
    def slb_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the SLB instance that is going to be bound.
        """
        return pulumi.get(self, "slb_id")

    @slb_id.setter
    def slb_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slb_id", value)

    @property
    @pulumi.getter(name="slbIp")
    def slb_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address that is allocated to the bound SLB instance.
        """
        return pulumi.get(self, "slb_ip")

    @slb_ip.setter
    def slb_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slb_ip", value)

    @property
    @pulumi.getter(name="slbStatus")
    def slb_status(self) -> Optional[pulumi.Input[str]]:
        """
        Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and forward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.
        """
        return pulumi.get(self, "slb_status")

    @slb_status.setter
    def slb_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slb_status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the bound SLB instance.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vserverGroupId")
    def vserver_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        """
        return pulumi.get(self, "vserver_group_id")

    @vserver_group_id.setter
    def vserver_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vserver_group_id", value)

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC related vswitch ID.
        """
        return pulumi.get(self, "vswitch_id")

    @vswitch_id.setter
    def vswitch_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vswitch_id", value)


class SlbAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 listener_port: Optional[pulumi.Input[int]] = None,
                 slb_id: Optional[pulumi.Input[str]] = None,
                 slb_ip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vserver_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Binds SLB to an EDAS application.

        > **NOTE:** Available in 1.82.0+

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.edas.SlbAttachment("default",
            app_id=var["app_id"],
            slb_id=var["slb_id"],
            slb_ip=var["slb_ip"],
            type=var["type"],
            listener_port=var["listener_port"],
            vserver_group_id=var["vserver_group_id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the application to which you want to bind an SLB instance.
        :param pulumi.Input[int] listener_port: The listening port for the bound SLB instance.
        :param pulumi.Input[str] slb_id: The ID of the SLB instance that is going to be bound.
        :param pulumi.Input[str] slb_ip: The IP address that is allocated to the bound SLB instance.
        :param pulumi.Input[str] type: The type of the bound SLB instance.
        :param pulumi.Input[str] vserver_group_id: The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SlbAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Binds SLB to an EDAS application.

        > **NOTE:** Available in 1.82.0+

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.edas.SlbAttachment("default",
            app_id=var["app_id"],
            slb_id=var["slb_id"],
            slb_ip=var["slb_ip"],
            type=var["type"],
            listener_port=var["listener_port"],
            vserver_group_id=var["vserver_group_id"])
        ```

        :param str resource_name: The name of the resource.
        :param SlbAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SlbAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 listener_port: Optional[pulumi.Input[int]] = None,
                 slb_id: Optional[pulumi.Input[str]] = None,
                 slb_ip: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 vserver_group_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SlbAttachmentArgs.__new__(SlbAttachmentArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["listener_port"] = listener_port
            if slb_id is None and not opts.urn:
                raise TypeError("Missing required property 'slb_id'")
            __props__.__dict__["slb_id"] = slb_id
            if slb_ip is None and not opts.urn:
                raise TypeError("Missing required property 'slb_ip'")
            __props__.__dict__["slb_ip"] = slb_ip
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["vserver_group_id"] = vserver_group_id
            __props__.__dict__["slb_status"] = None
            __props__.__dict__["vswitch_id"] = None
        super(SlbAttachment, __self__).__init__(
            'alicloud:edas/slbAttachment:SlbAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[str]] = None,
            listener_port: Optional[pulumi.Input[int]] = None,
            slb_id: Optional[pulumi.Input[str]] = None,
            slb_ip: Optional[pulumi.Input[str]] = None,
            slb_status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            vserver_group_id: Optional[pulumi.Input[str]] = None,
            vswitch_id: Optional[pulumi.Input[str]] = None) -> 'SlbAttachment':
        """
        Get an existing SlbAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] app_id: The ID of the application to which you want to bind an SLB instance.
        :param pulumi.Input[int] listener_port: The listening port for the bound SLB instance.
        :param pulumi.Input[str] slb_id: The ID of the SLB instance that is going to be bound.
        :param pulumi.Input[str] slb_ip: The IP address that is allocated to the bound SLB instance.
        :param pulumi.Input[str] slb_status: Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and forward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.
        :param pulumi.Input[str] type: The type of the bound SLB instance.
        :param pulumi.Input[str] vserver_group_id: The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        :param pulumi.Input[str] vswitch_id: VPC related vswitch ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SlbAttachmentState.__new__(_SlbAttachmentState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["listener_port"] = listener_port
        __props__.__dict__["slb_id"] = slb_id
        __props__.__dict__["slb_ip"] = slb_ip
        __props__.__dict__["slb_status"] = slb_status
        __props__.__dict__["type"] = type
        __props__.__dict__["vserver_group_id"] = vserver_group_id
        __props__.__dict__["vswitch_id"] = vswitch_id
        return SlbAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        """
        The ID of the application to which you want to bind an SLB instance.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="listenerPort")
    def listener_port(self) -> pulumi.Output[Optional[int]]:
        """
        The listening port for the bound SLB instance.
        """
        return pulumi.get(self, "listener_port")

    @property
    @pulumi.getter(name="slbId")
    def slb_id(self) -> pulumi.Output[str]:
        """
        The ID of the SLB instance that is going to be bound.
        """
        return pulumi.get(self, "slb_id")

    @property
    @pulumi.getter(name="slbIp")
    def slb_ip(self) -> pulumi.Output[str]:
        """
        The IP address that is allocated to the bound SLB instance.
        """
        return pulumi.get(self, "slb_ip")

    @property
    @pulumi.getter(name="slbStatus")
    def slb_status(self) -> pulumi.Output[str]:
        """
        Running Status of SLB instance. Inactive：The instance is stopped, and listener will not monitor and forward traffic. Active：The instance is running. After the instance is created, the default state is active. Locked：The instance is locked, the instance has been owed or locked by Alibaba Cloud. Expired: The instance has expired.
        """
        return pulumi.get(self, "slb_status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the bound SLB instance.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vserverGroupId")
    def vserver_group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the virtual server (VServer) group associated with the intranet SLB instance.
        """
        return pulumi.get(self, "vserver_group_id")

    @property
    @pulumi.getter(name="vswitchId")
    def vswitch_id(self) -> pulumi.Output[str]:
        """
        VPC related vswitch ID.
        """
        return pulumi.get(self, "vswitch_id")
