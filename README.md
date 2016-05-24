## topsl

** THIS PROJECT IS NO LONGER MAINTAINED **

The functionality here has been merged (with improvements) into the
tcell Views package. (go get github.com/gdamore/tcell/views)

Please use that package instead.

I'm leaving the rest of this package here for posterity.

[![Linux Status](https://img.shields.io/travis/gdamore/topsl.svg?label=linux)](https://travis-ci.org/gdamore/topsl)
[![Windows Status](https://img.shields.io/appveyor/ci/gdamore/topsl.svg?label=windows)](https://ci.appveyor.com/project/gdamore/topsl)
[![GitHub License](https://img.shields.io/github/license/gdamore/topsl.svg)](https://github.com/gdamore/topsl/blob/master/LICENSE)
[![Issues](https://img.shields.io/github/issues/gdamore/topsl.svg)](https://github.com/gdamore/topsl/issues)
[![Gitter](https://img.shields.io/badge/gitter-join-brightgreen.svg)](https://gitter.im/gdamore/topsl)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/gdamore/topsl)

> _Tops'l is a work in progress (Alpha).
> Please use with caution; at this
> time it is not suitable for production use._

Package topsl is a terminal oriented panels support library.  What it really
does is provide mid-level abstractions to make it easier to build richer
terminal applications, which can include editors, complex dialogs, etc.
It is built upon the excellent termbox-go library.

Note that this is super preliminary work, and was extracted out of a project
to create a nice interface for my govisor project. I wouldn't recommend
its use in other projects at this time.  Hopefully this will change quickly.
