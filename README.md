# GoTerm

Go terminal layout improved using [charm](https://github.com/charmbracelet) libraries
²
## Expected Result

```
┌────────────────────────────────────────────────────────────────────────────────────────┐
│  ┌───────DIR───────┐  ┌──────────────────────────TERMINAL───────────────────────────┐  │
│  │                 │  │ λ                                                           │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  │                 │  │                                                             │  │
│  └─────────────────┘  └─────────────────────────────────────────────────────────────┘  │
│                                                                                        │
│                                                                                        │
└────────────────────────────────────────────────────────────────────────────────────────┘

┌ ─ ┐
│   │
├ ┬ ┤
├ ┼ ┤
└ ┴ ┘

```

## TODOs

 - [ ] DIR
   - [x] Show current folder list
   - [x] "up" & "down" navigate through folder list
   - [x] "Enter" navigate in (or out) of the directory & update the list
   - [ ] Add some style
 - [ ] TERMINAL
   - [ ] Show the current directory on the terminal first line
   - [ ] Allow typing and executing commands
   - [ ] Return command result

### Advanced

 - [ ] MAN
   - [ ] On executing --help or man command output the result in the MAN view and replace the DIR view with the MAN view
   - [ ] pressing "m" switch to the MAN view
   - [ ] pressing "d" switch to the DIR view
 - [ ] TERMINAL
   - [ ] On certain command launch animation
     - [ ] mv -> animation of deplacing a file
     - [ ] cp -> aniamtion of factory duplication
     - [ ] rm -> animation of droping in bin
     - [ ] ...
 - [ ] INFO_BAR
   - [ ] Application version
   - [ ] Pressing "?" change the terminal window to an help menu (_what type of information_)
