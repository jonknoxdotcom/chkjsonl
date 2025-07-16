# chkjsonl
Check JSONL log file(s) for integrity

## Usage
* The command checks that files consist of valid JSONL records
* You can pass multiple filenames
* Valid JSON lines are indicated by a dot, invalid ones by their line number
* The 'dot stream' after the filename give you a relative idea where a bad record may lie
* Saves time when knee-deep in k8s audit logs

```
% chkjsonl example*
Processing example1.log (1)(2)(3).
Processing example2.log .................(18)..........
Processing example3.log ...........................
%
```
Following this run, I would probably discard `example1.log` as not being a valid log file, and I would vi `example2.log`, go to line 18, and fix what was broken. File `example3.log` is cool.

## Build
```
go build
```


