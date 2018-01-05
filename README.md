# File Tools

Use your favorite text editor to rename files in bulk!
 
Set your `EDITOR` environment variable to your favorite text editor and run `filetools edit`.

_Filetools is still in very early development and hasn't been thoroughly tested. Please use at  your own risk._

## Usage

```
$ filetools edit
 
  -desc
        Sort in descending order
  -path string
        Path (default "C:\\Users\\JohnB\\go\\src\\github.com\\JohnBrainard\\filetools")
  -r    Shorthand for -recursive
  -recursive
        Edit files in current and child directories
  -sort string
        Sort [name, date] (default "name")
  -verbose
        Display extra log messages
