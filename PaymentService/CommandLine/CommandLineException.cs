using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace CommandProcessing
{
#if NET40
    [Serializable]
#endif
    public class CommandLineException : Exception
    {
        public CommandLineException()
            : base()
        { }

        public CommandLineException(string message)
            : base(message)
        { }

        public CommandLineException(string identifier, string message)
            : base(String.Format("Error in option {0}: {1}", identifier, message))
        { }
    }
}
