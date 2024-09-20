# Intelscanner

A utility for reading & searching War in the Pacific: Admiral's Edition 
(copyrights not mine, if interested, you should buy it - www.matrixgames.com).

This utility will scan in all the sig int reports in your game, parse the fields,
and allow you to search on base name (for now). It's a wildcard regex search, so
`Shang` will match `Shanghai`; `Naga` will match `Naga Hills` & `Nagasaki`.

Use at your own risk, no warranty is given for this software. If you like it, I'm glad,
if you hate it, that's fine too. If you would like to improve it, please feel free to
raise a PR, I'll review as I get the chance.

## Building

You can build for your native operating system by running

```
go build
```

There is more info on building for other operating systems & architectures in the Go
documentation.

## Configuring

The dist directory has an example configuration file, `intelscanner.example.conf`. To
use it, first rename it to `intelscanner.conf` - this name is hardcoded. In this file
there are two values:

```
folder=/path/to/save/directory
side=Allies -OR- Japan
```

Set folder to the path to the archive folder in your SAVE drectory (you need daily 
archive saves to be set in your War in the Pacific: Admiral's Edition executable flags).

Set Side to be either `Allies` if playing the good guys or `Japan` if playing the evil
empire. This file needs to be placed in the same directory as the intelscanner executable
for it to be picked up.

Example (Windows):
```
folder=C:\Path\To\Witp:AE\Directory\SAVE\archive
side=Allies
```

## Running
The program runs in the command line as one run through - every search will reread & parse
all files in the directory. It has one required positional argument, the substring to search
on.

```
./intel-scanner.exe Shang
```

This will return all sigint events dealing with Shanghai