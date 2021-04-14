Terraform Provider for Cazoo AWS
==================

- Website: https://www.terraform.io

Maintainers
-----------

This provider plugin is maintained by:

* The Platform Team

Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12+


Using the provider
----------------------


Upgrading the provider
----------------------


Building the provider
---------------------

Clone repository to: `$GOPATH/src/github.com/Cazoo-uk/terraform-provider-cazoo-aws`

```sh
$ mkdir -p $GOPATH/src/github.com/Cazoo-uk; cd $GOPATH/src/github.com/Cazoo-uk
$ git clone git@github.com:Cazoo-uk/terraform-provider-cazoo-aws
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/Cazoo-uk/terraform-provider-cazoo-aws
$ make build
```

Developing the provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org)
installed on your machine (version 1.16.0+ is *required*). You can use [goenv](https://github.com/syndbg/goenv)
to manage your Go version. You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH),
as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`.
This will build the provider and put the provider binary in the `$GOPATH/bin`
directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-cazoo-aws
...
```