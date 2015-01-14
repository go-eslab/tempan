# HotSpot

The package provides an interface to [HotSpot][1], a thermal model and simulator
used in computer-architecture studies.

## [Documentation][doc]

## Installation

Fetch the package:

```bash
go get -d github.com/ready-steady/hotspot
```

Go to the directory of the package:

```bash
cd $GOPATH/src/github.com/ready-steady/hotspot
```

Finally, install the package:

```bash
make install
```

## References

* K. Skadron, M. Stan, K. Sankaranarayanan, W. Huang, S. Velusamy, and D.
  Tarjan, “[Temperature-Aware Microarchitecture: Modeling and
  Implementation][2],” ACM Transactions Architecture and Code Optimization, vol.
  1, pp. 94–125, March 2004.

## Contributing

1. Fork the project.
2. Implement your idea.
3. Create a pull request.

[1]: http://lava.cs.virginia.edu/HotSpot/
[2]: http://www.virginia.edu/cs/people/faculty/pdfs/p94-skadron.pdf

[doc]: http://godoc.org/github.com/ready-steady/hotspot
