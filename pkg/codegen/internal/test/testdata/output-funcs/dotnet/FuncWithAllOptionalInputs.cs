// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mypkg
{
    public static class FuncWithAllOptionalInputs
    {
        /// <summary>
        /// Check codegen of functions with all optional inputs.
        /// </summary>
        public static Task<FuncWithAllOptionalInputsResult> InvokeAsync(FuncWithAllOptionalInputsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<FuncWithAllOptionalInputsResult>("mypkg::funcWithAllOptionalInputs", args ?? new FuncWithAllOptionalInputsArgs(), options.WithVersion());
    }


    public sealed class FuncWithAllOptionalInputsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Property A
        /// </summary>
        [Input("a")]
        public string? A { get; set; }

        /// <summary>
        /// Property B
        /// </summary>
        [Input("b")]
        public string? B { get; set; }

        public FuncWithAllOptionalInputsArgs()
        {
        }
    }


    [OutputType]
    public sealed class FuncWithAllOptionalInputsResult
    {
        public readonly string R;

        [OutputConstructor]
        private FuncWithAllOptionalInputsResult(string r)
        {
            R = r;
        }
    }
}
