module github.com/davidroossien/modulemonkey/prototype

go 1.17

require github.com/davidroossien/modulemonkey/semanticsort v0.0.0
require github.com/davidroossien/modulemonkey/simplelog v0.0.0
require github.com/google/uuid v1.3.0 // indirect
replace github.com/davidroossien/modulemonkey/semanticsort => ../semanticsort
replace github.com/davidroossien/modulemonkey/simplelog => ../simplelog
