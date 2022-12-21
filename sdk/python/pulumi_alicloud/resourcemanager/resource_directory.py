# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ResourceDirectoryArgs', 'ResourceDirectory']

@pulumi.input_type
class ResourceDirectoryArgs:
    def __init__(__self__, *,
                 status: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ResourceDirectory resource.
        :param pulumi.Input[str] status: The status of control policy. Valid values:`Enabled` and `Disabled`.
        """
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of control policy. Valid values:`Enabled` and `Disabled`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


@pulumi.input_type
class _ResourceDirectoryState:
    def __init__(__self__, *,
                 master_account_id: Optional[pulumi.Input[str]] = None,
                 master_account_name: Optional[pulumi.Input[str]] = None,
                 root_folder_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ResourceDirectory resources.
        :param pulumi.Input[str] master_account_id: The ID of the master account.
        :param pulumi.Input[str] master_account_name: The name of the master account.
        :param pulumi.Input[str] root_folder_id: The ID of the root folder.
        :param pulumi.Input[str] status: The status of control policy. Valid values:`Enabled` and `Disabled`.
        """
        if master_account_id is not None:
            pulumi.set(__self__, "master_account_id", master_account_id)
        if master_account_name is not None:
            pulumi.set(__self__, "master_account_name", master_account_name)
        if root_folder_id is not None:
            pulumi.set(__self__, "root_folder_id", root_folder_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="masterAccountId")
    def master_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the master account.
        """
        return pulumi.get(self, "master_account_id")

    @master_account_id.setter
    def master_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_account_id", value)

    @property
    @pulumi.getter(name="masterAccountName")
    def master_account_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the master account.
        """
        return pulumi.get(self, "master_account_name")

    @master_account_name.setter
    def master_account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_account_name", value)

    @property
    @pulumi.getter(name="rootFolderId")
    def root_folder_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the root folder.
        """
        return pulumi.get(self, "root_folder_id")

    @root_folder_id.setter
    def root_folder_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_folder_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of control policy. Valid values:`Enabled` and `Disabled`.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ResourceDirectory(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Resource Manager Resource Directory resource. Resource Directory enables you to establish an organizational structure for the resources used by applications of your enterprise. You can plan, build, and manage the resources in a centralized manner by using only one resource directory.

        For information about Resource Manager Resource Directory and how to use it, see [What is Resource Manager Resource Directory](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).

        > **NOTE:** Available in v1.84.0+.

        > **NOTE:** An account can only be used to enable a resource directory after it passes enterprise real-name verification. An account that only passed individual real-name verification cannot be used to enable a resource directory.

        > **NOTE:** Before you destroy the resource, make sure that the following requirements are met:
          - All member accounts must be removed from the resource directory.
          - All folders except the root folder must be deleted from the resource directory.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.resourcemanager.ResourceDirectory("example", status="Enabled")
        ```

        ## Import

        Resource Manager Resource Directory can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:resourcemanager/resourceDirectory:ResourceDirectory example rd-s3****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] status: The status of control policy. Valid values:`Enabled` and `Disabled`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ResourceDirectoryArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Resource Manager Resource Directory resource. Resource Directory enables you to establish an organizational structure for the resources used by applications of your enterprise. You can plan, build, and manage the resources in a centralized manner by using only one resource directory.

        For information about Resource Manager Resource Directory and how to use it, see [What is Resource Manager Resource Directory](https://www.alibabacloud.com/help/en/doc-detail/94475.htm).

        > **NOTE:** Available in v1.84.0+.

        > **NOTE:** An account can only be used to enable a resource directory after it passes enterprise real-name verification. An account that only passed individual real-name verification cannot be used to enable a resource directory.

        > **NOTE:** Before you destroy the resource, make sure that the following requirements are met:
          - All member accounts must be removed from the resource directory.
          - All folders except the root folder must be deleted from the resource directory.

        ## Example Usage

        Basic Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        example = alicloud.resourcemanager.ResourceDirectory("example", status="Enabled")
        ```

        ## Import

        Resource Manager Resource Directory can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:resourcemanager/resourceDirectory:ResourceDirectory example rd-s3****
        ```

        :param str resource_name: The name of the resource.
        :param ResourceDirectoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourceDirectoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourceDirectoryArgs.__new__(ResourceDirectoryArgs)

            __props__.__dict__["status"] = status
            __props__.__dict__["master_account_id"] = None
            __props__.__dict__["master_account_name"] = None
            __props__.__dict__["root_folder_id"] = None
        super(ResourceDirectory, __self__).__init__(
            'alicloud:resourcemanager/resourceDirectory:ResourceDirectory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            master_account_id: Optional[pulumi.Input[str]] = None,
            master_account_name: Optional[pulumi.Input[str]] = None,
            root_folder_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ResourceDirectory':
        """
        Get an existing ResourceDirectory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] master_account_id: The ID of the master account.
        :param pulumi.Input[str] master_account_name: The name of the master account.
        :param pulumi.Input[str] root_folder_id: The ID of the root folder.
        :param pulumi.Input[str] status: The status of control policy. Valid values:`Enabled` and `Disabled`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ResourceDirectoryState.__new__(_ResourceDirectoryState)

        __props__.__dict__["master_account_id"] = master_account_id
        __props__.__dict__["master_account_name"] = master_account_name
        __props__.__dict__["root_folder_id"] = root_folder_id
        __props__.__dict__["status"] = status
        return ResourceDirectory(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="masterAccountId")
    def master_account_id(self) -> pulumi.Output[str]:
        """
        The ID of the master account.
        """
        return pulumi.get(self, "master_account_id")

    @property
    @pulumi.getter(name="masterAccountName")
    def master_account_name(self) -> pulumi.Output[str]:
        """
        The name of the master account.
        """
        return pulumi.get(self, "master_account_name")

    @property
    @pulumi.getter(name="rootFolderId")
    def root_folder_id(self) -> pulumi.Output[str]:
        """
        The ID of the root folder.
        """
        return pulumi.get(self, "root_folder_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of control policy. Valid values:`Enabled` and `Disabled`.
        """
        return pulumi.get(self, "status")
