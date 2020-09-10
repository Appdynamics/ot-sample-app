namespace OpenTelemetry.Exporter.Rest.Protocol
{
    /// <summary>
    /// SpanKind is the type of span. Can be used to specify additional relationships between spans
    /// in addition to a parent/child relationship.
    /// </summary>
    public enum SpanKind
    {
        /// <summary>
        /// Unspecified. Do NOT use as default.
        /// Implementations MAY assume SpanKind to be INTERNAL when receiving UNSPECIFIED.
        /// </summary>
        Unspecified = 0,
        /// <summary>
        /// Indicates that the span represents an internal operation within an application,
        /// as opposed to an operations happening at the boundaries. Default value.
        /// </summary>
        Internal = 1,
        /// <summary>
        /// Indicates that the span covers server-side handling of an RPC or other
        /// remote network request.
        /// </summary>
        Server = 2,
        /// <summary>
        /// Indicates that the span describes a request to some remote service.
        /// </summary>
        Client = 3,
        /// <summary>
        /// Indicates that the span describes a producer sending a message to a broker.
        /// Unlike CLIENT and SERVER, there is often no direct critical path latency relationship
        /// between producer and consumer spans. A PRODUCER span ends when the message was accepted
        /// by the broker while the logical processing of the message might span a much longer time.
        /// </summary>
        Producer = 4,
        /// <summary>
        /// Indicates that the span describes consumer receiving a message from a broker.
        /// Like the PRODUCER kind, there is often no direct critical path latency relationship
        /// between producer and consumer spans.
        /// </summary>
        Consumer = 5,
    }

    /// <summary>
    /// StatusCode mirrors the codes defined at
    /// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/api-tracing.md#statuscanonicalcode
    /// </summary>
    public enum StatusCode
    {
        Ok = 0,
        Cancelled = 1,
        UnknownError = 2,
        InvalidArgument = 3,
        DeadlineExceeded = 4,
        NotFound = 5,
        AlreadyExists = 6,
        PermissionDenied = 7,
        ResourceExhausted = 8,
        FailedPrecondition = 9,
        Aborted = 10,
        OutOfRange = 11,
        Unimplemented = 12,
        InternalError = 13,
        Unavailable = 14,
        DataLoss = 15,
        Unauthenticated = 16,
    }
}
