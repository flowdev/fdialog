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

### Keyword: `item`
* Keyword: `item`
* Type: `entry`
* Function: display a single line text entry of a form
* Children: none
###### Attributes:
* `label`: label of the text entry
  (required, string type, minimum length: 1)
* `disabled`: is the entry initially disabled?
  (optional, boolean type)
* `hint`: hint text for the entry
  (optional, string type, minimum length: 1)
* `placeHolder`: text initially shown in the entry area
  (optional, string type, minimum length: 1)
* `minLen`: minimum length of a valid entry
  (optional, integer type, minimum value: 0)
* `maxLen`: maximum length of a valid entry
  (optional, integer type, minimum value: 0)
* `regexp`: regular expression that a valid entry has to match
  (optional, string type, minimum length: 1)
* `failText`: text shown if the validation of entry text fails
  (optional, string type, minimum length: 1)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`
* Keyword: `item`
* Type: `multiLineEntry`
* Function: display a multiple line text entry of a form
* Children: none
###### Attributes:
* `label`: label of the text entry
  (required, string type, minimum length: 1)
* `disabled`: is the entry initially disabled?
  (optional, boolean type)
* `hint`: hint text for the entry
  (optional, string type, minimum length: 1)
* `placeHolder`: text initially shown in the entry area
  (optional, string type, minimum length: 1)
* `minLen`: minimum length of a valid entry
  (optional, integer type, minimum value: 0)
* `maxLen`: maximum length of a valid entry
  (optional, integer type, minimum value: 0)
* `regexp`: regular expression that a valid entry has to match
  (optional, string type, minimum length: 1)
* `failText`: text shown if the validation of entry text fails
  (optional, string type, minimum length: 1)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`
* Keyword: `item`
* Type: `passwordEntry`
* Function: display a text entry for passwords (hidden text) of a form
* Children: none
###### Attributes:
* `label`: label of the text entry
  (required, string type, minimum length: 1)
* `disabled`: is the entry initially disabled?
  (optional, boolean type)
* `hint`: hint text for the entry
  (optional, string type, minimum length: 1)
* `placeHolder`: text initially shown in the entry area
  (optional, string type, minimum length: 1)
* `minLen`: minimum length of a valid entry
  (optional, integer type, minimum value: 0)
* `maxLen`: maximum length of a valid entry
  (optional, integer type, minimum value: 0)
* `regexp`: regular expression that a valid entry has to match
  (optional, string type, minimum length: 1)
* `failText`: text shown if the validation of entry text fails
  (optional, string type, minimum length: 1)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`
* Keyword: `item`
* Type: `checkBox`
* Function: display a checkbox of a form
* Children: none
###### Attributes:
* `label`: label of the checkbox
  (required, string type, minimum length: 1)
* `disabled`: is the checkbox initially disabled?
  (optional, boolean type)
* `hint`: hint text for the checkbox
  (optional, string type, minimum length: 1)
* `subLabel`: text displayed next to the checkbox itself
  (optional, string type, minimum length: 1)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`
* Keyword: `item`
* Type: `checkGroup`
* Function: display a group of checkboxes of a form
* Children: none
###### Attributes:
* `label`: label of the checkboxes
  (required, string type, minimum length: 1)
* `disabled`: are the checkboxes initially disabled?
  (optional, boolean type)
* `hint`: hint text for the checkboxes
  (optional, string type, minimum length: 1)
* `options`: texts displayed next to the checkboxes
  (required, list of strings type, minimum length: 1)
* `initiallySelected`: initially selected checkboxes
  (optional, list of strings type, minimum length: 0)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`
* Keyword: `item`
* Type: `radioGroup`
* Function: display a group of radio buttons of a form
* Children: none
###### Attributes:
* `label`: label of the radio button group
  (required, string type, minimum length: 1)
* `disabled`: are the buttons initially disabled?
  (optional, boolean type)
* `hint`: hint text for the radio buttons
  (optional, string type, minimum length: 1)
* `options`: texts displayed next to the checkboxes
  (required, list of strings type, minimum length: 2)
* `initiallySelected`: initially selected checkboxes
  (optional, list of strings type, minimum length: 1)
* `horizontal`: are the buttons arranged horizontally?
  (optional, boolean type)
* `required`: has one button to be selected?
  (optional, boolean type)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`, Type `select`
* Keyword: `item`
* Type: `select`
* Function: display a select entry (a.k.a. drop-down list) of a form
* Children: none
###### Attributes:
* `label`: label of the select entry
  (required, string type, minimum length: 1)
* `disabled`: are the entry initially disabled?
  (optional, boolean type)
* `hint`: hint text for the select entry
  (optional, string type, minimum length: 1)
* `placeHolder`: text initially shown in the entry area
  (optional, string type, minimum length: 1)
* `options`: texts displayed next to the checkboxes
  (required, list of strings type, minimum length: 2)
* `initiallySelected`: initially selected checkboxes
  (optional, string type, minimum length: 1)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`, Type `selectEntry`
* Keyword: `item`
* Type: `selectEntry`
* Function: display a mix of a select and a text entry of a form
* Children: none
###### Attributes:
* `label`: label of the select entry
  (required, string type, minimum length: 1)
* `disabled`: is the entry initially disabled?
  (optional, boolean type)
* `hint`: hint text for the select entry
  (optional, string type, minimum length: 1)
* `placeHolder`: text initially shown in the entry area
  (optional, string type, minimum length: 1)
* `options`: texts displayed next to the checkboxes
  (required, list of strings type, minimum length: 0)
* `minLen`: minimum length of a valid entry
  (optional, integer type, minimum value: 0)
* `maxLen`: maximum length of a valid entry
  (optional, integer type, minimum value: 0)
* `regexp`: regular expression that a valid entry has to match
  (optional, string type, minimum length: 1)
* `failText`: text shown if the validation of entry text fails
  (optional, string type, minimum length: 1)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`, Type `slider`
* Keyword: `item`
* Type: `slider`
* Function: display a value slider in a form
* Children: none
###### Attributes:
* `label`: label of the slider
  (required, string type, minimum length: 1)
* `disabled`: is the slider initially disabled?
  (optional, boolean type)
* `hint`: hint text for the slider
  (optional, string type, minimum length: 1)
* `min`: minimum value of a valid entry
  (optional, float type, any value except `NaN`, `+inf` and `-inf`)
* `max`: maximum value of a valid entry
  (optional, float type, any value except `NaN`, `+inf` and `-inf`)
* `step`: the gap between valid values
  (optional, float type, minimum value: 0.0)
* `initialValue`: initial value of the slider
  (optional, float type, any value except `NaN`, `+inf` and `-inf`)
* `outputKey`: key of the entry data for writing to output
  (optional, string type, valid identifier)

### Keyword: `item`, Type `richText`
* Keyword: `item`
* Type: `richText`
* Function: display some formated text in a form
* Children: none
###### Attributes:
* `label`: label of the text
  (optional, string type, minimum length: 1)
* `hint`: hint text for the slider
  (optional, string type, minimum length: 1)
* `text`: text in MarkDown format to be displayed
  (required, string type, minimum length: 1)
* `scroll`: scrollbars for the text
  (optional, string type, `horizontal`, `vertical`, `both` or `none`)

### Keyword: `item`, Type `hyperlink`
* Keyword: `item`
* Type: `hyperlink`
* Function: display a hyperlink (HTTP or HTTPS) in a form
* Children: none
###### Attributes:
* `label`: label of the text
  (optional, string type, minimum length: 1)
* `hint`: hint text for the slider
  (optional, string type, minimum length: 1)
* `text`: text to be displayed
  (required, string type, minimum length: 1)
* `url`: destination of the link as HTTP or HTTPS
  (required, string type, valid HTTP or HTTPS URL)

### Keyword: `item`, Type `separator`
* Keyword: `item`
* Type: `separator`
* Function: display a separator in a form
* Children: none
###### Attributes:
None.
