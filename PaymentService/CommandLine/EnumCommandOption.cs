using System;
using System.Collections.Generic;
using System.Reflection;

namespace CommandProcessing
{
    public class EnumCommandOption<TEnum> : SingleValueCommandOption where TEnum : struct
    {
        public EnumCommandOption(string longIdentifier, string shortIdentifier, string description, bool isRequired = false)
            : base(longIdentifier, shortIdentifier, description, isRequired)
        {
            if (!typeof(TEnum).GetTypeInfo().IsEnum)
            {
                throw new ArgumentException("Must provide an enumeration type", "TEnum");
            }
        }

        protected override void ParseInternal(string input, ref CommandValue commandValue)
        {
            TEnum value;
            if (!Enum.TryParse<TEnum>(input, true, out value))
            {
                throw new CommandLineException(this.LongIdentifier, String.Format("Invalid option {0}", input));
            }
            commandValue = new CommandValue(this, value);
        }

        protected override void ValidateInternal(CommandValue commandValue)
        {
        }

        public override IEnumerable<string> Describe()
        {
            yield return String.Format("--{0} [-{1}] <{2}>", this.LongIdentifier, this.ShortIdentifier, typeof(TEnum).Name);
            yield return String.Format("\t{0}", this.Description);
            yield return String.Format("\tPossible values:");
            yield return String.Format("\t\t{0}", String.Join(", ", Enum.GetNames(typeof(TEnum))));
            if (IsRequired)
            {
                yield return "\tRequired";
            }
        }
    }
}
