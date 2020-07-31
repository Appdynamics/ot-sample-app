namespace CommandProcessing
{
    public abstract class MultiValueCommandOption : CommandOption
    {
        public int MinimumValues { get; private set; }
        public int MaximumValues { get; private set; }

        protected MultiValueCommandOption(string longIdentifier, string shortIdentifier, string description, int minValues, int maxValues)
            : base(longIdentifier, shortIdentifier, description, minValues > 0)
        {
            this.MinimumValues = minValues;
            this.MaximumValues = maxValues;
        }
    }
}
