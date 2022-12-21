# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 assume_role: Optional[pulumi.Input['ProviderAssumeRoleArgs']] = None,
                 client_connect_timeout: Optional[pulumi.Input[int]] = None,
                 client_read_timeout: Optional[pulumi.Input[int]] = None,
                 configuration_source: Optional[pulumi.Input[str]] = None,
                 credentials_uri: Optional[pulumi.Input[str]] = None,
                 ecs_role_name: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderEndpointArgs']]]] = None,
                 fc: Optional[pulumi.Input[str]] = None,
                 log_endpoint: Optional[pulumi.Input[str]] = None,
                 max_retry_timeout: Optional[pulumi.Input[int]] = None,
                 mns_endpoint: Optional[pulumi.Input[str]] = None,
                 ots_instance_name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 secure_transport: Optional[pulumi.Input[str]] = None,
                 security_token: Optional[pulumi.Input[str]] = None,
                 security_transport: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 skip_region_validation: Optional[pulumi.Input[bool]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] access_key: The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
               console.
        :param pulumi.Input[str] account_id: The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
               Alibaba Cloud console.
        :param pulumi.Input[int] client_connect_timeout: The maximum timeout of the client connection server.
        :param pulumi.Input[int] client_read_timeout: The maximum timeout of the client read request.
        :param pulumi.Input[str] configuration_source: Use this to mark a terraform configuration file source.
        :param pulumi.Input[str] credentials_uri: The URI of sidecar credentials service.
        :param pulumi.Input[str] ecs_role_name: The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
               of the Alibaba Cloud console.
        :param pulumi.Input[int] max_retry_timeout: The maximum retry timeout of the request.
        :param pulumi.Input[str] profile: The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
        :param pulumi.Input[str] region: The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
        :param pulumi.Input[str] secret_key: The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
               console.
        :param pulumi.Input[str] secure_transport: The security transport for the assume role invoking.
        :param pulumi.Input[str] security_token: security token. A security token is only required if you are using Security Token Service.
        :param pulumi.Input[str] shared_credentials_file: The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        :param pulumi.Input[bool] skip_region_validation: Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
               that are not public (yet).
        :param pulumi.Input[str] source_ip: The source ip for the assume role invoking.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if assume_role is not None:
            pulumi.set(__self__, "assume_role", assume_role)
        if client_connect_timeout is not None:
            pulumi.set(__self__, "client_connect_timeout", client_connect_timeout)
        if client_read_timeout is not None:
            pulumi.set(__self__, "client_read_timeout", client_read_timeout)
        if configuration_source is not None:
            pulumi.set(__self__, "configuration_source", configuration_source)
        if credentials_uri is not None:
            pulumi.set(__self__, "credentials_uri", credentials_uri)
        if ecs_role_name is None:
            ecs_role_name = _utilities.get_env('ALICLOUD_ECS_ROLE_NAME')
        if ecs_role_name is not None:
            pulumi.set(__self__, "ecs_role_name", ecs_role_name)
        if endpoints is not None:
            pulumi.set(__self__, "endpoints", endpoints)
        if fc is not None:
            warnings.warn("""Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.""", DeprecationWarning)
            pulumi.log.warn("""fc is deprecated: Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.""")
        if fc is not None:
            pulumi.set(__self__, "fc", fc)
        if log_endpoint is not None:
            warnings.warn("""Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.""", DeprecationWarning)
            pulumi.log.warn("""log_endpoint is deprecated: Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.""")
        if log_endpoint is not None:
            pulumi.set(__self__, "log_endpoint", log_endpoint)
        if max_retry_timeout is not None:
            pulumi.set(__self__, "max_retry_timeout", max_retry_timeout)
        if mns_endpoint is not None:
            warnings.warn("""Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.""", DeprecationWarning)
            pulumi.log.warn("""mns_endpoint is deprecated: Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.""")
        if mns_endpoint is not None:
            pulumi.set(__self__, "mns_endpoint", mns_endpoint)
        if ots_instance_name is not None:
            warnings.warn("""Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.""", DeprecationWarning)
            pulumi.log.warn("""ots_instance_name is deprecated: Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.""")
        if ots_instance_name is not None:
            pulumi.set(__self__, "ots_instance_name", ots_instance_name)
        if profile is None:
            profile = _utilities.get_env('ALICLOUD_PROFILE')
        if profile is not None:
            pulumi.set(__self__, "profile", profile)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if region is None:
            region = _utilities.get_env('ALICLOUD_REGION')
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if secure_transport is not None:
            pulumi.set(__self__, "secure_transport", secure_transport)
        if security_token is not None:
            pulumi.set(__self__, "security_token", security_token)
        if security_transport is not None:
            pulumi.set(__self__, "security_transport", security_transport)
        if shared_credentials_file is not None:
            pulumi.set(__self__, "shared_credentials_file", shared_credentials_file)
        if skip_region_validation is not None:
            pulumi.set(__self__, "skip_region_validation", skip_region_validation)
        if source_ip is not None:
            pulumi.set(__self__, "source_ip", source_ip)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        console.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
        Alibaba Cloud console.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="assumeRole")
    def assume_role(self) -> Optional[pulumi.Input['ProviderAssumeRoleArgs']]:
        return pulumi.get(self, "assume_role")

    @assume_role.setter
    def assume_role(self, value: Optional[pulumi.Input['ProviderAssumeRoleArgs']]):
        pulumi.set(self, "assume_role", value)

    @property
    @pulumi.getter(name="clientConnectTimeout")
    def client_connect_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum timeout of the client connection server.
        """
        return pulumi.get(self, "client_connect_timeout")

    @client_connect_timeout.setter
    def client_connect_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "client_connect_timeout", value)

    @property
    @pulumi.getter(name="clientReadTimeout")
    def client_read_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum timeout of the client read request.
        """
        return pulumi.get(self, "client_read_timeout")

    @client_read_timeout.setter
    def client_read_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "client_read_timeout", value)

    @property
    @pulumi.getter(name="configurationSource")
    def configuration_source(self) -> Optional[pulumi.Input[str]]:
        """
        Use this to mark a terraform configuration file source.
        """
        return pulumi.get(self, "configuration_source")

    @configuration_source.setter
    def configuration_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configuration_source", value)

    @property
    @pulumi.getter(name="credentialsUri")
    def credentials_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The URI of sidecar credentials service.
        """
        return pulumi.get(self, "credentials_uri")

    @credentials_uri.setter
    def credentials_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credentials_uri", value)

    @property
    @pulumi.getter(name="ecsRoleName")
    def ecs_role_name(self) -> Optional[pulumi.Input[str]]:
        """
        The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
        of the Alibaba Cloud console.
        """
        return pulumi.get(self, "ecs_role_name")

    @ecs_role_name.setter
    def ecs_role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ecs_role_name", value)

    @property
    @pulumi.getter
    def endpoints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ProviderEndpointArgs']]]]:
        return pulumi.get(self, "endpoints")

    @endpoints.setter
    def endpoints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ProviderEndpointArgs']]]]):
        pulumi.set(self, "endpoints", value)

    @property
    @pulumi.getter
    def fc(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fc")

    @fc.setter
    def fc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fc", value)

    @property
    @pulumi.getter(name="logEndpoint")
    def log_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "log_endpoint")

    @log_endpoint.setter
    def log_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_endpoint", value)

    @property
    @pulumi.getter(name="maxRetryTimeout")
    def max_retry_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum retry timeout of the request.
        """
        return pulumi.get(self, "max_retry_timeout")

    @max_retry_timeout.setter
    def max_retry_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retry_timeout", value)

    @property
    @pulumi.getter(name="mnsEndpoint")
    def mns_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mns_endpoint")

    @mns_endpoint.setter
    def mns_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mns_endpoint", value)

    @property
    @pulumi.getter(name="otsInstanceName")
    def ots_instance_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ots_instance_name")

    @ots_instance_name.setter
    def ots_instance_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ots_instance_name", value)

    @property
    @pulumi.getter
    def profile(self) -> Optional[pulumi.Input[str]]:
        """
        The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
        """
        return pulumi.get(self, "profile")

    @profile.setter
    def profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "profile", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        console.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="secureTransport")
    def secure_transport(self) -> Optional[pulumi.Input[str]]:
        """
        The security transport for the assume role invoking.
        """
        return pulumi.get(self, "secure_transport")

    @secure_transport.setter
    def secure_transport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secure_transport", value)

    @property
    @pulumi.getter(name="securityToken")
    def security_token(self) -> Optional[pulumi.Input[str]]:
        """
        security token. A security token is only required if you are using Security Token Service.
        """
        return pulumi.get(self, "security_token")

    @security_token.setter
    def security_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_token", value)

    @property
    @pulumi.getter(name="securityTransport")
    def security_transport(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "security_transport")

    @security_transport.setter
    def security_transport(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_transport", value)

    @property
    @pulumi.getter(name="sharedCredentialsFile")
    def shared_credentials_file(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        """
        return pulumi.get(self, "shared_credentials_file")

    @shared_credentials_file.setter
    def shared_credentials_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_credentials_file", value)

    @property
    @pulumi.getter(name="skipRegionValidation")
    def skip_region_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
        that are not public (yet).
        """
        return pulumi.get(self, "skip_region_validation")

    @skip_region_validation.setter
    def skip_region_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_region_validation", value)

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> Optional[pulumi.Input[str]]:
        """
        The source ip for the assume role invoking.
        """
        return pulumi.get(self, "source_ip")

    @source_ip.setter
    def source_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_ip", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 assume_role: Optional[pulumi.Input[pulumi.InputType['ProviderAssumeRoleArgs']]] = None,
                 client_connect_timeout: Optional[pulumi.Input[int]] = None,
                 client_read_timeout: Optional[pulumi.Input[int]] = None,
                 configuration_source: Optional[pulumi.Input[str]] = None,
                 credentials_uri: Optional[pulumi.Input[str]] = None,
                 ecs_role_name: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderEndpointArgs']]]]] = None,
                 fc: Optional[pulumi.Input[str]] = None,
                 log_endpoint: Optional[pulumi.Input[str]] = None,
                 max_retry_timeout: Optional[pulumi.Input[int]] = None,
                 mns_endpoint: Optional[pulumi.Input[str]] = None,
                 ots_instance_name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 secure_transport: Optional[pulumi.Input[str]] = None,
                 security_token: Optional[pulumi.Input[str]] = None,
                 security_transport: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 skip_region_validation: Optional[pulumi.Input[bool]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the alicloud package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
               console.
        :param pulumi.Input[str] account_id: The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
               Alibaba Cloud console.
        :param pulumi.Input[int] client_connect_timeout: The maximum timeout of the client connection server.
        :param pulumi.Input[int] client_read_timeout: The maximum timeout of the client read request.
        :param pulumi.Input[str] configuration_source: Use this to mark a terraform configuration file source.
        :param pulumi.Input[str] credentials_uri: The URI of sidecar credentials service.
        :param pulumi.Input[str] ecs_role_name: The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
               of the Alibaba Cloud console.
        :param pulumi.Input[int] max_retry_timeout: The maximum retry timeout of the request.
        :param pulumi.Input[str] profile: The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
        :param pulumi.Input[str] region: The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
        :param pulumi.Input[str] secret_key: The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
               console.
        :param pulumi.Input[str] secure_transport: The security transport for the assume role invoking.
        :param pulumi.Input[str] security_token: security token. A security token is only required if you are using Security Token Service.
        :param pulumi.Input[str] shared_credentials_file: The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        :param pulumi.Input[bool] skip_region_validation: Skip static validation of region ID. Used by users of alternative AlibabaCloud-like APIs or users w/ access to regions
               that are not public (yet).
        :param pulumi.Input[str] source_ip: The source ip for the assume role invoking.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the alicloud package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 assume_role: Optional[pulumi.Input[pulumi.InputType['ProviderAssumeRoleArgs']]] = None,
                 client_connect_timeout: Optional[pulumi.Input[int]] = None,
                 client_read_timeout: Optional[pulumi.Input[int]] = None,
                 configuration_source: Optional[pulumi.Input[str]] = None,
                 credentials_uri: Optional[pulumi.Input[str]] = None,
                 ecs_role_name: Optional[pulumi.Input[str]] = None,
                 endpoints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ProviderEndpointArgs']]]]] = None,
                 fc: Optional[pulumi.Input[str]] = None,
                 log_endpoint: Optional[pulumi.Input[str]] = None,
                 max_retry_timeout: Optional[pulumi.Input[int]] = None,
                 mns_endpoint: Optional[pulumi.Input[str]] = None,
                 ots_instance_name: Optional[pulumi.Input[str]] = None,
                 profile: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 secure_transport: Optional[pulumi.Input[str]] = None,
                 security_token: Optional[pulumi.Input[str]] = None,
                 security_transport: Optional[pulumi.Input[str]] = None,
                 shared_credentials_file: Optional[pulumi.Input[str]] = None,
                 skip_region_validation: Optional[pulumi.Input[bool]] = None,
                 source_ip: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["account_id"] = account_id
            __props__.__dict__["assume_role"] = pulumi.Output.from_input(assume_role).apply(pulumi.runtime.to_json) if assume_role is not None else None
            __props__.__dict__["client_connect_timeout"] = pulumi.Output.from_input(client_connect_timeout).apply(pulumi.runtime.to_json) if client_connect_timeout is not None else None
            __props__.__dict__["client_read_timeout"] = pulumi.Output.from_input(client_read_timeout).apply(pulumi.runtime.to_json) if client_read_timeout is not None else None
            __props__.__dict__["configuration_source"] = configuration_source
            __props__.__dict__["credentials_uri"] = credentials_uri
            if ecs_role_name is None:
                ecs_role_name = _utilities.get_env('ALICLOUD_ECS_ROLE_NAME')
            __props__.__dict__["ecs_role_name"] = ecs_role_name
            __props__.__dict__["endpoints"] = pulumi.Output.from_input(endpoints).apply(pulumi.runtime.to_json) if endpoints is not None else None
            if fc is not None and not opts.urn:
                warnings.warn("""Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.""", DeprecationWarning)
                pulumi.log.warn("""fc is deprecated: Field 'fc' has been deprecated from provider version 1.28.0. New field 'fc' which in nested endpoints instead.""")
            __props__.__dict__["fc"] = fc
            if log_endpoint is not None and not opts.urn:
                warnings.warn("""Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.""", DeprecationWarning)
                pulumi.log.warn("""log_endpoint is deprecated: Field 'log_endpoint' has been deprecated from provider version 1.28.0. New field 'log' which in nested endpoints instead.""")
            __props__.__dict__["log_endpoint"] = log_endpoint
            __props__.__dict__["max_retry_timeout"] = pulumi.Output.from_input(max_retry_timeout).apply(pulumi.runtime.to_json) if max_retry_timeout is not None else None
            if mns_endpoint is not None and not opts.urn:
                warnings.warn("""Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.""", DeprecationWarning)
                pulumi.log.warn("""mns_endpoint is deprecated: Field 'mns_endpoint' has been deprecated from provider version 1.28.0. New field 'mns' which in nested endpoints instead.""")
            __props__.__dict__["mns_endpoint"] = mns_endpoint
            if ots_instance_name is not None and not opts.urn:
                warnings.warn("""Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.""", DeprecationWarning)
                pulumi.log.warn("""ots_instance_name is deprecated: Field 'ots_instance_name' has been deprecated from provider version 1.10.0. New field 'instance_name' of resource 'alicloud_ots_table' instead.""")
            __props__.__dict__["ots_instance_name"] = ots_instance_name
            if profile is None:
                profile = _utilities.get_env('ALICLOUD_PROFILE')
            __props__.__dict__["profile"] = profile
            __props__.__dict__["protocol"] = protocol
            if region is None:
                region = _utilities.get_env('ALICLOUD_REGION')
            __props__.__dict__["region"] = region
            __props__.__dict__["secret_key"] = secret_key
            __props__.__dict__["secure_transport"] = secure_transport
            __props__.__dict__["security_token"] = security_token
            __props__.__dict__["security_transport"] = security_transport
            __props__.__dict__["shared_credentials_file"] = shared_credentials_file
            __props__.__dict__["skip_region_validation"] = pulumi.Output.from_input(skip_region_validation).apply(pulumi.runtime.to_json) if skip_region_validation is not None else None
            __props__.__dict__["source_ip"] = source_ip
        super(Provider, __self__).__init__(
            'alicloud',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[Optional[str]]:
        """
        The access key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        console.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The account ID for some service API operations. You can retrieve this from the 'Security Settings' section of the
        Alibaba Cloud console.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="configurationSource")
    def configuration_source(self) -> pulumi.Output[Optional[str]]:
        """
        Use this to mark a terraform configuration file source.
        """
        return pulumi.get(self, "configuration_source")

    @property
    @pulumi.getter(name="credentialsUri")
    def credentials_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The URI of sidecar credentials service.
        """
        return pulumi.get(self, "credentials_uri")

    @property
    @pulumi.getter(name="ecsRoleName")
    def ecs_role_name(self) -> pulumi.Output[Optional[str]]:
        """
        The RAM Role Name attached on a ECS instance for API operations. You can retrieve this from the 'Access Control' section
        of the Alibaba Cloud console.
        """
        return pulumi.get(self, "ecs_role_name")

    @property
    @pulumi.getter
    def fc(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "fc")

    @property
    @pulumi.getter(name="logEndpoint")
    def log_endpoint(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "log_endpoint")

    @property
    @pulumi.getter(name="mnsEndpoint")
    def mns_endpoint(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "mns_endpoint")

    @property
    @pulumi.getter(name="otsInstanceName")
    def ots_instance_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "ots_instance_name")

    @property
    @pulumi.getter
    def profile(self) -> pulumi.Output[Optional[str]]:
        """
        The profile for API operations. If not set, the default profile created with `aliyun configure` will be used.
        """
        return pulumi.get(self, "profile")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The region where Alibaba Cloud operations will take place. Examples are cn-beijing, cn-hangzhou, eu-central-1, etc.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[Optional[str]]:
        """
        The secret key for API operations. You can retrieve this from the 'Security Management' section of the Alibaba Cloud
        console.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="secureTransport")
    def secure_transport(self) -> pulumi.Output[Optional[str]]:
        """
        The security transport for the assume role invoking.
        """
        return pulumi.get(self, "secure_transport")

    @property
    @pulumi.getter(name="securityToken")
    def security_token(self) -> pulumi.Output[Optional[str]]:
        """
        security token. A security token is only required if you are using Security Token Service.
        """
        return pulumi.get(self, "security_token")

    @property
    @pulumi.getter(name="securityTransport")
    def security_transport(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "security_transport")

    @property
    @pulumi.getter(name="sharedCredentialsFile")
    def shared_credentials_file(self) -> pulumi.Output[Optional[str]]:
        """
        The path to the shared credentials file. If not set this defaults to ~/.aliyun/config.json
        """
        return pulumi.get(self, "shared_credentials_file")

    @property
    @pulumi.getter(name="sourceIp")
    def source_ip(self) -> pulumi.Output[Optional[str]]:
        """
        The source ip for the assume role invoking.
        """
        return pulumi.get(self, "source_ip")
