// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Vpc.Outputs
{

    [OutputType]
    public sealed class GetIpsecServersServerIkeConfigResult
    {
        /// <summary>
        /// The IKE authentication algorithm.
        /// </summary>
        public readonly string IkeAuthAlg;
        /// <summary>
        /// The IKE encryption algorithm.
        /// </summary>
        public readonly string IkeEncAlg;
        /// <summary>
        /// The IKE lifetime. Unit: seconds.
        /// </summary>
        public readonly int IkeLifetime;
        /// <summary>
        /// The IKE negotiation mode.
        /// </summary>
        public readonly string IkeMode;
        /// <summary>
        /// Diffie-Hellman key exchange algorithm.
        /// </summary>
        public readonly string IkePfs;
        /// <summary>
        /// The IKE version.
        /// </summary>
        public readonly string IkeVersion;
        /// <summary>
        /// IPsec server identifier. Supports the format of FQDN and IP address. The public IP address of the VPN gateway is selected by default.
        /// </summary>
        public readonly string LocalId;
        /// <summary>
        /// The peer identifier. Supports the format of FQDN and IP address, which is empty by default.
        /// </summary>
        public readonly string RemoteId;

        [OutputConstructor]
        private GetIpsecServersServerIkeConfigResult(
            string ikeAuthAlg,

            string ikeEncAlg,

            int ikeLifetime,

            string ikeMode,

            string ikePfs,

            string ikeVersion,

            string localId,

            string remoteId)
        {
            IkeAuthAlg = ikeAuthAlg;
            IkeEncAlg = ikeEncAlg;
            IkeLifetime = ikeLifetime;
            IkeMode = ikeMode;
            IkePfs = ikePfs;
            IkeVersion = ikeVersion;
            LocalId = localId;
            RemoteId = remoteId;
        }
    }
}