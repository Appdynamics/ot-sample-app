using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{

    /// <summary>
    /// Span represents a single operation within a trace. Spans can be
    /// nested to form a trace tree. Spans may also be linked to other spans
    /// from the same or different trace and form graphs. Often, a trace
    /// contains a root span that describes the end-to-end latency, and one
    /// or more subspans for its sub-operations. A trace can also contain
    /// multiple root spans, or none at all. Spans do not need to be
    /// contiguous - there may be gaps or overlaps between spans in a trace.
    ///
    /// The next available field id is 17.
    /// </summary>
    public sealed class Span
    {
        /// <summary>
        /// A unique identifier for a trace. All spans from the same trace share
        /// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes
        /// is considered invalid.
        ///
        /// This field is semantically required. Receiver should generate new
        /// random trace_id if empty or invalid trace_id was received.
        ///
        /// This field is required.
        /// </summary>
        public string TraceId { get; set; }
        /// <summary>
        /// A unique identifier for a span within a trace, assigned when the span
        /// is created. The ID is an 8-byte array. An ID with all zeroes is considered
        /// invalid.
        ///
        /// This field is semantically required. Receiver should generate new
        /// random span_id if empty or invalid span_id was received.
        ///
        /// This field is required.
        /// </summary>
        public string SpanId { get; set; }
        /// <summary>
        /// trace_state conveys information about request position in multiple distributed tracing graphs.
        /// It is a trace_state in w3c-trace-context format: https://www.w3.org/TR/trace-context/#tracestate-header
        /// See also https://github.com/w3c/distributed-tracing for more details about this field.
        /// </summary>
        public string TraceState { get; set; }
        /// <summary>
        /// The `span_id` of this span's parent span. If this is a root span, then this
        /// field must be empty. The ID is an 8-byte array.
        /// </summary>
        public string ParentSpanId { get; set; }
        /// <summary>
        /// A description of the span's operation.
        ///
        /// For example, the name can be a qualified method name or a file name
        /// and a line number where the operation is called. A best practice is to use
        /// the same display name at the same call point in an application.
        /// This makes it easier to correlate spans in different traces.
        ///
        /// This field is semantically required to be set to non-empty string.
        /// When null or empty string received - receiver may use string "name"
        /// as a replacement. There might be smarted algorithms implemented by
        /// receiver to fix the empty span name.
        ///
        /// This field is required.
        /// </summary>
        public string Name { get; set; }
        /// <summary>
        /// Distinguishes between spans generated in a particular context. For example,
        /// two spans with the same name may be distinguished using `CLIENT` (caller)
        /// and `SERVER` (callee) to identify queueing latency associated with the span.
        /// </summary>
        public SpanKind Kind { get; set; }
        /// <summary>
        /// start_time_unix_nano is the start time of the span. On the client side, this is the time
        /// kept by the local machine where the span execution starts. On the server side, this
        /// is the time when the server's application handler starts running.
        /// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
        ///
        /// This field is semantically required and it is expected that end_time >= start_time.
        /// </summary>
        public ulong StartTimeUnixNano { get; set; }
        /// <summary>
        /// end_time_unix_nano is the end time of the span. On the client side, this is the time
        /// kept by the local machine where the span execution ends. On the server side, this
        /// is the time when the server application handler stops running.
        /// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
        ///
        /// This field is semantically required and it is expected that end_time >= start_time.
        /// </summary>
        public ulong EndTimeUnixNano { get; set; }
        /// <summary>
        /// attributes is a collection of key/value pairs. The value can be a string,
        /// an integer, a double or the Boolean values `true` or `false`. Note, global attributes
        /// like server name can be set using the resource API. Examples of attributes:
        ///
        ///     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
        ///     "/http/server_latency": 300
        ///     "abc.com/myattribute": true
        ///     "abc.com/score": 10.239
        /// </summary>
        public List<KeyValue> Attributes { get; set; } = new List<KeyValue>();
        /// <summary>
        /// dropped_attributes_count is the number of attributes that were discarded. Attributes
        /// can be discarded because their keys are too long or because there are too many
        /// attributes. If this value is 0, then no attributes were dropped.
        /// </summary>
        public uint DroppedAttributeCount { get; set; }
        /// <summary>
        /// events is a collection of Event items.
        /// </summary>
        public List<Event> Events { get; set; } = new List<Event>();
        /// <summary>
        /// dropped_events_count is the number of dropped events. If the value is 0, then no
        /// events were dropped.
        /// </summary>
        public uint DroppedEventsCount { get; set; }
        /// <summary>
        /// links is a collection of Links, which are references from this span to a span
        /// in the same or different trace.
        /// </summary>
        public List<Link> Links { get; set; } = new List<Link>();
        /// <summary>
        /// dropped_links_count is the number of dropped links after the maximum size was
        /// enforced. If this value is 0, then no links were dropped.
        /// </summary>
        public uint DroppedLinksCount { get; set; }
        /// <summary>
        /// An optional final status for this span. Semantically when Status
        /// wasn't set it is means span ended without errors and assume
        /// Status.Ok (code = 0).
        /// </summary>
        public Status Status { get; set; }

    }
}
