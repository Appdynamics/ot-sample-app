using System;
using System.Collections.Generic;
using System.Text;
using System.Text.Json.Serialization;

namespace OpenTelemetry.Exporter.Rest.Protocol
{
    public enum AnyValueType
    {
        None,
        String,
        Bool,
        Int,
        Double,
        Array,
        KeyValueList
    }

    public sealed class AnyValue
    {
        private string stringValue;

        public string StringValue
        {
            get
            {
                if (ValueType == AnyValueType.String)
                    return stringValue;
                return null;
            }
            set
            {
                stringValue = value;
                ValueType = AnyValueType.String;
            }
        }

        private bool boolValue;
        public bool? BoolValue
        {
            get
            {
                if (ValueType == AnyValueType.Bool)
                    return boolValue;
                return null;
            }
            set
            {
                if (value.HasValue)
                {
                    boolValue = value.Value;
                    ValueType = AnyValueType.Bool;
                }
            }
        }

        private long intValue;
        public long? IntValue
        {
            get
            {
                if (ValueType == AnyValueType.Int)
                    return intValue;
                return null;
            }
            set
            {
                if (value.HasValue)
                {
                    intValue = value.Value;
                    ValueType = AnyValueType.Int;
                }
            }
        }

        private double doubleValue;
        public double? DoubleValue
        {
            get
            {
                if (ValueType == AnyValueType.Double)
                    return doubleValue;
                return null;
            }
            set
            {
                if (value.HasValue)
                {
                    doubleValue = value.Value;
                    ValueType = AnyValueType.Double;
                }
            }
        }

        private ArrayValue array;
        public ArrayValue ArrayValue
        {
            get
            {
                if (ValueType == AnyValueType.Array)
                    return array;
                return null;
            }
            set
            {
                array = value;
                ValueType = AnyValueType.Array;
            }
        }

        private KeyValueList kvlist;
        public KeyValueList KvlistValue
        {
            get
            {
                if (ValueType == AnyValueType.KeyValueList)
                    return kvlist;
                return null;
            }
            set
            {
                kvlist = value;
                ValueType = AnyValueType.KeyValueList;
            }
        }

        [JsonIgnore]
        public AnyValueType ValueType { get; private set; } = AnyValueType.None;
    }
}
