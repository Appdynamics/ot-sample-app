using System;
using System.Collections.Generic;

namespace CommandProcessing
{
    public class IntCommandOption : SingleValueCommandOption
    {
        public IntCommandOption(string longIdentifier, string shortIdentifier, string description, bool isRequired = false)
            : base(longIdentifier, shortIdentifier, description, isRequired)
        {
        }

        protected override void ParseInternal(string input, ref CommandValue commandValue)
        {
            int value;
            if (!int.TryParse(input, out value))
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
            yield return String.Format("--{0} [-{1}] <integer>", this.LongIdentifier, this.ShortIdentifier);
            yield return String.Format("\t{0}", this.Description);
            if (IsRequired)
            {
                yield return $"\tRequired";
            }
        }
    }
}
