using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    /// <summary>
    /// The Status type defines a logical error model that is suitable for different
    /// programming environments, including REST APIs and RPC APIs.
    /// </summary>
    public sealed class Status
    {
        /// <summary>
        /// The status code. This is optional field. It is safe to assume 0 (OK)
        /// when not set.
        /// </summary>
        public StatusCode Code { get; set; }
        /// <summary>
        /// A developer-facing human readable error message.
        /// </summary>
        public string Message { get; set; }
    }
}
