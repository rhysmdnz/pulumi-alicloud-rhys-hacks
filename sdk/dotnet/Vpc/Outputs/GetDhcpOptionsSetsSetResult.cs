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
    public sealed class GetDhcpOptionsSetsSetResult
    {
        /// <summary>
        /// The Number of VPCs bound by the DHCP option set.
        /// </summary>
        public readonly int AssociateVpcCount;
        /// <summary>
        /// The description of the DHCP options set. The description must be 2 to 256
        /// characters in length and cannot start with `http://` or `https://`.
        /// </summary>
        public readonly string DhcpOptionsSetDescription;
        public readonly string DhcpOptionsSetId;
        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a
        /// Virtual Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the
        /// ECS instances in the VPC network.
        /// </summary>
        public readonly string DhcpOptionsSetName;
        /// <summary>
        /// The root domain, for example, example.com. After a DHCP options set is associated with a Virtual
        /// Private Cloud (VPC) network, the root domain in the DHCP options set is automatically synchronized to the ECS
        /// instances in the VPC network.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The DNS server IP addresses. Up to four DNS server IP addresses can be specified. IP
        /// addresses must be separated with commas (,).
        /// </summary>
        public readonly string DomainNameServers;
        public readonly string Id;
        /// <summary>
        /// The ID of the account to which the DHCP options set belongs.
        /// </summary>
        public readonly string OwnerId;
        /// <summary>
        /// The status of the DHCP options set. Valid values: `Available`, `InUse` or `Pending`. `Available`: The DHCP options set is available for use. `InUse`: The DHCP options set is in use. `Pending`: The DHCP options set is being configured.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetDhcpOptionsSetsSetResult(
            int associateVpcCount,

            string dhcpOptionsSetDescription,

            string dhcpOptionsSetId,

            string dhcpOptionsSetName,

            string domainName,

            string domainNameServers,

            string id,

            string ownerId,

            string status)
        {
            AssociateVpcCount = associateVpcCount;
            DhcpOptionsSetDescription = dhcpOptionsSetDescription;
            DhcpOptionsSetId = dhcpOptionsSetId;
            DhcpOptionsSetName = dhcpOptionsSetName;
            DomainName = domainName;
            DomainNameServers = domainNameServers;
            Id = id;
            OwnerId = ownerId;
            Status = status;
        }
    }
}
