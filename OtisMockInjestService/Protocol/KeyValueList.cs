using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public sealed class KeyValueList
    {
        /// <summary>
        /// A collection of key/value pairs of key-value pairs. The list may be empty (may
        /// contain 0 elements).
        /// </summary>
        public List<KeyValue> Values { get; set; } = new List<KeyValue>();
    }
}
