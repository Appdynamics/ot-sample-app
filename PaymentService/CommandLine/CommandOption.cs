using System.Collections.Generic;
using System.Text;
using System.Threading.Tasks;

namespace CommandProcessing
{
    public abstract class CommandOption
    {
        public string LongIdentifier { get; private set; }
        public string ShortIdentifier { get; private set; }
        public string Description { get; private set; }
        public bool IsRequired { get; set; }
        protected string FullId => $"--{LongIdentifier} [-{ShortIdentifier}";

        protected CommandOption(string longIdentifier, string shortIdentifier, string description, bool isRequired)
        {
            this.LongIdentifier = longIdentifier;
            this.ShortIdentifier = shortIdentifier;
            this.Description = description;
            this.IsRequired = isRequired;
        }

        public void Parse(string input, ref CommandValue commandValue)
        {
            if (commandValue != null) throw new CommandLineException(this.LongIdentifier, "Not expecting another parameter: " + input);
            ParseInternal(input, ref commandValue);
        }

        public abstract void Validate(CommandValue commandValue);

        protected abstract void ParseInternal(string input, ref CommandValue commandValue);

        public abstract IEnumerable<string> Describe();
    }
}
