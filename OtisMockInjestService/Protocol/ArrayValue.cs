using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public sealed class ArrayValue
    {
        /// <summary>
        /// Array of values. The array may be empty (contain 0 elements).
        /// </summary>
        public List<AnyValue> Values { get; set; } = new List<AnyValue>();
    }
}
