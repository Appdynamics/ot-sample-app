using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public sealed class InstrumentationLibrarySpans
    {
        /// <summary>
        /// The instrumentation library information for the spans in this message.
        /// If this field is not set then no library info is known.
        /// </summary>
        public InstrumentationLibrary InstrumentationLibrary { get; set; }
        /// <summary>
        /// A list of Spans that originate from an instrumentation library.
        /// </summary>
        public List<Span> Spans { get; set; } = new List<Span>();
    }
}
