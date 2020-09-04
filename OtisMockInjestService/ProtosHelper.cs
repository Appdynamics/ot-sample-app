using Google.Protobuf;
using Grpc.Core;
using Otappd.Proto.Common.V1;
using Otappd.Proto.Resource.V1;
using Otappd.Proto.Trace.V1;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using static Otappd.Proto.Trace.V1.Span.Types;
using Protocol = OpenTelemetry.Exporter.Rest.Protocol;
using Status = Otappd.Proto.Trace.V1.Status;
using StatusCode = Otappd.Proto.Trace.V1.Status.Types.StatusCode; 

namespace OtisMockInjestService
{
    static class ResourceSpanHelper
    {
        public static Protocol.ResourceSpans Convert(this ResourceSpans proto)
        {
            return new Protocol.ResourceSpans()
            {
                InstrumentationLibrarySpans = proto.InstrumentationLibrarySpans.Select(s => s.Convert()).ToList(),
                Resource = proto.Resource?.Convert()
            };
        }

        public static Protocol.InstrumentationLibrarySpans Convert(this InstrumentationLibrarySpans proto)
        {
            return new Protocol.InstrumentationLibrarySpans()
            {
                InstrumentationLibrary = proto.InstrumentationLibrary?.Convert(),
                Spans = proto.Spans.Select(s => s.Convert()).ToList()
            };
        }

        public static Protocol.Resource Convert(this Resource proto)
        {
            return new Protocol.Resource()
            {
                DroppedAttributesCount = proto.DroppedAttributesCount,
                Attributes = proto.Attributes.Select(kv => kv.Convert()).ToList()
            };
        }

        public static Protocol.InstrumentationLibrary Convert(this InstrumentationLibrary proto)
        {
            return new Protocol.InstrumentationLibrary()
            {
                Name = proto.Name,
                Version = proto.Version
            };
        }

        public static Protocol.Span Convert(this Span proto)
        {
            return new Protocol.Span()
            {
                TraceId = proto.TraceId?.Convert(),
                SpanId = proto.SpanId?.Convert(),
                TraceState = proto.TraceState,
                ParentSpanId = proto.ParentSpanId?.Convert(),
                Name = proto.Name,
                Kind = proto.Kind.Convert(),
                StartTimeUnixNano = proto.StartTimeUnixNano,
                EndTimeUnixNano = proto.EndTimeUnixNano,
                Attributes = proto.Attributes.Select(a => a.Convert()).ToList(),
                DroppedAttributeCount = proto.DroppedAttributesCount,
                Events = proto.Events.Select(e => e.Convert()).ToList(),
                DroppedEventsCount = proto.DroppedEventsCount,
                Links = proto.Links.Select(l => l.Convert()).ToList(),
                DroppedLinksCount = proto.DroppedLinksCount,
                Status = proto.Status?.Convert()
            };
        }

        public static Protocol.SpanKind Convert(this SpanKind spanKind)
        {
            return (Protocol.SpanKind)((int)spanKind);
        }

        public static Protocol.Status Convert(this Status proto)
        {
            return new Protocol.Status()
            {
                Code = proto.Code.Convert(),
                Message = proto.Message
            };
        }

        public static Protocol.StatusCode Convert(this StatusCode statusCode)
        {
            return (Protocol.StatusCode)((int)statusCode);
        }

        public static Protocol.Event Convert(this Event proto)
        {
            return new Protocol.Event()
            {
                Name = proto.Name,
                TimeUnixNano = proto.TimeUnixNano,
                Attributes = proto.Attributes.Select(a => a.Convert()).ToList(),
                DroppedAttributesCount = proto.DroppedAttributesCount
            };
        }

        public static Protocol.Link Convert(this Link proto)
        {
            return new Protocol.Link()
            {
                TraceId = proto.TraceId?.Convert(),
                SpanId = proto.SpanId?.Convert(),
                TraceState = proto.TraceState,
                Attributes = proto.Attributes.Select(a => a.Convert()).ToList(),
                DroppedAttributesCount = proto.DroppedAttributesCount
            };
        }

        public static Protocol.KeyValue Convert(this KeyValue proto)
        {
            return new Protocol.KeyValue()
            {
                Key = proto.Key,
                Value = proto.Value?.Convert()
            };
        }

        public static Protocol.AnyValue Convert(this AnyValue proto)
        {
            switch (proto.ValueCase)
            {
                case AnyValue.ValueOneofCase.StringValue:
                    return new Protocol.AnyValue() { StringValue = proto.StringValue };

                case AnyValue.ValueOneofCase.IntValue:
                    return new Protocol.AnyValue() { IntValue = proto.IntValue };

                case AnyValue.ValueOneofCase.DoubleValue:
                    return new Protocol.AnyValue() { DoubleValue = proto.DoubleValue };

                case AnyValue.ValueOneofCase.BoolValue:
                    return new Protocol.AnyValue() { BoolValue = proto.BoolValue };

                case AnyValue.ValueOneofCase.ArrayValue:
                    return new Protocol.AnyValue() { ArrayValue = proto.ArrayValue.Convert() };

                case AnyValue.ValueOneofCase.KvlistValue:
                    return new Protocol.AnyValue() { KvlistValue = proto.KvlistValue.Convert() };

                case AnyValue.ValueOneofCase.None:
                default:
                    return new Protocol.AnyValue();
            }
        }

        public static Protocol.ArrayValue Convert(this ArrayValue proto)
        {
            return new Protocol.ArrayValue()
            {
                Values = proto.Values.Select(v => v.Convert()).ToList()
            };
        }

        public static Protocol.KeyValueList Convert(this KeyValueList proto)
        {
            return new Protocol.KeyValueList()
            {
                Values = proto.Values.Select(kv => kv.Convert()).ToList()
            };
        }

        private static string Convert(this ByteString byteString)
        {
            var bytes = byteString.ToByteArray();
            return String.Join("", bytes.ToList().Select(b => b.ToString("x2")));
        }
    }
}
