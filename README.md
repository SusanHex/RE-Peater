# RE-Peater
This tool is intended to print user defined text to the standard output at user defined intervals. Mainly intended to help debug [RE-Actor](https://github.com/SusanHex/RE-Actor "Reactors Github"):.

Here are the variables used to configure RE-Peater:

| Variable Name | Description |
|---------------|-------------|
| `MESSAGES` | This variable holds the messages that RE-Peater will print out. The format of the message should be like this `[<time designation><delay in ms>] <message>[<message separator>]`. Multiple messages may be included using the separator. Here is an example value using the defaut configuration: `<<1000 Hello world;;<<1000 Hello world again`. This will print `Hello world` after `1000` ms, then print `Hello world again` after another `1000` ms. This will go on forever until the program is terminated. |
| `MESSAGE_SEPARATOR` | This variable determines what will seperate each message. The default value is `;;`|
| `TIME_DESIGNATION_SEQUENCE` | This variable determines if a delay will follow. The default is `<<`|

