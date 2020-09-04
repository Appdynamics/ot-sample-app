using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    /// <summary>
    /// Event is a time-stamped annotation of the span, consisting of user-supplied
    /// text description and key-value pairs.
    /// </summary>
    public sealed class Event
    {
        /// <summary>
        /// time_unix_nano is the time the event occurred.
        /// </summary>
        public ulong TimeUnixNano { get; set; }
        /// <summary>
        /// name of the event.
        /// This field is semantically required to be set to non-empty string.
        /// </summary>
        public string Name { get; set; }
        /// <summary>
        /// attributes is a collection of attribute key/value pairs on the event.
        /// </summary>
        public List<KeyValue> Attributes { get; set; } = new List<KeyValue>();
        /// <summary>
        /// dropped_attributes_count is the number of dropped attributes. If the value is 0,
        /// then no attributes were dropped.
        /// </summary>
        public uint DroppedAttributesCount { get; set; }
    }
}
