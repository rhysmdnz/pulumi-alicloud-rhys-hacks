// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class ServerGroupServer
    {
        /// <summary>
        /// The description of the server.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The port that is used by the server. Valid values: `1` to `65535`. **Note:** This parameter is required if the `server_type` parameter is set to `Ecs`, `Eni`, `Eci`, or `Ip`. You do not need to configure this parameter if you set `server_type` to `Fc`.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Specifies whether to enable the remote IP address feature. You can specify up to 40 servers in each call. **Note:** If `server_type` is set to `Ip`, this parameter is available.
        /// </summary>
        public readonly bool? RemoteIpEnabled;
        /// <summary>
        /// The ID of the backend server.
        /// - If `server_group_type` is set to `Instance`, set the parameter to the ID of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. These backend servers are specified by Ecs, Eni, or Eci.
        /// - If `server_group_type` is set to `Ip`, set the parameter to an IP address specified in the server group.
        /// - If `server_group_type` is set to `Fc`, set the parameter to the Alibaba Cloud Resource Name (ARN) of a function specified in the server group.
        /// </summary>
        public readonly string ServerId;
        /// <summary>
        /// The IP address of an Elastic Compute Service (ECS) instance, an elastic network interface (ENI), or an elastic container instance. **Note:** If `server_group_type` is set to `Fc`, you do not need to configure parameters, otherwise this attribute is required. If `server_group_type` is set to `Ip`, the value of this property is the same as the `server_id` value.
        /// </summary>
        public readonly string? ServerIp;
        /// <summary>
        /// The type of the server. The type of the server. Valid values: 
        /// - Ecs: an ECS instance.
        /// - Eni: an ENI.
        /// - Eci: an elastic container instance.
        /// - Ip(Available in v1.194.0+): an IP address.
        /// - fc(Available in v1.194.0+): a function.
        /// </summary>
        public readonly string ServerType;
        /// <summary>
        /// The status of the resource.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// The weight of the server. Valid values: `0` to `100`. Default value: `100`. If the value is set to `0`, no
        /// requests are forwarded to the server. **Note:** You do not need to set this parameter if you set `server_type` to `Fc`.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private ServerGroupServer(
            string? description,

            int? port,

            bool? remoteIpEnabled,

            string serverId,

            string? serverIp,

            string serverType,

            string? status,

            int? weight)
        {
            Description = description;
            Port = port;
            RemoteIpEnabled = remoteIpEnabled;
            ServerId = serverId;
            ServerIp = serverIp;
            ServerType = serverType;
            Status = status;
            Weight = weight;
        }
    }
}