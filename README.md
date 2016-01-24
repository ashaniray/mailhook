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

To run mailhook (make sure mailhook is in $PATH) simply execute the `mailhook command`. For example to start mailhook at port 2025
execute the following command.

```
$ ./mailhook -p 2025
2016/01/25 02:48:56 starting SMTP endpoint on 0.0.0.0:2025
2016/01/25 02:48:56 starting filter ...
2016/01/25 02:48:56 starting admin web interface on 0.0.0.0:8080
2016/01/25 02:48:56 starting dispatcher

```

By default mailhook start the SMTP interface on port 25 and the admin interface on port 8080, however it can be run to listen
on other ports using the appropriate flag. The exhaustive list of commandline flags are listed below.


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

After running mailhook using the command described in previous section. You can configure mailhook using its web based
admin interface. If mailhook is run with its default flags, the admin interface will listen on port 8080. Open a web
browser and point http://localhost:8080 to access the admin interface. 

The following screen shows the admin UI when opened for the first time after installation.

![](https://github.com/gophergala2016/mailhook/blob/master/screenshots/home.png)

Now click on the "Add Rule" button to start ading rules and endpoints. the image below shows the screen to create 
rules.


![](https://github.com/gophergala2016/mailhook/blob/master/screenshots/create.png)

A sample rule with endpoints configures is shown in the following screenshot.

![](https://github.com/gophergala2016/mailhook/blob/master/screenshots/sample.png)

for more screenshots see [here](https://github.com/gophergala2016/mailhook/blob/master/screenshots/).

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

thus to access the various attributes of the mail object use `object.attribute` syntax. Make sure you access the mail attributes
with capitalized attribute names e.g: `mail.From`.

The following example script logs the attribute values of the mail object and returns false.

```
rule(function(mail) {

	console.log(mail.To);
	console.log(mail.From);
	console.log(mail.Body);
	return false;
});

```


# License

MIT License. Click [here](https://raw.githubusercontent.com/gophergala2016/mailhook/master/LICENSE) to view the license.


