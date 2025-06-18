bookmarkdown
------------

[![Static Badge](https://img.shields.io/badge/Donate-Support_this_Project-orange?style=for-the-badge&logo=buymeacoffee&logoColor=%23ffffff&labelColor=%23333&link=https%3A%2F%2Fxn--gckvb8fzb.com%2Fsupport%2F)](https://xn--gckvb8fzb.com/support/) [![Static Badge](https://img.shields.io/badge/Join_on_Matrix-green?style=for-the-badge&logo=element&logoColor=%23ffffff&label=Chat&labelColor=%23333&color=%230DBD8B&link=https%3A%2F%2Fmatrix.to%2F%23%2F%2521PHlbgZTdrhjkCJrfVY%253Amatrix.org)](https://matrix.to/#/%21PHlbgZTdrhjkCJrfVY%3Amatrix.org)

A super lightweight tool for parsing links from a markdown file, listing the 
titles for dmenu/bemenu/rofi/wofi and retrieving a title back to display the 
corresponding link.


## Build

```sh
go build .
```


## Run

```sh
open $(bookmarkdown "${JRNL}/content/bookmarks/index.md" \
  | bemenu \
  | bookmarkdown "${JRNL}/content/bookmarks/index.md" -)
```

