# eagerLaser
### Coding Assessment/Exercise

So here we go....

### Caveat Emptor
This program will retrieve a list of Pull Requests from a Github repository
and display the last 14 days worth.

Depending on the switches used, it will either display a simple console output or 
and HTML formatted suitable for embedded into emails.

I punted on sending the email part, instead to simply show what it looks like in 
the console.


### Options

```--console ``` will show a non-formatted output directly to the console.  If you 
do not choose this option, you will get the fancy HTML output.

```--sender <emailAddress>``` will send from that person.
```--destination <emailAddress>``` will send the email to that group/person.
```--project <owner/project>``` will pull the PRs from that repository.


### Config File
```yaml
---
project: spf13/viper
sender: brad@example.com
destination: developers@example.com
```