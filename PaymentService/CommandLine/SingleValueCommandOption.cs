namespace CommandProcessing
{
    public abstract class SingleValueCommandOption : CommandOption
    {
        protected SingleValueCommandOption(string longIdentifier, string shortIdentifier, string description, bool isRequired)
            : base(longIdentifier, shortIdentifier, description, isRequired)
        { }

        public override void Validate(CommandValue commandValue)
        {
            if (commandValue == null)
                throw new CommandLineException($"Expected a value for {this.FullId}");
            if (commandValue.Value == null)
                throw new CommandLineException($"No valid value found for {this.FullId}");
            ValidateInternal(commandValue);
        }

        protected abstract void ValidateInternal(CommandValue commandValue);
    }
}
