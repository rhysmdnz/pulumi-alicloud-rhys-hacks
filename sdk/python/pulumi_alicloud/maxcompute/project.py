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
from ._inputs import *

__all__ = ['ProjectArgs', 'Project']

@pulumi.input_type
class ProjectArgs:
    def __init__(__self__, *,
                 project_name: pulumi.Input[str],
                 comment: Optional[pulumi.Input[str]] = None,
                 default_quota: Optional[pulumi.Input[str]] = None,
                 ip_white_list: Optional[pulumi.Input['ProjectIpWhiteListArgs']] = None,
                 product_type: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input['ProjectPropertiesArgs']] = None,
                 security_properties: Optional[pulumi.Input['ProjectSecurityPropertiesArgs']] = None):
        """
        The set of arguments for constructing a Project resource.
        :param pulumi.Input[str] project_name: The name of the project
        :param pulumi.Input[str] comment: Comments of project
        :param pulumi.Input[str] default_quota: Default Computing Resource Group
        :param pulumi.Input['ProjectIpWhiteListArgs'] ip_white_list: IP whitelistSee the following `Block IpWhiteList`.
        :param pulumi.Input[str] product_type: Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
        :param pulumi.Input['ProjectPropertiesArgs'] properties: Project base attributesSee the following `Block Properties`.
        :param pulumi.Input['ProjectSecurityPropertiesArgs'] security_properties: Security-related attributesSee the following `Block SecurityProperties`.
        """
        pulumi.set(__self__, "project_name", project_name)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_quota is not None:
            pulumi.set(__self__, "default_quota", default_quota)
        if ip_white_list is not None:
            pulumi.set(__self__, "ip_white_list", ip_white_list)
        if product_type is not None:
            pulumi.set(__self__, "product_type", product_type)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if security_properties is not None:
            pulumi.set(__self__, "security_properties", security_properties)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Input[str]:
        """
        The name of the project
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comments of project
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> Optional[pulumi.Input[str]]:
        """
        Default Computing Resource Group
        """
        return pulumi.get(self, "default_quota")

    @default_quota.setter
    def default_quota(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_quota", value)

    @property
    @pulumi.getter(name="ipWhiteList")
    def ip_white_list(self) -> Optional[pulumi.Input['ProjectIpWhiteListArgs']]:
        """
        IP whitelistSee the following `Block IpWhiteList`.
        """
        return pulumi.get(self, "ip_white_list")

    @ip_white_list.setter
    def ip_white_list(self, value: Optional[pulumi.Input['ProjectIpWhiteListArgs']]):
        pulumi.set(self, "ip_white_list", value)

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> Optional[pulumi.Input[str]]:
        """
        Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
        """
        return pulumi.get(self, "product_type")

    @product_type.setter
    def product_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_type", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['ProjectPropertiesArgs']]:
        """
        Project base attributesSee the following `Block Properties`.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['ProjectPropertiesArgs']]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter(name="securityProperties")
    def security_properties(self) -> Optional[pulumi.Input['ProjectSecurityPropertiesArgs']]:
        """
        Security-related attributesSee the following `Block SecurityProperties`.
        """
        return pulumi.get(self, "security_properties")

    @security_properties.setter
    def security_properties(self, value: Optional[pulumi.Input['ProjectSecurityPropertiesArgs']]):
        pulumi.set(self, "security_properties", value)


@pulumi.input_type
class _ProjectState:
    def __init__(__self__, *,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_quota: Optional[pulumi.Input[str]] = None,
                 ip_white_list: Optional[pulumi.Input['ProjectIpWhiteListArgs']] = None,
                 owner: Optional[pulumi.Input[str]] = None,
                 product_type: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input['ProjectPropertiesArgs']] = None,
                 security_properties: Optional[pulumi.Input['ProjectSecurityPropertiesArgs']] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Project resources.
        :param pulumi.Input[str] comment: Comments of project
        :param pulumi.Input[str] default_quota: Default Computing Resource Group
        :param pulumi.Input['ProjectIpWhiteListArgs'] ip_white_list: IP whitelistSee the following `Block IpWhiteList`.
        :param pulumi.Input[str] owner: Project owner
        :param pulumi.Input[str] product_type: Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
        :param pulumi.Input[str] project_name: The name of the project
        :param pulumi.Input['ProjectPropertiesArgs'] properties: Project base attributesSee the following `Block Properties`.
        :param pulumi.Input['ProjectSecurityPropertiesArgs'] security_properties: Security-related attributesSee the following `Block SecurityProperties`.
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] type: Life cycle type.
        """
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if default_quota is not None:
            pulumi.set(__self__, "default_quota", default_quota)
        if ip_white_list is not None:
            pulumi.set(__self__, "ip_white_list", ip_white_list)
        if owner is not None:
            pulumi.set(__self__, "owner", owner)
        if product_type is not None:
            pulumi.set(__self__, "product_type", product_type)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if security_properties is not None:
            pulumi.set(__self__, "security_properties", security_properties)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def comment(self) -> Optional[pulumi.Input[str]]:
        """
        Comments of project
        """
        return pulumi.get(self, "comment")

    @comment.setter
    def comment(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "comment", value)

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> Optional[pulumi.Input[str]]:
        """
        Default Computing Resource Group
        """
        return pulumi.get(self, "default_quota")

    @default_quota.setter
    def default_quota(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_quota", value)

    @property
    @pulumi.getter(name="ipWhiteList")
    def ip_white_list(self) -> Optional[pulumi.Input['ProjectIpWhiteListArgs']]:
        """
        IP whitelistSee the following `Block IpWhiteList`.
        """
        return pulumi.get(self, "ip_white_list")

    @ip_white_list.setter
    def ip_white_list(self, value: Optional[pulumi.Input['ProjectIpWhiteListArgs']]):
        pulumi.set(self, "ip_white_list", value)

    @property
    @pulumi.getter
    def owner(self) -> Optional[pulumi.Input[str]]:
        """
        Project owner
        """
        return pulumi.get(self, "owner")

    @owner.setter
    def owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner", value)

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> Optional[pulumi.Input[str]]:
        """
        Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
        """
        return pulumi.get(self, "product_type")

    @product_type.setter
    def product_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_type", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the project
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input['ProjectPropertiesArgs']]:
        """
        Project base attributesSee the following `Block Properties`.
        """
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input['ProjectPropertiesArgs']]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter(name="securityProperties")
    def security_properties(self) -> Optional[pulumi.Input['ProjectSecurityPropertiesArgs']]:
        """
        Security-related attributesSee the following `Block SecurityProperties`.
        """
        return pulumi.get(self, "security_properties")

    @security_properties.setter
    def security_properties(self, value: Optional[pulumi.Input['ProjectSecurityPropertiesArgs']]):
        pulumi.set(self, "security_properties", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Life cycle type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Project(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_quota: Optional[pulumi.Input[str]] = None,
                 ip_white_list: Optional[pulumi.Input[pulumi.InputType['ProjectIpWhiteListArgs']]] = None,
                 product_type: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[pulumi.InputType['ProjectPropertiesArgs']]] = None,
                 security_properties: Optional[pulumi.Input[pulumi.InputType['ProjectSecurityPropertiesArgs']]] = None,
                 __props__=None):
        """
        Provides a Max Compute Project resource.

        For information about Max Compute Project and how to use it, see [What is Project](https://help.aliyun.com/document_detail/473237.html).

        > **NOTE:** Available in v1.77.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.maxcompute.Project("default",
            comment="test_for_terraform",
            default_quota="默认后付费Quota",
            product_type="PAYASYOUGO",
            project_name="test_create_spec_one")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comments of project
        :param pulumi.Input[str] default_quota: Default Computing Resource Group
        :param pulumi.Input[pulumi.InputType['ProjectIpWhiteListArgs']] ip_white_list: IP whitelistSee the following `Block IpWhiteList`.
        :param pulumi.Input[str] product_type: Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
        :param pulumi.Input[str] project_name: The name of the project
        :param pulumi.Input[pulumi.InputType['ProjectPropertiesArgs']] properties: Project base attributesSee the following `Block Properties`.
        :param pulumi.Input[pulumi.InputType['ProjectSecurityPropertiesArgs']] security_properties: Security-related attributesSee the following `Block SecurityProperties`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProjectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Max Compute Project resource.

        For information about Max Compute Project and how to use it, see [What is Project](https://help.aliyun.com/document_detail/473237.html).

        > **NOTE:** Available in v1.77.0+.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        default = alicloud.maxcompute.Project("default",
            comment="test_for_terraform",
            default_quota="默认后付费Quota",
            product_type="PAYASYOUGO",
            project_name="test_create_spec_one")
        ```

        :param str resource_name: The name of the resource.
        :param ProjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 comment: Optional[pulumi.Input[str]] = None,
                 default_quota: Optional[pulumi.Input[str]] = None,
                 ip_white_list: Optional[pulumi.Input[pulumi.InputType['ProjectIpWhiteListArgs']]] = None,
                 product_type: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[pulumi.InputType['ProjectPropertiesArgs']]] = None,
                 security_properties: Optional[pulumi.Input[pulumi.InputType['ProjectSecurityPropertiesArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProjectArgs.__new__(ProjectArgs)

            __props__.__dict__["comment"] = comment
            __props__.__dict__["default_quota"] = default_quota
            __props__.__dict__["ip_white_list"] = ip_white_list
            __props__.__dict__["product_type"] = product_type
            if project_name is None and not opts.urn:
                raise TypeError("Missing required property 'project_name'")
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["properties"] = properties
            __props__.__dict__["security_properties"] = security_properties
            __props__.__dict__["owner"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["type"] = None
        super(Project, __self__).__init__(
            'alicloud:maxcompute/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            comment: Optional[pulumi.Input[str]] = None,
            default_quota: Optional[pulumi.Input[str]] = None,
            ip_white_list: Optional[pulumi.Input[pulumi.InputType['ProjectIpWhiteListArgs']]] = None,
            owner: Optional[pulumi.Input[str]] = None,
            product_type: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            properties: Optional[pulumi.Input[pulumi.InputType['ProjectPropertiesArgs']]] = None,
            security_properties: Optional[pulumi.Input[pulumi.InputType['ProjectSecurityPropertiesArgs']]] = None,
            status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] comment: Comments of project
        :param pulumi.Input[str] default_quota: Default Computing Resource Group
        :param pulumi.Input[pulumi.InputType['ProjectIpWhiteListArgs']] ip_white_list: IP whitelistSee the following `Block IpWhiteList`.
        :param pulumi.Input[str] owner: Project owner
        :param pulumi.Input[str] product_type: Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
        :param pulumi.Input[str] project_name: The name of the project
        :param pulumi.Input[pulumi.InputType['ProjectPropertiesArgs']] properties: Project base attributesSee the following `Block Properties`.
        :param pulumi.Input[pulumi.InputType['ProjectSecurityPropertiesArgs']] security_properties: Security-related attributesSee the following `Block SecurityProperties`.
        :param pulumi.Input[str] status: The status of the resource
        :param pulumi.Input[str] type: Life cycle type.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ProjectState.__new__(_ProjectState)

        __props__.__dict__["comment"] = comment
        __props__.__dict__["default_quota"] = default_quota
        __props__.__dict__["ip_white_list"] = ip_white_list
        __props__.__dict__["owner"] = owner
        __props__.__dict__["product_type"] = product_type
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["properties"] = properties
        __props__.__dict__["security_properties"] = security_properties
        __props__.__dict__["status"] = status
        __props__.__dict__["type"] = type
        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def comment(self) -> pulumi.Output[Optional[str]]:
        """
        Comments of project
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter(name="defaultQuota")
    def default_quota(self) -> pulumi.Output[Optional[str]]:
        """
        Default Computing Resource Group
        """
        return pulumi.get(self, "default_quota")

    @property
    @pulumi.getter(name="ipWhiteList")
    def ip_white_list(self) -> pulumi.Output[Optional['outputs.ProjectIpWhiteList']]:
        """
        IP whitelistSee the following `Block IpWhiteList`.
        """
        return pulumi.get(self, "ip_white_list")

    @property
    @pulumi.getter
    def owner(self) -> pulumi.Output[str]:
        """
        Project owner
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> pulumi.Output[Optional[str]]:
        """
        Quota payment type, support `PayAsYouGo`, `Subscription`, `Dev`.
        """
        return pulumi.get(self, "product_type")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The name of the project
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.ProjectProperties']:
        """
        Project base attributesSee the following `Block Properties`.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter(name="securityProperties")
    def security_properties(self) -> pulumi.Output['outputs.ProjectSecurityProperties']:
        """
        Security-related attributesSee the following `Block SecurityProperties`.
        """
        return pulumi.get(self, "security_properties")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the resource
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Life cycle type.
        """
        return pulumi.get(self, "type")

