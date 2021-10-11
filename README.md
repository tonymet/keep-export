# Google Keep Export Utility

Google keep doesn't have an API.  To export notes to another app (e.g. One Note,
Evernote), you need to use Google Takeout.  Sometimes this is blocked by admins

This utility let's you export Google Keep --> Google Doc --> HTML --> Tidy --> Separate note HTML files.
Most note apps e.g. Apple Notes, One Note, Evernote import a directory of HTML files easily with bulk import


## Running the Tool

1. Open Google keep, and "eternal scroll" continuously until all notes are on screen. (this can take a while)
2. Select-All notes
1. Hamburger-menu --> Save to Google Doc. At this point all notes will be in a single doc
1. Go to the Doc, File Menu... --> export to HTML. now you have a large HTML with all notes, in GDOc html format




## Converting GDoc HTML into Separate HTML Files
```
brew install tidy-html5
tidy -gdoc GoogleKeepDocument.html > GoogleKeepDocument-tidy.html
go run  main.go < input/GoogleKeepDocument-tidy.html  
```


## Review & Import

Files are in `output` in chronological order e.g. 
```
git:(master) ✗ ls output|head
0000.html
0001.html
0002.html
```

To import, select the `output` directory and bulk import in your tool.