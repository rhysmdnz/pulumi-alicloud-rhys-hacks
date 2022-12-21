// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Brain.Outputs
{

    [OutputType]
    public sealed class GetIndustrialPidProjectsProjectResult
    {
        /// <summary>
        /// The ID of the Pid Project.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of Pid Organization.
        /// </summary>
        public readonly string PidOrganizationId;
        /// <summary>
        /// The description of Pid Project.
        /// </summary>
        public readonly string PidProjectDesc;
        /// <summary>
        /// The ID of Pid Project.
        /// </summary>
        public readonly string PidProjectId;
        /// <summary>
        /// The name of Pid Project.
        /// </summary>
        public readonly string PidProjectName;

        [OutputConstructor]
        private GetIndustrialPidProjectsProjectResult(
            string id,

            string pidOrganizationId,

            string pidProjectDesc,

            string pidProjectId,

            string pidProjectName)
        {
            Id = id;
            PidOrganizationId = pidOrganizationId;
            PidProjectDesc = pidProjectDesc;
            PidProjectId = pidProjectId;
            PidProjectName = pidProjectName;
        }
    }
}