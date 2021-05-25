KAD - Keyboard Automated Design
===============================

[![GoDoc](https://godoc.org/github.com/swill/kad?status.svg)](https://godoc.org/github.com/swill/kad)

KAD is the SVG CAD engine which powers the mechanical keyboard CAD generation site [builder.swillkb.com](http://builder.swillkb.com/).  If you are going to use this library, you should probably review the documentation site for the published builder in order to understand what all the different features are.  The documentation site is available and kept up to date at [builder-docs.swillkb.com](http://builder-docs.swillkb.com/).

KAD is a Golang library to aid in the design of mechanical keyboard plates and cases.  The keyboard layout uses the standard format developed by the [www.keyboard-layout-editor.com](http://www.keyboard-layout-editor.com/) project.  KAD is designed to produce SVG files which can be taken to a laser or water cutting fabrication shop to be cut.  KAD supports a huge number of features, including but not limited it; 4 switch types, 4 stabilizer types, 3 case types, rounded corners, padding, mount holes, usb cutouts, etc...


## Get Started

```
$ go get github.com/swill/kad
```

### Installing DXF dependencies on macOS

You need [Homebrew](https://brew.sh/) installed.

```
$ brew install pstoedit
$ brew install caskformula/caskformula/inkscape
```

## Example

### Usage

You can create a program that uses the KAD library to generate SVG files from the
settings defined in a JSON object. An example program is found in `./example/main.go`.
It was moved there from this page.

You can also build a command line program to easily try out
different settings.

```
$ go build cmd/cli.go
```
Then run it like so:
```
$ ./cli example/layout-01.json
```
It will create both SVG and EPS files for the layers specified in the `JSON` file.

There are also some example `JSON` files in the `./example/` folder.

*For more usage examples, check the `./test/` folder.*


### Output

The example program will produce these files:

| Top Layer | Switch Layer |
|:---------:|:------------:|
| ![Top Layer](https://swill.github.io/kad/usage/usage_example_top.svg) | ![Switch Layer](https://swill.github.io/kad/usage/usage_example_switch.svg) |

| Closed Layer | Open Layer |
|:------------:|:----------:|
| ![Closed Layer](https://swill.github.io/kad/usage/usage_example_closed.svg) | ![Open Layer](https://swill.github.io/kad/usage/usage_example_open.svg) |

| Bottom Layer |
|:------------:|
| ![Bottom Layer](https://swill.github.io/kad/usage/usage_example_bottom.svg) |


## License

```
KAD generates SVG CAD files based on a keyboard layout described in JSON.

Copyright (C) 2015-2016  Will Stevens (swill)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
```

