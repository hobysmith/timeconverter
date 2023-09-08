# Timeconverter
**Timeconverter** is a simple, cross-platform commandline utility for 
converting time and date values. 

**Timeconverter** provides support for 30+ standards and formats, as well as
two forms of custom format definitions.  **Timeconverter** also supports 
piping input and output via standard OS pipe chains.

While there are other utilities that provide similar functionality as **Timeconverter**,
the goal for this **Timeconverter** utility is to provide simple and consistent behavior with
a single utility across all common platforms.

## Overview
**Timeconverter** is a commandline utility, which uses a simple commandline structure to read a 
time value from the commandline using one format, then display that time in another format.

This is the basic structure for converting a time value:

    timeconverter timevalue [flags]  

The `timevalue` argument can be the word `now` or a time expressed in any of the supported formats.

For example, this command:

    timeconverter now -o UnixSecs
    
will display the current time in Unix seconds format:

    Converted Result: 1693668084

And this command:

    timeconverter 1693668084 -i UnixSecs -o RFC3339

will read the timevalue in UnixSecs, then convert and display the time in RFC3339 format:

    Converted Result: 2023-09-02T10:21:24-05:00

Basically, **Timeconverter** reads the time value using an input format, then outputs the time using an output format.

**Timeconverter** supports all standards defined in the **Go** standard library as of September 2, 2023.
It also supports a number of variants of those standards, plus two custom formats (Custom and CustomGO) 
that allow you to define any format of time and date construction you wish.

**Timeconverter** flags allow you to set defaults for input and output formats, 
so that you can use a basic syntax like `timeconverter timevalue` to convert 
input values using your default input format, then display them in your default output format.

**Timeconverter** supports piping input as well, so you can use it in a standard pipe chain like this:

    echo 1693668084 | timeconverter -i UnixSecs -o RFC3339

As well as piping outputs like this:

    timeconverter 1693668084 -v | cut -c1-10

## Status
This utility is still in a phase of ongoing development. It is feature complete at this
point, regarding basic expectations of the project. In addition, I have tested it with both
general unit testing coverage and manual testing.  All time conversion formats are covered with
unit tests. I believe it is stable and ready for use.

Nevertheless, there may be bugs, as well as missing or incomplete features. Please test for any 
specific use-cases you may require.  If you do run into issues or have ideas for feature changes,
please submit an issue following the instructions in the section 
[Reporting bugs and making feature requests](#reporting-bugs-and-making-feature-requests).

I do have a few more features that I want to add. Any breaking changes should follow versioning semantics,
unless someone finds a core fault I am not aware of.  But since I do not expect anyone would use this
code for anything other than the utility itself, I don't expect any dependencies to be affected.

Also, I hope to provide pre-built binaries for the various platforms, 
with possible installer type functionality as well, for folks that don't want to build the code themselves.

Finally, I will be updating this GitHub project in several ways.  

So, expect that there may be possible changes in the future. 

## Installing

***** _**Note**: I intend to provide more installation options.  For now, the following options are available._

### Installing from pre-compiled release binaries
You can download a pre-compiled binary or zip from 
the [releases](https://github.com/hobysmith/timeconverter/releases) page.
There are binaries for BSD, Mac, Linux and Windows. Both binaries and archives of the binaries
are available there.

### Installing remotely from GitHub repo
You use this to install the latest version directly from the repo:

        go install github.com/hobysmith/timeconverter@latest

### Installing from the local git repo
Clone the repo.  From the root path of the repo, simply run `go install`.

### Building a binary for local environment or cross-compiling a build using **make**
If you want to build a binary for the local environment:
1. Clone this repo to your local environment.
2. From the local repo root path, just run `make`.

Or, to cross-compile a build for a supported platform, run `make [target]`. The supported targets are:
- mac-arm64
- mac-amd64
- win-arm64
- win-amd64
- linux-arm64
- linux-amd64
- bsd

For example, to build a runtime for Linux on ARM64, run:

    make linux-arm64    

You can run `make list` to view a list of supported build targets and options.

## Usage
To see standard help info, you can use `timeconverter` or `timeconverter --help` or `timekeeper help`.

To see a list of supported time formats, you can use `timeconverter show -o`.

The **Timeconverter** [Usage Guide](GUIDE.md) provides detailed instructions on all **Timeconverter** functionality,
including setting local and global defaults, additional commands, custom formats, 
building for various platforms, etc.

## License
**Timeconverter** uses an MIT license.  For license details see [License](LICENSE).

You are welcome to use this utility and code as directed in the MIT License. 
My goal is just to provide tools that I actually use myself, with the hope that someone 
else may find them useful as well.

## Reporting bugs and making feature requests
Please create an Issue for any bugs you find or suggestions you may have relating to
**Timeconverter** functionality. I will try to respond to these as quickly as I can.

When creating issues for bugs, please prefix the title with "Bug:", like "Bug: Blah Blah feature is not working right."

And for feature requests, please prefix the title with "Feature Request:", like "Feature Request: Adding blah blah functionality would make this utility such the major hotness"

## Contributing
If you wish to contribute, you may fork and submit pull requests. 
Please follow this GitHub guide to do so: 
[GitHub: Contributing to Projects](https://docs.github.com/en/get-started/quickstart/contributing-to-projects) 

I will try to respond to those as I have time.