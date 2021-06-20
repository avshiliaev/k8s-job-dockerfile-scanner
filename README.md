# Implementation

The overall implementation is mainly based on the Bridge pattern. The components are divided into abstractions and
implementations. The main component, called `scanner`, is accepts dependency through injection at runtime. To achieve
this, the scanner is dependent on polymorphism of other components. In other words, every writer or reader, which
implements an accepted interface, can get injected.

The main interface of the scanner object is implemented as a chain of responsibility. There are multiple steps and each
of the steps can through an error while executing, thus stopping the execution of the program. It also makes the error
handling more centralized: the errors are being propagated to the top, and the scanner object is the only one which is
allowed to act upon them.

The application's frontend is built as a CLI. The app accepts command line arguments and decides on the flow of the
program. It makes up a kind of a state machine, where multiple transition between states get performed based on the
configuration passed. It is also a builder pattern, since the order of execution is offloaded to a separate object,
which is not aware of a concrete implementations.

All in all, it allows for further extension of the components. We can introduce different kinds of input readers, api
clients, vendor specific checks and serializers.

# Considerations for improvement

We can introduce a Helm chart to make the collection of manifests more consistent. 
