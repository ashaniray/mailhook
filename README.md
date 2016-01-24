# Mailhook

Mailhook is a smtp server which triggers webhooks with the mail content as payload. It's behaviour 
can be highly customized using its web based admin interface and Javascript based rules. 

# Installation

Before trying to install mailhook, make sure you have installed a recent version of [Go](https://golang.org/).

To install mailhook just run the following command.

``
$ go get github.com/gophergala2016/mailhook
``

This command will download mailhook in `$GOPATH/src/github.com/gophergala2016/mailhook`. Now change directory to this folder
and run `go build` to create the **mailhook** binary. Now place the mailhook binary to any folder in `$PATH` to run it form anywhere.


# Usage

```
$ mailhook -h
Usage of mailhook:
  -a string
    	web server bind address. (default "0.0.0.0")
  -d string
    	specify rules database file. (default "mailhook.db")
  -p int
    	smpt server port. (default 25)
  -q int
    	webserver port. (default 8080)
  -s string
    	smtp server bind address. (default "0.0.0.0")

```

# Configuring Mailhook

# Writing Mailhook rules
Mailhook can be customized by javascript based rules dispatch webhooks. A sample rule is shown below.

```
rule(function(mail) {

	return true;
});
```
this is the simplest possible rule which always evaluates to `true`. If a rule evaluates to `true` mailhook dispatches the 
webhooks configured with that rule. If the rule function evaluates to `false` mailhook ignores the dispatching of the
webhooks.

The rule function is an anonymous function which receives a mail object as its argument. The structure of the mail argument
is of the following form.

```
{
	From : "bob@example.com",
	To   : ["alice@example.com", "eve@example.com"],
	Body : "This is the mail body"

}
```




