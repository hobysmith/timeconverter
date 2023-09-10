# Timeconverter Usage Guide

<!-- TOC -->
* [Timeconverter Usage Guide](#timeconverter-usage-guide)
  * [1. Overview](#1-overview)
  * [2. Usage](#2-usage)
    * [2.1 Syntax](#21-syntax)
    * [2.2 Flags](#22-flags)
      * [--input-format, -i](#--input-format--i)
      * [--input-layout, -l](#--input-layout--l)
      * [--output-format, -o](#--output-format--o)
      * [--output-layout, -r](#--output-layout--r)
      * [--output-target, -t](#--output-target--t)
      * [--output-timezone, -z](#--output-timezone--z)
      * [--output-value-only, -v](#--output-value-only--v)
      * [--piped, -p](#--piped--p)
      * [--set-default](#--set-default)
      * [--set-global-default](#--set-global-default)
    * [2.3 Formats](#23-formats)
    * [2.4 Custom Formats](#24-custom-formats)
      * [2.4.1 Custom](#241-custom)
      * [2.4.2 CustomGO](#242-customgo)
    * [2.5 Output Timezones](#25-output-timezones)
    * [2.6 Piping Input](#26-piping-input)
    * [2.7 Piping output](#27-piping-output)
    * [2.8 Setting defaults](#28-setting-defaults)
      * [2.8.1. Local Defaults](#281-local-defaults)
      * [2.8.2 Global Defaults](#282-global-defaults)
      * [2.8.3 Order of defaults and flags](#283-order-of-defaults-and-flags)
      * [2.8.4 A few misc notes about defaults](#284-a-few-misc-notes-about-defaults)
        * [2.8.4.1 Seeing a warning when setting global defaults in a directory with local defaults](#2841-seeing-a-warning-when-setting-global-defaults-in-a-directory-with-local-defaults-)
        * [2.8.4.2 No shortcuts for setting defaults](#2842-no-shortcuts-for-setting-defaults)
  * [3. Commands](#3-commands)
    * [3.1 Clear](#31-clear)
    * [3.2 Help](#32-help)
    * [3.3 Version](#33-version)
    * [3.4 Show](#34-show)
      * [3.4.1 Custom Entities](#341-custom-entities)
      * [3.4.2 Defaults](#342-defaults)
      * [3.4.3 Formats](#343-formats)
  * [4. Building Timeconverter](#4-building-timeconverter)
    * [4.1 all](#41-all)
    * [4.3 build](#43-build)
    * [4.4 clean](#44-clean)
    * [4.5 install](#45-install)
    * [4.5 linux-amd64](#45-linux-amd64)
    * [4.6 linux-arm64](#46-linux-arm64)
    * [4.7 list](#47-list)
    * [4.8 mac-amd64](#48-mac-amd64)
    * [4.9 mac-arm64](#49-mac-arm64)
    * [4.10 test](#410-test)
    * [4.10 win-amd64](#410-win-amd64)
    * [4.11 win-arm64](#411-win-arm64)
<!-- TOC -->

## 1. Overview

**Timeconverter** is a commandline utility that converts time and date values from an input format into an output format.

Various flags are used to indicate the input and output formats, control defaults, and other behaviors.  

## 2. Usage

### 2.1 Syntax

The basic syntax for converting time values is:

    timeconverter timevalue [flags]

**Timeconverter** will read `timevalue` using the default input format, then convert and output
that time using the default output format. **Timeconverter** has built-in formats for input and output values.

For example, when using an input default of UnixSecs and an output default of USDateTimeZ...

    timeconverter 1693848600

Will output this...

    Converted Result: 2023-09-04 12:30:00 -0500

You can also explicitly specify the input and/or output formats like this:

    timeconverter timevalue [-i format] [-o format]

For example, this...

    timeconverter "2023-09-04 12:30:00 -0500" -i USDateTimeZ -o UnixDate

Will output this...

    Converted Result: Mon Sep  4 12:30:00 CDT 2023

Using flags, you can also set your own defaults. For more info on formats,
see [Formats](#formats).

### 2.2 Flags

**Timeconverter** provides flags for a number of uses.

Flags and their values are provided on the commandline by supplying the 
flag formal name or shortcut, followed by the value.  

You can use either form of `--flag=value` or
`--flag value`.  The equal sign is not required.

You can view a summary of all flags definitions in the help output. You can see the 
root level help info using just `timeconverter` or `timeconverter help`.  

To see command level help, you can use `timeconverter command` or `timeconverter command help`.
For example, `timeconverter show` will output the help info relating to the **show** command.

For each flag, there is a formal name and an optional shortcut. Not all flags have a shortcut.

Formal names are indicated on the commandline using two dashes, like `--input-format`.  
Shortcuts are single characters and are indicated with a single dash, like `-i`.

Here's a description of all root level flags.  Note that commands may have additional flags.
For flags that are specific to a command, see that command's info in [Commands](#commands).

#### --input-format, -i
`--input-format` specifies the input format to use when reading the input time value.
When no user defaults are set, **Timeconverter** uses a default format of ***USDateTimeZ***.
Use `--input-format` to indicate a specific time format that is different from the default.

For further info on supported formats, see [Formats](#formats).

#### --input-layout, -l
`--input-layout` specifies the expected formatting template when using a custom format for the input time value. 
For more info on using custom formats, see [Custom Formats](#custom-formats).

#### --output-format, -o
`--output-format` specifies the output format to use when outputting the converted time value.
When no user defaults are set, **Timeconverter** uses a default format of ***USDateTimeZ***.
Use `--output-format` to indicate a specific time format that is different from the default.

For further info on supported formats, see [Formats](#formats).

#### --output-layout, -r
`--output-layout` specifies the expected formatting template when using a custom format for the converted output time.
For more info on using custom formats, see [Custom Formats](#custom-formats).

#### --output-target, -t
`--output-target` indicates where **Timeconverter** should send output.

The options are "console" or "clipboard".
- The `console` option sends output to `stdout`. This is the default.
- The `clipboard` option sends output to the clipboard. 
  While some operating systems have support in some form for sending application output to 
  the clipboard, not all operating systems provide this or do so in a consistent way. This target option 
  provides a consistent mechanism for sending output to the clipboard for all supported operating systems.

***** _**Note**: When sending output to the clipboard, you probably also want to provide the flag `--output-value-only`_
_or `-v`.  This tells **Timeconverter** to only send the converted time value to the output, filtering out_
_any extraneous info.  See [--output-value-only, -v](#--output-value-only--v) for more info._

#### --output-timezone, -z
When converting the input time value, some formats may result in transforming the value into the
local system's timezone.  You can use this flag to explicitly define the timezone context 
of the converted time output.  

For more specific info on defining the output timezone, see
[2.5 Output Timezones](#25-output-timezones).

#### --output-value-only, -v
This flag tells **Timeconverter** to only output the actual time value itself.  Except for critical errors,
it will not output anything other text.  This is useful for sending the output to other apps via piping,
to the clipboard, a file, etc.

#### --piped, -p
This flag tells Timeconverter to expect the input to be piped in from another process.  **Timeconverter** 
automatically detects when input is being piped in, so you would only need this flag if **TimeConverter** 
is failing to detect this.

For more info relating to piping input and output, see 

#### --set-default
This flag will tell **Timeconverter** to save certain flag values as a **local default**.  This can prevent
you from having to enter certain flag settings on every invocation of **Timeconverter**.
See [2.8 Setting defaults](#28-setting-defaults) for more details on saving defaults.

***** _**Note**: The `--set-default` functionality has no shortcut character.  This is so that you cannot accidentally_
_set a local default by means of mistyping a shortcut character._

#### --set-global-default
This flag will tell **Timeconverter** to save certain flag values as a **global default**.  This can prevent
you from having to enter certain flag settings on every invocation of **Timeconverter**.
See [2.8 Setting defaults](#28-setting-defaults) for more details on saving defaults.

**** _**Note**: The `--set-global-default` functionality has no shortcut character.  This is so that you cannot accidentally_
_set a global default by means of mistyping a shortcut character._

### 2.3 Formats
Timeconverter is written in the **Go** language.  As such, it supports all time and date formats defined in 
**Go**'s time package as of Sept 4, 2023.  It also supports a few variants of those formats.

To see a list of all formats that are supported in **Timeconverter**, you can use `timeconverter show -f`.
That will list the name of the format, as well as the template layout for how the time value
is constructed or read.

If you are unfamiliar with **Go**'s time definition syntax, you can reference the **Go** fmt package's documentation
at [Go's time package layout constants](https://pkg.go.dev/time#pkg-constants).

### 2.4 Custom Formats
In addition to the standards based formats, **Timeconverter** allows you to define **custom formats**.
Using custom formats, you can define any layout for time and date components.

To use a custom format, set the format flag to `"custom"` and the layout flag to the layout
of time and date formation you wish.

For example, to set the input and output formats to custom layouts you case use...

    timeconverter "12 2023" -i custom -l "mm yyyy" -o custom -r "yyyy mm"

Which tells **Timeconverter** to read the input as a date composed of a 2-digit month followed by a 4-digit year,
and then output a date composed of a 4-digit year followed by a 2-digit month.  The output from the above
looks like:

    Converted Result: 2023 12

There are two types of custom layout syntax, which are referred to as "**Custom**" and "**CustomGO**".

#### 2.4.1 Custom
The format syntax "**Custom**" refers to **Timeconverter**'s own custom definition syntax.
It is based on the older time and date syntax used in languages like C/C++, Pascal, et al.
However, it varies where certain conventions were either confusing or did not exist.  Generally,
this is where older programming language conventions did not define constructions for international 
elements or more complex timezone constructions.

So, for example, there is a definition for USDateTimeZ format as well as one for EUDateTimeZ format. These
order the month and day in a more colloquial form.

For some users, that older format may feel more familiar or intuitive than Go's layout syntax.
So, this convention is provided for that purpose.

You can see a list of all date and time syntax components in the **Custom** syntax and their related 
time and date elements by using...

    timeconverter show -c

#### 2.4.2 CustomGO
The format syntax **CustomGo** refers to the time and date layout syntax used in the **Go** language.

**Go** has a fairly unique approach to defining time and date constructions. It is defined in the 
standard library at [Go's time package layout constants](https://pkg.go.dev/time#pkg-constants).

Some users may find the Go time and date syntax more familiar or intuitive.  So, this convention
is supplied for them.

### 2.5 Output Timezones
You can specify the timezone using a standard IANA identifier or by using a time offset.

***The IANA timezone identifiers*** can be found at [](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
An example using an IANA timezone identifier would be...

    timeconverter "2023-09-04 11:00:00" -i USDateTime -o RFC3339 -z America/Chicago

Which results in this output value: `2023-09-04T06:00:00-05:00`.

***Offset values*** must be defined in this format "+0000". You must provide a `+` or `-`, followed by
4 digits.  For UTC time, you would use +0000.

An example using a time offset would be...

    timeconverter "2023-09-04 11:00:00 +0000" -i USDateTimeZ -o USDateTimeZ -z America/Chicago

Which results in this output value: `2023-09-04 06:00:00 -0500`.

### 2.6 Piping Input
You can supply date and time values to **Timeconverter** using pipe sequences.
This allows you to read the time and date format from any app, assuming it can be parsed using a
format or a defined custom format that **Timeconverter** can parse.

You can use standard formats, as well as **Custom** and **CustomGo** formats, when piping input.

An example would be...

    echo 1693668084 | timeconverter -i UnixSecs -o RFC3339

If you have saved those input and output formats as defaults, then you can simply use...

    echo 1693668084 | timeconverter

**Timeconverter** should automatically detect when input is being piped in and read it automatically
for the input.  However, if there is a problem where **Timeconverter** is unable to detect this,
like with a particular operating system or a particular system configuration, then you can 
use the `--piped` flag to tell Timeconverter to expect the input from stdin and read it from there.

I have not encountered a scenario yet where the auto-detection does not work.  However,
just in case such a scenario exists, the `--piped` flag is provided as a workaround for that problem.

If neither the auto-detection nor the flag work, then please create an issue and define your specific
scenario.

### 2.7 Piping output
**Timeconverter** supports pipe-friendly output.  To remove any extraneous characters other than the
converted time value, you can use the `--output-value-only` flag. When that is flag is set,
**Timeconverter** will only output the time and date value itself.

This is with the exception of critical errors and warnings. Those will always be output.

An output pipe example would be...

    timeconverter 1693668084 -v -i UnixSecs -o USDateTime | cut -c1-10

Which would output `2023-09-02`.  If you have saved those input and output formats as defaults,
then you can just use `timeconverter 1693668084 -v | cut -c1-10` to do the same thing.

### 2.8 Setting defaults

**Timeconverter** allows you to save certain flag values as defaults.  This means you don't have to
supply those values when executing **Timeconverter**.  

The flags that can be saved to defaults are:

- input-format
- input-layout
- output-format
- output-layout
- output-target
- output-timezone
- output-value-only

There are two types of defaults:
- Local Defaults
- Global Defaults

To set defaults, enter the **Timeconverter** commandline with the values you want to save, then
add one of two flags, depending on what type of default you want to save.

Adding the flag `set-default`, will save the flag values as local defaults. 

Adding the flag `set-global-default` will save the flag values as global defaults.

Defaults are saved in a YAML formatted file on your system as hidden files.  For Windows environments, 
they are simply named `timeconverter.yaml`.  For all other environments, they are named `.timeconverter.yaml`.
The files are standard YAML text, which you can manually edit or remove if you wish.
They must maintain the same name, however.  So don't rename them.  

However, **Timeconverter** provies mechanisms for creating, updating and removing them, as well
as displaying their settings.  So, you don't have to edit them manually.

The command **CLEAR** will allow you to remove your local or global defaults. See [3.1 Clear](#31-clear)
for details on using `clear` to remove default definitions.

The command **SHOW** will show you the current local or global default values. See [3.4.2 Defaults](#342-defaults)
for details on using `show` for displaying your current default values.

WHERE default files are saved is what distinguishes local from global defaults.  The next two sections
will explain the difference between local and global default files.

**Timeconverter** loads your defaults automatically if they exist.  So that you can use...
    
    timeconverter "Fri, 12/1/2023"

Instead of...

    timeconverter "Fri, 12/1/2023" -i custom -l "ddd, mm/d/yyyy" -o custom -r "dddd, mmm d, yyyy"

Both of those executions will return...

    Converted Result: Friday, Dec 1, 2023

#### 2.8.1. Local Defaults
`Local` defaults are stored in the current directory when you run **Timeconverter** and set the
local defaults.  When you run Timeconverter in that directory, it will see and use the default file there.

This means that you can have a unique set of defaults for various directories.  This can
be of value if you do different things in different directories. 

#### 2.8.2 Global Defaults
`Global` defaults are stored in the user's config path.  This path differs per operating system.
But, it means that once you set global defaults, then **Timeconverter** will use those anywhere you run it.
This is convenient if you tend to use common constructions with **Timeconverter**. 

#### 2.8.3 Order of defaults and flags
You can mix and match local and global default settings, as well as use commandline flags in conjuction with
defaults.  Here's the order they are used in:

1. Commandline flags always override any defaults.  
   So, if you pass a value in a commandline flag, which is also set in a default, the commandline flag
   value will be used.  However, if other default values that you do not pass on the commandline,
   the default values will be used for those flags. Basically, this allows you to merge commandline 
   and default values into one Timekeeper execution.
   For example, if you set defaults for an output format of `custom` and also set the custom output layout,
   you can change the layout by simply providing a different output layout on the commandline.

2. Local defaults always override global defaults.
   If local defaults exist AND global defaults exist, then ONLY the local defaults will be loaded.
   Local and Global defaults are never merged together.  However, as described earlier, commandline 
   flags will be merged with whichever defaults are loaded; local or global.

3. If no local defaults exist, Timeconverter checks for global defaults.  If global defaults exist, 
   then they are loaded.  These will be merged with any commandline flags provided at runtime.

#### 2.8.4 A few misc notes about defaults

##### 2.8.4.1 Seeing a warning when setting global defaults in a directory with local defaults 
When setting global defaults using `--set-global-default`, if local defaults exist in that directory, 
you will this warning message:

    Warning: --set-global-default was provided when local defaults exist.
    Local defaults will not be loaded.
    Be sure to provide all required flags when setting global defaults where local defaults also exist.

This warning is shown so that you know that only flags provided on the commandline will be used when
setting the global default values.  Usually local defaults are loaded first, but in this situation, you 
could accidentally mix local defaults with the new global defaults.

To prevent this, **Timeconverter** does not load local defaults when you pass the --set-global-defaults flag.

##### 2.8.4.2 No shortcuts for setting defaults
It is easy to mistype a shortcut and accidentally invoke undesired behavior. In most situations, 
you just get an error from Timeconverter and there are no consequences.

However, when setting defaults accidentally, you can either get the wrong output; or worse, you can
override defaults unintentionally.  

To prevent this, no shortcuts are available for the `set-default` and `set-global-default` flags.
You must enter the formal name.  Since you won't be doing that very often, this shouldn't be a big issue.

## 3. Commands
**Timeconverter** provides several commands.  These use the syntax `timeconverter command [flags]`.

With the exception of `version` and `help`, which do not require flags, you can see help info and
available flags by just entering the command with no flags.

For example...

        timeconverter show

will list the available flag options.

The commands are described in the following sections.

### 3.1 Clear
The clear command will remove the files with default settings.

The flag `--local` will remove a local default file from the current directory as follows:

    timeconverter clear -l

And the flag `--global` will remove the global config file, like so...

    timeconverter clear -g


### 3.2 Help
Any timeconverter commandline can use a `help` command to obtain help about that particular commandline sequence.

For example, `timeconverter help` will show the complete help output for the root help elements.

While `timekeeper show help` will list the help about the show command.

However, if no flags are provided, generally `help` is assumed and the corresponding help output is displayed.

So, it is not required to use the word help to obtain help; but, it is supported if you wish to use it.

### 3.3 Version
Using `timeconverter version` will display the version info, including build time, version, and license reference.

### 3.4 Show
The `show` command displays various lists and info.  Use `timeconverter show` to see help on the available
info that can be provided.

#### 3.4.1 Custom Entities
`timeconverter show -c` will display a list of all the elements in the **Timeconverter** custom format's syntax, 
which are used by the "Custom" format type.  You can use these elements to construct your custom layouts
with the `Custom` format type.

#### 3.4.2 Defaults
`timeconverter show -l` will output the current local default values, if they are set.

`timeconverter show -g` will output the current global default values, if they are set.

These vaules are just simply listed in their YAML form.

#### 3.4.3 Formats
`timeconverter show -f` will display a list all available formats.  

This list will provide the name
of the format, which is what you use in the commandline flag for "-i" (input format) or "-o" (output format).
Also, it will show you the layout pattern that the format uses.  If the format does not use a layout pattern,
such as with Unix time variants, it will provide a brief description of the format's expectations.

## 4. Building Timeconverter
***** _**Note**: The following steps should work out of the box for all platforms, with the possible exception of Windows._
_Depending on your Windows configuration, the tools you have installed, and to some degree your_ 
_level of experience building these kinds of tools, you may have to install some things; specifically,_
_the `make` functionality.  However, it should defnitely be possible to use **make** to build these targets on_
_Windows if you install the necessary things.  Otherwise, you can build the runtime easily by just_
_using `go build`; but you won't have all the runtime metadata available if you do so._

A ***makefile*** definition is provided with this repository.  There are several targets and pseudo targets that
are available. 

Technically, you can build the project using just `go build` or `go install`.  However, the ***make*** targets
will inject the version metadata into the runtime.  So, if you build the project using the **Go** compiler,
just be aware that if you use `timeconverter version` with that build, it will display empty version data.

The **Timeconverter** ***makefile*** supports 8 platform build combinations: seven platform combinations
and the default local environment.  Each of the target platforms will have unique filename suffixes, except
for the default build, which will just be "timeconverter".

The ***makefile*** targets are described in the following sections.

### 4.1 all
`make all` will build all 8 targets.

### 4.3 build
`make build` will build the project using the current platform settings.  The file will be named "timeconverter".

Also, you can run just `make` with no explicit targer reference, which will be functionality equivalent
to `make build`.

### 4.4 clean
`make clear` will clear all **Go** artifacts using `go clean`, then it will remove all of the 8 build targets,
if they exist.

### 4.5 install
`make install` will build and install the Timeconverter binary locally using the **Go** install functionality.
After running this, Timeconverter should be available from any path on your system for the logged-in user.

### 4.5 linux-amd64
`make linux-amd64` will build an AMD64 target for Linux.

### 4.6 linux-arm64
`make linux-arm64` will build an ARM64 target for Linux.

### 4.7 list
`make list` will show a list of all the available targets.

### 4.8 mac-amd64
`make  mac-amd64` will build an AMD64 target for MacOS.

### 4.9 mac-arm64
`make mac-arm64` will build an ARM64 target for MacOS.

### 4.10 test
`make test` will run all **Go** tests.

### 4.10 win-amd64
`make  win-amd64` will build an AMD64 target for Windows.

### 4.11 win-arm64
`make win-arm64` will build an ARM64 target for Windows.