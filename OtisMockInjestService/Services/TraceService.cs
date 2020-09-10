using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;
using Grpc.Core;
using Microsoft.Extensions.Logging;
using Opentelemetry.Proto.Collector.Trace.V1;
using OpenTelemetry.Exporter.Rest.Protocol;
using static Opentelemetry.Proto.Collector.Trace.V1.TraceService;
using Protocol = OpenTelemetry.Exporter.Rest.Protocol;

namespace OtisMockInjestService
{
    public class TraceService : TraceServiceBase
    {
        private static JsonSerializerOptions SerializerOptions;

        static TraceService()
        {
            SerializerOptions = new JsonSerializerOptions()
            {
                IgnoreNullValues = true,
                WriteIndented = true,
                PropertyNamingPolicy = JsonNamingPolicy.CamelCase
            };
            SerializerOptions.Converters.Add(new JsonStringEnumConverter());
            SerializerOptions.Converters.Add(new AnyValueJsonConverter());
        }

        private readonly ILogger<TraceService> _logger;

        public TraceService(ILogger<TraceService> logger)
        {
            _logger = logger;
        }

        public override Task<ExportTraceServiceResponse> Export(ExportTraceServiceRequest request, ServerCallContext context)
        {
            var resourceSpans = request.ResourceSpans.Select(s => s.Convert()).ToList();

            Console.WriteLine("=====================================================================");
            Console.Write(DateTime.Now.ToString("o") + " : ");
            Console.WriteLine(JsonSerializer.Serialize(resourceSpans, typeof(List<Protocol.ResourceSpans>), SerializerOptions));

            return Task.FromResult(new ExportTraceServiceResponse
            {
            });
        }
    }
}
