using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Sockets;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;
using Grpc.Core;
using Microsoft.Extensions.Logging;
using OpenTelemetry.Exporter.Rest.Protocol;
using Otappd.Proto.Appdynamics.V1;
using static Otappd.Proto.Appdynamics.V1.SpanHandler;
using Protocol = OpenTelemetry.Exporter.Rest.Protocol;

namespace OtisMockInjestService
{
    public class SpanHandlerService : SpanHandlerBase
    {
        private static JsonSerializerOptions SerializerOptions;

        static SpanHandlerService()
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

        private readonly ILogger<SpanHandlerService> _logger;

        public SpanHandlerService(ILogger<SpanHandlerService> logger)
        {
            _logger = logger;
        }

        public override Task<SpansResponse> handleSpans(SpansRequest request, ServerCallContext context)
        {
            var resourceSpans = request.ResourceSpans.Select(s => s.Convert()).ToList();

            Console.WriteLine("=====================================================================");
            Console.WriteLine(JsonSerializer.Serialize(resourceSpans, typeof(List<Protocol.ResourceSpans>), SerializerOptions));

            return Task.FromResult(new SpansResponse
            {
            });
        }
    }
}
