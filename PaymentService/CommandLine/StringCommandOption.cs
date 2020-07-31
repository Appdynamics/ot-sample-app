using System;
using System.Collections.Generic;

namespace CommandProcessing
{
    public class StringCommandOption : SingleValueCommandOption
    {
        public StringCommandOption(string longIdentifier, string shortIdentifier, string description, bool isRequired = false)
            : base(longIdentifier, shortIdentifier, description, isRequired)
        {
        }

        protected override void ParseInternal(string input, ref CommandValue commandValue)
        {
            commandValue = new CommandValue(this, input);
        }

        protected override void ValidateInternal(CommandValue commandValue)
        {
        }

        public override IEnumerable<string> Describe()
        {
            yield return String.Format("--{0} [-{1}] <value>", this.LongIdentifier, this.ShortIdentifier);
            yield return String.Format("\t{0}", this.Description);
            if (IsRequired)
            {
                yield return $"\tRequired";
            }
        }
    }
}
