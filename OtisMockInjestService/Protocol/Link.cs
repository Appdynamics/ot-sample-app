using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    /// <summary>
    /// A pointer from the current span to another span in the same trace or in a
    /// different trace. For example, this can be used in batching operations,
    /// where a single batch handler processes multiple requests from different
    /// traces or when the handler receives a request from a different project.
    /// </summary>
    public sealed class Link
    {
        /// <summary>
        /// A unique identifier of a trace that this linked span is part of. The ID is a
        /// 16-byte array.
        /// </summary>
        public string TraceId { get; set; }
        /// <summary>
        /// A unique identifier for the linked span. The ID is an 8-byte array.
        /// </summary>
        public string SpanId { get; set; }
        /// <summary>
        /// The trace_state associated with the link.
        /// </summary>
        public string TraceState { get; set; }
        /// <summary>
        /// attributes is a collection of attribute key/value pairs on the link.
        /// </summary>
        public List<KeyValue> Attributes { get; set; } = new List<KeyValue>();
        /// <summary>
        /// dropped_attributes_count is the number of dropped attributes. If the value is 0,
        /// then no attributes were dropped.
        /// </summary>
        public uint DroppedAttributesCount { get; set; }

    }
}
