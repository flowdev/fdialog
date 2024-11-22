# UIDL - User Interface Description Language

UIDL is a small and simple language for describing user interfaces.
`fdialog` is the reference implementation.

The main element in UIDL is the command.
A command is a keyword followed by a name, any number of attributes in parentheses and
an optional block with more commands.
So a typical command looks like this:

```uidl
keyword name(attr1=123.456, attr2="value2", continue=false) {
    ....
}
```

So the UIDL commands form a tree structure quite like HTML elements.
As every command has to have a name, these names form a name path that can be
used to reference a command anywhere in a UIDL file: `name1.name2.name3`

The attributes of a command can have the following data types:

| Data Type | Examples                   |
|----------:|:---------------------------|
|   integer | `1`, `-3`, `+2`            |
|     float | `-1.3`, `+3.8e4` `7E82`    |
|   boolean | `true`, `false`            |
|    string | `"ab\ncd"`, ` raw string ` |
|      list | `[123, "abc", true]`       |

A UIDL list can contain a mix of different data types but lists of lists are not supported.

The full UIDL grammar with all exact definitions can be found in:
[grammar/UIDL.g4](https://./grammar/UIDL.g4)

The optional string attributes `id` and `group` are allowed for any keyword.
Each of them contains an identifier as value.
`id` is a short cut for it's full name path and can be used for references.
`group` marks the command as part of a group of commands.
This is used for storing input data from the user together as a group and writing it
all together to the output (as JSON object(s)).

The keyword decides together with the special string attribute `type` what the command is
and what attributes are allowed.

## Predefined Commands

The following commands are currently defined:

### Keyword: `window`
* Keyword: `window`
* Type: n/a
* Function: display a window with title bar
* Children: optional, content of the window
###### Attributes:
* `title`: displayed in the title bar
  (optional, string type, minimum length: 1)
* `width`: width of the window
  (optional, float type, minimum value: 50.0)
* `height`: height of the window
  (optional, float type, minimum value: 80.0)
* `appId`: (only main window) ID for loading preferences, etc.
  (optional, string type, minimum length: 1)
* `exitCode`: (only main window) exit code of the app when it ends unexpectedly
  (optional, integer type, values: 0 to 125)

### Keyword: `link`
* Keyword: `link`
* Type: n/a
* Function: link to another keyword in the UI description by using its full name path
* Children: none
###### Attributes:
* `destination`: destination keyword of the link; can be nested with dots (e.g.: main.confirm.dismiss)
  (required, string type, minimum length: 1, regex: `^[\pL\pN_]+(?:[.][\pL\pN_]+)*$`)

### Keyword: `action`, Type: `exit`
* Keyword: `action`
* Type: `exit`
* Function: ends the app and returns a exit code to the calling program
* Children: none
###### Attributes:
* `code`: exit code of the app
  (optional, integer type, values: 0 to 125)

### Keyword: `action`, Type: `close`
* Keyword: `action`
* Type: `close`
* Function: closes a dialog without doing anything else
* Children: none
###### Attributes:
None.

### Keyword: `action`, Type: `group`
* Keyword: `action`
* Type: `group`
* Function: executes a multiple child commands
* Children: at least one command
###### Attributes:
None.

### Keyword: `action`, Type: `write`
* Keyword: `action`
* Type: `write`
* Function: writes data (in JSON format) to standard output
* Children: none
###### Attributes:
* `outputKey`: key of a single data value written
  (optional, string type, valid identifier)
* `fullName`: full name path of the value to write
  (optional, string type, valid identifiers separated by dots (`.`))
* `id`: ID of the value to write
  (optional, string type, valid identifier)
* `group`: group of values to write
  (optional, string type, valid identifier)

All attributes are optional but either `group` or `outputKey` and one of
`id` and `fullName` mut be given.

### Keyword: `form`
* Keyword: `form`
* Type: n/a
* Function: display a form with submit and cancel buttons
* Children: a `submit` child, a `cancel` child and at least one more child for the
  content of the form are required.
###### Attributes:
* `submitText`: text of the submit button
  (optional, string type, minimum length: 1)
* `cancelText`: text of the cancel button
  (optional, string type, minimum length: 1)
