## font2css

Example usage:

```
./font2css file -f "Nunito Sans" -s normal -w 400 -l "Nunito Sans Regular" -l NunitoSans-Regular -i nunito-sans-regular-400.woff2 -o font.css
```

Output:

```
// font.css
@font-face {
  font-family: 'Nunito Sans';
  font-style: normal;
  font-weight: 400;
  font-display: swap;
  src: local('Nunito Sans Regular'), local('NunitoSans-Regular'), url(data:font/woff2;charset=utf-8;base64,d09GMgABAAAAAEIYABEAAAAAl9QAA[...]3g/VINsIn8Zr+3xS/ib3nn8lqCEkV) format('woff2');
                                                                                                           
}
```

Full usage:

```
usage: font2css [<flags>] <command> [<args> ...]

embeds a font-face as css

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.


  file --family=FAMILY --style=STYLE --weight=WEIGHT --input=INPUT [<flags>]
    make css from font file

    -f, --family=FAMILY      font-family
    -s, --style=STYLE        font-style [normal, italic]
    -w, --weight=WEIGHT      font-weight (numeric)
    -d, --display=swap       font-display [auto, block, swap, fallback, optional]
    -l, --locals=LOCALS ...  local install names, can specify multiple times. ex: -l "Nunito Sans Regular" -l NunitoSans-Regular (optional)
    -u, --unicode=UNICODE    unicode-range. ex -u "U+0000-00FF, U+0131, U+0152-0153, U+02BB-02BC, U+02C6, U+02DA, U+02DC, U+2000-206F, U+2074, U+20AC, U+2122, U+2191, U+2193, U+2212, U+2215, U+FEFF, U+FFFD" (optional)
    -i, --input=INPUT        input filename
    -o, --output=OUTPUT      output file name (optional)

  url --url=URL [<flags>]
    make css from google fonts url

    -u, --url=URL                Google Fonts URL ex: https://fonts.googleapis.com/css2?family=Nunito+Sans:ital,wght@0,200;0,300;0,400;0,600;0,700;0,800;0,900;1,200;1,300;1,400;1,600;1,700;1,800;1,900&display=swap
    -c, --charsets=CHARSETS ...  only do certain charsets, can specify multiple times. ex: -c latin -c latin-ext (optional)
    -s, --styles=STYLES ...      only do certain styles, can specify multiple times. ex: -s normal -s italic (optional)
    -w, --weights=WEIGHTS ...    weights to include, can specify multiple times. ex: -w 200 -w 400 -w 600 -w 900 (optional)
    -o, --output=OUTPUT          output file name (optional)
```
