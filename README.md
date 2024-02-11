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
│   gomder                                                              (press ? help)   │
└────────────────────────────────────────────────────────────────────────────────────────┘
```

## TODOs

 - [x] DIR
   - [x] Show current folder list
   - [x] "up" & "down" navigate through folder list
   - [x] "Enter" navigate in (or out) of the directory & update the list
   - [x] Add some style
 - [x] MULTI WINDOW
   - [x] Add a new window
   - [x] Add a tab navigation between windows
 - [x] TERMINAL
   - [x] Show the current directory on the terminal first line
   - [x] Allow typing and executing commands
   - [x] Return command result

### Advanced

 - [ ] DIR
   - [ ] Show some helps
 - [ ] MAN
   - [ ] On executing --help or man command output the result in the MAN view and replace the DIR view with the MAN view
   - [ ] pressing "m" switch to the MAN view
   - [ ] pressing "d" switch to the DIR view
 - [ ] TERMINAL
   - [ ] Improve command output render
   - [ ] Add command autocomplete on tab (change the way of switching between windows)
   - [ ] On certain command launch animation
     - [ ] mv -> animation of deplacing a file
     - [ ] cp -> aniamtion of factory duplication
     - [ ] rm -> animation of droping in bin
     - [ ] ...
 - [ ] INFO_BAR
   - [ ] Add InfoBar
   - [ ] Application version
   - [ ] Pressing "?" change the terminal window to an help menu (_what type of information_)
