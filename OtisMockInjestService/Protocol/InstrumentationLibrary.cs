using System;
using System.Collections.Generic;
using System.Text;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public sealed class InstrumentationLibrary
    {
        public string Name { get; set; }
        public string Version { get; set; }

    }
}
