using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace CommandProcessing
{
    /// <summary>
    /// Analyzes the command line inputs
    /// </summary>
    public class CommandLine
    {
        public static CommandLine Build(IEnumerable<CommandOption> commandOptions, string[] arguments)
        {
            var commandLine = new CommandLine(commandOptions);
            if (commandLine.Initialize(arguments))
            {
                return commandLine;
            }
            return null;
        }

        List<CommandOption> options;
        Dictionary<string, CommandValue> commandLineData = new Dictionary<string, CommandValue>();

        /// <summary>
        /// Returns the number of switches
        /// </summary>
        public int Count
        {
            get
            {
                return commandLineData.Count;
            }
        }

        /// <summary>
        /// Returns the arguments for a given command line switch. Returns null if the key was not found.
        /// </summary>
        /// <param name="key"></param>
        /// <returns></returns>
        public CommandValue this[string key]
        {
            get
            {
                if (commandLineData.ContainsKey(key))
                {
                    return commandLineData[key];
                }
                return null;
            }
        }

        public T Values<T>(string key)
        {
            var parameter = this[key];
            if (parameter != null)
            {
                return parameter.GetValue<T>();
            }
            return default(T);
        }

        /// <summary>
        /// Returns all the switches found
        /// </summary>
        public IEnumerable<string> Switches
        {
            get
            {
                return commandLineData.Keys;
            }
        }

        /// <summary>
        /// Returns if a command line parameter for a given command line switch exists.
        /// </summary>
        /// <param name="key"></param>
        /// <returns></returns>
        public bool Exists(string key)
        {
            return commandLineData.ContainsKey(key);
        }

        /// <summary>
        /// Construct the command line object with the raw input arguments.
        /// </summary>
        /// <param name="options"></param>
        private CommandLine(IEnumerable<CommandOption> options)
        {
            this.options = options.ToList();
        }

        #region Helper methods

        private static bool IsCommand(string arg)
        {
            return arg.StartsWith("-") || arg.StartsWith("/");
        }

        private static bool IsLongIdentifier(string arg)
        {
            return arg.StartsWith("--");
        }

        private static string GetLongIdentifier(string arg)
        {
            return arg.Substring(2);
        }

        private static string GetShortIdentifier(string arg)
        {
            return arg.Substring(1);
        }

        private bool Initialize(IList<string> args)
        {
            var hasRequiredArgs = this.options.Any(o => o.IsRequired);

            if ((hasRequiredArgs && args.Count == 0) || (args.Count == 1
                    && (String.Equals(args[0], "--Help", StringComparison.OrdinalIgnoreCase)
                        || String.Equals(args[0], "/Help", StringComparison.OrdinalIgnoreCase)
                        || String.Equals(args[0], "-h", StringComparison.OrdinalIgnoreCase)
                        || String.Equals(args[0], "/h", StringComparison.OrdinalIgnoreCase))))
            {
                Console.WriteLine("\tHelp:");
                Console.WriteLine();
                Console.WriteLine("\t\t--Help [-h]");
                Console.WriteLine("\t\t\tShow help on command line parameters");
                Console.WriteLine();
                Console.WriteLine("\t\t@<filename>");
                Console.WriteLine("\t\t\tProvide a command file where each line contains a command line parameter to use.");
                Console.WriteLine("\t\t\tThis must be the only command line parameter supplied if it is being used.");

                foreach (var option in this.options)
                {
                    Console.WriteLine();
                    foreach (var line in option.Describe())
                    {
                        Console.WriteLine("\t\t{0}", line);
                    }
                }
                return false;
            }

            if (args.Count > 0 && args[0].StartsWith("@"))
            {
                var additionalArgs = ReadCommandFile(args[0]);
                var allArgs = additionalArgs.Concat(args.Skip(1)).ToArray();
                InitializeInternal(allArgs);
            }
            else
            {
                InitializeInternal(args);
            }

            // Validation checks
            foreach (var option in this.options)
            {
                if (option.IsRequired && !this.Exists(option.LongIdentifier))
                {
                    throw new CommandLineException(String.Format("A value for command line parameter '{0}' is required.", option.LongIdentifier));
                }
            }

            return true;
        }

        private void InitializeInternal(IList<string> args)
        {
            if (args.Count > 0 && args[0].StartsWith("@"))
            {
                throw new CommandLineException("Invalid use of '@' not as the first command line argument or inside a command file");
            }

            CommandOption currentOption = null;
            CommandValue currentCommandValue = null;
            foreach (string arg in args)
            {
                if (IsCommand(arg))
                {
                    if (currentOption != null)
                    {
                        AddCommand(currentOption, currentCommandValue);
                    }

                    if (IsLongIdentifier(arg))
                    {
                        var longIdentifier = GetLongIdentifier(arg);
                        currentOption = this.options.Find(
                            o => String.Equals(o.LongIdentifier, longIdentifier, StringComparison.OrdinalIgnoreCase));
                    }
                    else
                    {
                        var shortIdentifier = GetShortIdentifier(arg);
                        currentOption = this.options.Find(
                            o => String.Equals(o.ShortIdentifier, shortIdentifier, StringComparison.OrdinalIgnoreCase));
                    }

                    if (currentOption == null)
                    {
                        throw new CommandLineException("Invalid command line option '" + arg + "'.");
                    }
                    currentCommandValue = null;
                }
                else
                {
                    if (currentOption == null)
                    {
                        throw new CommandLineException("No command line switch present.");
                    }

                    currentOption.Parse(arg, ref currentCommandValue);
                }
            }

            if (currentOption != null)
            {
                AddCommand(currentOption, currentCommandValue);
            }
        }

        private void AddCommand(CommandOption currentCommand, CommandValue currentCommandValue)
        {
            //// We were in a command so store the last values and reset
            //if (currentCommandValue.Count < currentCommand.MinParameters)
            //{
            //    throw new ArgumentException("Command option '" + currentCommand.LongIdentifier + "' has too few parameters. At least " + currentCommand.MinParameters.ToString() + " required.");
            //}
            //else if (currentCommandValue.Count > currentCommand.MaxParameters)
            //{
            //    throw new ArgumentException("Command option '" + currentCommand.LongIdentifier + "' has too many parameters. At most " + currentCommand.MinParameters.ToString() + " required.");
            //}
            this.commandLineData[currentCommand.LongIdentifier] = currentCommandValue;
        }

        private IEnumerable<string> ReadCommandFile(string arg)
        {
            var fileName = arg.TrimStart('@');
            if (!File.Exists(fileName))
            {
#if NET40
                fileName = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, arg.TrimStart('@'));
#else
                fileName = Directory.GetCurrentDirectory();
#endif
                if (!File.Exists(fileName))
                {
                    fileName = arg.TrimStart('@');
                }
            }
            var allLines = File.ReadAllLines(fileName);
            foreach (var line in allLines)
            {
                if (!String.IsNullOrWhiteSpace(line) && !line.StartsWith("//"))
                {
                    var spaceIndex = line.IndexOf(' ');
                    if (IsCommand(line) && spaceIndex > 0)
                    {
                        yield return line.Substring(0, spaceIndex);
                        if (line.Length > spaceIndex)
                            yield return System.Environment.ExpandEnvironmentVariables(line.Substring(spaceIndex + 1));
                    }
                    else
                    {
                        yield return System.Environment.ExpandEnvironmentVariables(line);
                    }
                }
            }
        }

#endregion
    }
}
