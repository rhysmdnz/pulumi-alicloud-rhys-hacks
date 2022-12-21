# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DomainAttachmentArgs', 'DomainAttachment']

@pulumi.input_type
class DomainAttachmentArgs:
    def __init__(__self__, *,
                 domain_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a DomainAttachment resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: The domain names bound to the DNS instance.
        :param pulumi.Input[str] instance_id: The id of the DNS instance.
        """
        pulumi.set(__self__, "domain_names", domain_names)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The domain names bound to the DNS instance.
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The id of the DNS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _DomainAttachmentState:
    def __init__(__self__, *,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DomainAttachment resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: The domain names bound to the DNS instance.
        :param pulumi.Input[str] instance_id: The id of the DNS instance.
        """
        if domain_names is not None:
            pulumi.set(__self__, "domain_names", domain_names)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The domain names bound to the DNS instance.
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the DNS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class DomainAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        dns = alicloud.dns.DomainAttachment("dns",
            domain_names=[
                "test111.abc",
                "test222.abc",
            ],
            instance_id="dns-cn-mp91lyq9xxxx")
        ```

        ## Import

        DNS domain attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:dns/domainAttachment:DomainAttachment example dns-cn-v0h1ldjhxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: The domain names bound to the DNS instance.
        :param pulumi.Input[str] instance_id: The id of the DNS instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        dns = alicloud.dns.DomainAttachment("dns",
            domain_names=[
                "test111.abc",
                "test222.abc",
            ],
            instance_id="dns-cn-mp91lyq9xxxx")
        ```

        ## Import

        DNS domain attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:dns/domainAttachment:DomainAttachment example dns-cn-v0h1ldjhxxx
        ```

        :param str resource_name: The name of the resource.
        :param DomainAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainAttachmentArgs.__new__(DomainAttachmentArgs)

            if domain_names is None and not opts.urn:
                raise TypeError("Missing required property 'domain_names'")
            __props__.__dict__["domain_names"] = domain_names
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(DomainAttachment, __self__).__init__(
            'alicloud:dns/domainAttachment:DomainAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'DomainAttachment':
        """
        Get an existing DomainAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: The domain names bound to the DNS instance.
        :param pulumi.Input[str] instance_id: The id of the DNS instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainAttachmentState.__new__(_DomainAttachmentState)

        __props__.__dict__["domain_names"] = domain_names
        __props__.__dict__["instance_id"] = instance_id
        return DomainAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> pulumi.Output[Sequence[str]]:
        """
        The domain names bound to the DNS instance.
        """
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The id of the DNS instance.
        """
        return pulumi.get(self, "instance_id")
