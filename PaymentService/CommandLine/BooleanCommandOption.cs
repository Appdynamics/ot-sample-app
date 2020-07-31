using System;
using System.Collections.Generic;

namespace CommandProcessing
{
    public class BooleanCommandOption : SingleValueCommandOption
    {
        public BooleanCommandOption(string longIdentifier, string shortIdentifier, string description, bool isRequired = false)
            : base(longIdentifier, shortIdentifier, description, isRequired)
        {
        }

        protected override void ParseInternal(string input, ref CommandValue commandValue)
        {
            if (String.Equals("true", input, StringComparison.CurrentCultureIgnoreCase)
                || String.Equals("yes", input, StringComparison.CurrentCultureIgnoreCase)
                || String.Equals("1", input))
            {
                commandValue = new CommandValue(this, true);
            }
            else if (String.Equals("false", input, StringComparison.CurrentCultureIgnoreCase)
                || String.Equals("no", input, StringComparison.CurrentCultureIgnoreCase)
                || String.Equals("0", input))
            {
                commandValue = new CommandValue(this, false);
            }
            else
            {
                throw new CommandLineException(this.LongIdentifier, String.Format("Invalid option {0}", input));
            }
        }

        protected override void ValidateInternal(CommandValue commandValue)
        {
        }

        public override IEnumerable<string> Describe()
        {
            yield return String.Format("--{0} [-{1}] [<true|1|false|0>]", this.LongIdentifier, this.ShortIdentifier);
            yield return String.Format("\t{0}", this.Description);
            if (IsRequired)
            {
                yield return $"\tRequired";
            }
        }
    }
}
