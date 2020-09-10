using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public sealed class Resource
    {
        /// <summary>
        /// Set of labels that describe the resource.
        /// </summary>
        public List<KeyValue> Attributes { get; set; } = new List<KeyValue>();
        /// <summary>
        /// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
        /// no attributes were dropped.
        /// </summary>
        public uint DroppedAttributesCount { get; set; }
    }
}
