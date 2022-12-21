// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ServiceMesh.Outputs
{

    [OutputType]
    public sealed class ServiceMeshMeshConfigProxy
    {
        /// <summary>
        /// The CPU resource  of the limitsOPA proxy container.
        /// </summary>
        public readonly string? LimitCpu;
        /// <summary>
        /// The memory resource limit of the OPA proxy container.
        /// </summary>
        public readonly string? LimitMemory;
        /// <summary>
        /// The CPU resource request of the OPA proxy container.
        /// </summary>
        public readonly string? RequestCpu;
        /// <summary>
        /// The memory resource request of the OPA proxy container.
        /// </summary>
        public readonly string? RequestMemory;

        [OutputConstructor]
        private ServiceMeshMeshConfigProxy(
            string? limitCpu,

            string? limitMemory,

            string? requestCpu,

            string? requestMemory)
        {
            LimitCpu = limitCpu;
            LimitMemory = limitMemory;
            RequestCpu = requestCpu;
            RequestMemory = requestMemory;
        }
    }
}