using System.Collections.Generic;

namespace CommandProcessing
{
    public class ActionCommandOption : CommandOption
    {
        public ActionCommandOption(string longIdentifier, string shortIdentifier, string description)
            : base(longIdentifier, shortIdentifier, description, false)
        {
        }

        public override IEnumerable<string> Describe()
        {
            yield return $"{this.FullId}";
            yield return $"\t{this.Description}";
        }

        public override void Validate(CommandValue commandValue)
        {
        }

        protected override void ParseInternal(string input, ref CommandValue commandValue)
        {
            throw new CommandLineException(this.LongIdentifier, $"Unexpected value '{input}'");
        }
    }
}
