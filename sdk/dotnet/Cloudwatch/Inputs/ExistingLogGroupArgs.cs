// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Awsx.Cloudwatch.Inputs
{

    /// <summary>
    /// Reference to an existing log group.
    /// </summary>
    public sealed class ExistingLogGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Region of the log group. If not specified, the provider region will be used.
        /// </summary>
        [Input("existing")]
        public Input<string>? Existing { get; set; }

        /// <summary>
        /// Name of the log group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ExistingLogGroupArgs()
        {
        }
    }
}