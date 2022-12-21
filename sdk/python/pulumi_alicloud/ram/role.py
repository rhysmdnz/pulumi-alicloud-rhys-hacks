# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ram_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[str] description: Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
        :param pulumi.Input[str] document: Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[int] max_session_duration: The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
        :param pulumi.Input[str] name: Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ram_users: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        :param pulumi.Input[str] version: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if document is not None:
            pulumi.set(__self__, "document", document)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if max_session_duration is not None:
            pulumi.set(__self__, "max_session_duration", max_session_duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ram_users is not None:
            warnings.warn("""Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
            pulumi.log.warn("""ram_users is deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
        if ram_users is not None:
            pulumi.set(__self__, "ram_users", ram_users)
        if services is not None:
            warnings.warn("""Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
            pulumi.log.warn("""services is deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if version is not None:
            warnings.warn("""Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
            pulumi.log.warn("""version is deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def document(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        """
        return pulumi.get(self, "document")

    @document.setter
    def document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        This parameter is used for resource destroy. Default value is `false`.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
        """
        return pulumi.get(self, "max_session_duration")

    @max_session_duration.setter
    def max_session_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_session_duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ramUsers")
    def ram_users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        """
        return pulumi.get(self, "ram_users")

    @ram_users.setter
    def ram_users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ram_users", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ram_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[str] arn: The role arn.
        :param pulumi.Input[str] description: Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
        :param pulumi.Input[str] document: Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[int] max_session_duration: The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
        :param pulumi.Input[str] name: Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ram_users: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        :param pulumi.Input[str] role_id: The role ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        :param pulumi.Input[str] version: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if document is not None:
            pulumi.set(__self__, "document", document)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if max_session_duration is not None:
            pulumi.set(__self__, "max_session_duration", max_session_duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if ram_users is not None:
            warnings.warn("""Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
            pulumi.log.warn("""ram_users is deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
        if ram_users is not None:
            pulumi.set(__self__, "ram_users", ram_users)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if services is not None:
            warnings.warn("""Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
            pulumi.log.warn("""services is deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
        if services is not None:
            pulumi.set(__self__, "services", services)
        if version is not None:
            warnings.warn("""Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
            pulumi.log.warn("""version is deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The role arn.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def document(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        """
        return pulumi.get(self, "document")

    @document.setter
    def document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "document", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        """
        This parameter is used for resource destroy. Default value is `false`.
        """
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
        """
        return pulumi.get(self, "max_session_duration")

    @max_session_duration.setter
    def max_session_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_session_duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="ramUsers")
    def ram_users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        """
        return pulumi.get(self, "ram_users")

    @ram_users.setter
    def ram_users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ram_users", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        """
        The role ID.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def services(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        """
        return pulumi.get(self, "services")

    @services.setter
    def services(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "services", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ram_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a new RAM Role.
        role = alicloud.ram.Role("role",
            description="this is a role test.",
            document=\"\"\"  {
            "Statement": [
              {
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Principal": {
                  "Service": [
                    "apigateway.aliyuncs.com", 
                    "ecs.aliyuncs.com"
                  ]
                }
              }
            ],
            "Version": "1"
          }
          
        \"\"\",
            force=True)
        ```

        ## Import

        RAM role can be imported using the id or name, e.g.

        ```sh
         $ pulumi import alicloud:ram/role:Role example my-role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
        :param pulumi.Input[str] document: Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[int] max_session_duration: The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
        :param pulumi.Input[str] name: Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ram_users: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        :param pulumi.Input[str] version: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RoleArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_alicloud as alicloud

        # Create a new RAM Role.
        role = alicloud.ram.Role("role",
            description="this is a role test.",
            document=\"\"\"  {
            "Statement": [
              {
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Principal": {
                  "Service": [
                    "apigateway.aliyuncs.com", 
                    "ecs.aliyuncs.com"
                  ]
                }
              }
            ],
            "Version": "1"
          }
          
        \"\"\",
            force=True)
        ```

        ## Import

        RAM role can be imported using the id or name, e.g.

        ```sh
         $ pulumi import alicloud:ram/role:Role example my-role
        ```

        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 document: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ram_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["document"] = document
            __props__.__dict__["force"] = force
            __props__.__dict__["max_session_duration"] = max_session_duration
            __props__.__dict__["name"] = name
            if ram_users is not None and not opts.urn:
                warnings.warn("""Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
                pulumi.log.warn("""ram_users is deprecated: Field 'ram_users' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
            __props__.__dict__["ram_users"] = ram_users
            if services is not None and not opts.urn:
                warnings.warn("""Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
                pulumi.log.warn("""services is deprecated: Field 'services' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
            __props__.__dict__["services"] = services
            if version is not None and not opts.urn:
                warnings.warn("""Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace. """, DeprecationWarning)
                pulumi.log.warn("""version is deprecated: Field 'version' has been deprecated from version 1.49.0, and use field 'document' to replace. """)
            __props__.__dict__["version"] = version
            __props__.__dict__["arn"] = None
            __props__.__dict__["role_id"] = None
        super(Role, __self__).__init__(
            'alicloud:ram/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            document: Optional[pulumi.Input[str]] = None,
            force: Optional[pulumi.Input[bool]] = None,
            max_session_duration: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ram_users: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role_id: Optional[pulumi.Input[str]] = None,
            services: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The role arn.
        :param pulumi.Input[str] description: Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
        :param pulumi.Input[str] document: Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        :param pulumi.Input[bool] force: This parameter is used for resource destroy. Default value is `false`.
        :param pulumi.Input[int] max_session_duration: The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
        :param pulumi.Input[str] name: Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ram_users: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        :param pulumi.Input[str] role_id: The role ID.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] services: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        :param pulumi.Input[str] version: (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["document"] = document
        __props__.__dict__["force"] = force
        __props__.__dict__["max_session_duration"] = max_session_duration
        __props__.__dict__["name"] = name
        __props__.__dict__["ram_users"] = ram_users
        __props__.__dict__["role_id"] = role_id
        __props__.__dict__["services"] = services
        __props__.__dict__["version"] = version
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The role arn.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the RAM role. This name can have a string of 1 to 1024 characters. **NOTE:** The `description` supports modification since V1.144.0.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def document(self) -> pulumi.Output[str]:
        """
        Authorization strategy of the RAM role. It is required when the `services` and `ram_users` are not specified.
        """
        return pulumi.get(self, "document")

    @property
    @pulumi.getter
    def force(self) -> pulumi.Output[Optional[bool]]:
        """
        This parameter is used for resource destroy. Default value is `false`.
        """
        return pulumi.get(self, "force")

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum session duration of the RAM role. Valid values: 3600 to 43200. Unit: seconds. Default value: 3600. The default value is used if the parameter is not specified.
        """
        return pulumi.get(self, "max_session_duration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the RAM role. This name can have a string of 1 to 64 characters, must contain only alphanumeric characters or hyphens, such as "-", "_", and must not begin with a hyphen.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ramUsers")
    def ram_users(self) -> pulumi.Output[Sequence[str]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of ram users who can assume the RAM role. The format of each item in this list is `acs:ram::${account_id}:root` or `acs:ram::${account_id}:user/${user_name}`, such as `acs:ram::1234567890000:root` and `acs:ram::1234567890001:user/Mary`. The `${user_name}` is the name of a RAM user which must exists in the Alicloud account indicated by the `${account_id}`.
        """
        return pulumi.get(self, "ram_users")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The role ID.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter
    def services(self) -> pulumi.Output[Sequence[str]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) List of services which can assume the RAM role. The format of each item in this list is `${service}.aliyuncs.com` or `${account_id}@${service}.aliyuncs.com`, such as `ecs.aliyuncs.com` and `1234567890000@ots.aliyuncs.com`. The `${service}` can be `ecs`, `log`, `apigateway` and so on, the `${account_id}` refers to someone's Alicloud account id.
        """
        return pulumi.get(self, "services")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        """
        (It has been deprecated from version 1.49.0, and use field 'document' to replace.) Version of the RAM role policy document. Valid value is `1`. Default value is `1`.
        """
        return pulumi.get(self, "version")
