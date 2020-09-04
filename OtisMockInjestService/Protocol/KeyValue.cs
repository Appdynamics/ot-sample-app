using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public sealed class KeyValue
    {
        public string Key { get; set; }
        public AnyValue Value { get; set; }
    }
}
