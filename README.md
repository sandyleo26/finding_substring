Finding Substring
--
## Usage
`$ go run main.go text_to_search subtext`

## Example:
```
$ go run main.go "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!" "Peter"
1, 20, 75
```

```
$ go run main.go "上AB上ab"
1, 4
```

## Description
`main.go` - main file which also implements `Search` and `ToLowerCase`
`main_test.go` - unit tests

The search is case insensitive by first converting both strings to lower case and do a search. The complexity for space and time are O(n) and O(n*m)  respectively. It also considers unicode.