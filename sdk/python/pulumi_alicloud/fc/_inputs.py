# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AliasRoutingConfigArgs',
    'CustomDomainCertConfigArgs',
    'CustomDomainRouteConfigArgs',
    'FunctionAsyncInvokeConfigDestinationConfigArgs',
    'FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs',
    'FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs',
    'FunctionCustomContainerConfigArgs',
    'ServiceLogConfigArgs',
    'ServiceNasConfigArgs',
    'ServiceNasConfigMountPointArgs',
    'ServiceTracingConfigArgs',
    'ServiceVpcConfigArgs',
]

@pulumi.input_type
class AliasRoutingConfigArgs:
    def __init__(__self__, *,
                 additional_version_weights: Optional[pulumi.Input[Mapping[str, pulumi.Input[float]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[float]]] additional_version_weights: A map that defines the proportion of events that should be sent to different versions of a Function Compute service.
        """
        if additional_version_weights is not None:
            pulumi.set(__self__, "additional_version_weights", additional_version_weights)

    @property
    @pulumi.getter(name="additionalVersionWeights")
    def additional_version_weights(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[float]]]]:
        """
        A map that defines the proportion of events that should be sent to different versions of a Function Compute service.
        """
        return pulumi.get(self, "additional_version_weights")

    @additional_version_weights.setter
    def additional_version_weights(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[float]]]]):
        pulumi.set(self, "additional_version_weights", value)


@pulumi.input_type
class CustomDomainCertConfigArgs:
    def __init__(__self__, *,
                 cert_name: pulumi.Input[str],
                 certificate: pulumi.Input[str],
                 private_key: pulumi.Input[str]):
        """
        :param pulumi.Input[str] cert_name: The name of the certificate, used to distinguish different certificates.
        :param pulumi.Input[str] certificate: Certificate data of the HTTPS certificates, follow the 'pem' format.
        :param pulumi.Input[str] private_key: Private key of the HTTPS certificates, follow the 'pem' format.
        """
        pulumi.set(__self__, "cert_name", cert_name)
        pulumi.set(__self__, "certificate", certificate)
        pulumi.set(__self__, "private_key", private_key)

    @property
    @pulumi.getter(name="certName")
    def cert_name(self) -> pulumi.Input[str]:
        """
        The name of the certificate, used to distinguish different certificates.
        """
        return pulumi.get(self, "cert_name")

    @cert_name.setter
    def cert_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert_name", value)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Input[str]:
        """
        Certificate data of the HTTPS certificates, follow the 'pem' format.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[str]:
        """
        Private key of the HTTPS certificates, follow the 'pem' format.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key", value)


@pulumi.input_type
class CustomDomainRouteConfigArgs:
    def __init__(__self__, *,
                 function_name: pulumi.Input[str],
                 path: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 methods: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] function_name: The name of the Function Compute function that requests are routed to.
        :param pulumi.Input[str] path: The path that requests are routed from.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] methods: The requests of the specified HTTP methos are routed from. Valid method: GET, POST, DELETE, HEAD, PUT and PATCH. For example, "GET, HEAD" methods indicate that only requests from GET and HEAD methods are routed.
        :param pulumi.Input[str] qualifier: The version or alias of the Function Compute service that requests are routed to. For example, qualifier v1 indicates that the requests are routed to the version 1 Function Compute service. For detail information about version and alias, please refer to the [developer guide](https://www.alibabacloud.com/help/doc-detail/96464.htm).
        """
        pulumi.set(__self__, "function_name", function_name)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "service_name", service_name)
        if methods is not None:
            pulumi.set(__self__, "methods", methods)
        if qualifier is not None:
            pulumi.set(__self__, "qualifier", qualifier)

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Input[str]:
        """
        The name of the Function Compute function that requests are routed to.
        """
        return pulumi.get(self, "function_name")

    @function_name.setter
    def function_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_name", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The path that requests are routed from.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def methods(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The requests of the specified HTTP methos are routed from. Valid method: GET, POST, DELETE, HEAD, PUT and PATCH. For example, "GET, HEAD" methods indicate that only requests from GET and HEAD methods are routed.
        """
        return pulumi.get(self, "methods")

    @methods.setter
    def methods(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "methods", value)

    @property
    @pulumi.getter
    def qualifier(self) -> Optional[pulumi.Input[str]]:
        """
        The version or alias of the Function Compute service that requests are routed to. For example, qualifier v1 indicates that the requests are routed to the version 1 Function Compute service. For detail information about version and alias, please refer to the [developer guide](https://www.alibabacloud.com/help/doc-detail/96464.htm).
        """
        return pulumi.get(self, "qualifier")

    @qualifier.setter
    def qualifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "qualifier", value)


@pulumi.input_type
class FunctionAsyncInvokeConfigDestinationConfigArgs:
    def __init__(__self__, *,
                 on_failure: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs']] = None,
                 on_success: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs']] = None):
        """
        :param pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs'] on_failure: Configuration block with destination configuration for failed asynchronous invocations. See below for details.
        :param pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs'] on_success: Configuration block with destination configuration for successful asynchronous invocations. See below for details.
        """
        if on_failure is not None:
            pulumi.set(__self__, "on_failure", on_failure)
        if on_success is not None:
            pulumi.set(__self__, "on_success", on_success)

    @property
    @pulumi.getter(name="onFailure")
    def on_failure(self) -> Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs']]:
        """
        Configuration block with destination configuration for failed asynchronous invocations. See below for details.
        """
        return pulumi.get(self, "on_failure")

    @on_failure.setter
    def on_failure(self, value: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs']]):
        pulumi.set(self, "on_failure", value)

    @property
    @pulumi.getter(name="onSuccess")
    def on_success(self) -> Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs']]:
        """
        Configuration block with destination configuration for successful asynchronous invocations. See below for details.
        """
        return pulumi.get(self, "on_success")

    @on_success.setter
    def on_success(self, value: Optional[pulumi.Input['FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs']]):
        pulumi.set(self, "on_success", value)


@pulumi.input_type
class FunctionAsyncInvokeConfigDestinationConfigOnFailureArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str]):
        """
        :param pulumi.Input[str] destination: Alicloud Resource Name (ARN) of the destination resource. See the [Developer Guide](https://www.alibabacloud.com/help/doc-detail/181866.htm) for acceptable resource types and associated RAM permissions.
        """
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        Alicloud Resource Name (ARN) of the destination resource. See the [Developer Guide](https://www.alibabacloud.com/help/doc-detail/181866.htm) for acceptable resource types and associated RAM permissions.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)


@pulumi.input_type
class FunctionAsyncInvokeConfigDestinationConfigOnSuccessArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str]):
        """
        :param pulumi.Input[str] destination: Alicloud Resource Name (ARN) of the destination resource. See the [Developer Guide](https://www.alibabacloud.com/help/doc-detail/181866.htm) for acceptable resource types and associated RAM permissions.
        """
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        Alicloud Resource Name (ARN) of the destination resource. See the [Developer Guide](https://www.alibabacloud.com/help/doc-detail/181866.htm) for acceptable resource types and associated RAM permissions.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)


@pulumi.input_type
class FunctionCustomContainerConfigArgs:
    def __init__(__self__, *,
                 image: pulumi.Input[str],
                 args: Optional[pulumi.Input[str]] = None,
                 command: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] image: The container image address.
        :param pulumi.Input[str] args: The args field specifies the arguments passed to the command.
        :param pulumi.Input[str] command: The entry point of the container, which specifies the actual command run by the container.
        """
        pulumi.set(__self__, "image", image)
        if args is not None:
            pulumi.set(__self__, "args", args)
        if command is not None:
            pulumi.set(__self__, "command", command)

    @property
    @pulumi.getter
    def image(self) -> pulumi.Input[str]:
        """
        The container image address.
        """
        return pulumi.get(self, "image")

    @image.setter
    def image(self, value: pulumi.Input[str]):
        pulumi.set(self, "image", value)

    @property
    @pulumi.getter
    def args(self) -> Optional[pulumi.Input[str]]:
        """
        The args field specifies the arguments passed to the command.
        """
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter
    def command(self) -> Optional[pulumi.Input[str]]:
        """
        The entry point of the container, which specifies the actual command run by the container.
        """
        return pulumi.get(self, "command")

    @command.setter
    def command(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "command", value)


@pulumi.input_type
class ServiceLogConfigArgs:
    def __init__(__self__, *,
                 logstore: pulumi.Input[str],
                 project: pulumi.Input[str],
                 enable_instance_metrics: Optional[pulumi.Input[bool]] = None,
                 enable_request_metrics: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] logstore: The log store name of Alicloud Simple Log Service.
        :param pulumi.Input[str] project: The project name of the Alicloud Simple Log Service.
        :param pulumi.Input[bool] enable_instance_metrics: Enable instance level metrics.
        :param pulumi.Input[bool] enable_request_metrics: Enable request level metrics.
        """
        pulumi.set(__self__, "logstore", logstore)
        pulumi.set(__self__, "project", project)
        if enable_instance_metrics is not None:
            pulumi.set(__self__, "enable_instance_metrics", enable_instance_metrics)
        if enable_request_metrics is not None:
            pulumi.set(__self__, "enable_request_metrics", enable_request_metrics)

    @property
    @pulumi.getter
    def logstore(self) -> pulumi.Input[str]:
        """
        The log store name of Alicloud Simple Log Service.
        """
        return pulumi.get(self, "logstore")

    @logstore.setter
    def logstore(self, value: pulumi.Input[str]):
        pulumi.set(self, "logstore", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The project name of the Alicloud Simple Log Service.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="enableInstanceMetrics")
    def enable_instance_metrics(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable instance level metrics.
        """
        return pulumi.get(self, "enable_instance_metrics")

    @enable_instance_metrics.setter
    def enable_instance_metrics(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_instance_metrics", value)

    @property
    @pulumi.getter(name="enableRequestMetrics")
    def enable_request_metrics(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable request level metrics.
        """
        return pulumi.get(self, "enable_request_metrics")

    @enable_request_metrics.setter
    def enable_request_metrics(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_request_metrics", value)


@pulumi.input_type
class ServiceNasConfigArgs:
    def __init__(__self__, *,
                 group_id: pulumi.Input[int],
                 mount_points: pulumi.Input[Sequence[pulumi.Input['ServiceNasConfigMountPointArgs']]],
                 user_id: pulumi.Input[int]):
        """
        :param pulumi.Input[int] group_id: The group id of your NAS file system.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceNasConfigMountPointArgs']]] mount_points: Config the NAS mount points, including following attributes:
        :param pulumi.Input[int] user_id: The user id of your NAS file system.
        """
        pulumi.set(__self__, "group_id", group_id)
        pulumi.set(__self__, "mount_points", mount_points)
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Input[int]:
        """
        The group id of your NAS file system.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter(name="mountPoints")
    def mount_points(self) -> pulumi.Input[Sequence[pulumi.Input['ServiceNasConfigMountPointArgs']]]:
        """
        Config the NAS mount points, including following attributes:
        """
        return pulumi.get(self, "mount_points")

    @mount_points.setter
    def mount_points(self, value: pulumi.Input[Sequence[pulumi.Input['ServiceNasConfigMountPointArgs']]]):
        pulumi.set(self, "mount_points", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[int]:
        """
        The user id of your NAS file system.
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[int]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class ServiceNasConfigMountPointArgs:
    def __init__(__self__, *,
                 mount_dir: pulumi.Input[str],
                 server_addr: pulumi.Input[str]):
        """
        :param pulumi.Input[str] mount_dir: The local address where to mount your remote NAS directory.
        :param pulumi.Input[str] server_addr: The address of the remote NAS directory.
        """
        pulumi.set(__self__, "mount_dir", mount_dir)
        pulumi.set(__self__, "server_addr", server_addr)

    @property
    @pulumi.getter(name="mountDir")
    def mount_dir(self) -> pulumi.Input[str]:
        """
        The local address where to mount your remote NAS directory.
        """
        return pulumi.get(self, "mount_dir")

    @mount_dir.setter
    def mount_dir(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount_dir", value)

    @property
    @pulumi.getter(name="serverAddr")
    def server_addr(self) -> pulumi.Input[str]:
        """
        The address of the remote NAS directory.
        """
        return pulumi.get(self, "server_addr")

    @server_addr.setter
    def server_addr(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_addr", value)


@pulumi.input_type
class ServiceTracingConfigArgs:
    def __init__(__self__, *,
                 params: pulumi.Input[Mapping[str, Any]],
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[Mapping[str, Any]] params: Tracing parameters, which type is map[string]string. When the protocol type is Jaeger, the key is "endpoint" and the value is your tracing intranet endpoint. For example endpoint: http://tracing-analysis-dc-hz.aliyuncs.com/adapt_xxx/api/traces.
        :param pulumi.Input[str] type: Tracing protocol type. Currently, only Jaeger is supported.
        """
        pulumi.set(__self__, "params", params)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def params(self) -> pulumi.Input[Mapping[str, Any]]:
        """
        Tracing parameters, which type is map[string]string. When the protocol type is Jaeger, the key is "endpoint" and the value is your tracing intranet endpoint. For example endpoint: http://tracing-analysis-dc-hz.aliyuncs.com/adapt_xxx/api/traces.
        """
        return pulumi.get(self, "params")

    @params.setter
    def params(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "params", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Tracing protocol type. Currently, only Jaeger is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ServiceVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_id: pulumi.Input[str],
                 vswitch_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] security_group_id: A security group ID associated with the Function Compute Service.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] vswitch_ids: A list of vswitch IDs associated with the Function Compute Service.
        """
        pulumi.set(__self__, "security_group_id", security_group_id)
        pulumi.set(__self__, "vswitch_ids", vswitch_ids)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupId")
    def security_group_id(self) -> pulumi.Input[str]:
        """
        A security group ID associated with the Function Compute Service.
        """
        return pulumi.get(self, "security_group_id")

    @security_group_id.setter
    def security_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "security_group_id", value)

    @property
    @pulumi.getter(name="vswitchIds")
    def vswitch_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of vswitch IDs associated with the Function Compute Service.
        """
        return pulumi.get(self, "vswitch_ids")

    @vswitch_ids.setter
    def vswitch_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "vswitch_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

