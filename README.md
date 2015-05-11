EKM is EC2 Keypair Manager
===========================

Installation
--------------

```sh
$ go get github.com/pocke/ekm
```

Usage
-----------

### New

(Not implemented)

```sh
$ ekm new <keypair-name> --region <region>         # => create and output pem to STDOUT
$ ekm new <keypair-name> --region <region> --file  # => create and save pem as <Keypair-name>.pem 
```

### Import

(Not implemented)

```sh
$ ekm import --region <region> --file=<keypair file path>
$ ekm import <keypair-name> --region <region> --file=<keypair file path>
$ cat <keypair file> | ekm import <keypair-name> --region <region>
```

### List

(Not implemented)

```sh
$ ekm list --region <region>  # => list of keypair at <region>
```

### Show

(Not implemented)

```sh
$ ekm show <keypair-name> --region <region> # => Show keypair information
```

### Delete

(Not implemented)

```sh
$ ekm delete <keypair-name> --region <region> # Delete keypair from AWS.
```
