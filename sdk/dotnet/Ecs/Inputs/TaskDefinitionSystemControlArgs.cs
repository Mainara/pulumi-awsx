// *** WARNING: this file was generated by pulumi-gen-awsx. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Awsx.Ecs.Inputs
{

    public sealed class TaskDefinitionSystemControlArgs : Pulumi.ResourceArgs
    {
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("value")]
        public Input<string>? Value { get; set; }

        public TaskDefinitionSystemControlArgs()
        {
        }
    }
}
