using System;
using System.Collections.Generic;
using System.Linq;

namespace CommandProcessing
{
    public class ListCommandOption<T, E> : CommandOption where T : CommandOption
    {
        private CommandOption elementOption;
        private CommandValue commandValue;
        private List<E> entries;

        public ListCommandOption(T elementOption)
            : base(elementOption.LongIdentifier, elementOption.ShortIdentifier,
                  elementOption.Description, elementOption.IsRequired)
        {
            this.elementOption = elementOption;
        }

        public override void Validate(CommandValue commandValue)
        {
            throw new NotImplementedException();
        }

        public override IEnumerable<string> Describe()
        {
            return this.elementOption.Describe()
                .Append("\tMultiple allowed");
        }

        protected override void ParseInternal(string input, ref CommandValue commandValue)
        {
            CommandValue innerValue = null;
            this.elementOption.Parse(input, ref innerValue);

            if (this.commandValue == null)
            {
                this.entries = new List<E>();
                this.commandValue = new CommandValue(this, this.entries);
            }

            commandValue = this.commandValue;                
            this.entries.Add(innerValue.GetValue<E>());
        }
    }
}
