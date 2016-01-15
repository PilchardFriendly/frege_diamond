# frege-diamond

Example of pure Frege language project.

## Usage


Run the Frege compiler and then run the Diamond main function:

    ./go
Or...

    lein fregec :run Diamond

Run the Frege compiler and then run the tests:

    ./go test
Or...

    lein fregec :test -v DiamondTest

You can also package up the Frege code and its runtime:

    lein uberjar

That produces a JAR file which you can run:

    java -cp target/frege-diamond-0.1.0-SNAPSHOT-standalone.jar Diamond

## License

Copyright (c) 2015 Your Name

Distributed under the BSD 3-clause License, the same as Frege.
