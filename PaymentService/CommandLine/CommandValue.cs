using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CommandProcessing
{
    public class CommandValue
    {
        public CommandOption Option { get; private set; }

        public object Value { get; protected set; }

        public CommandValue(CommandOption option, object value)
        {
            this.Option = option;
            this.Value = value;
        }

        public T GetValue<T>()
        {
            return (T)Value;
        }
    }
}
