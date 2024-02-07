module main-package

require (
    subtask1/subtask1 v0.0.0
    subtask2/subtask2 v0.0.0
)

replace subtask1/subtask1 => ./subtask1
replace subtask2/subtask2 => ./subtask2

go 1.21.6
