using OpenTelemetry.Exporter.Rest.Protocol;
using System;
using System.Collections.Generic;
using System.Diagnostics.CodeAnalysis;
using System.Linq;
using System.Text.Json;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public class AnyValueJsonConverter : JsonConverter<AnyValue>
    {
        public override AnyValue Read(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            if (reader.TokenType != JsonTokenType.StartObject)
            {
                throw new JsonException("Expected object start");
            }

            reader.Read();
            if (reader.TokenType != JsonTokenType.PropertyName)
            {
                if (reader.TokenType == JsonTokenType.EndObject)
                {
                    // No property to read so None value
                    return new AnyValue();
                }
                throw new JsonException("Expected property name");
            }

            string propertyName = reader.GetString();
            AnyValue value = new AnyValue();
            switch (propertyName)
            {
                case "StringValue":
                    value.StringValue = ReadProperty<string>(ref reader, typeToConvert, options);
                    break;

                case "BoolValue":
                    value.BoolValue = ReadProperty<bool>(ref reader, typeToConvert, options);
                    break;

                case "IntValue":
                    value.IntValue = ReadProperty<long>(ref reader, typeToConvert, options);
                    break;

                case "DoubleValue":
                    value.DoubleValue = ReadProperty<double>(ref reader, typeToConvert, options);
                    break;

                case "ArrayValue":
                    value.ArrayValue = ReadProperty<ArrayValue>(ref reader, typeToConvert, options);
                    break;

                case "KvlistValue":
                    value.KvlistValue = ReadProperty<KeyValueList>(ref reader, typeToConvert, options);
                    break;

                default:
                    throw new JsonException($"Unexpected property name {propertyName}");
            }

            return value;
        }

        public override void Write(Utf8JsonWriter writer, AnyValue value, JsonSerializerOptions options)
        {
            writer.WriteStartObject();
            switch (value.ValueType)
            {
                case AnyValueType.String:
                    WriteProperty(writer, value.StringValue, JsonEncodedText.Encode("StringValue", encoder: null), options);
                    break;

                case AnyValueType.Bool:
                    WriteProperty(writer, value.BoolValue, JsonEncodedText.Encode("BoolValue", encoder: null), options);
                    break;

                case AnyValueType.Int:
                    WriteProperty(writer, value.IntValue, JsonEncodedText.Encode("IntValue", encoder: null), options);
                    break;

                case AnyValueType.Double:
                    WriteProperty(writer, value.DoubleValue, JsonEncodedText.Encode("DoubleValue", encoder: null), options);
                    break;

                case AnyValueType.Array:
                    WriteProperty(writer, value.ArrayValue, JsonEncodedText.Encode("ArrayValue", encoder: null), options);
                    break;

                case AnyValueType.KeyValueList:
                    WriteProperty(writer, value.KvlistValue, JsonEncodedText.Encode("KvlistValue", encoder: null), options);
                    break;

                case AnyValueType.None:
                    // Nothing to write
                    break;

                default:
                    throw new JsonException($"Unsupported/invalid ValueType");
            }
            writer.WriteEndObject();
        }

        private T ReadProperty<T>(ref Utf8JsonReader reader, Type typeToConvert, JsonSerializerOptions options)
        {
            T k;

            // Attempt to use existing converter first before re-entering through JsonSerializer.Deserialize().
            // The default converter for objects does not parse null objects as null, so it is not used here.
            if (typeToConvert != typeof(object) && (options?.GetConverter(typeToConvert) is JsonConverter<T> keyConverter))
            {
                reader.Read();
                k = keyConverter.Read(ref reader, typeToConvert, options);
            }
            else
            {
                k = JsonSerializer.Deserialize<T>(ref reader, options);
            }

            return k;
        }

        private void WriteProperty<T>(Utf8JsonWriter writer, T value, JsonEncodedText name, JsonSerializerOptions options)
        {
            Type typeToConvert = typeof(T);

            writer.WritePropertyName(name);

            // Attempt to use existing converter first before re-entering through JsonSerializer.Serialize().
            // The default converter for object does not support writing.
            if (typeToConvert != typeof(object) && (options?.GetConverter(typeToConvert) is JsonConverter<T> keyConverter))
            {
                keyConverter.Write(writer, value, options);
            }
            else
            {
                JsonSerializer.Serialize<T>(writer, value, options);
            }
        }
    }
}
