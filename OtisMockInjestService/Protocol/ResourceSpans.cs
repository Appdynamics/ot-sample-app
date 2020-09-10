using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    /// <summary>
    /// A collection of InstrumentationLibrarySpans from a Resource.
    /// </summary>
    public sealed class ResourceSpans
    {
        /// <summary>
        /// The resource for the spans in this message.
        /// If this field is not set then no resource info is known.
        /// </summary>
        public Resource Resource { get; set; }
        /// <summary>
        /// A list of InstrumentationLibrarySpans that originate from a resource.
        /// </summary>
        public List<InstrumentationLibrarySpans> InstrumentationLibrarySpans { get; set; } = new List<InstrumentationLibrarySpans>();
    }
}
