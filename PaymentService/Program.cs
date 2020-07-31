using System.Collections.Generic;
using System.Diagnostics;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using CommandProcessing;
using Microsoft.Extensions.DependencyInjection;

namespace PaymentService
{
    public enum Exporter
    {
        AppD,
        OTC,
        Jaeger,
        Console
    }

    public class Program
    {
        static IReadOnlyList<CommandOption> _CommandOptions = new List<CommandOption>
        {
            new EnumCommandOption<Exporter>("Exporter", "e", "OpenTelemetry exporter to use")
        };

        public static Exporter SelectedExporter { get; private set; } = Exporter.Console;
        public static string OtcHost { get; private set; }
        public static string JaegerHost { get; private set; }

        public static void Main(string[] args)
        {
            Activity.DefaultIdFormat = ActivityIdFormat.W3C;

            var commandLine = CommandLine.Build(_CommandOptions, args);
            SelectedExporter = commandLine["Exporter"]?.GetValue<Exporter>() ?? Exporter.Console;

            OtcHost = System.Environment.GetEnvironmentVariable("OTC_HOST");
            JaegerHost = System.Environment.GetEnvironmentVariable("JAEGER_HOST");

            var host = CreateHostBuilder(args).Build();
            var logger = host.Services.GetRequiredService<ILogger<Program>>();
            logger.LogInformation("Payment service created.");
            logger.LogInformation($" Args = {string.Join(' ', args)}");
            logger.LogInformation($" * Exporter = {SelectedExporter}");
            logger.LogInformation($" * OTC_HOST = {OtcHost}");
            logger.LogInformation($" * JAEGER_HOST = {JaegerHost}");

            host.Run();
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureLogging(logging =>
                {
                    logging.AddConsole((options) => { options.IncludeScopes = true; });
                })
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.UseUrls("http://*:5000");
                    webBuilder.UseStartup<Startup>();
                });
    }
}
