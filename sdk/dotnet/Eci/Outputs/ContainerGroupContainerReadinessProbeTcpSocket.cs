// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eci.Outputs
{

    [OutputType]
    public sealed class ContainerGroupContainerReadinessProbeTcpSocket
    {
        /// <summary>
        /// The port number. Valid values: 1 to 65535.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private ContainerGroupContainerReadinessProbeTcpSocket(int? port)
        {
            Port = port;
        }
    }
}