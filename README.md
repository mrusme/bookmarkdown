bookmarkdown
------------

A super lightweight tool for parsing links from a markdown file, listing the 
titles for dmenu/bemenu/rofi/wofi and retrieving a title back to display the 
corresponding link.


## Build

```sh
go build .
```


## Run

```sh
open $(./bookmarkdown ${JRNL}/content/bookmarks/index.md \
  | bemenu \
  | ./bookmarkdown ${JRNL}/content/bookmarks/index.md -)
```

