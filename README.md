# HotSpot [![Build Status][travis-svg]][travis-url]

The package provides an interface to [HotSpot][1].

## [Documentation][doc]

## Installation

Fetch the package:

```bash
go get -d github.com/simulated-reality/hotspot
```

Go to the directory of the package:

```bash
cd $GOPATH/src/github.com/simulated-reality/hotspot
```

Finally, install the package:

```bash
make install
```

## References

* K. Skadron, M. Stan, K. Sankaranarayanan, W. Huang, S. Velusamy, and D.
  Tarjan, “[Temperature-aware microarchitecture: modeling and
  implementation][2],” ACM Transactions Architecture and Code Optimization, vol.
  1, pp. 94–125, March 2004.

## Contributing

1. Fork the project.
2. Implement your idea.
3. Open a pull request.

[1]: http://lava.cs.virginia.edu/HotSpot
[2]: http://www.virginia.edu/cs/people/faculty/pdfs/p94-skadron.pdf

[travis-svg]: https://travis-ci.org/simulated-reality/hotspot.svg?branch=master
[travis-url]: https://travis-ci.org/simulated-reality/hotspot
[doc]: http://godoc.org/github.com/simulated-reality/hotspot
